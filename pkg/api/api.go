package api

import (
	"net/http"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"

	"github.com/aeekayy/go-api-base/pkg/api/handlers"
)

func StartHTTP(config *Config, db *gorm.DB) {
	log.Infof("Setting up routers")

	r := NewRouter(db)

	if config.EnableMetrics {
		log.Info("Prometheus metrics are on")
		metricsHandler := handlers.NewPrometheusHandler()
		r.HandleFunc(metricsHandler.Path, metricsHandler.Handle)
	}

	// IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
	r.HandleFunc("/ping", pingHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))
	
	log.Fatal(http.ListenAndServe(":8080", r))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	if r.Method == http.MethodOptions {
		return
	}

	w.Write([]byte(`{"alive": true}`))
}
