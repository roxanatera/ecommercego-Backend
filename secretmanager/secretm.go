package secretmanager

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"ecommerce-Backend/models"
)

// Función que obtiene el secreto de AWS Secret Manager
func GetSecret(ctx context.Context, nameSecret string) (models.SecretRDSJson, error) { 
	var datosSecret models.SecretRDSJson
	fmt.Println(" > GetSecret " + nameSecret)

	// Cargar configuración AWS usando el contexto
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return datosSecret, fmt.Errorf("error cargando configuración AWS: %v", err)
	}

	// Crear cliente de Secrets Manager con la configuración
	svc := secretsmanager.NewFromConfig(cfg)
	
	// Obtener el secreto usando el contexto
	clave, err := svc.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nameSecret),
	})
	if err != nil {
		return datosSecret, fmt.Errorf("error al obtener el secreto: %v", err)
	}

	// Decodificar el JSON del secreto
	err = json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	if err != nil {
		return datosSecret, fmt.Errorf("error decodificando secreto: %v", err)
	}

	fmt.Println(" > Secreto obtenido: " + nameSecret)
	return datosSecret, nil
}