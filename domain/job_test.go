package domain

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewJob(t *testing.T) {
	video := NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	_, err := NewJob("path", "converted", video)

	require.Nil(t, err)
}
