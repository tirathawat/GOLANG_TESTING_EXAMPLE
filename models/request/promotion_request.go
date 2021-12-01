package request

import "example/errs"

type PromotionRequest struct {
	PromotionID *int     `json:"promotion_id"`
	Amount      *float64 `json:"amount"`
}

func (r *PromotionRequest) Validate() error {
	if r.PromotionID == nil {
		return errs.ErrPromoIDNotFound
	} else if r.Amount == nil {
		return errs.ErrPurchaseNotFound
	} else if *r.Amount <= 0 {
		return errs.ErrPurchaseZero
	}
	return nil
}
