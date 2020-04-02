package handler

import (
	"net/http"

	"github.com/yuhando/simpleapi/app/lib"
)

// HeartBeat ...
func HeartBeat(w http.ResponseWriter, r *http.Request) {
	lib.RespondNoContent(w, http.StatusNoContent, r)
}

// GetHealtCheck ...
func GetHealtCheck(w http.ResponseWriter, r *http.Request) {
	lib.RespondBlank(w, http.StatusOK)
}
