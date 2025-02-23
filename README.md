# siren-net


## File structure

```
siren-net/
├── backend/
│   ├── cmd/
│   │   └── web/
│   │       └── main.go
│   ├── internal/
│   │   ├── app/
│   │   │   ├── app.go
│   │   ├── config/
│   │   │   ├── config.go
│   │   ├── handlers/
│   │   │   ├── auth_handler.go
│   │   ├── middleware/
│   │   │   ├── auth_middleware.go
│   │   ├── models/
│   │   │   ├── user.go
│   │   ├── routes/
│   │   │   ├── routes.go
│   │   ├── services/
│   │   │   ├── auth_service.go
│   │   │   ├── errors.go
│   │   │   ├── user_service.go
│   │   └── utils/
│   │       ├── password.go
│   ├── pkg/
│   │   └── database/
│   │       ├── database.go
│   ├── go.mod
│   ├── go.sum
│   ├── .env
│   └── Dockerfile
├── frontend/
│   └── ...
├── stable-diffusion/
│   ├── models/  -- stable diffusion models and loras
│   ├── outputs/ -- stable diffusion outputs
│   └── Dockerfile
├── ollama/
│   ├── models/  -- deepseek models
│   ├── Dockerfile
│   ├── Modelfile
│   └── start.sh
├── .gitignore
├── .gitmodules
├── docker-compose.yml
├── LICENSE
└── README.md
```