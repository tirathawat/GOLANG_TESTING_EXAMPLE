package storages

import "time"

type PromotionDB struct {
	ID               int       `gorm:"primaryKey;column:promotion_id" json:"promotion_id"`
	Name             string    `gorm:"column:promotion_name" json:"promotion_name"`
	PurchaseMin      float64   `gorm:"purchase_min" json:"purchase_min"`
	DiscountPercent  float64   `gorm:"discount_percents" json:"discount_percents"`
	ExpiredDate      time.Time `gorm:"expired_date" json:"expired_date"`
	CreatedTimestamp time.Time `gorm:"created_timestamp" json:"created_timestamp"`
	UpdateTimestamp  time.Time `gorm:"update_timestamp" json:"update_timestamp"`
}
