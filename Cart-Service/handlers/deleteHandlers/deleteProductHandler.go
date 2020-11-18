package deleteHandlers

import (
	"Case-Study/Cart-Service/helpers"
	"Case-Study/Cart-Service/interfaces"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type DeleteHandler struct{}

func (g *DeleteHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userID")
	log.Println(userID)
	w.Header().Set("Content-Type", "application/json;,charset=UTF-8")
	params := mux.Vars(r)
	productID := params["productID"]
	log.Println(productID)

	err := interfaces.DBClient.DeleteProduct(userID, productID)
	if err != nil {
		response := helpers.ResponseMapper(400, "Database error")
		json.NewEncoder(w).Encode(response)
	} else {
		response := helpers.ResponseMapper(200, "OK")
		json.NewEncoder(w).Encode(response)
	}
}
