package service

import (
	"github.com/gin-gonic/gin"
)

type NoContent struct {
}

type StatusResponse struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Cause string `json:"cause"`
}

type response struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

type Meta struct {
	Code       int                    `json:"code"`
	Error      interface{}            `json:"error"`
	Key        string                 `json:"key"`
	Pagination PaginationResp         `json:"pagination"`
	Data       map[string]interface{} `json:"data"`
}

type PaginationResp struct {
	Last int16 `json:"last"`
}

type GinGoincRespParams struct {
	Code       int
	Key        string
	Pagination PaginationResp
	Data       interface{}
}

// GinGonicResp
// if response has error or something wrong during request
// this method will be used to return failure response
func GinGonicResp(
	c *gin.Context,
	params GinGoincRespParams) {
	resp := response{
		Meta: Meta{
			Code:       params.Code,
			Key:        params.Key,
			Pagination: params.Pagination,
		},
		Data: params.Data,
	}

	c.JSON(params.Code, resp)
}
