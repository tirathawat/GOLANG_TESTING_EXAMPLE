package repositories

import (
	"example/models/storages"

	"github.com/stretchr/testify/mock"
)

type promotionRepositoryMock struct {
	mock.Mock
}

func NewPromotionRepositoryMock() *promotionRepositoryMock {
	return &promotionRepositoryMock{}
}

func (m promotionRepositoryMock) GetPromotion(id int) (storages.PromotionDB, error) {
	args := m.Called()
	return args.Get(0).(storages.PromotionDB), args.Error(1)
}
