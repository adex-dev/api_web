package controllers

import (
	"apiweb/config"
	"apiweb/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func Categoriesshow(c *fiber.Ctx) error {
	var cate []models.Categories_tb
	brand := c.Params("brands")

	if err := config.DB.Where("brand=? AND status=?", brand, "ACTIVE").Order("name ASC").Find(&cate).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Terjadi kesalahan server",
			})
		}
	}

	return c.JSON(fiber.Map{
		"data": cate,
	})
}
func Promosishow(c *fiber.Ctx) error {
	var promosi []models.Promosi
	brand := c.FormValue("brands")

	if err := config.DB.Where("brands=? AND status=?", brand, "ACTIVE").Order("addons DESC").Find(&promosi).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"data": "Not Fond",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": promosi,
	})
}
func StoreIsoide(c *fiber.Ctx) error {
	var store []models.Store
	brand := "ISOIDE"

	if err := config.DBMASTER.Where("brands=? AND status=?", brand, "AKTIF").Order("nama_outlet ASC").Find(&store).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"data": "Not Fond",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": store,
	})
}

func EventIsoide(c *fiber.Ctx) error {
	var eventlist []models.Evenisoide
	start, err := strconv.Atoi(c.Params("start"))
	end, enderr := strconv.Atoi(c.Params("end"))
	if err != nil || enderr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid 'start' parameter",
		})
	}
	if err := config.DB.Where("status=?", "ACTIVE").Offset(start).Limit(end).Order("created_at DESC").Find(&eventlist).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"data": "Not Fond",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data": eventlist,
	})
}
