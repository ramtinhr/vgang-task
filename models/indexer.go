package models

import (
	"github.com/ramtinhr/vgang-task/utils"
	"gorm.io/gorm/clause"
)

type Indexer struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Username     string `json:"username" gorm:"unique"`
	Password     string `json""password"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func AddIndexer(indexer *Indexer) (*Indexer, error) {
	tx := utils.PgsqlDB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "username"}},
		DoUpdates: clause.AssignmentColumns([]string{"access_token", "refresh_token"}),
	}).Create(&indexer)

	return indexer, tx.Error
}
