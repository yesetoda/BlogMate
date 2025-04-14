package infrastructure

import (
	"log"
	"net/smtp"

	"github.com/yesetoda/BlogMate/config"
)

func SendEmail(toEmail string, title string, body string, link string) error {
	cfg, err := config.LoadConfig() // Renamed variable to `cfg`
	if err != nil {
		log.Println("Failed to load config:", err)
		return err // Return error to avoid using a nil/incomplete config
	}

	// Construct HTML email message
	message := `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>` + title + `</title>
	</head>
	<body>
		<h1>` + title + `</h1>
		<p>` + body + `</p>
		<a href="` + link + `">Click the Link</a>
	</body>
	</html>
	`

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	key := cfg.Email.Key
	host := "smtp.gmail.com"
	auth := smtp.PlainAuth("", "yeneineh.seiba@a2sv.org", key, host)

	port := "587"
	address := host + ":" + port
	messages := []byte(mime + message)

	// Send the email via Gmail's SMTP server
	err = smtp.SendMail(address, auth, "yeneineh.seiba@a2sv.org", []string{toEmail}, messages)
	if err != nil {
		log.Println("send mail error:", err)
		return err
	}
	return nil
}
