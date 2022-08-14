package database

import (
	"database/sql"
	"strconv"

	"github.com/viper-00/nothing/internal/fileops"
	"github.com/viper-00/nothing/internal/logger"
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
