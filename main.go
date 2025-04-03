package main

import (
	"context"
	"ecommerce-Backend/awsgo"
	"ecommerce-Backend/dbconfig"
	"ecommerce-Backend/handlers"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(CodigoLambda)
}

func CodigoLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	awsgo.InitAWS()

	if !ValidarParametros() {
		panic("Error: Debe enviar 'SecretName', 'UserPoolId', 'Region' y 'UrlPrefix' como variables de entorno")
	}

	// 1. Normalización de headers (case-insensitive)
	normalizedHeaders := make(map[string]string)
	for k, v := range request.Headers {
		normalizedHeaders[strings.ToLower(k)] = v
	}

	// 2. Debug inicial
	fmt.Printf("DEBUG - Incoming request: %+v\n", request)

	// 3. Procesamiento principal
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method

	dbconfig.ReadSecret(ctx)

	status, message := handlers.Manejadores(path, method, request.Body, normalizedHeaders, request)

	// 4. Construcción de respuesta con CORS
	res := &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       message,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	}

	return res, nil
}

func ValidarParametros() bool{
	_,traeParametros := os.LookupEnv("SecretName")
	if !traeParametros {
		return traeParametros
	}

	_,traeParametros = os.LookupEnv("UserPoolId")
	if !traeParametros {
		return traeParametros
	}

	_,traeParametros = os.LookupEnv("Region")
	if !traeParametros {
		return traeParametros
	}

	_,traeParametros = os.LookupEnv("UrlPrefix")
	if !traeParametros {		
		return traeParametros
	}

	return traeParametros


}