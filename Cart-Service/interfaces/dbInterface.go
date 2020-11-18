package interfaces

import "Case-Study/Cart-Service/model"

//DBClient object
var DBClient DBInteractions

//-----------------------------------------------------------------------------------------------------------------------//

//DBInteractions interface contains database methods
type DBInteractions interface {
	DBConnect(config model.DBConfig) error
	GetCartQuery(string) (model.Cart, error)
	InsertProduct(string, model.Products) error
	UpdateQuantity(string, string, int, float64) error
	DeleteProduct(string, string) error
}
