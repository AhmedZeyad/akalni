Akilni Todo List
🗄️ Database

- [X] Run new schema migration (restaurants, categories, menu_items, orders, order_items)
- [X] Add admins table

📁 Project Structure

- [x] Create folders: models/, handlers/, repository/, middleware/
- [x] Separate routes: /api/client/, /api/admin/, /api/restaurant/

🔐 Auth & Middleware

- [x] Extend existing middleware to support admin role
- [ ] Add restaurant role middleware (for later)
- [ ] Export the finished document

📦 Models (Go structs)
- [X] users
- [X] restaurants
- [X] categories
- [X] products
- [X] Order
- [X] OrderItem

🛣️ Endpoints

- [ ] GET /api/client/restaurants — list active + open restaurants
- [ ] GET /api/client/restaurants/:id/menu — menu grouped by category
- [ ] POST /api/client/orders — create order (validate items, calc total)
- [ ] GET /api/client/orders/:id — get order (client owns it)
- [ ] PATCH /api/admin/orders/:id/status — update order status

🔧 Admin Endpoints (bonus)

- [ ] POST /api/admin/restaurants — create restaurant
- [ ] POST /api/admin/restaurants/:id/menu-items — add menu item
- [ ] PATCH /api/admin/restaurants/:id — update restaurant (activate/deactivate)

🧹 Cleanup

- [ ] Remove .env / hardcoded secrets
- [ ] Add .env.example
- [ ] Add README.md
