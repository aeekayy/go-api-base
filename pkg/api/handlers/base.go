package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aeekayy/go-api-base/pkg/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// BaseHandler This is the base handler all the handlers extend
type BaseHandler struct {
	http.Handler

	Name     string
	Category string

	DB     *gorm.DB
	CORS   map[string]bool
	Config *config.HTTPConfig
}

// NewBaseHandler returns a new base handler
func NewBaseHandler(db *gorm.DB, cors map[string]bool) BaseHandler {
	return BaseHandler{
		DB:   db,
		CORS: cors,
	}
}

// ReqContext object for the context of a request
type ReqContext struct {
	ErrorCode   string `json:"errorCode"`
	ErrorDetail string `json:"errorDetail"`

	HTTPStatusCode int `json:"httpStatus"`
	// Reply can be any type convertible to valid JSON
	HTTPReply interface{}
}

// WriteJSON writes a JSON response or an error if mashalling the object fails.
func (h *BaseHandler) WriteJSON(w http.ResponseWriter, status int, obj interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// TODO: Add correlation trail

	b, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, `{"error": %s}`, strconv.Quote(err.Error()))
	} else {
		w.WriteHeader(status)
		_, err = w.Write(b)

		if err != nil {
			log.Errorf("error while writing response %s: %s", h.Name, err)
		}
	}
}
