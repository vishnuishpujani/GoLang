package getHandlers

import (
	"Case-Study/Cart-Service/helpers"
	"encoding/json"
	"net/http"
)

type GetHandler struct{}

func (g *GetHandler) PingHandler(w http.ResponseWriter, r *http.Request) {
	response := helpers.ResponseMapper(200, "OK")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(response)
}
