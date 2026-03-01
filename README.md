# 🧠 Todo-Api (Go + React + Docker)

Full-stack task manager built with Go, React, and Docker.
Demonstrates REST API design and layered backend architecture.

---

## 🚀 Features

* ✅ RESTful API (CRUD)
* 🐳 Dockerized backend
* ⚛️ React frontend (Vite)
* 🕒 Task timestamps (created / completed)
* 🎨 Simple UI
* 🌐 CORS support

---

## 📂 Project Structure

```
todo-api/
 ├─ cmd/api          # entry point
 ├─ internal
 │   ├─ app          # app wiring
 │   ├─ handler      # HTTP handlers
 │   ├─ service      # business logic
 │   ├─ storage      # data layer
 │   └─ model        # domain models
 ├─ web              # React frontend
 ├─ README.md
 ├─ Dockerfile
 └─ go.mod 
```

---

## ⚙️ Running Locally

### 1️⃣ Backend

```bash
go run ./cmd/api
```

Server runs on:

```
http://localhost:8080
```

---

### 2️⃣ Frontend

```bash
cd web
npm install
npm run dev
```

Frontend runs on:

```
http://localhost:5173
```

---

## 🐳 Run with Docker

```bash
docker build -t todo-api .
docker run -p 8080:8080 todo-api
```

---

## 📡 API Endpoints

| Method | Endpoint    | Description   |
| ------ | ----------- | ------------- |
| GET    | /tasks      | List tasks    |
| POST   | /tasks      | Create task   |
| PATCH  | /tasks/{id} | Complete task |
| DELETE | /tasks/{id} | Delete task   |

---

## 🕒 Task Model

```json
{
  "id": 1,
  "title": "Buy milk",
  "completed": true,
  "created_at": "2026-03-01T06:40:00Z",
  "completed_at": "2026-03-01T06:41:10Z"
}
```
