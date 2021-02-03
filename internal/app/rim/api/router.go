package api

import "net/http"

func (s *Server) setupRouter() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		switch r.Method {
		case http.MethodGet:
			s.queryImages(w, r)
		case http.MethodPut:
			s.addImage(w, r)
		}
	})
}
