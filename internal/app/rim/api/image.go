package api

import (
	"context"
	"net/url"
	"rim-server/internal/app/rim/event"
	"rim-server/internal/app/rim/model"
	"rim-server/internal/app/rim/s3"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func imageRoute() {
	r.GET("/image", queryImages)
	r.GET("/image/:id", getImage)
	r.GET("/imageprocess/:fileId", waitForImageProcessed)
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
	presignedURL, err := s3.Client.PresignedGetObject(context.Background(), "test-img", image.FileID, time.Second*24*60*60, url.Values{})
	image.URL = presignedURL.String()

	if err != nil {
		c.Err()
	}

	c.JSON(200, image)
}

type addImageParam struct {
	model.Image
	Immediate bool `json:"immediate" gorm:"-"`
}

type addImageResponse struct {
	model.Image
	UploadURL string `json:"uploadUrl" gorm:"-"`
}

func addImage(c *gin.Context) {
	var image addImageParam
	c.ShouldBindJSON(&image)
	image.FileID = uuid.New().String() + ".jpg"
	image.Create()
	var respImage addImageResponse
	respImage.ID = image.ID
	respImage.FileID = image.FileID

	presignedURL, err := s3.Client.PresignedPutObject(context.Background(), "test-img", image.FileID, time.Second*3*60)
	imageStatus := event.ImageProcessStatus{Image: model.Image{FileID: image.FileID}}
	imageStatus.ImageCreated()
	respImage.UploadURL = presignedURL.String()
	if err != nil {
		c.Err()
	}
	c.JSON(200, respImage)
}

func waitForImageProcessed(c *gin.Context) {
	fileID := c.Param("fileId")
	imageStatus := event.ImageProcessStatus{Image: model.Image{FileID: fileID}}
	imageStatus.WaitForImageProcessed()
	c.JSON(200, gin.H{
		"status": "completed",
	})
}
