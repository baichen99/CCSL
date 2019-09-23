package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

type MemberInterface interface {
	GetMemberList(parameters utils.GetMemberListParameters) (members []models.Member, count int, err error)
	GetMember(memberID string) (member models.Member, err error)
	CreateMember(member models.Member) (err error)
	UpdateMember(memberID string, updateData map[string]interface{}) (err error)
	DeleteMember(memberID string) (err error)
}

type MemberService struct {
	PG *gorm.DB
}

func NewMemberService(pg *gorm.DB) MemberInterface {
	return &MemberService{
		PG: pg,
	}
}

func (s *MemberService) GetMemberList(parameters utils.GetMemberListParameters) (members []models.Member, count int, err error) {
	// TODO: find out what needs here?
	db := s.PG.Scopes(
		utils.FilterByColumn("name_zh", parameters.NameZh),
		utils.FilterByColumn("name_en", parameters.NameEn),
	)
	// TODO: orderQuery?
	err = db.Find(&members).Count(&count).Error
	return
}

func (s *MemberService) GetMember(memberID string) (member models.Member, err error) {
	err = s.PG.Where("id = ?", memberID).Take(&member).Error
	return
}

func (s *MemberService) CreateMember(member models.Member) (err error) {
	err = s.PG.Create(&member).Error
	return
}

func (s *MemberService) UpdateMember(memberID string, updateData map[string]interface{}) (err error) {
	var member models.Member
	err = s.PG.Model(&member).Where("id = ?", memberID).Update(updateData).Error
	return
}

func (s *MemberService) DeleteMember(memberID string) (err error) {
	var member models.Member
	err = s.PG.Where("id = ?", memberID).Delete(&member).Error
	return
}
