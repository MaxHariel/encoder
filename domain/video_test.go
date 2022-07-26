package domain_test

import (
	"testing"
	"time"

	"github.com/MaxHariel/encoder/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIsNotValidUUID(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "asd"
	video.FilePath = "path/video"
	video.ResourceID = "1"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoIsValidate(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "path/video"
	video.ResourceID = "1"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}
