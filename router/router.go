package router

import (
	"example/delivery/http/promotion"

	"github.com/gofiber/fiber/v2"
)

type router struct {
	App              *fiber.App
	Router           fiber.Router
	PromotionHandler promotion.IPromotionHandler
}

var instantiated *router = nil

func New(
	app *fiber.App,
	demoHandler promotion.IPromotionHandler,
) *router {
	if instantiated == nil {
		instantiated = &router{
			App:              app,
			Router:           app.Group("api/v1"),
			PromotionHandler: demoHandler,
		}
		instantiated.init()
	}
	return instantiated
}

func (r *router) init() {
	r.setupDemo()
}

func (r *router) setupDemo() {
	api := r.Router.Group("promotion")
	{
		api.Post("/discount", r.PromotionHandler.CalculateDiscount)
	}
}
