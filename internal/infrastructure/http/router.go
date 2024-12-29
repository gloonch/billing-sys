package http

import (
	"billing-sys/internal/infrastructure/metrics"
	"net/http"
)

func NewRouter(handlers *Handlers) http.Handler {
	r := http.NewServeMux()

	metrics.Init()
	r.Handle("/", metrics.MetricsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})))
	r.Handle("/metrics", metrics.MetricsHandler())

	r.HandleFunc("GET /buildings/{id}", handlers.GetBuildingHandler)
	r.HandleFunc("POST /buildings", handlers.CreateBuildingHandler)
	r.HandleFunc("GET /list_buildings", handlers.ListBuildingHandler)
	r.HandleFunc("PUT /buildings/{id}", handlers.UpdateBuildingHandler)
	r.HandleFunc("DELETE /buildings/{id}", handlers.DeleteBuildingHandler)
	r.HandleFunc("GET /buildings/{id}/charges/{strategy}", handlers.CalculateBuildingChargeHandler)

	r.HandleFunc("POST /units", handlers.CreateUnitHandler)
	r.HandleFunc("GET /units/{id}", handlers.GetUnitHandler)
	r.HandleFunc("GET /units", handlers.ListUnitHandler)
	r.HandleFunc("PUT /units/{id}", handlers.UpdateUnitHandler)
	r.HandleFunc("DELETE /units/{id}", handlers.DeleteUnitHandler)

	r.HandleFunc("POST /payments", handlers.CreatePaymentHandler)                    // POST
	r.HandleFunc("GET /payments/unit/{unit_id}", handlers.ListPaymentsByUnitHandler) // GET
	r.HandleFunc("DELETE /payments/{id}", handlers.DeletePaymentHandler)             // DELETE

	return r
}
