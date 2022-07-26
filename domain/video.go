package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Video struct {
	ID         string    `json:"encoded_video_folder" valid:"uuid" gorm:"type:uuid;primaryKey"`
	ResourceID string    `json:"resource_id" valid:"notnull" gorm:"type:string"`
	FilePath   string    `json:"file_path" valid:"notnull" gorm:"type:string"`
	Jobs       []*Job    `json:"-" valid:"-"`
	CreatedAt  time.Time `json:"-" valid:"-" gorm:"foreignKey:VideoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func NewVideo() *Video {
	return &Video{}
}

func (video *Video) Validate() error {
	_, err := govalidator.ValidateStruct(video)
	if err != nil {
		return err
	}
	return nil
}
