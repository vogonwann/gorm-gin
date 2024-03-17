package service

import (
	"github.com/vogonwann/gorm-gin/request"
	"github.com/vogonwann/gorm-gin/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
