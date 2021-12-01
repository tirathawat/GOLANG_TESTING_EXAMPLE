package repositories

import (
	"example/database"
	"example/models/storages"
)

type promotionRepository struct {
	Database database.IDatabase
}

type IPromotionRepository interface {
	GetPromotion(id int) (storages.PromotionDB, error)
}

func NewPromotionRepository(db database.IDatabase) promotionRepository {
	return promotionRepository{Database: db}
}

func (r promotionRepository) GetPromotion(id int) (storages.PromotionDB, error) {
	promotionDB := storages.PromotionDB{}
	err := r.Database.GetDB().
		Table("promotion").
		Where("promotion_id = ?", id).
		Find(&promotionDB).
		Error
	return promotionDB, err
}
