package models

type Product struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ProductID uint   `json:"product_id" gorm:"unique;not null"`
	ShortUrl  string `json:"short_url" gorm:"unique;not null"`
}
