# Event Management System

A professional RESTful API built with **Go** and the **Gin** web framework. This system allows users to create, manage, and register for events with robust authentication and authorization.

## 🚀 Features

- **User Authentication**: Secure signup and login using JWT (JSON Web Tokens) and bcrypt password hashing.
- **Event Management**: Full CRUD (Create, Read, Update, Delete) operations for events.
- **Ownership Authorization**: Only the creator of an event can update or delete it.
- **Event Registration**: Users can register for upcoming events and cancel their registrations.
- **Structured Architecture**: Clean separation of concerns with Models, Routes, Middlewares, and Database layers.

## 🛠️ Technology Stack

- **Language**: [Go (Golang)](https://golang.org/)
- **Web Framework**: [Gin Gonic](https://gin-gonic.com/)
- **Database**: [SQLite](https://www.sqlite.org/)
- **Authentication**: JWT-go
- **Security**: Bcrypt for password hashing

## 📂 Project Structure

```text
.
├── db/             # Database initialization and schema
├── middlewares/    # Custom middlewares (Auth)
├── models/         # Data models and business logic (User, Event, Registration)
├── routes/         # API handlers and route definitions
├── utils/          # Utility functions (JWT, Hashing)
├── Diagrams/       # Architectural diagrams (Class, Sequence, Control Flow)
├── main.go         # Application entry point
└── go.mod          # Dependency management
```

## 🏁 Getting Started

### Prerequisites

- Go 1.24+ installed on your machine.

### Installation

1. Clone the repository or navigate to the project directory:
   ```bash
   cd EventManagementSystem
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run .
   ```
   *The server will start at `http://localhost:8080`.*

## 🛣️ API Endpoints

### Public Routes
| Method | Endpoint | Description |
| :--- | :--- | :--- |
| POST | `/signup` | Create a new user account |
| POST | `/login` | Login and receive a JWT token |
| GET | `/events` | List all events |
| GET | `/events/:id` | Get details of a single event |

### Protected Routes (Requires JWT in `Authorization` header)
| Method | Endpoint | Description |
| :--- | :--- | :--- |
| POST | `/events` | Create a new event |
| PUT | `/events/:id` | Update an event (Owner only) |
| DELETE | `/events/:id` | Delete an event (Owner only) |
| POST | `/events/:id/register` | Register for an event |
| DELETE | `/events/:id/register` | Cancel event registration |

## 📊 Documentation

Architectural diagrams are available in the `Diagrams/` folder:
- **Class Diagram**: Visualizes model structures and relationships.
- **Sequence Diagram**: Detailed flow of event registration.
- **Control Flow**: Overall request lifecycle through the Gin engine.
- **UML Diagram**: Visualizes the system architecture.
- **Use Case Diagram**: Visualizes the system use cases.

