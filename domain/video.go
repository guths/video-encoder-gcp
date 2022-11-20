package domain

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Video struct {
	ID         string    `valid:"uuid"`    // id do sistema atual
	ResourceID string    `valid:"notnull"` // id do sistema que envia o video
	FilePath   string    `valid:"notnull"`
	CreatedAt  time.Time `valid:"-"`
	Jobs       []*Job
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
