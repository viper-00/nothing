package alertapi

import (
	"context"

	"github.com/viper-00/nothing/internal/alertstatus"
	"github.com/viper-00/nothing/internal/email"
	"github.com/viper-00/nothing/internal/logger"
	"github.com/viper-00/nothing/internal/monitor"
	"github.com/viper-00/nothing/internal/pagerduty"
	"github.com/viper-00/nothing/internal/slack"
	"github.com/viper-00/nothing/pkg/memdb"
)

type Server struct {
	Database *memdb.Database
}

func (s *Server) HandleAlerts(ctx context.Context, alert *Alert) (*Response, error) {
	metricName := ""
	sendPagerDuty := alert.Pagerduty
	sendEmail := alert.Email
	sendSlack := alert.Slack

	if alert.MetricName == monitor.DISKS {
		metricName = alert.Disk
	}

	if alert.MetricName == monitor.SERVICES {
		metricName = alert.Service
	}

	res := s.Database.Tables["alert"].Where("server_name", "==", alert.ServerName).Add("metric_type", "==", alert.MetricName).Add("metric_name", "==", metricName).Add("resolved", "==", false)
	if res.RowCount == 0 {
		err := s.Database.Tables["alert"].Insert(
			"server_name, metric_type, metric_name, log_id, subject, content, status, timestamp, resolved, pg_incident_id, slack_msg_ts",
			alert.ServerName,
			alert.MetricName,
			metricName,
			alert.LogId,
			alert.Subject,
			alert.Content,
			int(alert.Status),
			alert.Timestamp,
			alert.Resolved,
			"",
			"",
		)
		if err != nil {
			logger.Log("error", "notification_tracker: "+err.Error())
		}
	} else {
		if !res.Rows[0].Columns["resolved"].BoolVal && (alert.Status != int32(alertstatus.Warning) && alert.Status != int32(alertstatus.Critical)) {
			alert.Content += "\n" + res.Rows[0].Columns["content"].StringVal
			res.Update("resolved", true)
			if res.Rows[0].Columns["pg_incident_id"].StringVal != "" {
				err := pagerduty.UpdateIncident(res.Rows[0].Columns["pg_incident_id"].StringVal)
				if err != nil {
					logger.Log("error", err.Error())
				}
				sendPagerDuty = false
			}
			if res.Rows[0].Columns["slack_msg_ts"].StringVal != "" {
				msgTs := res.Rows[0].Columns["slack_msg_ts"].StringVal
				_, err := slack.SendSlackMessage(alert.Subject, alert.Content, alert.SlackChannel, true, msgTs)
				if err != nil {
					logger.Log("error", err.Error())
				}
				sendSlack = false
			}
		}
	}

	if sendEmail {
		err := email.SendEmail(alert.Subject, alert.Content)
		if err != nil {
			logger.Log("error", err.Error())
		}
	}

	if sendPagerDuty {
		id, err := pagerduty.CreateIncident(createIncident(alert.Subject, alert.Content))
		if err != nil {
			logger.Log("error", err.Error())
		}
		res := s.Database.Tables["alert"].Where("server_name", "==", alert.ServerName).Add("metric_type", "==", alert.MetricName).Add("metric_name", "==", metricName).Add("resolved", "==", false)
		res.Update("pg_incident_id", id)
	}

	if sendSlack {
		msgTs, err := slack.SendSlackMessage(alert.Subject, alert.Content, alert.SlackChannel, false, "")
		if err != nil {
			logger.Log("error", err.Error())
		}
		res := s.Database.Tables["alert"].Where("server_name", "==", alert.ServerName).Add("metric_type", "==", alert.MetricName).Add("metric_name", "==", metricName).Add("resolved", "==", false)
		res.Update("slack_msg_ts", msgTs)
	}

	return &Response{Success: true, Msg: "alert processed"}, nil
}

func (s *Server) AlertRequest(ctx context.Context, request *Request) (*AlertArray, error) {
	alerts := AlertArray{}
	res := s.Database.Tables["alert"].Where("server_name", "==", request.ServerName)

	for _, row := range res.Rows {
		alerts.Alerts = append(alerts.Alerts, &Alert{
			ServerName: row.Columns["server_name"].StringVal,
			MetricName: row.Columns["metric_type"].StringVal,
			Subject:    row.Columns["subject"].StringVal,
			Content:    row.Columns["content"].StringVal,
			Timestamp:  row.Columns["timestamp"].StringVal,
			Resolved:   row.Columns["resolved"].BoolVal,
			Disk:       row.Columns["metric_name"].StringVal,
			Service:    row.Columns["metric_name"].StringVal,
		})
	}

	return &alerts, nil
}

func (s *Server) mustEmbedUnimplementedAlertServiceServer() {}

func createIncident(subject, content string) pagerduty.Incident {
	incident := pagerduty.Incident{}
	incident.Incident.Type = "incident"
	incident.Incident.Urgency = "high"
	incident.Incident.Body.Type = "incident_body"
	incident.Incident.Service.Type = "service_reference"
	incident.Incident.Title = subject
	incident.Incident.Body.Details = content
	return incident
}
