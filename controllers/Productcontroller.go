package controllers

import (
	"apiweb/config"
	"apiweb/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func Product(c *fiber.Ctx) error {
	var pd []models.Product
	start, err := strconv.Atoi(c.Params("start"))
	end, enderr := strconv.Atoi(c.Params("end"))
	brand := c.Params("brand")
	kategori := c.Params("kategori")

	if err != nil || enderr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid 'start' parameter",
		})
	}

	var query = config.DB.Where("brands=?", brand).Order("id DESC").Offset(start).Limit(end)

	if kategori != "" {
		query = query.Where("categories=?", kategori)
	}

	if err := query.Find(&pd).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data not found",
				"data":    nil,
			})
		}

		// Handle other errors
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": pd,
	})
}
