package dbconfig

import (
	"context"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"ecommerce-Backend/models"
	"ecommerce-Backend/secretmanager"
	"os"
	"fmt"
)

var SecretModel models.SecretRDSJson

var Db *sql.DB
var err error

func ReadSecret(ctx context.Context) error {
	secretName := os.Getenv("SecretName")
	if secretName == "" {
		return fmt.Errorf("SecretName no está definido en las variables de entorno")
	}

	var err error
	SecretModel, err = secretmanager.GetSecret(ctx, secretName)
	return err
}

// Cambia esta función
func DbConnect() error {  
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println("Error al conectar a la base de datos: " + err.Error())
		return err
	}
	
	err = Db.Ping()
	if err != nil {
		fmt.Println("Error al verificar conexión: " + err.Error())
		return err
	}
	
	fmt.Println("Conexión exitosa a la base de datos")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {
	dbUser := claves.Username
	authToken := claves.Password
	dbEndpoint := claves.Host
	dbName := "ecommercego"
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?allowCleartextPasswords=true", 
		dbUser, 
		authToken, 
		dbEndpoint, 
		claves.Port, 
		
		dbName,
	)
	fmt.Println("DSN:", dsn)
	return dsn
}
func UserIsAdmin(UserUUID string)(bool, string){
	fmt.Println(" > UserIsAdmin start > ", UserUUID)

	err := DbConnect()
	if err != nil {
		return false, "Error al conectar a la base de datos: " + err.Error()
	}
	defer Db.Close()

	sentencia := "SELECT 1 FROM users WHERE userUUID ='"+UserUUID+"' AND user_Status = 1"
	fmt.Println(" > UserIsAdmin sentencia > ", sentencia)
	rows, err := Db.Query(sentencia)
	if err != nil {
		return false, "Error al ejecutar la consulta: " + err.Error()
	}
	
	var valor string
	rows.Next()
	rows.Scan(&valor)

	fmt.Println(" > UserIsAdmin valor > Ejecucion Exitosa- valor devuelto " + valor)
	if valor == "1" {
		return true, "User es Admin"
	}
	return false, "User no es Admin"
}
