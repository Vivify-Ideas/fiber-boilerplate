package email

import (
	"crypto/tls"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/bytebufferpool"
	"github.com/vivify-ideas/fiber_boilerplate/config"
	"github.com/vivify-ideas/fiber_boilerplate/notifications/definitions"
	mail "github.com/xhit/go-simple-mail/v2"
)

// Send - send email
func Send(notification definitions.Notification) {
	env := config.App.Env
	smtpClient := getClient()

	htmlBody := prepareHTML(notification.Email.Template, notification.Context)

	emails := []string{}
	for _, user := range notification.Users {
		emails = append(emails, user.Email)
	}
	emails = append(emails, notification.Emails...)

	email := mail.NewMSG()
	email.SetFrom(env["EMAIL_FROM"]).
		AddTo(emails...).
		SetSubject(notification.Email.Subject)

	email.SetBody(mail.TextHTML, htmlBody)

	// Call Send and pass the client
	err := email.Send(smtpClient)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent", email.GetRecipients())
	}
}

func getClient() *mail.SMTPClient {
	server := mail.NewSMTPClient()
	env := config.App.Env

	// SMTP Config
	server.Host = env["EMAIL_HOST"]
	serverPort, err := strconv.ParseInt(env["EMAIL_PORT"], 0, 32)
	server.Port = int(serverPort)

	serverUsername := env["EMAIL_USERNAME"]
	if serverUsername != "" {
		server.Username = serverUsername
	}

	serverPassword := env["EMAIL_PASSWORD"]
	if serverUsername != "" {
		server.Password = serverPassword
	}

	server.Encryption = mail.EncryptionTLS

	// Variable to keep alive connection
	server.KeepAlive = false

	// Timeout for connect to SMTP Server
	server.ConnectTimeout = 10 * time.Second

	// Timeout for send the data and wait respond
	server.SendTimeout = 10 * time.Second

	// Set TLSConfig to provide custom TLS configuration. For example,
	// to skip TLS verification (useful for testing):
	server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// SMTP client
	smtpClient, err := server.Connect()

	if err != nil {
		log.Panic(err)
	}

	return smtpClient
}

func prepareHTML(view string, context fiber.Map) string {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	log.Println(view)

	if err := config.App.Views.Render(buf, view, context, "emails/base_layout"); err != nil {
		log.Println(err)
	}
	return buf.String()
}
