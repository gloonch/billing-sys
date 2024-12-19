package http

import "net/http"

func NewRouter(handlers *Handlers) http.Handler {
	r := http.NewServeMux()

	// use middleware here
	// ...

	r.HandleFunc("GET /buildings/{id}", handlers.GetBuildingHandler)
	r.HandleFunc("POST /buildings", handlers.CreateBuildingHandler)
	r.HandleFunc("GET /list_buildings", handlers.ListBuildingHandler)
	r.HandleFunc("UPDATE /buildings", handlers.CreateBuildingHandler)
	r.HandleFunc("DELETE /buildings", handlers.CreateBuildingHandler)

	return r
}
