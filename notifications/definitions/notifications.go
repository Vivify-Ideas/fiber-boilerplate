package definitions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vivify-ideas/fiber_boilerplate/models"
)

// **Notification keys start**

// PasswordReset - reset password notification
const PasswordReset = "password_reset"

// **Notification keys end**

// EmailNotification - email notification
type EmailNotification struct {
	Template string
	Subject  string
}

// InAppNotification - in-app notification
type InAppNotification struct {
	Title   string
	Message string
}

// Notification - single notification fields
type Notification struct {
	Email   EmailNotification
	InApp   InAppNotification
	Context fiber.Map
	Users   []models.User
	Emails  []string
}

// Notifications list
var Notifications = map[string]Notification{
	PasswordReset: {
		Email: EmailNotification{
			Template: "emails/password_reset",
			Subject:  "Password Reset",
		},
	},
}
