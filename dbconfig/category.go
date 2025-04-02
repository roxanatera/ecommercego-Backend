package dbconfig

import (
	
	"database/sql"
	"fmt"
	//"strconv"
	//"strings"

	_"github.com/go-sql-driver/mysql"
	"ecommerce-Backend/models"
	//"ecommerce-Backend/tools"
)

func InsertCategory(c models.Category)(int64, error){
	fmt.Println("Comienza Registro de Categoria")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentencia := "INSERT INTO category (Categ_Name, Categ_Path) VALUES ('" + c.CategName + "', '" + c.CategPath + "')"

	var result sql.Result
	result, err = Db.Exec(sentencia)
	if err != nil {
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}
	fmt.Println("Insert Category > Ejecucion Exitosa")
	return LastInsertId, nil
}