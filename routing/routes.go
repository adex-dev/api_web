package routing

import (
	"apiweb/controllers"
	"github.com/gofiber/fiber/v2"
)

func Routes(r *fiber.App) {
	app := r.Group("api")
	isoide := app.Group("/isoide")

	isoide.Get("instagram/:start/:end", controllers.Instagramshow)
	isoide.Get("categories/:brands", controllers.Categoriesshow)
	isoide.Get("product/:start/:end/:brand", controllers.Product)
	isoide.Get("product/:start/:end/:brand/:kategori", controllers.Product)
	isoide.Get("event/:start/:end", controllers.EventIsoide)
	isoide.Get("outlet", controllers.StoreIsoide)
	isoide.Post("promosi", controllers.Promosishow)
}
