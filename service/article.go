package service

import (
	"server/global"
	"server/model/elasticsearch"
	"server/model/request"
)

type ArticleService struct {
}

func (a *ArticleService) ArticleCreate(req request.ArticleCreate) error {

	article := elasticsearch.Article{
		Cover:    req.Cover,
		Title:    req.Title,
		Category: req.Category,
		Tags:     req.Tags,
		Abstract: req.Abstract,
		Content:  req.Content,
	}
	err := global.DB.Create(&article).Error
	return err
}
