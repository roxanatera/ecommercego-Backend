package main

import (
	"context"
	"ecommerce-Backend/awsgo"
	"ecommerce-Backend/dbconfig"
	"ecommerce-Backend/handlers"
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
		panic("Error: Debe enviar 'SecretName', 'UserPoolId', 'Region' y 'UrlPrefix' como variables de entorno.")
	}

	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), ""	, -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	dbconfig.ReadSecret(ctx)

	status, message := handlers.Manejadores(path, method, body, header, request)
	
	headersResp := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:      string(message),
		Headers: headersResp,
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