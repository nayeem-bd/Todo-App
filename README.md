# Todo App

A RESTful Todo application built with Go, featuring clean architecture principles, Docker support, and modern Go development practices.

## 🚀 Features

- **CRUD Operations**: Create, read, update, and delete todos
- **Clean Architecture**: Organized with domain-driven design patterns
- **PostgreSQL Database**: Persistent storage with GORM ORM
- **Docker Support**: Containerized application with Docker Compose
- **Environment Configuration**: Flexible configuration via environment variables
- **Validation**: Request validation using go-playground/validator
- **Configuration**: YAML-based configuration with Viper
- **HTTP Router**: Fast routing with Chi router
- **Graceful Shutdown**: Proper server shutdown handling
- **Health Checks**: Application and database health monitoring
- **Hot Reload**: Development support with Air
- **Structured Logging**: Logrus for production-ready logging

## 🏗️ Architecture

The application follows **Clean Architecture** principles with hexagonal architecture patterns:

```
todo-app/
├── cmd/                    # Application entry point
│   └── main.go            # Main application bootstrap
├── domain/                 # Business entities and core logic
│   ├── todo.go            # Todo entity/model
│   └── dto/               # Data Transfer Objects
│       └── todo.go        # Todo request/response DTOs
├── modules/               # Feature modules (Hexagonal Architecture)
│   └── todo/              # Todo bounded context
│       ├── delivery/      # Delivery layer (adapters)
│       │   └── http/      # HTTP handlers
│       ├── repository/    # Data access layer
│       │   └── todo.go    # Todo repository implementation
│       └── usecase/       # Business logic layer
│           └── todo.go    # Todo use cases
├── http/                  # HTTP infrastructure
│   ├── handlers.go        # HTTP route handlers
│   └── routes.go          # Route definitions
├── internal/              # Private application packages
│   ├── config/            # Configuration management
│   │   ├── config.go      # Config structures and loading
│   │   └── database.go    # Database configuration
│   ├── logger/            # Logging utilities
│   │   └── logger.go      # Structured logging setup
│   ├── middleware/        # HTTP middleware
│   │   └── logger.go      # Request logging middleware
│   ├── store/             # Database connection
│   │   └── store.go       # Database initialization
│   └── utils/             # Shared utilities
│       ├── response.go    # HTTP response helpers
│       └── validator.go   # Request validation
├── config.yaml           # Application configuration
├── docker-compose.yaml   # Docker orchestration
├── Dockerfile            # Container definition
├── .air.toml             # Hot reload configuration
├── .dockerignore         # Docker build exclusions
└── Makefile              # Build automation
```

### Architecture Layers

**1. Domain Layer** (`domain/`)
- Core business entities and rules
- Independent of external frameworks
- Contains Todo entity and DTOs

**2. Use Case Layer** (`modules/*/usecase/`)
- Application-specific business logic
- Orchestrates data flow between entities
- Independent of delivery mechanisms

**3. Interface Adapters** (`modules/*/delivery/`, `modules/*/repository/`)
- **Delivery**: HTTP handlers, API controllers
- **Repository**: Data access implementations
- Converts data between use cases and external interfaces

**4. Infrastructure Layer** (`internal/`, `http/`)
- Database connections, logging, configuration
- HTTP routing and middleware
- External service integrations

## 📋 Prerequisites

- Go 1.23.4 or higher
- Docker and Docker Compose
- Git
- Make (optional, for build automation)

## 🛠️ Installation & Setup

### Option 1: Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone https://github.com/nayeem-bd/Todo-App.git
   cd Todo-App
   ```

2. **Start the application with Docker Compose**
   ```bash
   # Build and start services
   docker-compose up --build

   # Or run in background
   docker-compose up -d --build
   ```

3. **Access the application**
   - API: http://localhost:8080

### Option 2: Local Development with Hot Reload

1. **Clone and install dependencies**
   ```bash
   git clone https://github.com/nayeem-bd/Todo-App.git
   cd Todo-App
   go mod download
   ```

2. **Setup PostgreSQL database**
   ```bash
   # Using Docker for database only
   docker run --name postgres-todo \
     -e POSTGRES_DB=todoapp \
     -e POSTGRES_USER=root \
     -e POSTGRES_PASSWORD=secret \
     -p 5432:5432 \
     -d postgres:15-alpine
   ```

3. **Install Air for hot reload** (optional)
   ```bash
   go install github.com/cosmtrek/air@latest
   ```

4. **Run with hot reload**
   ```bash
   # Using Air (recommended for development)
   air

   # Or run directly
   go run cmd/main.go
   ```

### Option 3: Using Makefile

```bash
# View available commands
make help

# Build the application
make build

# Run with Docker
make docker-up

# Clean build artifacts
make clean
```

## 🐳 Docker Configuration

### Services Architecture
- **app**: Go application container
  - Port: 8080
  - Health checks enabled
  - Resource limits configured
- **postgres**: PostgreSQL database
  - Port: 5432 (internal only by default)
  - Persistent volume storage
  - Health checks for reliability

### Environment Variables
```bash
# Database Configuration
DB_HOST=postgres              # Database host (postgres for Docker)
DB_PORT=5432                 # Database port
DB_NAME=todoapp              # Database name
DB_USER=root                 # Database user
DB_PASSWORD=secret           # Database password

# Application Configuration
APP_PORT=8080               # Application port
POSTGRES_PORT=5432          # Exposed PostgreSQL port

# Resource Management
APP_CPU_LIMIT=0.5           # CPU limit
APP_MEMORY_LIMIT=512M       # Memory limit
APP_CPU_RESERVATION=0.1     # CPU reservation
APP_MEMORY_RESERVATION=256M # Memory reservation
```

### Docker Commands
```bash
# Development
docker-compose up --build    # Build and start
docker-compose logs -f app   # Follow application logs
docker-compose exec app sh   # Access app container

# Production
docker-compose -f docker-compose.prod.yml up -d

# Maintenance
docker-compose down          # Stop services
docker-compose down -v       # Stop and remove volumes
docker system prune          # Clean up unused resources
```

## 🔗 API Endpoints

### Todo Operations
| Method | Endpoint | Description | Request Body | Response |
|--------|----------|-------------|--------------|----------|
| GET    | `/api/v1/todos` | List all todos | - | Array of todos |
| GET    | `/api/v1/todos/{id}` | Get todo by ID | - | Single todo object |
| POST   | `/api/v1/todos` | Create new todo | Todo object | Created todo |

### Request/Response Format

#### Create Todo Request
```json
{
  "title": "string (required, 3-150 chars)",
  "description": "string (required, 5-500 chars)", 
  "category": "string (optional, max 50 chars)"
}
```

#### Todo Response Object
```json
{
  "id": 1,
  "title": "Learn Clean Architecture",
  "description": "Study hexagonal architecture patterns in Go",
  "category": "learning",
  "created_at": "2025-01-07T10:30:00Z",
  "updated_at": "2025-01-07T10:30:00Z",
  "done_at": null
}
```

#### Success Response Format
```json
{
  "success": true,
  "message": "Todo created successfully",
  "data": {
    // Todo object or array of todos
  }
}
```

#### Error Response Format
```json
{
  "error": true,
  "message": "Validation failed",
  "status_code": 400,
  "errors": {
    // Validation errors or error details
  }
}
```

### Example API Usage

#### Create a Todo
```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Clean Architecture",
    "description": "Study hexagonal architecture patterns in Go",
    "category": "learning"
  }'
```

#### Get All Todos
```bash
curl http://localhost:8080/api/v1/todos
```

#### Get Specific Todo
```bash
curl http://localhost:8080/api/v1/todos/1
```

### Status Codes
- `200 OK` - Successful GET requests
- `201 Created` - Successful POST requests
- `400 Bad Request` - Invalid request body or parameters
- `404 Not Found` - Todo not found
- `500 Internal Server Error` - Server errors

### Current Limitations
- **Update Operations**: PUT/PATCH endpoints not yet implemented
- **Delete Operations**: DELETE endpoint not yet implemented  
- **Pagination**: All todos returned without pagination
- **Filtering**: No filtering or search capabilities
- **Authentication**: No authentication/authorization implemented

### Planned Enhancements
- [ ] Add UPDATE todo endpoint (`PUT /api/v1/todos/{id}`)
- [ ] Add DELETE todo endpoint (`DELETE /api/v1/todos/{id}`)
- [ ] Add mark as complete/incomplete functionality
- [ ] Implement pagination with query parameters
- [ ] Add filtering and search capabilities
- [ ] Add authentication and authorization
- [ ] Add todo categories and tags

## 🧪 Development Workflow

### Hot Reload Development
```bash
# Start with Air (automatic restart on file changes)
air

# Configuration in .air.toml
# - Watches Go files
# - Ignores vendor/, tmp/ directories
# - Builds to tmp/app
```

### Testing
```bash
# Run specific module tests
go test ./modules/todo/usecase -v
```

### Code Quality
```bash
# Format code
go fmt ./...

# Vet code for issues
go vet ./...

# Run golangci-lint (if installed)
golangci-lint run
```

### Database Migrations
```bash
# Auto-migration is handled by GORM on startup
# Manual migration commands can be added to Makefile
```

## 🔧 Configuration Management

### Hierarchy (highest to lowest priority)
1. **Environment Variables** (Docker/Production)
2. **config.yaml** (Development)
3. **Default Values** (Fallback)

### Configuration Structure
```yaml
server:
  port: "8080"

database:
  host: "127.0.0.1"
  port: 5432
  name: "todoapp"
  username: "root"
  password: "secret"
  options:
    sslmode: ["disable"]
  max_idle_connection: 2
  max_open_connection: 2
  max_connection_lifetime: 300
  batch_size: 10
  slow_threshold: 10
```

## 📊 Monitoring & Observability

### Health Checks
- **Application Health**: `GET /health`
- **Database Health**: Automatic connection testing
- **Container Health**: Docker health check probes

### Logging
- **Structured Logging**: JSON format with Logrus
- **Request Logging**: HTTP middleware
- **Error Tracking**: Contextual error information
- **Log Levels**: Debug, Info, Warn, Error, Fatal

### Metrics (Future Enhancement)
- Response time monitoring
- Database query performance
- Error rate tracking
- Resource utilization

## 🛡️ Security & Best Practices

### Container Security
- **Non-root user**: Application runs as `appuser`
- **Minimal base image**: Alpine Linux
- **Resource limits**: CPU and memory constraints
- **Health checks**: Automated failure detection

### Application Security
- **Input validation**: Request body validation
- **SQL injection prevention**: GORM ORM protection
- **Environment variables**: Secure configuration
- **Error handling**: No sensitive data exposure

### Production Readiness
- **Graceful shutdown**: Signal handling
- **Connection pooling**: Database optimization
- **Request timeout**: HTTP server configuration
- **Rate limiting**: (Future enhancement)

## 🔧 Troubleshooting

### Common Issues

**Database Connection Failed**
```bash
# Check database status
docker-compose logs postgres

# Verify environment variables
docker-compose config

# Reset database
docker-compose down -v && docker-compose up --build
```

**Application Not Starting**
```bash
# Check application logs
docker-compose logs app

# Verify configuration
cat config.yaml

# Check port conflicts
lsof -i :8080
```

**Build Issues**
```bash
# Clean and rebuild
make clean && make build

# Docker build with no cache
docker-compose build --no-cache app

# Check Go modules
go mod tidy && go mod verify
```

### Development Tips
- Use `air` for hot reload during development
- Check `.air.toml` for reload configuration
- Monitor logs with `docker-compose logs -f`
- Use `make` commands for consistent builds

## 📚 Dependencies & Tech Stack

### Core Dependencies
- **Chi Router** (v5.2.2): Fast HTTP router
- **GORM** (v1.30.0): ORM for database operations
- **Viper** (v1.20.1): Configuration management
- **Validator** (v10.27.0): Request validation
- **Logrus** (v1.9.3): Structured logging

### Database
- **PostgreSQL Driver** (v1.6.0): Database connectivity
- **Connection Pooling**: Optimized database access

### Development Tools
- **Air**: Hot reload for development
- **Docker**: Containerization
- **Make**: Build automation

[//]: # (## 🚀 Deployment)

[//]: # ()
[//]: # (### Production Deployment)

[//]: # (```bash)

[//]: # (# Build production image)

[//]: # (docker build -t todo-app:prod .)

[//]: # ()
[//]: # (# Deploy with Docker Compose)

[//]: # (docker-compose -f docker-compose.prod.yml up -d)

[//]: # ()
[//]: # (# Or deploy to Kubernetes)

[//]: # (kubectl apply -f k8s/)

[//]: # (```)

[//]: # ()
[//]: # (### Environment-Specific Configs)

[//]: # (- **Development**: `config.yaml` + local environment)

[//]: # (- **Staging**: Environment variables + Docker)

[//]: # (- **Production**: Environment variables + orchestration)

## 🤝 Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Follow coding standards (`go fmt`, `go vet`)
4. Add tests for new functionality
5. Commit changes (`git commit -m 'feat: add amazing feature'`)
6. Push to branch (`git push origin feature/amazing-feature`)
7. Open Pull Request

### Code Standards
- Follow Go conventions and idioms
- Use meaningful variable and function names
- Add comments for exported functions
- Maintain test coverage above 80%
- Follow Clean Architecture principles

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

**Nayeem Ahmed**
- GitHub: [@nayeem-bd](https://github.com/nayeem-bd)

## 🙏 Acknowledgments

- **Clean Architecture** by Robert C. Martin
- **Hexagonal Architecture** principles
- **Go community** for excellent tooling
- **Docker** for containerization platform
