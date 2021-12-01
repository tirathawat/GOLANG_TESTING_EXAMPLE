package promotion

import (
	"example/errs"
	"example/logs"
	"example/models/request"
	"example/repositories"
	"time"
)

type promotionService struct {
	Repo repositories.IPromotionRepository
}

type IPromotionService interface {
	CalculateDiscount(request request.PromotionRequest) (float64, error)
}

func NewPromotionService(repo repositories.IPromotionRepository) promotionService {
	return promotionService{Repo: repo}
}

func (s promotionService) CalculateDiscount(request request.PromotionRequest) (float64, error) {

	discountedAmount := 0.0

	promotionDB, err := s.Repo.GetPromotion(*request.PromotionID)
	if err != nil {
		logs.New().Error(err)
		return discountedAmount, errs.ErrGetError
	}

	if promotionDB.ExpiredDate.Before(time.Now().Local()) {
		return discountedAmount, errs.ErrPromoExpired
	}

	if *request.Amount >= promotionDB.PurchaseMin {
		discountedAmount = *request.Amount - (promotionDB.DiscountPercent * *request.Amount / 100)
	}

	return discountedAmount, nil
}
