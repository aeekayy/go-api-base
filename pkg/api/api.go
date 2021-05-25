package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/aeekayy/go-api-base/pkg/api/handlers"
	"github.com/spf13/viper"
)

type Server struct {
	*http.Server
	DB     *gorm.DB
	Config *Config
}

func NewServer(config *Config, db *gorm.DB) (*Server, error) {
	log.Info("Creating a new server")

	config.EnableCORS = viper.GetBool("http.enable_cors")

	api, err := AddRouter(config, db)
	if err != nil {
		return nil, err
	}

	port := viper.GetString("http.port")

	if port != "" {
		// allow port to be set as localhost:3000 in env during development to avoid "accept incoming network connection" request on restarts
		if strings.Contains(port, ":") {
			config.Port = port
		} else {
			config.Port = ":" + port
		}
	}

	srv := http.Server{
		Addr:    config.Port,
		Handler: api,
	}

	httpSrv := &Server{Server: &srv, DB: db, Config: config}

	return httpSrv, nil
}

func AddRouter(config *Config, db *gorm.DB) (*mux.Router, error) {
	log.Infof("Setting up routers")

	r := NewRouter(db)

	if config.EnableMetrics {
		log.Info("Prometheus metrics are on")
		metricsHandler := handlers.NewPrometheusHandler()
		r.HandleFunc(metricsHandler.Path, metricsHandler.Handle)
	}

	// IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
	r.Use(mux.CORSMethodMiddleware(r))

	cors := gorillaHandlers.CORS(
		gorillaHandlers.AllowedHeaders([]string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}),
		gorillaHandlers.AllowedOrigins([]string{"api.aeekay.co", "app.aeekay.co"}),
		gorillaHandlers.AllowCredentials(),
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	)
	r.Use(cors)

	r.HandleFunc("/ping", pingHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)

	return r, nil
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (srv *Server) Start() {
	log.Println("starting server...")
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("Listening on %s\n", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Println("Shutting down server... Reason:", sig)
	// teardown logic...

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Println("Server gracefully stopped")
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
