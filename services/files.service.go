package services

import (
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vivify-ideas/fiber_boilerplate/database"
	"github.com/vivify-ideas/fiber_boilerplate/models"
)

// UploadFiles -> upload files
func UploadFiles(c *fiber.Ctx, files []*multipart.FileHeader, user *models.User) ([]*models.File, error) {
	db := database.Init()

	storedFiles := []*models.File{}

	// Loop through files:
	for _, file := range files {
		// Save the files to disk:
		relativePath := fmt.Sprintf("public/uploads/%d_%s", time.Now().Unix(), file.Filename)
		err := c.SaveFile(file, relativePath)
		if err != nil {
			return nil, err
		}

		file := models.File{
			Path:   strings.Replace(relativePath, "public", "", 1),
			UserId: int(user.ID),
			User:   *user,
		}
		storedFiles = append(storedFiles, &file)

		if err := db.Create(&file).Error; err != nil {
			return nil, err
		}
	}
	return storedFiles, nil
}
