package provider

import "github.com/ramtinhr/vgang-task/models"

type IProvider interface {
	// Login to vgang account
	Login() (models.Indexer, error)

	// CreateShortUrl for each vgang product in a specific category
	FetchProductsByCat(catId uint) error
}
