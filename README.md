# User Management System

A simple User Management System built using the Gin framework in Go. This project allows for user registration, login, and basic CRUD operations for user data.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)

## Features

- User registration
- User login
- User CRUD operations (Create, Read)
- Password hashing and validation
- SQLite database integration

## Technologies Used

- **Go**: The programming language used for the backend.
- **Gin**: A web framework for Go, which provides routing and middleware support.
- **GORM**: An ORM library for Go to manage database operations.
- **SQLite**: A lightweight database used for storing user data.

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/ialaminpro/go-gin-user-crud.git
   cd go-gin-user-crud
   ```

2. **Install Go**: Ensure you have Go installed on your machine. You can download it from the [official Go website](https://golang.org/dl/).

3. **Install dependencies**:
   ```bash
   go get -u gin-gonic/gin
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/sqlite
   ```

## Usage

1. **Run the application**:
   ```bash
   go run main.go
   ```

2. **Access the API**: The application will run on `http://localhost:8080`.

3. **Use a tool like Postman or cURL to interact with the API**.

## API Endpoints

| Method | Endpoint         | Description                   |
|--------|------------------|-------------------------------|
| POST   | /register        | Register a new user           |
| POST   | /login           | Login with user credentials    |
| GET    | /users           | Get a list of users           |


### Example Requests

- **Register User**:
  ```bash
  curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username": "john_doe", "password": "securepassword"}'
  ```

- **Login User**:
  ```bash
  curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username": "john_doe", "password": "securepassword"}'
  ```

- **Get User List**:
  ```bash
  curl -X GET http://localhost:8080/users
  ```

## Acknowledgments

- [Gin](https://github.com/gin-gonic/gin) - Web framework for Go
- [GORM](https://gorm.io/) - ORM library for Go
- [SQLite](https://www.sqlite.org/index.html) - Lightweight database
```
