# 🖥️ NearbyBite – Backend Development Plan
**Tech Stack:** Golang (Gin), PostgreSQL, PostGIS  
**Role:** Backend Developer  
**API Style:** REST  
**Auth:** JWT (Access Token)

---

## 🟦 Phase 0: Project Setup & Foundations

### ✅ Initial Setup
- [x] Initialize Go module
- [x] Setup Gin framework
- [x] Setup project folder structure
  - [x] `/cmd`
  - [x] `/config`
  - [x] `/database`
  - [x] `/models`
  - [x] `/handlers`
  - [x] `/services`
  - [x] `/middlewares`
  - [x] `/routes`
  - [x] `/utils`
- [x] Environment variables support
- [x] Logging setup
- [x] Graceful server shutdown

---

## 🟦 Phase 1: Database & PostGIS

### 🗄 Database Setup
- [x] Setup PostgreSQL connection
- [x] Enable PostGIS extension
- [SKIP] Connection pooling
- [SKIP] Migration tool setup (golang-migrate / goose)

### 📐 Schema Design
- [x] Users table
- [SKIP] Restaurants table with `GEOGRAPHY(Point)`
- [SKIP] Menu items table
- [SKIP] Orders table
- [SKIP] Order items table
- [SKIP] Indexes for performance

---

## 🟦 Phase 2: Authentication & Authorization

### 🔐 Auth Core
- [x] Password hashing (bcrypt)
- [x] JWT token generation
- [x] JWT token validation
- [] Token expiration handling

### 📡 Auth Endpoints
- [x] `POST /api/v1/auth/register`
- [x] `POST /api/v1/auth/login`
- [ ] `GET /api/v1/auth/me`
- [ ] `POST /api/v1/auth/logout` (optional)

### 🧠 Middleware
- [ ] Auth middleware (Bearer token)
- [ ] Protect private routes
- [ ] Inject user context into request

### ❌ Error Handling
- [ ] Standard auth error responses
- [ ] Invalid token handling
- [ ] Unauthorized access handling

---

## 🟦 Phase 3: Restaurant Management

### 🏪 Restaurant APIs
- [ ] `GET /api/v1/restaurants`
- [ ] `GET /api/v1/restaurants/:id`
- [ ] `POST /api/v1/restaurants` (admin)
- [ ] `PUT /api/v1/restaurants/:id` (admin)
- [ ] `DELETE /api/v1/restaurants/:id` (admin)

### 📍 Geospatial Logic
- [ ] Accept latitude & longitude query params
- [ ] Use `ST_DWithin` for radius search
- [ ] Sort restaurants by distance
- [ ] Return distance in response

---

## 🟦 Phase 4: Menu Management

### 🍽 Menu APIs
- [ ] `GET /api/v1/restaurants/:id/menu`
- [ ] `POST /api/v1/menu-items` (admin)
- [ ] `PUT /api/v1/menu-items/:id` (admin)
- [ ] `DELETE /api/v1/menu-items/:id` (admin)

### 🧩 Menu Logic
- [ ] Link menu items to restaurant
- [ ] Category support
- [ ] Price validation
- [ ] Availability flag

---

## 🟦 Phase 5: Order System

### 🛒 Order APIs
- [ ] `POST /api/v1/orders`
- [ ] `GET /api/v1/orders/:id`
- [ ] `GET /api/v1/orders/my`

### 📦 Order Logic
- [ ] Validate authenticated user
- [ ] Validate restaurant exists
- [ ] Validate menu items
- [ ] Enforce single-restaurant rule
- [ ] Calculate total price
- [ ] Store order & order items

### 🔄 Order Status
- [ ] Pending
- [ ] Accepted
- [ ] Preparing
- [ ] Delivered
- [ ] Cancelled

---

## 🟦 Phase 6: Validation & Error Handling

### 🛡 Validation
- [ ] Request body validation
- [ ] Query parameter validation
- [ ] Consistent validation errors

### ⚠️ Error Handling
- [ ] Global error handler
- [ ] Standard error response format
- [ ] 400 / 401 / 403 / 404 / 500 handling

---

## 🟦 Phase 7: Performance & Security

### ⚡ Performance
- [ ] Query optimization
- [ ] Index tuning
- [ ] Pagination for list endpoints
- [ ] Avoid N+1 queries

### 🔒 Security
- [ ] CORS configuration
- [ ] Rate limiting
- [ ] SQL injection prevention
- [ ] Secure headers

---

## 🟦 Phase 8: Testing & Documentation

### 🧪 Testing
- [ ] Unit tests for services
- [ ] Integration tests for handlers
- [ ] Mock database tests

### 📄 Documentation
- [ ] API documentation (Swagger / OpenAPI)
- [ ] Example requests & responses
- [ ] Auth flow documentation

---

## 🟦 Phase 9: Deployment Readiness

- [ ] Environment-based config
- [ ] Production build
- [ ] Dockerfile
- [ ] Health check endpoint
- [ ] Logging & monitoring readiness

---

## ⭐ Optional / Future Enhancements
- [ ] Refresh tokens
- [ ] Role-based access (Admin / User)
- [ ] WebSocket order tracking
- [ ] Payment gateway integration
