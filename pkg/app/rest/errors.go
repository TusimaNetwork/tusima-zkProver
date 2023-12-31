package rest

import (
	"fmt"
	"github.com/LimeChain/crc-prover/pkg/log"
	"github.com/go-chi/render"
	"net/http"
)

// ErrorJSON makes json and respond with error
func ErrorJSON(w http.ResponseWriter, r *http.Request, httpStatusCode int, err error, details string, errCode int) {
	log.WithContext(r.Context()).Error(fmt.Sprintf("%d - %d - %v - %s", httpStatusCode, errCode, err, details))
	render.Status(r, httpStatusCode)
	render.JSON(w, r, map[string]interface{}{"code": errCode, "error": err.Error(), "details": details})
}
