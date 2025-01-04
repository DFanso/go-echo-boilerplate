# Go Echo Boilerplate

A production-ready boilerplate for building RESTful APIs using Go Echo framework with MongoDB.

## Features

- **Echo Framework**: High performance, minimalist Go web framework
- **MongoDB Integration**: Using qmgo as ORM
- **Clean Architecture**: Follows clean architecture principles with proper separation of concerns
- **Environment Management**: Multiple environment support with proper configuration management
- **Structured Logging**: Custom logging middleware with colored output
- **API Versioning**: Built-in support for API versioning
- **Input Validation**: Request validation using go-playground/validator
- **Error Handling**: Centralized error handling with proper HTTP status codes
- **Hot Reload**: Support for hot reload during development using Air

## Project Structure

```
.
├── cmd/
│   └── main.go                 # Application entry point
├── config/
│   └── config.go              # Configuration management
├── internal/
│   ├── controllers/           # Request handlers
│   ├── models/                # Data models
│   ├── repositories/          # Data access layer
│   ├── routes/                # Route definitions
│   │   └── v1/               # API version 1 routes
│   └── services/             # Business logic
├── pkg/
│   ├── database/             # Database connections
│   ├── middleware/           # Custom middleware
│   └── utils/                # Utility functions
├── .air.toml                 # Air configuration for hot reload
├── .env.development          # Development environment variables
├── .env.production          # Production environment variables
├── go.mod                    # Go module file
└── README.md                # Project documentation
```

## Prerequisites

- Go 1.21 or higher
- MongoDB 4.4 or higher
- Air (for hot reload)

## Getting Started

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/go-echo-boilerplate.git
   cd go-echo-boilerplate
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   ```bash
   cp .env.development .env
   ```
   Edit the `.env` file with your configuration.

4. **Install Air for hot reload**
   ```bash
   go install github.com/cosmtrek/air@latest
   ```

5. **Run the application**
   
   Development mode with hot reload:
   ```bash
   air
   ```

   Production mode:
   ```bash
   go run cmd/main.go
   ```

## Environment Variables

```env
# Server Configuration
SERVER_PORT=8080
APP_ENV=development
APP_NAME=go-echo-api
LOG_LEVEL=debug

# MongoDB Configuration
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=go_echo_db
MONGODB_TIMEOUT_SECONDS=30
```

## API Endpoints

### User Routes
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create new user
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Health Check
- `GET /health` - Service health check

## Request/Response Examples

### Create User
```bash
# Request
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com"
  }'

# Response
{
  "status": "success",
  "message": "User created successfully",
  "data": {
    "id": "60d5ecb8e3c8768b3c4b4b1e",
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2025-01-04T10:30:00Z",
    "updated_at": "2025-01-04T10:30:00Z"
  }
}
```

## Development

### Adding New Routes

1. Create a new controller in `internal/controllers/`
2. Add corresponding service and repository if needed
3. Register routes in `internal/routes/v1/`

Example:
```go
// internal/routes/v1/new_route.go
package v1

func RegisterNewRoutes(v1 *echo.Group, controller *controllers.NewController) {
    group := v1.Group("/resource")
    group.GET("", controller.GetAll)
    group.POST("", controller.Create)
}
```

### Error Handling

The boilerplate includes a centralized error handling system. Use the utility functions in `pkg/utils/response.go`:

```go
if err != nil {
    return utils.ErrorResponse(ctx, http.StatusBadRequest, "Error message", err)
}
return utils.SuccessResponse(ctx, http.StatusOK, "Success message", data)
```

## Testing - not implemented yet

Run tests:
```bash
go test ./...
```

With coverage:
```bash
go test ./... -cover
```

## Production Deployment

1. Build the application:
   ```bash
   go build -o app cmd/main.go
   ```

2. Set up production environment:
   ```bash
   cp .env.production .env
   ```

3. Run the application:
   ```bash
   ./app
   ```


## License

This project is licensed under the MIT License - see the LICENSE file for details