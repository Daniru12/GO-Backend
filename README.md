
# Personal Profile Management System (Go CRUD API)

A **CRUD API** built in **Go** for managing personal profiles, including **soft delete functionality**. This project demonstrates clean architecture using **Go-Kit**, **MySQL**, and **Gorilla Mux** for routing.

---

## Features

* ✅ Create, Read, Update personal profiles
* ✅ Soft delete functionality (`status = "D"`) instead of actual deletion
* ✅ Fetch single or all profiles
* ✅ Handles errors gracefully with structured JSON responses
* ✅ Logging for all operations
* ✅ Metrics endpoint using **Prometheus**

---

## Tech Stack

* **Language:** Go
* **Architecture:** Clean Architecture (Repository → Use Case → Service → Endpoint → HTTP Handler)
* **Database:** MySQL
* **Routing:** Gorilla Mux
* **Transport:** Go-Kit HTTP Server
* **Logging:** Custom logger (project1/logger)
* **Monitoring:** Prometheus

---

## Project Structure

```
project1/
├── config/                  # Configuration for app (port, DB)
├── database/                # DB connection
├── error-handler/           # Application errors and responses
├── logger/                  # Logging utilities
├── repositories/            # DB repository layer
├── services/                # Service layer
├── usecases/                # Use case / interactor layer
├── transport/
│   ├── endpoints/           # Go-Kit endpoints
│   ├── request/             # Request decoders
│   └── response/            # Response encoders
├── util/                    # Utilities (CustomTime, StringPtr)
└── main.go                  # Application entry point
```

---

## API Endpoints

| Method | Endpoint                                | Description                            |
| ------ | --------------------------------------- | -------------------------------------- |
| GET    | `/personal/profile/{personal_id}`       | Fetch single profile by ID             |
| GET    | `/personal/studentprofiles`             | Fetch all profiles (excluding deleted) |
| POST   | `/personal/createprofile`               | Create new profile                     |
| PATCH  | `/personal/updateprofile/{personal_id}` | Update profile fields (partial update) |
| PATCH  | `/personal/deleteprofile/{personal_id}` | Soft delete profile (`status = "D"`)   |
| GET    | `/metrics`                              | Prometheus metrics                     |

---

## Request & Response Examples

**Create Profile**

```bash
POST /personal/createprofile
Content-Type: application/json

{
    "name": "John Doe",
    "description": "Software Engineer",
    "status": "A"
}
```

**Response**

```json
{
    "request": {
        "name": "John Doe",
        "description": "Software Engineer",
        "status": "A"
    },
    "response": {
        "id": 1,
        "name": "John Doe",
        "description": "Software Engineer",
        "status": "A",
        "create_time": "2025-08-19 11:00:00",
        "update_time": "2025-08-19 11:00:00"
    }
}
```

**Soft Delete**

```bash
PATCH /personal/deleteprofile/1
```

**Response**

```json
{
    "message": "Profile deleted successfully"
}
```

---

## Setup Instructions

1. **Clone the repository**

```bash
git clone https://github.com/yourusername/personal-profile-crud.git
cd personal-profile-crud
```

2. **Set up MySQL database**

```sql
CREATE DATABASE personal_profiles_db;
USE personal_profiles_db;

CREATE TABLE personal_profiles (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    status CHAR(1) DEFAULT 'A',
    create_time DATETIME,
    update_time DATETIME
);
```

3. **Configure your project**

* Edit `config/config.go` with your MySQL credentials and server port.

4. **Run the project**

```bash
go mod tidy
go run main.go
```

5. **Access API**

* Base URL: `http://localhost:1234`
* Metrics: `http://localhost:1234/metrics`

---

## Notes

* Soft delete: Updates `status` to `"D"` instead of removing the record.
* Partial updates supported via PATCH method.
* All errors returned as JSON with proper HTTP status codes.

---

## License

MIT License © 2025

