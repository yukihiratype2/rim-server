package model_test

import (
	"rim-server/internal/app/rim/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveImage(t *testing.T) {
	err := model.Connect()
	if err != nil {
		t.Error(err)
	}
	image := model.Image{}
	err = image.Create()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, image.Favorite, false)
	image.Favorite = true
	err = image.Update()
	if err != nil {
		t.Error(err)
	}
	var result model.Image
	result.ID = image.ID
	result.First()
	assert.Equal(t, result.Favorite, true)
}
