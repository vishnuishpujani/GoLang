// InsertProductHandler godoc
// @Summary Insert a new Product
// @Description Insert a new Product with the input paylod
// @Tags product
// @Accept  json
// @Produce  json
// @Param product body model.Products true "insert product"
// @Success 200 {object} model.Products
// @Router /cart [put]
package handlers

import (
	"Case-Study/Cart-Service/helpers"
	"Case-Study/Cart-Service/interfaces"
	"Case-Study/Cart-Service/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type PutHandler struct{}

func (p *PutHandler) InsertProductHandler(w http.ResponseWriter, r *http.Request) {
	var product model.Products
	userID := r.Header.Get("userID")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &product)

	if err != nil {
		response := helpers.ResponseMapper(400, "error in getting response")
		json.NewEncoder(w).Encode(response)
	}

	err = interfaces.DBClient.InsertProduct(userID, product)
	if err != nil {
		response := helpers.ResponseMapper(400, "Request error")
		json.NewEncoder(w).Encode(response)
	}
	response := helpers.ResponseMapper(200, "OK")
	json.NewEncoder(w).Encode(response)

}
