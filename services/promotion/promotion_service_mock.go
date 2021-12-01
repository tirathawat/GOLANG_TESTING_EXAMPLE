package promotion

import (
	"example/models/request"

	"github.com/stretchr/testify/mock"
)

type promotionServiceMock struct {
	mock.Mock
}

func NewPromotionServiceMock() *promotionServiceMock {
	return &promotionServiceMock{}
}

func (m *promotionServiceMock) CalculateDiscount(request request.PromotionRequest) (float64, error) {
	args := m.Called()
	return args.Get(0).(float64), args.Error(1)
}
