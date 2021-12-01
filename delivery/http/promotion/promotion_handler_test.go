package promotion_test

import (
	"bytes"
	"encoding/json"
	"errors"
	handlers "example/delivery/http/promotion"
	"example/errs"
	services "example/services/promotion"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {

	type testCase struct {
		name           string
		requestBody    map[string]interface{}
		expectedResult float64
		expectedErr    *map[string]string
	}

	ErrCases := []testCase{

		{
			name: "dont have promotion id",
			requestBody: map[string]interface{}{
				"amount": 100,
			},
			expectedResult: 0,
			expectedErr: &map[string]string{
				"th_message": errs.PROMO_ID_NOT_FOUND_TH,
				"en_message": errs.PROMO_ID_NOT_FOUND_EN,
			},
		},

		{
			name: "dont have amount",
			requestBody: map[string]interface{}{
				"promotion_id": 1,
			},
			expectedResult: 0,
			expectedErr: &map[string]string{
				"th_message": errs.PURCHASE_NOT_FOUND_TH,
				"en_message": errs.PURCHASE_NOT_FOUND_EN,
			},
		},

		{
			name: "bad request (incorrect type)",
			requestBody: map[string]interface{}{
				"promotion_id": "1",
				"amount":       "100",
			},
			expectedResult: 0,
			expectedErr: &map[string]string{
				"th_message": errs.BAD_REQUEST_ERROR_TH,
				"en_message": errs.BAD_REQUEST_ERROR_EN,
			},
		},
	}

	for _, c := range ErrCases {
		t.Run(c.name, func(t *testing.T) {
			//Arrange
			requestBody, _ := json.Marshal(c.requestBody)

			promoService := services.NewPromotionServiceMock()
			promoService.On("CalculateDiscount").Return(c.expectedResult, nil)
			promoHandler := handlers.NewPromotionHandler(promoService)

			//{{url}}/discount
			app := fiber.New()
			app.Post("/discount", promoHandler.CalculateDiscount)

			req := httptest.NewRequest(http.MethodPost, "/discount", bytes.NewBuffer(requestBody))
			req.Header.Set("Content-Type", "application/json")

			//Act
			res, _ := app.Test(req)
			defer res.Body.Close()

			//Assert
			if assert.Equal(t, fiber.StatusBadRequest, res.StatusCode) {
				body, _ := io.ReadAll(res.Body)
				m := make(map[string]string)
				_ = json.Unmarshal(body, &m)
				assert.Equal(t, *c.expectedErr, m)
			}
		})
	}

	t.Run("success", func(t *testing.T) {

		//Arrange
		requestBody, _ := json.Marshal(map[string]interface{}{
			"promotion_id": 1,
			"amount":       100,
		})

		expected := 80.0

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount").Return(expected, nil)
		promoHandler := handlers.NewPromotionHandler(promoService)

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
			assert.Equal(t, fmt.Sprintf("%v", expected), string(body))
		}
	})

	t.Run("service error", func(t *testing.T) {

		//Arrange
		requestBody, _ := json.Marshal(map[string]interface{}{
			"promotion_id": 1,
			"amount":       100,
		})

		expected := 80.0

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount").Return(expected, errors.New(""))
		promoHandler := handlers.NewPromotionHandler(promoService)

		//{{url}}/discount
		app := fiber.New()
		app.Post("/discount", promoHandler.CalculateDiscount)

		req := httptest.NewRequest(http.MethodPost, "/discount", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		//Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		//Assert
		assert.NotEqual(t, fiber.StatusOK, res.StatusCode)
	})
}
