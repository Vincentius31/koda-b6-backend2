# User & Product Backend Service

A high-performance, production-ready RESTful API backend service built with Go. This project implements a clean architecture pattern with a robust dependency injection container, utilizing PostgreSQL for persistent storage and Redis for high-speed data caching.

## 🚀 Tech Stack

* **Language:** [Go (Golang)](https://go.dev/)
* **Web Framework:** [Gin HTTP web framework](https://gin-gonic.com/)
* **Database:** [PostgreSQL](https://www.postgresql.org/)
* **Database Driver:** [pgx/v5](https://github.com/jackc/pgx) (Connection Pooler)
* **Cache:** [Redis](https://redis.io/) via [go-redis/v9](https://github.com/redis/go-redis)
* **Environment Management:** [Godotenv](https://github.com/joho/godotenv)

## 🏗 Architecture

This project follows a layered architecture to ensure separation of concerns, testability, and maintainability:

* **`models/`**: Defines the data structures and domain entities.
* **`repository/`**: Handles direct interactions with the database (PostgreSQL).
* **`service/`**: Contains the core business logic and orchestration, including implementing Redis caching strategies (Cache-Aside pattern).
* **`handlers/`**: Manages HTTP requests/responses, payload validation, and interacts with the service layer.
* **`di/`**: Dependency Injection container that wires up repositories, services, and handlers.

## ⚙️ Prerequisites

Before running this project, ensure you have the following installed:
* Go (v1.18 or newer)
* PostgreSQL
* Redis server running locally or remotely

## 🛠 Setup & Installation

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/Vincentius31/koda-b6-backend2.git](https://github.com/Vincentius31/koda-b6-backend2.git)
    cd koda-b6-backend2
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Environment Variables:**
    Create a `.env` file in the root directory and configure your credentials:

    ```env
    # Server
    PORT=8888

    # PostgreSQL Database
    PGHOST=localhost
    PGPORT=5432
    PGUSER=postgres
    PGPASSWORD=your_db_password
    PGDATABASE=your_db_name
    PGSSLMODE=disable

    # Redis Cache
    REDIS_ADDR=localhost:6379
    ```

4.  **Database Migration/Seeding:**
    Ensure you have created the required tables (`users` and `products`) in your PostgreSQL database before running the application.

5.  **Run the application:**
    ```bash
    go run main.go
    ```
    *The server will start at `http://localhost:8888`.*

## 📡 API Endpoints

### Users
| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/users` | Get all users |
| `GET` | `/users/:email` | Get user by email (Cached) |
| `POST` | `/users` | Create a new user |
| `PUT` | `/users/:email` | Update user password |
| `DELETE`| `/users/:email` | Delete a user |

### Products
| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/products` | Get all products |
| `GET` | `/products/:id` | Get product by ID (Cached) |
| `POST` | `/products` | Create a new product |
| `PUT` | `/products/:id` | Update product details |
| `DELETE`| `/products/:id` | Delete a product |

## 💡 Caching Strategy
This API utilizes **Redis** to cache single-read operations (e.g., `GetByEmail` and `GetByID`). 
* **Cache Hit:** Data is served directly from Redis, reducing DB load and response time.
* **Cache Miss:** Data is queried from PostgreSQL, returned to the user, and stored in Redis with a 1-hour expiration.
* **Cache Invalidation:** Write operations (`PUT`, `DELETE`) automatically invalidate the respective cache keys to prevent stale data.