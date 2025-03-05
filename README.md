# Siren-Net Platform


A scalable platform for creating and managing AI-powered social interaction agents with adaptive communication capabilities.

## 📑 Table of Contents
- [📝 Project Overview](#-project-overview)
- [🚀 Key Features](#-key-features)
- [🏗️ Technical Architecture](#-technical-architecture)
- [🔧 Development Setup](#-development-setup)
- [⚙ Configuration](#-configuration)
- [📁 File Structure](#-file-structure)
- [📜 License](#-license)


## 📝 Project Overview

Siren-Net enables users to:
- Create personalized AI agents with configurable personalities
- Manage complex social interaction workflows
- Generate dynamic content based on contextual analysis
- Prioritize engagement using interaction metrics
- Deploy AI agents across multiple communication channels

## 🚀 Key Features

### 🧠 AI Personality Engine
- Context-aware conversation patterns
- Memory retention system for relationship building
- Multi-modal response generation (text + images)

### 🤖 Agent Management
- Web-based dashboard for agent configuration
- Real-time interaction monitoring
- Batch operations for agent groups

### ⚡ Smart Prioritization
- Interaction frequency scoring
- Relationship depth analysis
- Dynamic attention allocation
- Priority queue management

### 🎨 Content Generation System
- Contextual image creation API
- Style transfer capabilities
- Template-based content workflows
- Multi-model inference support

### 🔒 Security Features
- JWT-based authentication
- Role-based access control
- Request validation middleware
- Encrypted communication channels

## 🏗️ Technical Architecture

### 🛠 Backend (Go/Gin)
- **Gin** web framework with custom middleware
- Layered architecture (handlers ↔ services ↔ repositories)
- GORM with SQLite for data persistence

### 🖥 Frontend (React/JavaScript)
- React 18+ with functional components
- Shadcn UI component library

### 🧩 AI Services
| Service              | Technology     | Functionality                     |
|----------------------|----------------|-----------------------------------|
| Conversation Engine  | Ollama (DeepSeek) | NLP processing, dialog management |
| Image Generation     | Stable Diffusion | Context-aware image synthesis     |
| Priority Engine      | Custom Go      | Interaction scoring algorithms    |

## 🔧 Development Setup

### ⚙️ Prerequisites
- Go 1.21+
- Node 18+
- Docker 24+
- NVIDIA GPU (recommended)

```bash
# Clone repository
git clone https://github.com/yourusername/siren-net.git

# Start services
docker-compose up -d --build

# Access endpoints
http://localhost:5173  # Frontend
http://localhost:8080  # Backend API
http://localhost:11434 # Ollama
http://localhost:7860  # Stable Diffusion
```


### ⚙ Configuration

1. Copy backend/.env.example to backend/.env
2. Set required values:

```ini

# JWT Configuration
JWT_SECRET = your_secure_secret
DATABASE_URL = your_url_to_sqlite_database
```

### 📁 File Structure

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
│   │   │   ├── auth_handler_test.go
│   │   ├── middleware/
│   │   │   ├── auth_middleware.go
│   │   ├── models/
│   │   │   ├── agent.go
│   │   │   ├── client.go
│   │   │   ├── message.go
│   │   │   ├── transaction.go
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
│   ├── tests/
│   │   └── integration/
│   │       ├── auth_test.go
│   ├── go.mod
│   ├── go.sum
│   ├── .env
│   ├── .env.example
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

## 📜 License

MIT License - See [LICENSE](LICENSE) for details