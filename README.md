# Golang Blog API

A production-ready RESTful Blog API built with **Go**, **Gin**, **JWT Authentication**, **MySQL**, and **GORM**.

---

## Features

✅ User Registration
✅ User Login (JWT Authentication)
✅ User Profile Management
✅ Create Blog Posts
✅ Update Blog Posts
✅ Delete Blog Posts
✅ Publish Posts (`draft → published`)
✅ Protected Routes with Middleware
✅ MySQL Database Integration
✅ UTF-8 Persian Language Support
✅ RESTful API Architecture

---

## Tech Stack

* **Go (Golang)**
* **Gin Framework**
* **GORM ORM**
* **MySQL**
* **JWT Authentication**
* **REST API**

---

## Project Structure

```txt
golang-blog-api/
│── cmd/
│   └── api/
│       └── main.go
│
│── config/
│   └── config.go
│
│── internal/
│   ├── database/
│   ├── handler/
│   ├── middleware/
│   ├── models/
│   └── routes/
│
│── pkg/
│   └── utils/
│
│── .env
│── go.mod
│── go.sum
│── README.md
```

---

## Installation

### 1. Clone Repository

```bash
git clone https://github.com/Mahdicip/golang-blog-api.git
cd golang-blog-api
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Configure Environment Variables

Create a `.env` file:

```env
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=blog_api

JWT_SECRET=your_secret_key
PORT=3000
```

---

## Run Project

```bash
go run ./cmd/api
```

Server starts at:

```txt
http://localhost:3000
```

---

## API Endpoints

### Authentication

#### Register User

```http
POST /api/v1/auth/register
```

Request Body:

```json
{
  "name": "mahdi",
  "email": "mahdi@test.com",
  "password": "123456"
}
```

---

#### Login User

```http
POST /api/v1/auth/login
```

Request Body:

```json
{
  "email": "mahdi@test.com",
  "password": "123456"
}
```

Response:

```json
{
  "token": "JWT_TOKEN"
}
```

---

### User Profile

#### Get Current User

```http
GET /api/v1/me
```

Headers:

```txt
Authorization: Bearer JWT_TOKEN
```

---

#### Update Profile

```http
PUT /api/v1/me
```

---

### Posts

#### Get All Posts

```http
GET /api/v1/posts
```

---

#### Get Single Post

```http
GET /api/v1/posts/:slug
```

---

#### Create Post

```http
POST /api/v1/posts
```

Headers:

```txt
Authorization: Bearer JWT_TOKEN
```

Body:

```json
{
  "title": "سلام دنیا",
  "content": "فارسی تست",
  "slug": "persian-post"
}
```

---

#### Update Post

```http
PUT /api/v1/posts/:id
```

---

#### Delete Post

```http
DELETE /api/v1/posts/:id
```

---

#### Publish Post

```http
PATCH /api/v1/posts/:id/publish
```

---

## Database

Database engine:

```txt
MySQL 8+
```

Encoding:

```txt
utf8mb4
```

Persian/Farsi language is fully supported.

---

## Example Response

```json
{
  "id": 1,
  "title": "سلام دنیا",
  "content": "فارسی",
  "slug": "persian-post",
  "status": "draft",
  "user_id": 1
}
```

---

## Future Improvements

* Pagination
* Search API
* Swagger Documentation
* Docker Support
* Role-Based Authorization
* File Upload
* Unit Testing

---

## Author

**Mahdi**

GitHub: https://github.com/Mahdicip
