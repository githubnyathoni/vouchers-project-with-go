# Voucher Management System

The Voucher Management System is a REST API-based application for managing vouchers, users, and transactions, developed using Golang and the Gin framework.

## Features

- Brand Management
  - Create new brand.
- Voucher Management
  - Create new vouchers.
  - Retrieve voucher details by ID.
  - Retrieve all vouchers by brand_id.
- Transaction Management
  - Create new transactions based on vouchers.
  - Retrieve transaction details by ID.

## Technologies

- Backend: Golang (Gin Framework)
- Database: PostgreSQL
- ORM: GORM
- Database Migration: golang-migrate
- Testing: Testify (Mocking & Assertions)

## Entity Relationship Diagram

![Entity Relationship Diagram Voucher Management](https://imgur.com/bMDP7Xj.png)

## Getting Started

### Installation

1. Clone the repository
   ```bash
   git clone https://github.com/githubnyathoni/vouchers-project-with-go.git
   cd vouchers-project-with-go
   ```
2. Install dependencies
   ```bash
   go mod tidy
   ```

### Environment Variables

Set up environment variables by creating a .env file (use .env.example as a reference):

```bash
cp .env.example .env
```

### Database Setup

1. Install PostgreSQL on your system and create a new database for the project.

2. Install the golang-migrate tool:
   ```bash
   go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```
3. Run the migrations to create database tables:
   ```bash
   migrate -database "postgres://<username>:<password>@<host>:<port>/<database_name>?sslmode=disable" -path ./migrations up
   ```

### Running the Application

```bash
go run cmd/server/main.go
```

### API Documentation

The OpenAPI specification for the APIs is defined in apispec.json. It provides the structure and details of all the available endpoints, request parameters, and response formats.

### Testing

Run unit tests

```bash
go test ./... -v
```
