package routers

import (
	"encoding/json"
	"ecommerce-Backend/models"
	"ecommerce-Backend/dbconfig"
	"strconv"
)

func InsertCategory(body string, User string) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en el body de la categoria" + err.Error()
	}
	if len(t.CategName) == 0 {
		return 400, "Error en el body de la categoria Name, faltan datos"
	}
	if len(t.CategPath) == 0 {
		return 400, "Error en el body de la categoria Path, faltan datos"
	}

	isAdmin, msg := dbconfig.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err := dbconfig.InsertCategory(t)
	if err != nil {
		return 400, "Error al insertar la categoria: " + err.Error()
	}
	if result == 0 {
		return 400, "Error al insertar la categoria"
	}
	return 200, "{ CategID: " + strconv.Itoa(int(result)) + "}"
	
}