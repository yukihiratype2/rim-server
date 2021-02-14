package api

import (
	"net/http"
	"rim-server/internal/app/rim/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func tagRoute() {
	r.GET("/tag", queryTags)
	r.PUT("/tag", addTag)
}

func queryTags(c *gin.Context) {
	var tag model.Tag
	c.ShouldBindQuery(&tag)
	var tags []model.Tag
	tag.Find(&tags)
	c.JSON(http.StatusOK, tags)
	// s.model.DB.Preload("Images").Find(&tag)
}

func addTag(c *gin.Context) {
	var tag model.Tag
	err := c.MustBindWith(&tag, binding.JSON)
	err = tag.Create()
	if err != nil {
		c.Err()
	}
	c.JSON(http.StatusOK, tag)
	// s.model.DB.Create(&tag)
}
