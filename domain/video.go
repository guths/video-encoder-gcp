package domain

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	gorm.Model
	ID         string    `json:"encoded_video_folder" valid:"uuid" gorm:"type:uuid;primary_key"` // id do sistema atual
	ResourceID string    `json:"resource_id" valid:"notnull" gorm:"type:varchar(255)"`           // id do sistema que envia o video
	FilePath   string    `json:"file_path" valid:"notnull" gorm:"type:varchar(255)"`
	CreatedAt  time.Time `json:"-" valid:"-"`
	Jobs       []*Job    `json:"-" valid:"-" gorm:"ForeignKey:VideoID"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo() *Video {
	return &Video{}
}

func (video *Video) Validate() error {
	isValid, err := govalidator.ValidateStruct(video)

	if !isValid {
		return err
	}

	return nil
}
