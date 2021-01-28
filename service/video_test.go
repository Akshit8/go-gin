package service

import (
	"testing"

	"github.com/Akshit8/go-gin/entity"
	"github.com/Akshit8/go-gin/repository"
	"github.com/stretchr/testify/require"
)

const (
	sampleTitle       = "cool title"
	sampleDescription = "video description"
	sampleURL         = "https://www.youtube.com/embed/JLQMZzzM4"
	sampleFirstName   = "Akshit"
	sampleLastName    = "Sadana"
	sampleAge         = uint8(20)
	sampleEmail       = "akshit@gmail.com"
)

func createSampleVideo() entity.Video {
	return entity.Video{
		Title:       sampleTitle,
		Description: sampleDescription,
		URL:         sampleURL,
		Author: entity.Person{
			FirstName: sampleFirstName,
			LastName:  sampleLastName,
			Age:       sampleAge,
			Email:     sampleEmail,
		},
	}
}

func TestFindAll(t *testing.T) {
	t.Parallel()
	videoRepository := repository.NewVideoRepository("../sqlite/test.db")
	service := NewVideoService(videoRepository)

	service.Save(createSampleVideo())

	videos := service.FindAll()

	firstVideo := videos[0]
	require.NotEmpty(t, firstVideo)
	require.Equal(t, sampleTitle, firstVideo.Title)
	require.Equal(t, sampleDescription, firstVideo.Description)
	require.Equal(t, sampleURL, firstVideo.URL)
	require.Equal(t, sampleFirstName, firstVideo.Author.FirstName)
	require.Equal(t, sampleLastName, firstVideo.Author.LastName)
	require.Equal(t, sampleAge, firstVideo.Author.Age)
	require.Equal(t, sampleEmail, firstVideo.Author.Email)
}
