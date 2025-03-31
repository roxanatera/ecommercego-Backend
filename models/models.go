package models

import (

)

type SecretRDSJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Engine string `json:"engine"`
	Host string `json:"host"`
	Port int `json:"port"`
	Dbname string `json:"dbname"`
}


type SignUp struct {
    UserUUID      string `json:"uuid"`
    UserEmail     string `json:"email"`
    UserFirstName string `json:"first_name"`
    UserLastName  string `json:"last_name"`
}