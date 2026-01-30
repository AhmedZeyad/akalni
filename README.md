# Akelni – Backend Tech Stack

## Overview
Akelni is a food delivery backend system designed for scalability, security, and performance, built using **Go** and modern backend best practices.

---

## 🛠 Technical Stack

| Category | Technology | Key Features & Benefits |
| :--- | :--- | :--- |
| **Language** | **Go (Golang)** | Strong performance, built-in concurrency, and simple syntax. |
| **Framework** | **Gin** | Lightweight, high-speed HTTP framework with middleware support. |
| **Database** | **PostgreSQL** | Reliable relational database with strong transactional support. |
| **DB Access** | **sqlx & pgx** | High-performance drivers; keeps SQL control without ORM overhead. |
| **Auth** | **JWT** | Stateless authentication using Access/Refresh token strategy. |
| **API Style** | **RESTful** | Resource-based endpoints with JSON request/response formats. |

---

## 🏗 Development Practices

To ensure long-term maintainability, the project adheres to the following standards:

* **Clean Architecture:** Clear separation of concerns between business logic and external layers.
* **RBAC:** Role-Based Access Control for **Admin, User, Restaurant, and Driver**.
* **Configuration:** Environment-based settings for different deployment stages.
* **Robustness:** Strict input validation and standardized error handling with proper HTTP status codes.

---

## 📚 Documentation

* **API Documentation:** [Swagger](https://swagger.io/) for detailed API specifications.
* **Code Documentation:** [GoDoc](https://godoc.org/) for inline code comments.

---
