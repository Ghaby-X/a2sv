# Task Manager API

This is a simple Task Manager API built with Go, Gin, and MongoDB, following Clean Architecture principles. It allows you to create, read, update, and delete tasks, with user authentication and role-based authorization.

## Architecture

The application is structured using Clean Architecture, separating concerns into the following layers:

-   **Domain**: Contains the core business entities (Task, User).
-   **Usecases**: Contains the application-specific business rules.
-   **Repositories**: Abstracts the data access logic.
-   **Infrastructure**: Implements external dependencies and services (e.g., JWT, password hashing).
-   **Delivery**: Contains files related to the delivery layer, handling incoming requests and responses (e.g., main.go, controllers, routers).

## Features

*   User registration and login with JWT authentication.
*   Role-based access control (admin, user).
*   Create, read, update, and delete tasks.
*   Promote users to admin.

## API Documentation

For detailed information about the API endpoints, authentication, and authorization, please see the [API Documentation](docs/api_documentation.md).

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/Ghaby-X/task_manager.git
cd task_manager
```

### 2. Install Go

Make sure you have Go installed on your system. You can download it from the official website: [https://golang.org/](https://golang.org/)

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Set up MongoDB

You need a running MongoDB instance. You can either install it locally or use a cloud service like MongoDB Atlas.

**Using Docker:**

If you have Docker installed, you can easily start a MongoDB container:

```bash
docker run -d -p 27017:27017 --name mongodb mongo
```

This will start a MongoDB instance on `localhost:27017`.

### 5. Create a `.env` File

Create a `.env` file in the root of the project and add your MongoDB connection URI and a JWT secret:

```
MONGO_URI=mongodb://localhost:27017
DB_NAME=task_manager
JWT_SECRET=your_jwt_secret
```

## How to Run

To run the application, use the following command from the project root:

```bash
go run Delivery/main.go
```

The server will start on the default Gin port (`:8080`).

## Environment Variables

*   `MONGO_URI`: The connection URI for your MongoDB instance.
*   `DB_NAME`: The name of the database.
*   `JWT_SECRET`: A secret key for signing JWT tokens.