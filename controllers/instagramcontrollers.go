package controllers

import (
	"apiweb/config"
	"apiweb/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

func Instagramshow(c *fiber.Ctx) error {
	var instagram []models.Insta
	startStr := c.Params("start")
	endsStr := c.Params("end")

	if startStr == "" || endsStr == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing 'start' or 'end' parameters",
		})

	}
	starts, err := strconv.Atoi(startStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid 'start' parameter",
		})
	}

	ends, err := strconv.Atoi(endsStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid 'end' parameter",
		})
	}
	config.DB.Order("timestamp DESC").Offset(starts).Limit(ends).Find(&instagram)
	return c.JSON(fiber.Map{
		"data": instagram,
	})
}
