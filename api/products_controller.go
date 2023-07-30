package api

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ramtinhr/vgang-task/models"
	"github.com/ramtinhr/vgang-task/service"
)

type prodTransformer struct {
	ProdId   uint   `json:"productId"`
	ShortUrl string `json:"shortUrl"`
}

// GetShortUrls fetch products and return short url and product id
func GetShortUrls(c *gin.Context) {
	var data []*prodTransformer
	prods, err := models.FetchProducts()

	if err != nil {
		c.JSON(http.StatusBadRequest, service.ErrorResponse{
			Code:  http.StatusBadRequest,
			Cause: err.Error(),
		})
	}

	for _, prod := range prods {
		data = append(data, &prodTransformer{
			ProdId:   prod.ProductID,
			ShortUrl: fmt.Sprintf("%s/%s", os.Getenv("IP_ADD"), prod.Hash),
		})
	}

	service.GinGonicResp(c, service.GinGoincRespParams{
		Code: http.StatusOK,
		Data: data,
	})
}

// UserShortUrl
// Redirect to vgang website but it won't work without authentication
// You should already logged in to the vgang dashboard
func UseShortUrl(c *gin.Context) {
	hash := c.Param("hash")
	if hash != "" {
		prod, err := models.FindProdByHash(hash)
		if err != nil {
			c.JSON(http.StatusBadRequest, service.ErrorResponse{
				Code:  http.StatusBadRequest,
				Cause: err.Error(),
			})
		}

		c.Redirect(http.StatusFound, fmt.Sprintf("%s/retailer/list-of-products/%v", os.Getenv("VGANG_URL"), prod.ProductID))
	}

	c.JSON(http.StatusBadRequest, service.ErrorResponse{
		Code:  http.StatusBadRequest,
		Cause: errors.New("hash is not defined").Error(),
	})
}
