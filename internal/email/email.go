package email

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/viper-00/nothing/internal/logger"
)

func SendEmail(subject, content string) error {
	err := godotenv.Load()
	if err != nil {
		logger.Log("Error", "Error loading .env file")
		return err
	}

	emailUser := os.Getenv("EMAIL_USER")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailHost := os.Getenv("EMAIL_HOST")
	emailPort := os.Getenv("EMAIL_PORT")
	emailTo := os.Getenv("ALERT_EMAIL_LIST")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailAuth := smtp.PlainAuth("", emailUser, emailPassword, emailHost)

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	header := "From: " + emailFrom + "\r\n" +
		"To: " + emailTo + "\r\n" +
		"Date: " + time.Now().UTC().Format("Mon Jan 01 00:00:00 -0700 2006") + "\r\n" +
		"Subject: " + subject + "\r\n" + mime + "\r\n"
	msg := []byte(header + "\n" + "<pre>" + content + "</pre>")
	addr := fmt.Sprintf("%s:%s", emailHost, emailPort)
	to := strings.Split(emailTo, ",")

	if err := smtp.SendMail(addr, emailAuth, emailFrom, to, msg); err != nil {
		return err
	}
	return nil
}
