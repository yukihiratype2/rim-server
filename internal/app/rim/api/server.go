package api

import (
	"net/http"
	"rim-server/internal/app/rim/model"

	"github.com/minio/minio-go/v7"
)

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

type Server struct {
	model *model.Model
	s3    *minio.Client
}

func New(model *model.Model, s3 *minio.Client) *Server {
	server := &Server{model, s3}
	return server
}

func (s *Server) Start() {
	s.setupRouter()
	http.ListenAndServe(":8080", nil)
}
