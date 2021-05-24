package handlers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	//log "github.com/sirupsen/logrus"
)

// BaseHandler This is the base handler all the handlers extend
type BaseHandler struct {
	http.Handler

	Name		string
	Category	string

	DB		*gorm.DB
}

func NewBaseHandler(db *gorm.DB) BaseHandler {
	return BaseHandler{
		DB:	db,
	}
}

func (b *BaseHandler) WriteResponse(w http.ResponseWriter, reqCtx *ReqContext) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.Header().Add("X-Correlation-Id", reqCtx.Correlation)

	//if reqCtx.HTTPReply != nil {
	//	if err := json.NewEncoder(w).Encode(reqCtx.HTTPReply); err != nil {
	//		log.Errorf("%s %s Error encoding response: %v", b.Name, reqCtx.Correlation, err)
	//	}
	//}
}

type ReqContext struct {
	ErrorCode   string `json:"errorCode"`
	ErrorDetail string `json:"errorDetail"`

	HTTPStatusCode int `json:"httpStatus"`
	// Reply can be any type convertible to valid JSON
	HTTPReply interface{}
}
