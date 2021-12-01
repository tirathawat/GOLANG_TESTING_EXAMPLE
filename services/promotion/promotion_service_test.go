package promotion_test

import (
	"errors"
	"example/errs"
	"example/models/request"
	"example/models/storages"
	"example/repositories"
	"example/services/promotion"
	"example/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {

	typeUtil := utils.NewType()

	type testCase struct {
		name        string
		promotionDB storages.PromotionDB
		request     request.PromotionRequest
		expected    float64
	}

	promotion1 := storages.PromotionDB{
		ID:               1,
		Name:             "promotion1",
		PurchaseMin:      100,
		DiscountPercent:  20,
		ExpiredDate:      time.Now().Local().Add(time.Hour * 12),
		CreatedTimestamp: time.Now().Local(),
		UpdateTimestamp:  time.Now().Local(),
	}

	promotion2 := storages.PromotionDB{
		ID:               2,
		Name:             "promotion2",
		PurchaseMin:      100,
		DiscountPercent:  20,
		ExpiredDate:      time.Now().Local().Add(time.Hour * 12 * -1),
		CreatedTimestamp: time.Now().Local(),
		UpdateTimestamp:  time.Now().Local(),
	}

	cases := []testCase{
		{
			name:        "equal purchase min",
			promotionDB: promotion1,
			request: request.PromotionRequest{
				PromotionID: typeUtil.GetAddressInt(1),
				Amount:      typeUtil.GetAddressFloat64(100),
			},
			expected: 80,
		},

		{
			name:        "more than purchase min",
			promotionDB: promotion1,
			request: request.PromotionRequest{
				PromotionID: typeUtil.GetAddressInt(1),
				Amount:      typeUtil.GetAddressFloat64(200),
			},
			expected: 160,
		},

		{
			name:        "less than purchase min",
			promotionDB: promotion1,
			request: request.PromotionRequest{
				PromotionID: typeUtil.GetAddressInt(1),
				Amount:      typeUtil.GetAddressFloat64(100),
			},
			expected: 80,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			//Arrage
			promoRepo := repositories.NewPromotionRepositoryMock()
			promoRepo.On("GetPromotion").Return(c.promotionDB, nil)
			promoService := promotion.NewPromotionService(promoRepo)

			//Act
			discount, _ := promoService.CalculateDiscount(c.request)
			expected := c.expected

			//Assert
			assert.Equal(t, expected, discount)
		})
	}

	t.Run("expired promotion error", func(t *testing.T) {
		//Arrage
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(promotion2, nil)
		promoService := promotion.NewPromotionService(promoRepo)

		//Act
		_, err := promoService.CalculateDiscount(request.PromotionRequest{
			PromotionID: typeUtil.GetAddressInt(1),
			Amount:      typeUtil.GetAddressFloat64(200),
		})

		//Assert
		assert.ErrorIs(t, err, errs.ErrPromoExpired)
	})

	t.Run("get promotion error", func(t *testing.T) {
		//Arrage
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(storages.PromotionDB{}, errors.New(""))
		promoService := promotion.NewPromotionService(promoRepo)

		//Act
		_, err := promoService.CalculateDiscount(request.PromotionRequest{
			PromotionID: typeUtil.GetAddressInt(1),
			Amount:      typeUtil.GetAddressFloat64(200),
		})

		//Assert
		assert.ErrorIs(t, err, errs.ErrGetError)
	})

}
