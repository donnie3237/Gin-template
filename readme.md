# 🧪 Gin Template with `internal` Packages

A starter template for building Go applications using the Gin web framework. This structure follows the Clean Architecture approach and uses Go's `internal` packages to organize code modularly. The `main.go` entry point is placed in the root directory.

---

## 📁 Project Structure

```
gin-template/
├── internal/
│   ├── config/         # Load configuration (from .env, etc.)
│   ├── handler/        # HTTP handlers (controller layer)
│   ├── middleware/     # Custom middleware (e.g. logger, auth)
│   ├── model/          # Request/response structs and DB models
│   ├── repository/     # Database interaction layer
│   ├── router/         # Route setup and grouping
│   └── service/        # Business logic layer
├── .env                # Environment variables
├── go.mod
├── go.sum
└── main.go             # Application entry point
```

---

## 🚀 Getting Started


### 1. Create a `.env` file

```env
PORT = 8081
MONGO_URI = ''
MONGO_DB = ''
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the server

```bash
go run main.go
```

## 🧩 Suitable For

- RESTful API backends
- Microservice architecture
- Web/mobile app backends

