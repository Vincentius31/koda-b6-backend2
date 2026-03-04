# Go CRUD API with Dependency Injection (DI)

This project is a simple RESTful API implementation using the **Go** programming language and the **Gin Gonic** framework. The project architecture utilizes the **Dependency Injection (DI)** pattern to decouple responsibilities between the *Repository*, *Service*, and *Handler* layers.

## 🚀 Features
* **User Management**: Create, Read, Update, Delete (CRUD) based on Email.
* **Product Management**: Create, Read, Update, Delete (CRUD) based on auto-incrementing IDs.
* **In-Memory Storage**: Data is stored temporarily in slices (RAM).
* **Dependency Injection**: Centralized dependency management located in the `internal/di` folder.

## 📁 Directory Structure
```text
.
├── cmd
│   └── main.go              # Application entry point
├── internal
│   ├── di
│   │   └── container.go     # Dependency Injection container/orchestrator
│   ├── handlers             # Layer for handling HTTP Requests/Responses
│   ├── models               # Data struct and DTO definitions
│   ├── repository           # Data access layer (In-memory)
│   └── service              # Business logic layer
└── api_test.http            # API testing file (REST Client)
```

## 🛠️ Getting Started
1. Ensure you have Go installed (version 1.18 or higher is recommended).
2. Clone this repository.
3. Install the Gin dependency:
```
go get [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
```
4. Run the application:
```
go run cmd/main.go
```
5. The server will be running at: `http://localhost:8888`