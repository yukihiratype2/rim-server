package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) setupRouter() {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/image/list", func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		s.queryImages(w, r)
	})

	r.HandleFunc("/image", s.addImage).Methods("PUT")
	r.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
	}).Methods("OPTIONS")
	r.HandleFunc("/image/{ID}", func(w http.ResponseWriter, r *http.Request) {
		id := (mux.Vars(r))["ID"]
		s.getImage(w, r, id)
	}).Methods("GET")

	r.HandleFunc("/tag", func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		switch r.Method {
		case http.MethodGet:
			s.queryTags(w, r)
		case http.MethodPut:
			s.addTag(w, r)
		}
	})
}
