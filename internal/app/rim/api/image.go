package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"rim-server/internal/app/rim/model"
	"time"

	"github.com/google/uuid"
)

func (s *Server) queryImages(w http.ResponseWriter, req *http.Request) {
	var image []model.Image

	s.model.DB.Find(&image)

	result, err := json.Marshal(image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (s *Server) addImage(w http.ResponseWriter, req *http.Request) {
	var image model.Image

	err := json.NewDecoder(req.Body).Decode(&image)

	s.model.DB.Create(image)
	res, err := json.Marshal(image)
	presignedURL, err := s.s3.PresignedPutObject(context.Background(), "test-img", uuid.New().String(), time.Second*24*60*60)
	fmt.Println(presignedURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
