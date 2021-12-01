package promotion

import (
	"example/models/request"
	"example/services/promotion"
	"example/utils"

	"github.com/gofiber/fiber/v2"
)

type promotionHandler struct {
	Service promotion.IPromotionService
}
type IPromotionHandler interface {
	CalculateDiscount(c *fiber.Ctx) error
}

func NewPromotionHandler(service promotion.IPromotionService) promotionHandler {
	return promotionHandler{Service: service}
}

func (h promotionHandler) CalculateDiscount(c *fiber.Ctx) error {
	handleUtil := utils.NewHandle()
	request := request.PromotionRequest{}

	err := handleUtil.BindRequest(c, &request)
	if err != nil {
		return handleUtil.HandleError(c, err)
	}

	err = request.Validate()
	if err != nil {
		return handleUtil.HandleError(c, err)
	}

	response, err := h.Service.CalculateDiscount(request)
	if err != nil {
		return handleUtil.HandleError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
