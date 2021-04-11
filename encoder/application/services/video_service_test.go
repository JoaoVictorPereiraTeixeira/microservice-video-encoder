package services_test

import (
	"log"
	"time"

	"github.com.br/JoaoVictorPereiraTeixeira/application/repositories"
	"github.com.br/JoaoVictorPereiraTeixeira/domain"
	"github.com.br/JoaoVictorPereiraTeixeira/framework/database"

	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func prepare() (*domain.Video, repositories.VideoRepository) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "convite.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	return video, repo
}

// func TeteVideoServiceDownload(t *testing.T) {

// 	video, repo := prepare()

// 	videoService := services.NewVideoService()
// 	videoService.Video = video
// 	videoService.VideoRepository = repo

// 	err := videoService.Download("video-to-upload")
// 	require.Nil(t, err)

// 	err = videoService.Fragment()
// 	require.Nil(t, err)

// 	err = videoService.Encode()
// 	require.Nil(t, err)

// 	err = videoService.Finish()
// 	require.Nil(t, err)
// }
