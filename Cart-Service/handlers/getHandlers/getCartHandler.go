package getHandlers

import (
	"Case-Study/Cart-Service/helpers"
	"Case-Study/Cart-Service/interfaces"
	"encoding/json"
	"log"
	"net/http"
)

func (g *GetHandler) GetCartHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userID")
	log.Println(userID)

	w.Header().Set("Content-Type", "application/json; charset= UTF-8")
	cartItems, err := interfaces.DBClient.GetCartQuery(userID)
	if err != nil {
		response := helpers.ResponseMapper(400, "Bad Request")
		json.NewEncoder(w).Encode(response)

	}
	json.NewEncoder(w).Encode(cartItems)
}
