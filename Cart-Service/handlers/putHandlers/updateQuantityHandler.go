package handlers

import (
	"Case-Study/Cart-Service/helpers"
	"Case-Study/Cart-Service/interfaces"
	"Case-Study/Cart-Service/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (g *PutHandler) UpdateQuantityHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userID")
	log.Println(userID)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var quantity model.Quantity

	params := mux.Vars(r)
	productID := params["productID"]
	log.Println(productID)
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &quantity)
	if err != nil {
		response := helpers.ResponseMapper(400, "error in getting response")
		json.NewEncoder(w).Encode(response)
	}

	err = interfaces.DBClient.UpdateQuantity(userID, productID, quantity.Quantity, quantity.Price)
	if err != nil {
		log.Println(err)
		response := helpers.ResponseMapper(400, "Internal Error")
		json.NewEncoder(w).Encode(response)
	} else {
		response := helpers.ResponseMapper(200, "OK")
		json.NewEncoder(w).Encode(response)
	}
}
