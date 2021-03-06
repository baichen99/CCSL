package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// MemberInterface struct
type MemberInterface interface {
	GetMemberList(parameters utils.GetMemberListParameters) (members []models.Member, count int, err error)
	GetMember(id string) (member models.Member, err error)
	CreateMember(member models.Member) (err error)
	UpdateMember(id string, updateData map[string]interface{}) (err error)
	DeleteMember(id string) (err error)
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
	db := s.PG.
		Scopes(
			utils.SearchByColumn("members.name_zh", parameters.NameZh),
			utils.SearchByColumn("members.name_en", parameters.NameEn),
		)

	err = db.
		Model(&models.Member{}).
		Count(&count).
		Error

	if err != nil {
		return
	}
	typeOrder :=
		`array_position(array['consultant','researchFellow','assistantResearchFellow','postgraduate','signLanguageTranslator','researchAssistantDeaf'], members.type) asc`

	err = db.
		Order(typeOrder).
		Scopes(utils.FilterByListParameters(parameters.GetListParameters)).
		Find(&members).
		Error

	return
}

// GetMember returns member with given id
func (s *MemberService) GetMember(id string) (member models.Member, err error) {
	err = s.PG.
		Where("id = ?", id).
		Take(&member).
		Error
	return
}

// CreateMember creates a new member
func (s *MemberService) CreateMember(member models.Member) (err error) {
	err = s.PG.
		Create(&member).
		Error
	return
}

// UpdateMember updates member with given id
func (s *MemberService) UpdateMember(id string, updateData map[string]interface{}) (err error) {
	err = s.PG.
		Model(&models.Member{}).
		Where("id = ?", id).
		Updates(updateData).
		Error
	return
}

// DeleteMember soft deletes a member with given id
func (s *MemberService) DeleteMember(id string) (err error) {
	err = s.PG.
		Where("id = ?", id).
		Delete(&models.Member{}).
		Error
	return
}
