package api

import "net/http"

func (s *Server) setupRouter() {

	http.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		switch r.Method {
		case http.MethodGet:
			s.queryImages(w, r)
		case http.MethodPut:
			s.addImage(w, r)
		}
	})

	http.HandleFunc("/tag", func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		switch r.Method {
		case http.MethodGet:
			s.queryTags(w, r)
		case http.MethodPut:
			s.addTag(w, r)
		}
	})
}
