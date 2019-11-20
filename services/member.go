package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// MemberInterface struct
type MemberInterface interface {
	GetMemberList(parameters utils.GetMemberListParameters) (members []models.Member, count int, err error)
	GetMember(memberID string) (member models.Member, err error)
	CreateMember(member models.Member) (err error)
	UpdateMember(memberID string, updateData map[string]interface{}) (err error)
	DeleteMember(memberID string) (err error)
}

// MemberService implements member interface
type MemberService struct {
	PG *gorm.DB
}

// NewMemberService returns new member serivce
func NewMemberService(pg *gorm.DB) MemberInterface {
	return &MemberService{
		PG: pg,
	}
}

// GetMemberList returns members list
func (s *MemberService) GetMemberList(parameters utils.GetMemberListParameters) (members []models.Member, count int, err error) {
	db := s.PG.LogMode(false).Scopes(
		utils.SearchByColumn("members.name_zh", parameters.NameZh),
		utils.SearchByColumn("members.name_en", parameters.NameEn),
	)
	err = db.Find(&members).Count(&count).Error

	if err != nil {
		return
	}

	orderQuery := parameters.OrderBy + " " + parameters.Order
	if parameters.Limit != 0 {
		err = db.Order(orderQuery).Limit(parameters.Limit).Offset(parameters.Limit * (parameters.Page - 1)).Find(&members).Error
	} else {
		err = db.Order(orderQuery).Find(&members).Error
	}
	return
}

// GetMember returns member with given id
func (s *MemberService) GetMember(memberID string) (member models.Member, err error) {
	err = s.PG.LogMode(false).Where("id = ?", memberID).Take(&member).Error
	return
}

// CreateMember creates a new member
func (s *MemberService) CreateMember(member models.Member) (err error) {
	err = s.PG.LogMode(true).Create(&member).Error
	return
}

// UpdateMember updates member with given id
func (s *MemberService) UpdateMember(memberID string, updateData map[string]interface{}) (err error) {
	var member models.Member
	err = s.PG.LogMode(true).Model(&member).Where("id = ?", memberID).Update(updateData).Error
	return
}

// DeleteMember soft deletes a member with given id
func (s *MemberService) DeleteMember(memberID string) (err error) {
	var member models.Member
	err = s.PG.LogMode(true).Where("id = ?", memberID).Delete(&member).Error
	return
}
