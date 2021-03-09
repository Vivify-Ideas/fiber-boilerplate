package services

import (
	"errors"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	Config "github.com/vivify-ideas/fiber_boilerplate/config"
	"github.com/vivify-ideas/fiber_boilerplate/database"
	"github.com/vivify-ideas/fiber_boilerplate/helpers"
	"github.com/vivify-ideas/fiber_boilerplate/models"
	"github.com/vivify-ideas/fiber_boilerplate/notifications"
	"github.com/vivify-ideas/fiber_boilerplate/notifications/definitions"
	"github.com/vivify-ideas/fiber_boilerplate/requests"
	"golang.org/x/crypto/bcrypt"
)

// Result transformation struct
type Result struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

// CreateUser -> user creation
func CreateUser(userData *requests.RegisterRequest) (Result, error) {
	db := database.Init()

	hash, err := hashPassword(string(userData.Password))
	if err != nil {
		return Result{}, err
	}

	user := models.User{
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Email:     userData.Email,
		Password:  hash,
	}
	if err := db.Create(&user).Error; err != nil {
		return Result{}, err
	}

	token, err := GetJWTFromUser(&user)
	if err != nil {
		return Result{}, err
	}

	newUser := Result{
		User:  user,
		Token: token,
	}

	return newUser, nil
}

// Login -> user login
func Login(email string, password string) (Result, error) {
	db := database.Init()
	user := models.User{Email: email}

	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return Result{}, err
	}

	if !checkPasswordHash(password, user.Password) {
		return Result{}, errors.New("Wrong credentials")
	}

	token, err := GetJWTFromUser(&user)
	if err != nil {
		return Result{}, err
	}

	return Result{
		User:  user,
		Token: token,
	}, nil
}

// GetUserFromJWT -> get logged-in user
func GetUserFromJWT(token jwt.Token) (models.User, error) {
	claims := token.Claims.(jwt.MapClaims)
	user := new(models.User)
	user.ID = uint(claims["identity"].(float64))
	db := database.Init()
	err := db.Where("id = ?", user.ID).First(&user).Error
	if err != nil {
		return *user, err
	}

	return *user, err
}

// ResetPassword - reset user password initial step
// Reset password email sent
func ResetPassword(email string) (bool, error) {
	user := new(models.User)
	// user.Email = email

	db := database.Init()
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return false, err
	}
	token, _ := helpers.GenerateRandomStringURLSafe(12)
	user.ResetPasswordToken = token
	db.Save(&user)

	link := Config.App.Env["APP_URL"] + "/api/v1/auth/reset-password-complete?token=" + token

	notifications.Send(notifications.NotifyParams{
		Key: definitions.PasswordReset,
		Context: fiber.Map{
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"link":      link,
		},
		Users: []models.User{*user},
	})

	return true, nil
}

// ResetPasswordComplete - reset user password final step
func ResetPasswordComplete(token string, password string) (bool, error) {
	user := new(models.User)

	db := database.Init()
	err := db.Where("reset_password_token = ?", token).First(&user).Error
	if err != nil {
		return false, err
	}

	hash, err := hashPassword(password)
	if err != nil {
		return false, err
	}
	user.Password = hash
	// reset token
	newToken, _ := helpers.GenerateRandomStringURLSafe(12)
	user.ResetPasswordToken = newToken

	db.Save(&user)

	return true, nil
}

// GetJWTFromUser -> getting token from passed user model
func GetJWTFromUser(user *models.User) (string, error) {
	config := Config.App.Env
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config["APP_SECRET"]))
	return t, err
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
