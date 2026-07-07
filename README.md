# Concurrent Cinema Seat Reservation Simulator

<p align="left">

<img src="https://img.shields.io/badge/Go-1.24-00ADD8?style=flat-square&logo=go&logoColor=white"/>
<img src="https://img.shields.io/badge/Redis-7-DC382D?style=flat-square&logo=redis&logoColor=white"/>
<img src="https://img.shields.io/badge/Docker-2496ED?style=flat-square&logo=docker&logoColor=white"/>
<img src="https://img.shields.io/badge/REST%20API-Backend-4CAF50?style=flat-square"/>
<img src="https://img.shields.io/badge/Concurrency-Thread%20Safe-7C3AED?style=flat-square"/>

</p>

A backend project built with Go that demonstrates atomic seat reservation, race condition prevention, and high-concurrency request handling using Redis. The project simulates how real-world ticket booking platforms safely process simultaneous reservation requests for the same seat.

---

# Application Demo

This demonstration shows the complete booking workflow, including concurrent seat reservation, temporary seat holding, booking confirmation, and automatic seat release.

> **Replace the link below with your uploaded GitHub video.**


https://github.com/user-attachments/assets/5cdc6c45-248d-4013-979b-7b40f5eebc8a



---

# Concurrency Stress Test

The booking service is validated using a concurrency stress test where **100,000 goroutines** simultaneously attempt to reserve the same seat.

Only one booking succeeds while all remaining requests are rejected, demonstrating thread-safe reservation logic.



https://github.com/user-attachments/assets/ab78f2cb-3f4d-4a75-b925-04e80958a5b9



# Features

- Concurrent seat reservation
- Temporary seat hold
- Booking confirmation
- Automatic hold expiration
- Seat release
- RESTful API
- Redis-backed storage
- Thread-safe booking logic
- Concurrency stress testing

---

# Problem Statement

In an online booking system, multiple users may attempt to reserve the same seat simultaneously.

Without proper synchronization, duplicate bookings may occur.

This project ensures that:

- Only one user can reserve a seat at a time.
- Remaining requests are safely rejected.
- Seats automatically become available again if the reservation expires.

---

# Booking Workflow

```text
                 Available
                     │
                     ▼
                Hold Seat
                     │
        ┌────────────┴────────────┐
        │                         │
        ▼                         ▼
Confirm Booking            Hold Expires
        │                         │
        ▼                         ▼
     Booked                 Available
```

---

# Concurrency Workflow

```text
          User A
             │
             │
             ▼
       Reserve Seat A1
             │
             ▼
        Seat Locked
             │
     ┌───────┴────────┐
     │                │
     ▼                ▼
 Success         User B Request
                      │
                      ▼
             Seat Already Held
```

The booking service guarantees that only one reservation succeeds for a seat under concurrent access.

---

# System Architecture

```text
                  Browser

                     │

                     ▼

              HTTP REST API

                     │

                     ▼

             Booking Service

                     │

                     ▼

                   Redis
```

---

# Project Structure

```text
.
├── cmd
│   └── main.go
│
├── internal
│   ├── adapters
│   │   └── redis
│   ├── booking
│   └── utils
│
├── static
│   └── index.html
│
├── docker-compose.yaml
├── go.mod
└── README.md
```

---

# Getting Started

## Clone the repository

```bash
git clone https://github.com/AnshuKashyap01/Concurrent-Cinema-Seat-Reservation-Simulator.git

cd Concurrent-Cinema-Seat-Reservation-Simulator
```

---

## Start Redis

```bash
docker compose up -d
```

---

## Run the application

```bash
go run ./cmd
```

---

## Access the application

Application

```
http://localhost:8080
```

Redis Commander

```
http://localhost:8081
```

---

# API Endpoints

| Method | Endpoint | Description |
|---------|----------|-------------|
| GET | `/movies` | List available movies |
| GET | `/movies/{movieID}/seats` | Retrieve seat availability |
| POST | `/movies/{movieID}/seats/{seatID}/hold` | Hold a seat |
| PUT | `/sessions/{sessionID}/confirm` | Confirm reservation |
| DELETE | `/sessions/{sessionID}` | Release reservation |

---

# Concurrency Testing

To verify the booking logic under heavy concurrent load, the project includes a stress test.

The test launches **100,000 concurrent goroutines**, each attempting to reserve the **same seat** simultaneously.

Expected outcome:

- 1 successful booking
- 99,999 rejected booking attempts
- No duplicate reservations

Run the test:

```bash
go test -v ./internal/booking
```

Example output:

```text
=== RUN   TestConcurrentBooking_ExactlyOneWins

Processed 100000 concurrent booking requests

Successful bookings : 1

Failed bookings     : 99999

--- PASS: TestConcurrentBooking_ExactlyOneWins (2.31s)

PASS
ok      github.com/AnshuKashyap01/Concurrent-Cinema-Seat-Reservation-Simulator/internal/booking
```

---

# Stress Test Scenario

```text
                 Seat A1

User 1  ───────────────┐
User 2  ───────────────┤
User 3  ───────────────┤
       .               │
       .               │
       .               │
User100000─────────────┘

            │
            ▼

      Booking Service

            │

            ▼

 Successful Booking : 1

 Failed Requests    : 99,999

 Duplicate Booking  : 0
```

---

# Technologies Used

- Go
- Redis
- Docker
- HTML
- CSS
- JavaScript

---

# Future Improvements

- WebSocket-based real-time seat updates
- Multiple cinema screens
- Multiple show timings
- User authentication
- Booking history
- Performance metrics
- Distributed deployment

---

# Author

**Anshu Kashyap**

GitHub: https://github.com/AnshuKashyap01

LinkedIn: *Add your LinkedIn profile*

---

If you found this project useful, consider giving the repository a star.
