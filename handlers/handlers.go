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
    // 1. Obtener token del header personalizado
    token := ""
    for key, value := range request.Headers {
        if strings.EqualFold(key, "x-auth") {
            token = strings.TrimPrefix(value, "Bearer ")
            break
        }
    }

    // 2. Debug mejorado
    fmt.Printf("DEBUG - Request completa: %+v\n", request)
    
    if token == "" {
        return 401, "Token requerido. Envía el token en el header 'x-auth'"
    }

    // 3. Validación del token
    isOk, statusCode, user := validoAuthorization(token, path, method)
    if !isOk {
        return statusCode, user
    }
    // Procesar ID de path
    id := request.PathParameters["id"]
    idn, _ := strconv.Atoi(id)

    

    // Routing existente
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

// Función validoAuthorization actualizada
func validoAuthorization(token string, path string, method string) (bool, int, string) {
    // Rutas públicas
    if (path == "/product" && method == "GET") || 
       (path == "/category" && method == "GET") {
        return true, 200, "public"
    }

    // Validar el token recibido
    todoOk, err, msg := auth.ValidoToken(token)
    if !todoOk {
        fmt.Printf("ERROR Validando token: %v - %s\n", err, msg)
        return false, 401, msg
    }

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
