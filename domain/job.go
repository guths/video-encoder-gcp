package domain

import "time"

type Job struct {
	ID               string
	OutputBucketPath string //caminho de saida do video convertido
	Status           string //status do processamento do video
	Video            *Video
	Error            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
