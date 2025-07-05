# Todo App

A RESTful Todo application built with Go, featuring clean architecture principles and modern Go development practices.

## 🚀 Features

- **CRUD Operations**: Create, read, update, and delete todos
- **Clean Architecture**: Organized with domain-driven design patterns
- **PostgreSQL Database**: Persistent storage with GORM ORM
- **Validation**: Request validation using go-playground/validator
- **Configuration**: YAML-based configuration with Viper
- **HTTP Router**: Fast routing with Chi router
- **Graceful Shutdown**: Proper server shutdown handling

## 🏗️ Architecture

The application follows Clean Architecture principles with the following structure:

```
├── cmd/                    # Application entry point
├── domain/                 # Business entities and interfaces
├── modules/todo/           # Feature-specific modules
│   ├── delivery/http/      # HTTP handlers
│   ├── repository/         # Data access layer
│   └── usecase/           # Business logic
├── internal/               # Private application code
│   ├── config/            # Configuration management
│   ├── store/             # Database connection
│   └── utils/             # Utility functions
└── http/                  # HTTP layer (routes, handlers)
```

## 📋 Prerequisites

- Go 1.23.4 or higher
- PostgreSQL 12 or higher
- Git

## 🛠️ Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/nayeem-bd/Todo-App.git
   cd Todo-App
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up PostgreSQL database**
   ```sql
   CREATE DATABASE todoapp;
   CREATE USER root WITH PASSWORD 'secret';
   GRANT ALL PRIVILEGES ON DATABASE todoapp TO root;
   ```

4. **Configure the application**
   
   Update `config.yaml` with your database credentials:
   ```yaml
   server:
     port: "8080"
   
   database:
     host: 127.0.0.1
     port: 5432
     name: todoapp
     username: root
     password: secret
     options:
       sslmode:
         - disable
   ```

5. **Run the application**
   ```bash
   go run cmd/main.go
   ```

The server will start on `http://localhost:8080`

## 📡 API Endpoints

### Todos

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | `/todos` | Get all todos |
| GET    | `/todos/{id}` | Get a specific todo |
| POST   | `/todos` | Create a new todo |
| PUT    | `/todos/{id}` | Update an existing todo |
| DELETE | `/todos/{id}` | Delete a todo |

### Request/Response Examples

**Create Todo**
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go",
    "description": "Complete the Go tutorial",
    "category": "learning"
  }'
```

**Response**
```json
{
  "id": 1,
  "title": "Learn Go",
  "description": "Complete the Go tutorial",
  "category": "learning",
  "created_at": "2025-07-06T10:30:00Z",
  "updated_at": "2025-07-06T10:30:00Z",
  "done_at": null
}
```

## 🗄️ Database Schema

### Todos Table

```sql
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    description VARCHAR(500) NOT NULL,
    category VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    done_at TIMESTAMP NULL
);
```

## 🔧 Configuration

The application uses YAML configuration located in `config.yaml`:

```yaml
server:
  port: "8080"                    # Server port

database:
  host: 127.0.0.1                # Database host
  port: 5432                     # Database port
  name: todoapp                  # Database name
  username: root                 # Database username
  password: secret               # Database password
  max_idle_connection: 2         # Maximum idle connections
  max_open_connection: 2         # Maximum open connections
  max_connection_lifetime: 300   # Connection lifetime in seconds
  batch_size: 10                 # Batch size for operations
```

## 🚀 Development

### Hot Reload (Optional)

For development with hot reload, you can use Air:

1. **Install Air**
   ```bash
   go install github.com/cosmtrek/air@latest
   ```

2. **Run with hot reload**
   ```bash
   air
   ```

### Building for Production

```bash
# Build binary
go build -o todo-app cmd/main.go

# Run binary
./todo-app
```

## 📦 Dependencies

- **Chi Router**: Fast HTTP router
- **GORM**: Go ORM for database operations
- **PostgreSQL Driver**: Database driver for PostgreSQL
- **Viper**: Configuration management
- **Validator**: Request validation

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- **Nayeem** - *Initial work* - [nayeem-bd](https://github.com/nayeem-bd)

## 🙏 Acknowledgments

- Clean Architecture principles by Robert C. Martin
- Go community for excellent libraries and tools
- PostgreSQL team for the robust database system
