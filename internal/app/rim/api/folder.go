package api

import (
	"rim-server/internal/app/rim/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func folderRoute() {
	r.PUT("/folder", addFolder)
	r.GET("/folder", queryFolder)
	r.DELETE("/folder/:id", deleteFolder)
}

func queryFolder(c *gin.Context) {
	var folder model.Folder
	var folders []model.Folder
	folder.Find(&folders)
	c.JSON(200, folders)
}

func addFolder(c *gin.Context) {
	var folder model.Folder
	c.MustBindWith(&folder, binding.JSON)
	folder.Create()
	c.JSON(200, folder)
}

func deleteFolder(c *gin.Context) {
	var folder model.Folder
	ID, err := strconv.ParseUint(c.Param("id"), 10, 8)
	folder.ID = uint(ID)
	err = folder.Delete()
	if err != nil {
		c.Err()
	}
	c.JSON(200, folder)
}
