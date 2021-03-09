package notifications

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vivify-ideas/fiber_boilerplate/config"
	"github.com/vivify-ideas/fiber_boilerplate/models"
	"github.com/vivify-ideas/fiber_boilerplate/notifications/channels/email"
	"github.com/vivify-ideas/fiber_boilerplate/notifications/definitions"
)

type NotifyParams struct {
	Key     string
	Context fiber.Map
	Users   []models.User
	Emails  []string
}

func Send(notifyParams NotifyParams) {
	notification := definitions.Notifications[notifyParams.Key]
	notification.Users = notifyParams.Users
	notification.Emails = notifyParams.Emails

	notifyParams.Context["APP_URL"] = config.App.Env["APP_URL"]
	notification.Context = notifyParams.Context

	if notification.Email.Template != "" {
		email.Send(notification)
	}
}
