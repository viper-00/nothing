package database

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/viper-00/nothing/internal/alertstatus"
	"github.com/viper-00/nothing/internal/fileops"
	"github.com/viper-00/nothing/internal/logger"
	"github.com/viper-00/nothing/internal/monitor"
)

type MySql struct {
	DB        *sql.DB
	SqlErr    error
	Connected bool
	Host      string
	Database  string
	User      string
	Password  string
}

func (mysql *MySql) Connect(host, database, user, password string, isMultiStatement bool) {
	mysql.Host = host
	mysql.Database = database
	connStr := user + ":" + password + "@" + "tcp(" + host + ")/" + database
	if isMultiStatement {
		connStr += "?multiStatement=true"
	}
	mysql.DB, mysql.SqlErr = sql.Open("mysql", connStr)
	// Test ping connecting
	mysql.SqlErr = mysql.DB.Ping()
	if mysql.SqlErr != nil {
		mysql.Connected = false
		logger.Log("error", "cannot connect to mysql database "+mysql.SqlErr.Error())
		return
	}
	mysql.Connected = true
}

func (mysql *MySql) Close() {
	mysql.DB.Close()
}

func (mysql *MySql) Init() error {
	q := fileops.ReadFile("init.sql")
	_, err := mysql.DB.Exec(q)
	if err != nil {
		mysql.SqlErr = err
		logger.Log("ERROR", err.Error())
		return err
	}
	return nil
}

func (mysql *MySql) AgentIDExists(agentID string) bool {
	countStr := mysql.monitorDataSelect("SELECT COUNT(*) FROM server WHERE name = ?", agentID)[0]
	count, err := strconv.ParseInt(countStr, 10, 64)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func (mysql *MySql) monitorDataSelect(query string, args ...interface{}) []string {
	rows, err := mysql.DB.Query(query, args...)
	mysql.SqlErr = err
	if err != nil {
		return nil
	}
	defer rows.Close()

	out := []string{}
	for rows.Next() {
		var logText string
		rows.Scan(&logText)
		out = append(out, logText)
	}
	return out
}

func (mysql *MySql) RemoveAgent(agentID string) error {
	stmt, err := mysql.DB.Prepare("DELETE FROM server WHERE name = ?")
	if err != nil {
		mysql.SqlErr = err
		logger.Log("ERROR", err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(agentID)
	return err
}

func (mysql *MySql) ClearAllAlertsWithNullEnd() error {
	query := "UPDATE alert SET end_log_id = 0 WHERE end_log_id IS NULL"
	stmt, err := mysql.DB.Prepare(query)
	if err != nil {
		mysql.SqlErr = err
		logger.Log("ERROR", err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		mysql.SqlErr = err
		logger.Log("ERROR", err.Error())
		return err
	}
	return nil
}

func (mysql *MySql) GetLogFromDBWithId(serverName, logType, logName string, from, to int64) [][]string {
	serverId := mysql.getServerId(serverName)
	if from > 0 && to > 0 {
		query := "SELECT id, log_text FROM system_metrics WHERE server_id = ? AND log_type = ? AND log_name = ? AND log_time BETWEEN ? AND ? ORDER BY log_time"
		res, _ := mysql.Select(query, serverId, logType, logName, from, to)
		return res.Data
	} else {
		query := "SELECT id, log_text FROM system_metrics WHERE server_id = ? AND log_type = ? AND log_name = ? ORDER BY log_time DESC LIMIT 1"
		res, _ := mysql.Select(query, serverId, logType, logName)
		return res.Data
	}
}

func (mysql *MySql) getServerId(serverName string) string {
	res := mysql.monitorDataSelect("SELECT id FROM server WHERE name = ?", serverName)
	if len(res) == 0 {
		return ""
	}
	return res[0]
}

func (mysql *MySql) Select(query string, args ...interface{}) (Table, error) {
	table := Table{}
	row, err := mysql.DB.Query(query, args...)
	if err != nil {
		return table, err
	}
	defer row.Close()

	columns, err := row.Columns()
	mysql.SqlErr = err
	if err != nil {
		return table, err
	}

	output := make([][]string, 0)
	rawResult := make([][]byte, len(columns))
	dest := make([]interface{}, len(columns))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for row.Next() {
		row.Scan(dest...)
		res := make([]string, 0)
		for _, raw := range rawResult {
			if raw != nil {
				res = append(res, string(raw))
			}
		}
		output = append(output, res)
	}

	table.Headers = columns
	table.Data = output
	return table, mysql.SqlErr
}

func (mysql *MySql) GetAlertByStartEvent(logId string) []string {
	t, err := mysql.Select("SELECT * FROM alert WHERE start_log_id = ?", logId)
	if err != nil {
		logger.Log("error", err.Error())
	}
	if len(t.Data) == 0 {
		return nil
	}
	return t.Data[0]
}

func (mysql *MySql) GetPreviousOpenAlert(alertStatus *alertstatus.AlertStatus) []string {
	serverId := mysql.getServerId(alertStatus.Server)
	query := "SELECT * FROM alert As a JOIN system_metrics As m ON a.start_log_id = m.id WHERE a.end_log_id IS NULL AND a.time < ? AND a.server_id = ? AND m.log_type = ?"

	if alertStatus.Alert.MetricName == monitor.DISKS || alertStatus.Alert.MetricName == monitor.NETWORKS || alertStatus.Alert.MetricName == monitor.SERVICES {
		query += "AND m.log_name = ?"
	}

	var (
		table Table
		err   error
	)

	switch alertStatus.Alert.MetricName {
	case monitor.DISKS:
		table, err = mysql.Select(query, alertStatus.UnixTime, serverId, alertStatus.Alert.MetricName, alertStatus.Alert.Disk)
	case monitor.SERVICES:
		table, err = mysql.Select(query, alertStatus.UnixTime, serverId, alertStatus.Alert.MetricName, alertStatus.Alert.Servers)
	default:
		table, err = mysql.Select(query, alertStatus.UnixTime, serverId, alertStatus.Alert.MetricName)
	}

	if err != nil {
		logger.Log("error", "GetPreviousOpenAlert"+err.Error())
	}

	if len(table.Data) == 0 {
		return nil
	}

	return table.Data[0]
}

func (mysql *MySql) SetAlertEndLog(alertStatue *alertstatus.AlertStatus, startEventId string) error {
	serverId := mysql.getServerId(alertStatue.Server)
	if len(serverId) == 0 {
		err := fmt.Errorf("server %s not registered", alertStatue.Server)
		logger.Log("ERROR", err.Error())
		return err
	}

	stmt, err := mysql.DB.Prepare("UPDATE alert SET end_log_id = ? WHERE server_id = ? AND start_log_id = ?")
	if err != nil {
		mysql.SqlErr = err
		logger.Log("ERROR", err.Error())
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(alertStatue.StartEvent, serverId, startEventId)
	if err != nil {
		mysql.SqlErr = err
		logger.Log("ERROR", err.Error())
		return err
	}

	return nil
}

func (mysql *MySql) UpdateAlert(alertStatus *alertstatus.AlertStatus, startEventId string) error {
	serverId := mysql.getServerId(alertStatus.Server)
	if len(serverId) == 0 {
		err := fmt.Errorf("server %s not registered", alertStatus.Server)
		logger.Log("ERROR", err.Error())
		return err
	}

	stmt, err := mysql.DB.Prepare("UPDATE alert SET type = ?, expected = ?, actual = ? WHERE server_id = ? AND start_log_id = ?")
	if err != nil {
		mysql.SqlErr = err
		logger.Log("ERROR", err.Error())
		return err
	}

	defer stmt.Close()

	expectedValue := alertStatus.Alert.WarnThreshold
	if alertStatus.Type == alertstatus.Critical {
		expectedValue = alertStatus.Alert.CriticalThreshold
	}

	_, err = stmt.Exec(alertStatus.Type, expectedValue, alertStatus.Value, serverId, startEventId)
	if err != nil {
		mysql.SqlErr = err
		logger.Log("ERROR", err.Error())
		return err
	}

	return nil
}

func (mysql *MySql) AddAlert(alertStatus *alertstatus.AlertStatus) error {
	serverId := mysql.getServerId(alertStatus.Server)
	if len(serverId) == 0 {
		err := fmt.Errorf("server %s not registered", alertStatus.Server)
		logger.Log("ERROR", err.Error())
		return err
	}

	stmt, err := mysql.DB.Prepare("INSERT INTO alert (server_id, type, expected, actual, time, start_log_id) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		mysql.SqlErr = err
		logger.Log("ERROR", err.Error())
		return err
	}

	defer stmt.Close()

	expectedValue := alertStatus.Alert.WarnThreshold
	if alertStatus.Type == alertstatus.Critical {
		expectedValue = alertStatus.Alert.CriticalThreshold
	}

	_, err = stmt.Exec(serverId, alertStatus.Type, expectedValue, alertStatus.Value, alertStatus.UnixTime, alertStatus.StartEvent)
	if err != nil {
		mysql.SqlErr = err
		logger.Log("ERROR", err.Error())
		return err
	}

	return nil
}

func (mysql *MySql) PurgeMonitorDataOlderThan(unixTime string) (int64, error) {
	query := "DELETE FROM system_metrics WHERE log_time < ?"

	stmt, err := mysql.DB.Prepare(query)
	if err != nil {
		mysql.SqlErr = err
		logger.Log("ERROR", err.Error())
		return -1, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(unixTime)
	if err != nil {
		mysql.SqlErr = err
		logger.Log("ERROR", err.Error())
		return -1, err
	}

	return res.RowsAffected()
}
