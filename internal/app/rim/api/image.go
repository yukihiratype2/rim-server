package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"rim-server/internal/app/rim/model"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func (s *Server) queryImages(w http.ResponseWriter, req *http.Request) {
	var image []model.Image

	s.model.DB.Preload("Tags").Find(&image)

	result, err := json.Marshal(image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (s *Server) getImage(w http.ResponseWriter, req *http.Request, id string) {
	setupResponse(&w, req)
	var image model.Image
	ID, err := strconv.ParseUint(id, 10, 8)
	image.ID = uint(ID)
	s.model.DB.Preload("Tags").First(&image)
	presignedURL, err := s.s3.PresignedGetObject(context.Background(), "test-img", image.FileID, time.Second*24*60*60, url.Values{})
	image.URL = presignedURL.String()
	result, err := json.Marshal(image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

type uploadURL struct {
	URL string `json:"url"`
}

func (s *Server) addImage(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	var image model.Image
	image.FileID = uuid.New().String() + ".jpg"
	s.model.DB.Create(&image)
	presignedURL, err := s.s3.PresignedPutObject(context.Background(), "test-img", image.FileID, time.Second*24*60*60)
	var url uploadURL
	url.URL = presignedURL.String()
	res, err := json.Marshal(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
