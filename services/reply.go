package services

import (
	"ccsl/models"
	"ccsl/utils"
	"github.com/jinzhu/gorm"
)

type ReplyInterface interface {
	GetReplyList(parameters utils.GetReplyListParameters) (reply []models.Reply, count int, err error)
	GetReply(id string) (reply models.Reply, err error)
	CreateReply(reply models.Reply) (err error)
	UpdateReply(id string, updateData map[string]interface{}) (err error)
	DeleteReply(id string) (err error)
}
type ReplyService struct {
	PG *gorm.DB
}

func NewReplyService(pg *gorm.DB) ReplyInterface {
	return &ReplyService{
		PG: pg,
	}
}

// GetReplyList returns all replies
func (r ReplyService) GetReplyList(parameters utils.GetReplyListParameters) (reply []models.Reply, count int, err error) {
	db := r.PG.
		Scopes(
			utils.FilterByColumn("replies.creator_id", parameters.CreatorID),
			utils.FilterByColumn("replies.post_id", parameters.PostID),
		)

	err = db.
		Model(&models.Reply{}).
		Count(&count).
		Error

	if err != nil {
		return
	}

	err = db.
		Preload("Creator", func(pg *gorm.DB) *gorm.DB {
			return pg.Select("id, name, avatar, username")
		}).
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&reply).
		Error

	return
}

// GetReply get reply by id
func (r ReplyService) GetReply(id string) (reply models.Reply, err error) {
	err = r.PG.
		Preload("Creator", func(pg *gorm.DB) *gorm.DB {
			return pg.Select("id, name, avatar, username")
		}).
		Where("id = ?", id).
		Take(&reply).
		Error

	return
}

// CreateReply create reply
func (r ReplyService) CreateReply(reply models.Reply) (err error) {
	err = r.PG.
		Create(&reply).
		Error
	return
}

// UpdateReply update reply with given id
func (r ReplyService) UpdateReply(id string, updateData map[string]interface{}) (err error) {
	err = r.PG.
		Model(&models.Reply{}).
		Where("id = ?", id).
		Update(updateData).
		Error
	return
}

// DeleteReply delete reply with given id
func (r ReplyService) DeleteReply(id string) (err error) {
	err = r.PG.
		Where("id = ?", id).
		Delete(&models.Reply{}).
		Error
	return
}
