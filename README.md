# Todo App

A RESTful Todo application built with Go, featuring clean architecture principles, Docker support, and modern Go development practices.

## 🚀 Features

- **RESTful API**: Create, read, and mark todos as complete
- **Clean Architecture**: Organized with domain-driven design and hexagonal architecture patterns
- **PostgreSQL Database**: Persistent storage with GORM ORM
- **Redis Caching**: High-performance caching for improved response times
- **RabbitMQ Messaging**: Event-driven architecture for todo completion notifications
- **Docker Support**: Containerized application with Docker Compose
- **Environment Configuration**: Flexible configuration via YAML files and environment variables
- **Input Validation**: Request validation using go-playground/validator
- **Prometheus Metrics**: Built-in metrics endpoint for monitoring
- **Chi Router**: Fast and lightweight HTTP routing
- **Graceful Shutdown**: Proper server shutdown handling with context cancellation
- **Structured Logging**: Production-ready logging with Logrus
- **Request Middleware**: Logger, recovery, and real IP detection

## 🏗️ Architecture

```
todo-app/
├── cmd/                    # Application commands
│   ├── serve.go           # HTTP server command
│   └── worker.go          # Background worker command
├── domain/                # Domain entities and DTOs
│   ├── todo.go           # Todo domain model
│   └── dto/              # Data transfer objects
├── http/                 # HTTP layer
│   ├── handlers.go       # HTTP handlers
│   └── routes.go         # Route definitions
├── internal/             # Internal packages
│   ├── config/          # Configuration management
│   ├── logger/          # Logging utilities
│   ├── middleware/      # HTTP middleware
│   ├── migrations/      # Database migrations
│   ├── queue/          # Queue worker implementation
│   ├── store/          # Data store interfaces
│   └── utils/          # Utility functions
├── modules/             # Feature modules
│   └── todo/           # Todo module
│       ├── delivery/   # Delivery layer (HTTP, Queue)
│       ├── repository/ # Data persistence layer
│       └── usecase/    # Business logic layer
└── deployment/         # Kubernetes deployment files
```

## 🛠️ Prerequisites

- Go 1.19 or higher
- Docker and Docker Compose
- PostgreSQL 15+
- Redis 7+
- RabbitMQ 3+

## 🚀 Quick Start

### Using Docker Compose (Recommended)

1. Clone the repository:
```bash
git clone <repository-url>
cd todo-app
```

2. Start the services:
```bash
docker-compose up -d
```

3. Run the application:
```bash
go run main.go serve
```

4. Start the background worker:
```bash
go run main.go work
```

### Manual Setup

1. Start PostgreSQL, Redis, and RabbitMQ services
2. Configure environment variables (see Configuration section)
3. Run database migrations
4. Start the application

## ⚙️ Configuration

The application uses environment variables and YAML configuration files. Key environment variables:

```bash
# Database
DB_NAME=todoapp
DB_USER=root
DB_PASSWORD=secret
DB_HOST=localhost
DB_PORT=5432

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=""

# RabbitMQ
RABBITMQ_USER=guest
RABBITMQ_PASSWORD=guest
RABBITMQ_HOST=localhost
RABBITMQ_PORT=5672
```

## 📋 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | `/api/todos` | Get all todos |
| POST   | `/api/todos` | Create a new todo |
| GET    | `/api/todos/{id}` | Get a specific todo |
| PUT    | `/api/todos/{id}` | Update a todo |
| DELETE | `/api/todos/{id}` | Delete a todo |
| PATCH  | `/api/todos/{id}/complete` | Mark todo as complete |
| GET    | `/metrics` | Prometheus metrics |
| GET    | `/health` | Health check endpoint |

## 🔨 Development

### Hot Reloading

For development with hot reloading:

```bash
make watch
```

### Running Tests

```bash
go test ./...
```

### Building the Application

```bash
go build -o bin/todo-app main.go
```

## 📦 Deployment

### Docker

Build and run with Docker:

```bash
docker build -t todo-app .
docker-compose up -d
```

### Kubernetes

Deploy to Kubernetes:

```bash
kubectl apply -f deployment/
```

## 🐛 Troubleshooting

### Common Issues

1. **Database Connection Issues**: Ensure PostgreSQL is running and credentials are correct
2. **Redis Connection Issues**: Verify Redis service is accessible
3. **RabbitMQ Connection Issues**: Check RabbitMQ service status and credentials

### Logs

Check application logs for detailed error information:

```bash
docker-compose logs -f app
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- Built with [Chi](https://github.com/go-chi/chi) router
- Database ORM powered by [GORM](https://gorm.io/)
- Caching with [go-redis](https://github.com/redis/go-redis)
- Message queuing with [amqp091-go](https://github.com/rabbitmq/amqp091-go)
- Configuration management with [Viper](https://github.com/spf13/viper)
