# 📧 Async Email Service (RabbitMQ + Golang)

An event-driven email service built using **Golang**, **RabbitMQ**, and **SMTP**.
This project demonstrates how to decouple email sending using a **producer-consumer architecture**.

---

## 🚀 Features

* 📤 REST API to queue emails
* 📨 RabbitMQ message broker
* ⚙️ Background email worker (consumer)
* 🔄 Asynchronous processing
* 🔐 Environment-based configuration
* 🐳 Docker support for RabbitMQ

---

## 🧠 Architecture Overview

```
        +-------------------+
        |   Client / Postman|
        +--------+----------+
                 |
                 | HTTP POST (/send-email)
                 v
        +-------------------+
        |   Producer (API)  |
        |  Golang Server    |
        +--------+----------+
                 |
                 | Publish Message
                 v
        +-------------------+
        |    RabbitMQ       |
        |   (Message Queue) |
        +--------+----------+
                 |
                 | Consume Message
                 v
        +-------------------+
        |   Consumer Worker |
        |   Email Sender    |
        +--------+----------+
                 |
                 | SMTP (Gmail)
                 v
        +-------------------+
        |   Email Delivered |
        +-------------------+
```

---

## 📁 Project Structure

```
.
├── producer/          # HTTP API to publish email jobs
├── consumer/          # Worker that processes emails
├── shared/            # Shared models (Email struct)
├── docker-compose.yml # RabbitMQ setup
└── .env               # Environment variables
```

---

## ⚙️ Setup Instructions

### 1️⃣ Clone the repo

```
git clone https://github.com/your-username/async-email-service.git
cd async-email-service
```

---

### 2️⃣ Start RabbitMQ using Docker

```
docker-compose up -d
```

RabbitMQ UI:
👉 http://localhost:15672
Default credentials:

```
username: guest
password: guest
```

---

### 3️⃣ Setup Environment Variables

Create `.env` file:

```
APP_EMAIL=your_email@gmail.com
APP_PASSWORD=your_app_password
```

⚠️ Use Gmail App Password (not your real password)

---

### 4️⃣ Run Consumer (Email Worker)

```
cd consumer
go run main.go
```

---

### 5️⃣ Run Producer (API Server)

```
cd producer
go run main.go
```

Server runs on:

```
http://localhost:8080
```

---

## 📬 API Usage

### Send Email

**POST** `/send-email`

### Request Body:

```json
{
  "to": "test@example.com",
  "subject": "Hello 🚀",
  "body": "<h1>This is a test email</h1>"
}
```

### Response:

```
Email queued successfully ✅
```

---

## 🔥 How It Works

1. Client sends request to API
2. API publishes message to RabbitMQ
3. Consumer listens to queue
4. Consumer sends email using SMTP
5. Fully async → fast + scalable

---

## 💡 Why This Project Matters

* Demonstrates **event-driven architecture**
* Solves **tight coupling problem**
* Mimics **real production systems**
* Great for **SDE interviews (LLD + Backend)**

---

## 🛠 Tech Stack

* Golang
* RabbitMQ
* Docker
* SMTP (Gmail)
* REST APIs

---

## 🚧 Future Improvements

* Retry mechanism (dead-letter queue)
* Email templates
* Rate limiting
* Logging + monitoring
* Kafka integration
* Microservices separation

---

## 👨‍💻 Author

Sahil Kumar
Full Stack / Backend Engineer 🚀

---

## ⭐ If you like this project

Give it a star ⭐ on GitHub!
