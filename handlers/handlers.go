package handlers

import (
	"ecommerce-Backend/auth"
	"fmt"
	"strconv"
	"strings"

	"ecommerce-Backend/routers"

	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	// Debug: Verificar headers recibidos
	fmt.Printf("DEBUG - Request Headers: %+v\n", request.Headers)
	fmt.Printf("DEBUG - Path Parameters: %+v\n", request.PathParameters)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	// Usar request.Headers en lugar del parámetro headers
	isOk, statusCode, user := validoAuthorization(path, method, body, request.Headers)
	if !isOk {
		return statusCode, user
	}

	switch path[1:5] {
	case "user":
		return ProcesoUsers(body, path, method, user, id, request)
	case "product":
		return ProcesoProducts(body, path, method, user, idn, request)
	case "stock":
		return ProcesoStock(body, path, method, user, idn, request)
	case "address":
		return ProcesoAddress(body, path, method, user, idn, request)
	case "order":
		return ProcesoOrder(body, path, method, user, idn, request)
	case "cate":
		return ProcesoCategory(body, path, method, user, idn, request)
	}

	return 400, "Method Invalid"
}

func validoAuthorization(path string, method string, body string, headers map[string]string) (bool, int, string) {
	// Debug: Verificar todos los headers disponibles
	fmt.Printf("DEBUG - Validating auth with headers: %+v\n", headers)

	// Rutas públicas
	if (path == "/product" && method == "GET") ||
		(path == "/category" && method == "GET") {
		return true, 200, "OK"
	}

	// Obtener token (case-insensitive)
	token := ""
	for k, v := range headers {
		if strings.EqualFold(k, "authorization") {
			token = strings.TrimPrefix(v, "Bearer ")
			break
		}
	}

	if token == "" {
		fmt.Println("DEBUG - No authorization header found")
		return false, 401, "Token Requerido"
	}

	todoOk, err, msg := auth.ValidoToken(token)
	if !todoOk {
		fmt.Printf("DEBUG - Token validation failed. Error: %v, Msg: %s\n", err, msg)
		return false, 401, msg
	}

	fmt.Println("DEBUG - Token validado correctamente")
	return true, 200, msg
}

func ProcesoUsers(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {

	return 400, "Method Invalid"
}
func ProcesoProducts(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	return 400, "Method Invalid"
}
func ProcesoStock(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}
func ProcesoAddress(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func ProcesoOrder(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	return 400, "Method Invalid"
}
func ProcesoCategory(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertCategory(body, user)
	}
	return 400, "Method Invalid"
}
