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
	err := c.ShouldBindQuery(&tag)
	if err != nil {
		c.Error(err)
	}
	var tags []model.Tag
	err = tag.Find(&tags)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, tags)
	// s.model.DB.Preload("Images").Find(&tag)
}

func addTag(c *gin.Context) {
	var tag model.Tag
	err := c.MustBindWith(&tag, binding.JSON)
	if err != nil {
		return
	}
	err = tag.Create()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, tag)
}
