package models

import (
	"github.com/ramtinhr/vgang-task/utils"
)

type Product struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ProductID uint   `json:"product_id" gorm:"unique;not null"`
	Hash      string `json:"short_url" gorm:"unique;not null"`
}

func AddProducts(products []*Product) error {
	err := utils.PgsqlDB.Create(&products).Error

	return err
}

func FetchProducts() ([]*Product, error) {
	var prods []*Product
	err := utils.PgsqlDB.Find(&prods).Error

	return prods, err
}

func FindProdByHash(hash string) (*Product, error) {
	var prod *Product
	err := utils.PgsqlDB.First(&prod, "hash = ?", hash).Error

	return prod, err
}
