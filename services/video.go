package services

import (
	"ccsl/models"
	"ccsl/utils"

	"github.com/jinzhu/gorm"
)

// VideoInterface struct
type VideoInterface interface {
	GetVideosList(parameters utils.GetVideoListParameters) (videos []models.VideoResult, count int, err error)
	CreateVideo(video models.Video) (err error)
	GetVideo(videoID string) (video models.Video, err error)
	UpdateVideo(videoID string, updatedData models.Video) (err error)
	DeleteVideo(videoID string) (err error)
}

// VideoService implements Video interface
type VideoService struct {
	PG *gorm.DB
}

// NewVideoService returns new Video serivce
func NewVideoService(pg *gorm.DB) VideoInterface {
	return &VideoService{
		PG: pg,
	}
}

// GetVideosList returns video list
func (s *VideoService) GetVideosList(parameters utils.GetVideoListParameters) (videos []models.VideoResult, count int, err error) {
	// Adding custom scopes to the query based on get list parameters.
	db := s.PG.Scopes(
		utils.FilterByColumn("initial", parameters.Initial),
		utils.SearchByColumn("chinese", parameters.Chinese),
		utils.FilterByArray("english", parameters.English, " "),
		utils.FilterByColumn("type", parameters.Type),
		utils.SearchByColumn("name", parameters.Name),
		utils.SearchByColumn("region", parameters.Region),
		utils.FilterByColumn("gender", parameters.Gender),
		utils.FilterByArray("left_sign", parameters.LeftSign, ","),
		utils.FilterByArray("right_sign", parameters.RightSign, ","),
		utils.FilterByColumn("construct_type", parameters.ConstructType),
		utils.SearchByColumn("construct_words", parameters.ConstructWords),
		utils.FilterByColumn("performer_id", parameters.PerformerID),
		utils.FilterByColumn("word_id", parameters.WordID),
	)
	// Fetching the total number of rows based on the conditions provided.
	query := db.Table("videos").Select("videos.id, words.initial, words.chinese, words.english, words.type, performers.name, performers.region, performers.gender, videos.construct_type, videos.construct_words,videos.left_sign, videos.right_sign, videos.video_path").Joins("LEFT JOIN words ON words.id = videos.word_id").Joins("LEFT JOIN performers ON performers.id = videos.performer_id")

	err = query.Count(&count).Error

	regionOrder := "array_position(array['通用手语','浦东新区','黄浦区','静安区','徐汇区','长宁区','普陀区','虹口区','杨浦区','宝山区','闵行区','嘉定区','金山区','松江区','青浦区','奉贤区','崇明区'], region) asc"

	if err != nil {
		return
	}
	// Fetching the items to be returned by the query.
	orderQuery := parameters.OrderBy + " " + parameters.Order
	if parameters.Limit != 0 {
		err = query.Order(regionOrder).Order(orderQuery).Limit(parameters.Limit).Offset(parameters.Limit * (parameters.Page - 1)).Scan(&videos).Error

	} else {
		err = query.Order(regionOrder).Order(orderQuery).Scan(&videos).Error
	}
	return
}

func (s *VideoService) CreateVideo(video models.Video) (err error) {
	err = s.PG.Create(&video).Error
	return
}
func (s *VideoService) GetVideo(videoID string) (video models.Video, err error) {
	err = s.PG.Where("id = ?", videoID).Take(&video).Error
	return
}
func (s *VideoService) UpdateVideo(videoID string, updatedData models.Video) (err error) {
	var video models.Video
	err = s.PG.Where("id = ?", videoID).Take(&video).Model(&video).Updates(updatedData).Error
	return
}
func (s *VideoService) DeleteVideo(videoID string) (err error) {
	var video models.Video
	err = s.PG.Where("id = ?", videoID).Delete(&video).Error
	return
}
