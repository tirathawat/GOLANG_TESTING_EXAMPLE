package promotion_test

import (
	"bytes"
	"encoding/json"
	"time"

	handler "example/delivery/http/promotion"
	"example/models/storages"
	"example/repositories"
	services "example/services/promotion"

	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscountIntegrationService(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(storages.PromotionDB{
			ID:               1,
			Name:             "promotion1",
			PurchaseMin:      100,
			DiscountPercent:  20,
			ExpiredDate:      time.Now().Local().Add(time.Hour * 12),
			CreatedTimestamp: time.Now().Local(),
			UpdateTimestamp:  time.Now().Local(),
		}, nil)

		promoService := services.NewPromotionService(promoRepo)
		promoHandler := handler.NewPromotionHandler(promoService)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"promotion_id": 1,
			"amount":       100,
		})

		//{{url}}/discount
		app := fiber.New()
		app.Post("/discount", promoHandler.CalculateDiscount)

		req := httptest.NewRequest(http.MethodPost, "/discount", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		//Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		//Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, "80", string(body))
		}
	})
}
