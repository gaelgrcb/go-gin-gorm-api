# Go REST API &middot; Gin & GORM Exploration

![Go](https://img.shields.io/badge/Go-1.27+-00ADD8?style=flat&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Framework-Gin-00ADD8?style=flat&logo=go&logoColor=white)
![GORM](https://img.shields.io/badge/ORM-GORM-blue?style=flat)
![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL-336791?style=flat&logo=postgresql&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=flat)

> **Portfolio Project**: A modern, lightweight RESTful service built to gain hands-on experience and familiarity with the Go backend ecosystem, focusing primarily on the **Gin** web framework and **GORM** with PostgreSQL.

---

## 📌 Project Purpose

The primary goal of this project is to explore and master idiomatic Go design patterns for backend development:
- **Routing & HTTP handling**: Implementing high-performance routing, context management, and middleware with [Gin](https://github.com/gin-gonic/gin).
- **Persistence Layer**: Structuring database connections, connection pooling, and ORM operations using [GORM](https://gorm.io/) and PostgreSQL.
- **Application Configuration**: Managing environment variables and configurations dynamically via [Viper](https://github.com/spf13/viper).
- **Standardized API Responses**: Establishing consistent JSON payload formats for success, warning, and error responses across endpoints.

---

## 🛠️ Tech Stack

- **Language**: Go (v1.27+)
- **Web Framework**: [Gin Web Framework](https://gin-gonic.com/)
- **ORM**: [GORM](https://gorm.io/)
- **Database Driver**: [pgx / PostgreSQL Driver](https://github.com/jackc/pgx)
- **Configuration Management**: [Viper](https://github.com/spf13/viper)
- **Database**: PostgreSQL

---

## 📂 Project Structure

```text
.
├── cmd/
│   └── api/
│       └── main.go          # Application entrypoint & HTTP server
├── config/
│   └── config.go            # Viper configuration parser & validation
├── pkg/
│   ├── database/
│   │   └── database.go      # GORM database connection & connection pool
│   └── response/
│       └── response.go      # Unified JSON response helpers (Ok, Warn, Error)
├── .env.example             # Template for environment variables
├── .gitignore               # Git ignore rules
├── go.mod                   # Go module definitions
├── go.sum                   # Dependency checksums
└── README.md                # Project documentation
```

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) installed (v1.22+ or v1.27+)
- [PostgreSQL](https://www.postgresql.org/) instance running locally or via Docker

### 1. Clone & Navigate

```bash
git clone https://github.com/gaelgrcb/go-gin-gorm-api.git
cd go-gin-gorm-api
```

### 2. Configure Environment

Copy `.env.example` into a new `.env` file and adjust your database credentials:

```bash
cp .env.example .env
```

Example configuration:

```env
# Server
APP_NAME=crud-api
APP_ENV=development
APP_PORT=8080

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=midb

# Security
JWT_SECRET=your_jwt_secret_key_here
JWT_EXPIRES_HOURS=24

# CORS
CORS_ALLOWED_ORIGINS=http://localhost:3000
```

### 3. Install Dependencies

```bash
go mod download
```

### 4. Run the Application

```bash
go run cmd/api/main.go
```

The server will start listening on port `:8080`.

---

## 📡 API Endpoints

| Method | Endpoint | Description | Status |
| :--- | :--- | :--- | :--- |
| `GET` | `/` | Healthcheck / Smoke test (`pong`) | ✅ Active |

### Standard Response Format

Endpoints return a consistent response structure:

```json
{
  "status": "OK",
  "msg": "Operation successful",
  "data": {}
}
```

---

## 🗺️ Roadmap & Next Steps

- [ ] Define domain models and auto-migrations with GORM
- [ ] Implement full CRUD operations for resource entities
- [ ] Add JWT authentication & authorization middleware
- [ ] Containerize application with Docker & Docker Compose
- [ ] Write unit and integration tests

---

## 👤 Author

- **Gael García** - [@gaelgrcb](https://github.com/gaelgrcb)
