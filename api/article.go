package api

import (
	"server/model/request"
	"server/model/response"

	"github.com/gin-gonic/gin"
)

type ArticleApi struct {
}

func (a *ArticleApi) ArticleCreate(c *gin.Context) {
	var req request.ArticleCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = articleService.ArticleCreate(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("Article created successfully", c)
}
