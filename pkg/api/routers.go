/*
 * Veritone Build and Release API
 *
 * Build and release API for Veritone
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	handlers "github.com/aeekayy/go-api-base/pkg/api/handlers"
	"github.com/aeekayy/go-api-base/pkg/config"
)

const (
	apiVersion = "v2"
)

var noCacheHeaders = map[string]string{
	"Expires":         time.Unix(0, 0).Format(time.RFC1123),
	"Cache-Control":   "no-cache, private, max-age=0",
	"Pragma":          "no-cache",
	"X-Accel-Expires": "0",
}

var allowedOriginHosts = map[string]bool{
	"app.aeekay.co": true,
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(config *config.HTTPConfig, db *gorm.DB) *mux.Router {
	routes := getRoutes(config, db)

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func getRoutes(config *config.HTTPConfig, db *gorm.DB) Routes {
	return Routes{
		Route{
			"Index",
			"GET",
			fmt.Sprintf("/%s/", apiVersion),
			handlers.Index,
		},

		Route{
			"Ping",
			http.MethodGet,
			fmt.Sprintf("/%s/%s", apiVersion, "ping"),
			pingHandler,
		},

		Route{
			"GetEvents",
			http.MethodGet,
			fmt.Sprintf("/%s/%s", apiVersion, "events"),
			handlers.GetEvents{handlers.BaseHandler{DB: db, CORS: allowedOriginHosts}}.ServeHTTP,
		},

		Route{
			"Login",
			http.MethodPost,
			fmt.Sprintf("/%s/%s/%s", apiVersion, "auth", "login"),
			handlers.PostLogin{handlers.BaseHandler{Config: config, DB: db, CORS: allowedOriginHosts}}.ServeHTTP,
		},

		Route{
			"Signup",
			http.MethodPost,
			fmt.Sprintf("/%s/%s/%s", apiVersion, "user", "signup"),
			handlers.PostSignup{handlers.BaseHandler{Config: config, DB: db, CORS: allowedOriginHosts}}.ServeHTTP,
		},
	}
}
