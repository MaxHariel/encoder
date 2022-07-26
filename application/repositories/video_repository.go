package repositories

import (
	"fmt"

	"github.com/MaxHariel/encoder/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Insert(video *domain.Video) (*domain.Video, error)
	Find(id string) (*domain.Video, error)
}

type VideoRepositoryDb struct {
	Db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepositoryDb {
	return &VideoRepositoryDb{Db: db}
}

func (repo *VideoRepositoryDb) Insert(video *domain.Video) (*domain.Video, error) {
	if video.ID == "" {
		video.ID = uuid.New().String()
	}

	err := repo.Db.Create(video).Error
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (repo *VideoRepositoryDb) Find(id string) (*domain.Video, error) {
	var video domain.Video
	err := repo.Db.First(&video, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	if video.ID == "" {
		return nil, fmt.Errorf("video does not exist")
	}

	return &video, nil
}
