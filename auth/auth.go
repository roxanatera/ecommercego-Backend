package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub       string
	Event_Id  string
	Token_use string
	Scope     string
	Auth_time int64
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func ValidoToken(token string) (bool, error, string) {

	// Split the token into its parts
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		fmt.Println("invalid token format")
		return false, nil, "El token no es valido"

	}

	// Decode the payload part of the token
	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("failed to decode payload part of token:", err.Error())
		return false, err, err.Error()
	}

	// Unmarshal the JSON payload into a TokenJSON struct
	var tokenJSON TokenJSON
	err = json.Unmarshal(userInfo, &tokenJSON)
	if err != nil {
		fmt.Println("No se puede decodificar la estructura JSON:  ", err.Error())
		return false, err, err.Error()
	}

	// Check if the token is expired

	ahora := time.Now()
	tm := time.Unix(int64(tokenJSON.Exp), 0)
	if ahora.Before(ahora) {
		fmt.Println("El token ha expirado = " + tm.String())
		fmt.Println("El token ha expirado")
		return false, err, "El token ha expirado"
	}

	return true, nil, string(tokenJSON.Username)
}
