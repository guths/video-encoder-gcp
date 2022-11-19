package domain

import "time"

type Video struct {
	ID         string // id do sistema atual
	ResourceID string // id do sistema que envia o video
	FilePath   string
	CreatedAt  time.Time
}
