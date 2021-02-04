package api

func setupRouter() {
	imageRoute()
	// r.HandleFunc("/tag", func(w http.ResponseWriter, r *http.Request) {
	// 	setupResponse(&w, r)
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		s.queryTags(w, r)
	// 	case http.MethodPut:
	// 		s.addTag(w, r)
	// 	}
	// })
}
