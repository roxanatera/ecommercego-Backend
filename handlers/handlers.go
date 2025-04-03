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
    // Debug mejorado
    fmt.Printf("DEBUG - Full Request Data:\nPath: %s\nMethod: %s\nHeaders: %+v\nPathParams: %+v\nBody: %s\n",
        path, method, request.Headers, request.PathParameters, body)

    // Obtener token de x-authorization (nuevo) o authorization (fallback)
    token := ""
    for key, value := range request.Headers {
        if strings.EqualFold(key, "x-authorization") {
            token = strings.TrimPrefix(value, "Bearer ")
            break
        }
        // Fallback para authorization si x-authorization no existe
        if strings.EqualFold(key, "authorization") && token == "" {
            token = strings.TrimPrefix(value, "Bearer ")
        }
    }

    // Validación del token (actualizada)
    if token == "" {
        return 401, "Token requerido (Headers recibidos: " + fmt.Sprintf("%v", request.Headers)
    }

    // Procesar ID de path
    id := request.PathParameters["id"]
    idn, _ := strconv.Atoi(id)

    // Validación de autorización (modificada para usar token directamente)
    isOk, statusCode, user := validoAuthorization(token, path, method)
    if !isOk {
        return statusCode, user
    }

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
