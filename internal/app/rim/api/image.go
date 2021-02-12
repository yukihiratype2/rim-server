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
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

func imageRoute() {
	r.GET("/image", queryImages)
	r.GET("/image/:id", getImage)
	r.POST("image", updateImage)
	r.GET("/imageprocess/:fileId", waitForImageProcessed)
	r.PUT("image", addImage)
}

func queryImages(c *gin.Context) {
	var query model.Image
	c.ShouldBindQuery(&query)
	var images []model.Image

	query.Find(&images)

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

type addImageResponse struct {
	model.Image
	UploadURL string `json:"uploadUrl" gorm:"-"`
}

func addImage(c *gin.Context) {
	var image model.Image
	c.ShouldBindJSON(&image)
	image.FileID = uuid.New().String() + ".jpg"
	image.Create()
	respImage := addImageResponse{Image: image}

	presignedURL, err := s3.Client.PresignedPutObject(context.Background(), "test-img", image.FileID, time.Second*3*60)
	imageStatus := event.ImageProcessStatus{Image: model.Image{FileID: image.FileID}}
	imageStatus.ImageCreated()
	respImage.UploadURL = presignedURL.String()
	if err != nil {
		c.Err()
	}
	c.JSON(200, respImage)
}

func updateImage(c *gin.Context) {
	var image model.Image
	c.MustBindWith(&image, binding.JSON)
	image.Update()
	c.JSON(200, image)
}

func waitForImageProcessed(c *gin.Context) {
	fileID := c.Param("fileId")
	imageStatus := event.ImageProcessStatus{Image: model.Image{FileID: fileID}}
	imageStatus.WaitForImageProcessed()
	c.JSON(200, gin.H{
		"status": "completed",
	})
}
