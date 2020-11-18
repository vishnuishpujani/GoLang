package reciever

import (
	delete "Case-Study/Cart-Service/handlers/deleteHandlers"
	get "Case-Study/Cart-Service/handlers/getHandlers"
	update "Case-Study/Cart-Service/handlers/putHandlers"
)

var Update update.PutHandler
var Get get.GetHandler
var Delete delete.DeleteHandler
