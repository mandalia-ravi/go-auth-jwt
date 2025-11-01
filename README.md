# Golang Authentication with JWT (Gin + MongoDB)

A minimal and production-ready implementation of authentication in Go using **Gin-Gonic**, **MongoDB**, and **JWT**.  
Includes user registration, login, protected routes, refresh tokens, and password hashing using bcrypt.

---

## 🚀 Features

- User Registration with hashed passwords (bcrypt)
- Login with Email + Password
- Issues **Access Token + Refresh Token**
- Middleware to protect authenticated routes
- Refresh token endpoint to renew access token
- MongoDB for user persistence
- Clean and idiomatic Go project structure
- Example Dockerfile + Multi-arch build instructions

---

## 🛠 Tech Stack

| Component | Library / Version |
|----------|------------------|
| Language | Go 1.21+ |
| Web Framework | Gin (`github.com/gin-gonic/gin`) |
| Database | MongoDB (`go.mongodb.org/mongo-driver`) |
| Auth | JWT (`github.com/golang-jwt/jwt`) |
| Password Hashing | Bcrypt (`golang.org/x/crypto/bcrypt`) |

---

## 📦 Prerequisites

- Go installed (1.21 or higher)
- MongoDB (Local or hosted on Atlas)
- Docker (optional, for container deployment)
- Postman / curl for testing