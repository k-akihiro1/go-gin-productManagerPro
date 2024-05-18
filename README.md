# go-gin-productManagerPro
This repository is implemented to refresh my memory of Go/Gin.

# Development Environment
Language: Go
Framework: Gin
ORM: Gorm
Database: MySQL
Authentication: JWT (gin-jwt)

# Adopted Architecture
Layered Architecture
Controller
- Handling request data
- Setting response
Service
- Functionality (implementation of business logic)
- Implemented functionalities:
  - Products: search all, search by ID, register, update, delete
  - Users: login, authentication features (applicable to: search by ID, register, update, delete)

Router - IController
         ↑
    Controller → IService
                 ↑
            Service → IRepository
                          ↑
                  Repository → DB
References:
https://qiita.com/fghyuhi/items/8d5c0f7f8aec643e5907
https://zenn.dev/taiyou/articles/747ab00a61a2f2
https://qiita.com/koji0705/items/49172d713e13fa554ba7