# 🚀 REST API in Go

A simple REST API built with Go and Gin, featuring JWT authentication, event management, and SQLite database support.

## 🌟 Features
- 🔑 User authentication (signup, login, JWT-based)
- 📅 CRUD operations for events
- 📝 User event registration and cancellation
- 🛡️ Authentication middleware
- 🗄️ SQLite database integration

## 🛠️ Tech Stack
- 🐹 **Go**
- ⚡ **Gin** (HTTP framework)
- 🔐 **JWT** (Authentication)
- 🔑 **Bcrypt** (Password hashing)
- 🗄️ **SQLite** (Database)

## 📦 Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/marcusvramos/go-rest-api.git
   cd your-repo
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Run the application:
   ```sh
   go run main.go
   ```

## 📡 API Endpoints

### 🔐 Auth
- `POST /signup` - 🆕 Register a new user
- `POST /login` - 🔑 Authenticate and get a JWT token

### 📅 Events
- `GET /events` - 📜 List all events
- `GET /events/:id` - 🔍 Get a specific event
- `POST /events` - ✏️ Create a new event (auth required)
- `PUT /events/:id` - 🛠️ Update an event (auth required)
- `DELETE /events/:id` - ❌ Delete an event (auth required)

### 📝 Event Registration
- `POST /events/:id/register` - ✅ Register for an event (auth required)
- `DELETE /events/:id/register` - 🔄 Cancel event registration (auth required)

