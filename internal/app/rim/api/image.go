package api

import (
	"context"
	"net/url"
	"rim-server/internal/app/rim/imageservice"
	"rim-server/internal/app/rim/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func imageRoute() {
	r.GET("/image", queryImages)
	r.GET("/image/:id", getImage)
	r.PUT("image", addImage)
}

func queryImages(c *gin.Context) {
	var images []model.Image

	model.Find(&images)

	c.JSON(200, images)
}

func getImage(c *gin.Context) {
	var image model.Image
	ID, err := strconv.ParseUint(c.Param("id"), 10, 8)
	image.ID = uint(ID)
	image.First()
	presignedURL, err := imageservice.S3.PresignedGetObject(context.Background(), "test-img", image.FileID, time.Second*24*60*60, url.Values{})
	image.URL = presignedURL.String()

	if err != nil {
		c.Err()
	}

	c.JSON(200, image)
}

type uploadURL struct {
	URL string `json:"url"`
}

func addImage(c *gin.Context) {
	var image model.Image
	c.ShouldBindJSON(&image)
	image.FileID = uuid.New().String() + ".jpg"
	image.Create()

	presignedURL, err := imageservice.S3.PresignedPutObject(context.Background(), "test-img", image.FileID, time.Second*3*60)
	var url uploadURL
	url.URL = presignedURL.String()
	if err != nil {
		c.Err()
	}
	c.JSON(200, url)
}
