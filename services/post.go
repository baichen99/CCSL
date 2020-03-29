package services

import (
	"ccsl/models"
	"ccsl/utils"
	"github.com/jinzhu/gorm"
)

type PostInterface interface {
	GetPostList(parameters utils.GetPostListParameters) (post []models.Post, count int, err error)
	GetPost(id string) (post models.Post, err error)
	CreatePost(post models.Post) (err error)
	UpdatePost(id string, updateData map[string]interface{}) (err error)
	DeletePost(id string) (err error)
}

type PostService struct {
	PG *gorm.DB
}

func NewPostService(pg *gorm.DB) PostInterface {
	return &PostService{
		PG: pg,
	}
}

// GetPostList returns all posts
func (p PostService) GetPostList(parameters utils.GetPostListParameters) (post []models.Post, count int, err error) {

	db := p.PG.
		Scopes(
			utils.FilterByColumn("posts.creator_id", parameters.CreatorID),
			utils.SearchByColumn("posts.title", parameters.Title),
		)

	err = db.
		Model(&models.Post{}).
		Count(&count).
		Error

	if err != nil {
		return
	}

	err = db.
		Preload("Creator", func(pg *gorm.DB) *gorm.DB {
			return pg.Select("id, name, avatar, state, roles")
		}).
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&post).
		Error

	return
}

// GetPost get post by id
func (p PostService) GetPost(id string) (post models.Post, err error) {
	err = p.PG.
		Preload("Creator").
		Preload("Replies").
		Where("id = ?", id).
		Take(&post).
		Error

	return
}

// CreatePost creates post
func (p PostService) CreatePost(post models.Post) (err error) {
	err = p.PG.
		Create(&post).
		Error
	return
}

// UpdatePost updates post with given id
func (p PostService) UpdatePost(id string, updateData map[string]interface{}) (err error) {
	err = p.PG.
		Model(&models.Post{}).
		Where("id = ?", id).
		Update(updateData).
		Error

	return
}

// DeletePost delete post by id
func (p PostService) DeletePost(id string) (err error) {
	err = p.PG.
		Where("id = ?", id).
		Delete(&models.Post{}).
		Error
	return
}
