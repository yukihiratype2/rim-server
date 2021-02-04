package api

import (
	"encoding/json"
	"net/http"
	"rim-server/internal/app/rim/model"
)

func queryTags(w http.ResponseWriter, req *http.Request) {
	var tag []model.Tag

	// s.model.DB.Preload("Images").Find(&tag)

	result, err := json.Marshal(tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func addTag(w http.ResponseWriter, req *http.Request) {
	var tag model.Tag

	err := json.NewDecoder(req.Body).Decode(&tag)

	// s.model.DB.Create(&tag)
	res, err := json.Marshal(tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
