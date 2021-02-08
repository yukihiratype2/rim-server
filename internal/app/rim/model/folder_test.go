package model_test

import (
	"rim-server/internal/app/rim/model"
	"testing"

	"github.com/dchest/uniuri"
	"github.com/stretchr/testify/assert"
)

func TestAddFolder(t *testing.T) {
	err := model.Connect()
	if err != nil {
		t.Error(err)
	}
	folder := model.Folder{Label: uniuri.New()}
	folder.Create()
	assert.NotEqual(t, folder.ID, 0)
}
