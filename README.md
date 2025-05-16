![Go](https://img.shields.io/badge/Go-1.21-blue?logo=go)
![MySQL](https://img.shields.io/badge/MySQL-8.0-blue?logo=mysql)
![Gin](https://img.shields.io/badge/Gin_Framework-RESTful-lightgrey?logo=go)
![Docker](https://img.shields.io/badge/Docker-Containerized-blue?logo=docker)
![JWT](https://img.shields.io/badge/JWT-Secure%20Auth-orange?logo=jsonwebtokens)
![License](https://img.shields.io/badge/License-MIT-green)

# Inventory Management System (IMS)

A role-based inventory management system built using **Golang**, **Gin**, **GORM**, and **MySQL**, featuring secure JWT authentication and role-based access control.

---

## Features

* User registration and JWT-based login
* Role-based access control: `admin`, `manager`, `seller`, and `user`
* Product CRUD operations with ownership enforcement for sellers
* Order creation and product linking
* Role-specific route protection (e.g., only admins can delete)
* Dockerized backend and MySQL stack for easy setup

---

## Framework Used

* **Golang** with Gin (web framework)
* **GORM** (ORM for Golang)
* **MySQL** (via Docker)
* **JWT** for authentication
* **Docker & Docker Compose** for containerization
* **Postman** for API testing

---

## Project Structure

```
├── controllers/           # Request handlers
├── middleware/            # JWT and role-based auth middleware
├── models/                # GORM models (User, Product, Order)
├── routes/                # Route definitions
├── config/                # DB config and JWT setup
├── main.go                # Entry point
├── Dockerfile             # Dockerized Go app
├── docker-compose.yml     # Multi-service container definition
├── .env                   # Local secrets (ignored by git)
├── .env.example           # Template for environment variables
```

---

## Setup & Initialization

### 1. Clone and Setup

```bash
git clone https://github.com/yourusername/inventory-system.git
cd inventory-system
cp .env.example .env  # then fill in your secrets
```

### 2. Environment Variables

Edit your `.env` file:

```
DB_HOST=db
DB_PORT=3306
DB_USER=inventory_user
DB_PASSWORD=yourpassword
DB_NAME=inventory
JWT_SECRET=yourjwtsecret
```

### 3. Run with Docker Compose

```bash
docker-compose up --build
```

App runs on [http://localhost:8080](http://localhost:8080)

---

## Role-Based Access Control

| Role    | View | Add | Update       | Delete |
| ------- | ---- | --- | ------------ | ------ |
| user    | ✅    | ❌   | ❌            | ❌      |
| seller  | ✅    | ✅   | ✅ (own only) | ❌      |
| manager | ✅    | ✅   | ✅            | ❌      |
| admin   | ✅    | ✅   | ✅            | ✅      |

> Middleware checks JWT token for `role` and `user_id`. Sellers can only edit their own products.

---

## Database Design

* `users`: `id`, `username`, `password`, `role`
* `products`: `id`, `name`, `price`, `quantity`, `owner_id`
* `orders`: `id`, `product_id`, `quantity`, `buyer_id`

---

## Key API Endpoints

### Authentication

```
POST /auth/register
POST /auth/login
```

### Products

```
GET    /products                   # Public
POST   /manage/products           # Admin/Manager/Seller
PUT    /manage/products/:id       # Admin/Manager/Seller (own only if seller)
DELETE /admin/products/:id        # Admin only
```

### Orders (for future expansion)

```
GET /orders
POST /orders
```

---

## Docker Notes

* `docker-compose.yml` runs both Go app and MySQL
* `host.docker.internal` used to access host DB if needed
* Data is not persisted unless volume is added

---

## Testing

Use **Postman** to test APIs:

* Login to get JWT
* Add `Authorization: Bearer <token>` header
* Test restricted and public routes

---

## Future Improvements

* Add order tracking and product stock updates
* Add audit logs for admin actions
* Add Swagger docs for API endpoints
* Add unit & integration testing
