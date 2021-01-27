package service

import (
	"testing"

	"github.com/Akshit8/go-gin/entity"
	"github.com/stretchr/testify/require"
)

const (
	sampleTitle       = "cool title"
	sampleDescription = "video description"
	sampleURL         = "https://www.youtube.com/embed/JLQMZzzM4"
)

func createSampleVideo() *entity.Video {
	return &entity.Video{
		Title:       sampleTitle,
		Description: sampleDescription,
		URL:         sampleURL,
	}
}

func TestFindAll(t *testing.T) {
	t.Parallel()
	service := NewVideoService()

	service.Save(createSampleVideo())

	videos := service.FindAll()

	firstVideo := videos[0]
	require.NotEmpty(t, firstVideo)
	require.Equal(t, sampleTitle, firstVideo.Title)
	require.Equal(t, sampleDescription, firstVideo.Description)
	require.Equal(t, sampleURL, firstVideo.URL)
}
