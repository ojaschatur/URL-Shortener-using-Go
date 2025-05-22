# 🔗 URL Shortener with Rate Limiting and PostgreSQL

This is a simple yet powerful URL shortener service built using **Go**, **PostgreSQL**, and **Docker**. It allows users to shorten long URLs, stores them in a database, and supports expiration and rate limiting based on IP.

---

## 📆 Features

- 🔐 Unique short URL generation
- ⏳ Expiry time for each shortened URL
- 🚫 Rate limiting per IP (default: 5 requests/minute)
- 📂 PostgreSQL integration for persistent storage
- 🐳 Dockerized for easy setup and deployment
- 📂 Clean code structure (modular design using packages)

---

## 🏗️ Tech Stack

- **Backend:** Go (Golang)
- **Database:** PostgreSQL
- **Runtime Environment:** Docker & Docker Compose
- **Packages Used:**
  - [`github.com/lib/pq`](https://github.com/lib/pq) – PostgreSQL driver for Go
  - [`github.com/joho/godotenv`](https://github.com/joho/godotenv) – Loads environment variables from `.env`

---

## 📁 Project Structure

```
.
├── db/                 # Database connection logic
│   └── connect.go
├── handlers/           # HTTP handlers
│   └── url.go
├── models/             # DB models and queries
│   └── url.go
├── utils/              # Utility functions
│   └── rate_limiter.go
├── main.go             # Application entry point
├── Dockerfile          # Multi-stage Docker build
├── docker-compose.yml  # Docker Compose setup
├── go.mod              # Go module definition
├── go.sum              # Go dependency checksums
└── README.md           # Documentation
```

---

## 🚀 Getting Started

### 1️⃣ Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

### 2️⃣ Clone the Repository

```bash
git clone https://github.com/your-username/url-shortener.git
cd url-shortener
```

---

### 3️⃣ Environment Variables

Create a `.env` file in the root directory (optional for local dev, values are already set in `docker-compose.yml`):

```env
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=urlshortener
```

---

### 4️⃣ Run with Docker

```bash
docker-compose up --build
```

This command:

- Builds the Go app using a multi-stage Dockerfile.
- Spins up a PostgreSQL container.
- Starts the app on `http://localhost:3000`.

---

---

### ▶️ Run Locally Without Docker (Using Go)

> Make sure you have **Go** and **PostgreSQL** installed locally.

#### 1️⃣ Setup Database

Create a PostgreSQL database named `urlshortener` and user `postgres` with password `password`, or change these in your `.env`.

Create the necessary tables:

```sql
-- URLs table
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_url VARCHAR(10) UNIQUE NOT NULL,
    creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expiry_date TIMESTAMP NOT NULL
);

-- Rate limiting table
CREATE TABLE rate_limits (
    ip VARCHAR(50) PRIMARY KEY,
    request_count INT NOT NULL,
    reset_time TIMESTAMP NOT NULL
);
```

Alternatively, execute the SQL manually (see schema above).

#### 2️⃣ Create `.env` File

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=urlshortener
```

#### 3️⃣ Run the Application

```bash
go run main.go
```

App will start on: [http://localhost:3000](http://localhost:3000)


## 🔌 API Endpoints

### 👈 `POST /shorten`

Shortens a URL and returns the shortened version.

#### Request Body

```json
{
  "url": "https://example.com/very/long/url",
  "expiry": 5
}
```

#### Response

```json
{
  "original_url": "https://example.com/very/long/url",
  "short_url": "http://localhost:3000/redirect/abc12345",
  "expiry": "2025-05-22T10:50:00Z",
  "rate_limit": 5,
  "rate_limit_reset": "2025-05-22T10:51:00Z"
}
```

---

### 🚫 Rate Limiting

- Each IP is allowed a maximum of **5 requests per minute**.
- If exceeded, a `429 Too Many Requests` error is returned.

---

### 🛠️ Database Schema

Create these tables manually in your PostgreSQL instance:

```sql
-- URLs table
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_url VARCHAR(10) UNIQUE NOT NULL,
    creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expiry_date TIMESTAMP NOT NULL
);

-- Rate limiting table
CREATE TABLE rate_limits (
    ip VARCHAR(50) PRIMARY KEY,
    request_count INT NOT NULL,
    reset_time TIMESTAMP NOT NULL
);
```

---

## 📷 Example Curl Request

```bash
curl -X POST http://localhost:3000/shorten \
-H "Content-Type: application/json" \
-d '{"url":"https://github.com/ojaschatur", "expiry": 2}'
```

---

## 🪣 To Do

- Add URL redirect handler (`/redirect/{code}`)
- Add unit tests
- Implement frontend UI
- Add authentication (optional)

---

## 👨‍💼 Author

- **Ojas Chatur** – [GitHub Profile](https://github.com/ojaschatur)

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
