# Go Fiber CRM Basic

A simple CRM project built with **Go**, **Fiber v2**, **GORM**, and **MySQL**.

## Features

- Create, Read, Update, Delete (CRUD) for Leads
- RESTful API endpoints
- MySQL database with GORM auto-migration
- Modular project structure using Go modules

## Project Structure

```

go-fiber-crm-basic/
├── cmd/
│   └── server/
│       └── main.go        # Entry point
├── internal/
│   ├── database/
│   │   └── db.go          # Database connection
│   └── lead/
│       ├── handler.go     # API handlers
│       ├── model.go       # Lead model
│       └── routes.go      # API routes
├── go.mod
├── README.md
└── .gitignore

````

## Installation

1. Clone the repo:

```bash
git clone https://github.com/Shashank-raj1907/go-fiber-crm-basic.git
cd go-fiber-crm-basic
````

2. Install dependencies:

```bash
go mod tidy
```

3. Set up MySQL database and user:

```sql
CREATE DATABASE crm_db;
GRANT ALL PRIVILEGES ON crm_db.* TO 'shashank'@'localhost' IDENTIFIED BY 'your_password';
FLUSH PRIVILEGES;
```

4. Configure `internal/database/db.go` with your DB credentials.

5. Run the server:

```bash
go run ./cmd/server/main.go
```

Server runs on `http://localhost:3000`.

## API Endpoints

| Method | Endpoint        | Description       |
| ------ | --------------- | ----------------- |
| POST   | /api/leads      | Create a new lead |
| GET    | /api/leads      | Get all leads     |
| GET    | /api/leads/\:id | Get lead by ID    |
| PUT    | /api/leads/\:id | Update lead by ID |
| DELETE | /api/leads/\:id | Delete lead by ID |

License

MIT License
