package domain_test

import (
	"testing"
	"time"

	"github.com/MaxHariel/encoder/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "path/video"
	video.ResourceID = "1"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("path", "Conveter", video)

	require.NotNil(t, job)
	require.Nil(t, err)

}
