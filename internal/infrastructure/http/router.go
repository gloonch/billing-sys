package http

import "net/http"

func NewRouter(handlers *Handlers) http.Handler {
	r := http.NewServeMux()

	// use middleware here
	// ...

	r.HandleFunc("/buildings/{id}", handlers.GetBuildingHandler)
	r.HandleFunc("/buildings", handlers.CreateBuildingHandler)

	return r
}
