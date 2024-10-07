# Go Fiber Project

This is a Go Fiber project setup that includes Fiber as the web framework, SQLx for database interaction with PostgreSQL, JWT for token management, Go-i18n for localization, and Goose for database migrations. Environment variables and YAML config files are also supported for easy configuration management.

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or higher)
- [PostgreSQL](https://www.postgresql.org/)
- A terminal to run commands

## Project Setup

### 1. Install Project Dependencies

To install the necessary Go modules, run the following commands:

```bash
# Install Fiber (Web framework)
go get github.com/gofiber/fiber/v2

# Install SQLX (Database interaction)
go get github.com/jmoiron/sqlx

# Install PostgreSQL driver for SQLX
go get github.com/lib/pq

# Install JWT (For generating and validating tokens)
go get github.com/dgrijalva/jwt-go

# Install Goose (Database migrations)
go get github.com/pressly/goose/v3

# Install Go-i18n (Internationalization for translations)
go get github.com/nicksnyder/go-i18n/v2/i18n

# Install YAML package (For reading .yaml config files)
go get gopkg.in/yaml.v2

# Optional: For managing environment variables
go get github.com/joho/godotenv

# Create a new migration file
go install github.com/pressly/goose/v3/cmd/goose@latest
goose create add_some_column_or_table sql

# Example output:
# Created new file: 20170506082420_add_some_column.sql
# go_project_setup
