
# ⚙️ E-commerce Backend (Go + AWS Serverless)

Backend desarrollado en **Go** y desplegado en una arquitectura completamente **serverless** en AWS. Este servicio gestiona usuarios, productos, categorías y órdenes mediante funciones Lambda, usando **Cognito** para autenticación y **Secrets Manager** para el manejo seguro de credenciales.

---

<p align="center">
  <img src="https://github.com/roxanatera/ecommercego-Backend/blob/main/diagram-backend-go.png" width="600" alt="Arquitectura Backend AWS">
</p>

## 📦 Tecnologías utilizadas

- **Go (Golang)** – Lenguaje de programación principal.
- **AWS Lambda** – Lógica backend serverless.
- **Amazon API Gateway** – Exposición de endpoints HTTP REST.
- **Amazon Cognito** – Autenticación de usuarios con tokens JWT.
- **AWS Secrets Manager** – Gestión segura de credenciales y configuración.
- **Amazon RDS (MySQL)** – Base de datos relacional.
- **AWS SDK v2 for Go** – Cliente para interactuar con servicios AWS.

---


## 🧾 Endpoints principales

Rutas como `GET /product` y `GET /category` son públicas. Todas las demás requieren validación de token.

---

| Método | Ruta         | Descripción                        |
|--------|--------------|------------------------------------|
| POST   | /signup      | Registro de nuevo usuario          |
| POST   | /category    | Crear nueva categoría (admin only) |
| GET    | /product     | Obtener productos (público)        |
| GET    | /category    | Obtener categorías (público)       |

---

## 🧬 Variables de entorno requeridas

```env
SecretName=nombre_del_secreto_en_secrets_manager
UserPoolId=eu-west-1_xxxxx
Region=eu-west-1
UrlPrefix=/api


### Crear una categoría (requiere token válido)

POST /category
Headers:
x-auth: Bearer eyJraWQiOiJ... (token JWT válido)
Body:
{
"CategName": "Deportes",
"CategPath": "/deportes"
}


Respuesta:

```json
{
  "CategID": 6
}´´´

## Variables de entorno necesarias
SecretName=my-secret
UserPoolId=eu-west-1_example
Region=eu-west-1
UrlPrefix=/api
Usar herramientas como sam local invoke o Postman para pruebas de rutas.

