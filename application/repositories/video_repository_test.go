package repositories_test

import (
	"testing"
	"time"

	"github.com/MaxHariel/encoder/application/repositories"
	"github.com/MaxHariel/encoder/domain"
	"github.com/MaxHariel/encoder/framework/database"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewVideoRepository(t *testing.T) {
	db := database.NewDbTest()

	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "path/video"
	video.ResourceID = "afaf"
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepository(db)
	v, err := repo.Insert(video)
	require.Nil(t, err)
	require.NotEmpty(t, v)

	v, err = repo.Find(video.ID)
	require.Nil(t, err)
	require.NotEmpty(t, v)
	require.Equal(t, video.ID, v.ID)
}
