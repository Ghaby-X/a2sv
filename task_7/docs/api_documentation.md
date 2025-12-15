# Task Management API

This document provides an overview of the Task Management API, including user authentication, authorization, and task management endpoints.

## Postman Collection

A Postman collection is available in this directory to test the API. You can import the `TaskManager.postman_collection.json` file into Postman to get started.

## Authentication and Authorization

The API uses JSON Web Tokens (JWT) for authentication and role-based access control for authorization.

### User Roles

*   **admin**: Has full access to all API endpoints, including creating, updating, deleting tasks, and promoting other users.
*   **user**: Can retrieve all tasks and retrieve tasks by their ID.

### Endpoints

#### Register User

*   **Method:** `POST`
*   **Path:** `/register`
*   **Description:** Creates a new user account. The first user registered will automatically be assigned the "admin" role. Subsequent users will be assigned the "user" role.
*   **Request Body:**
    ```json
    {
        "username": "newuser",
        "email": "newuser@example.com",
        "password": "securepassword"
    }
    ```
*   **Response:**
    ```json
    {
        "message": "User registered successfully"
    }
    ```

#### Login User

*   **Method:** `POST`
*   **Path:** `/login`
*   **Description:** Authenticates a user and returns a JWT token. This token must be included in the `Authorization` header of subsequent requests to protected routes.
*   **Request Body:**
    ```json
    {
        "username": "newuser",
        "password": "securepassword"
    }
    ```
*   **Response:**
    ```json
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
    ```

#### Promote User to Admin (Admin Only)

*   **Method:** `POST`
*   **Path:** `/users/promote/:id`
*   **Description:** Promotes a specified user to the "admin" role. Requires "admin" role.
*   **Authorization:** Bearer Token (Admin)
*   **Response:**
    ```json
    {
        "message": "User promoted to admin successfully"
    }
    ```

## API Endpoints

All task management endpoints, except for `GET /tasks` and `GET /tasks/:id`, require authentication. Endpoints that modify tasks (`POST`, `PUT`, `DELETE`) require the "admin" role.

### Get All Tasks

*   **Method:** `GET`
*   **Path:** `/tasks`
*   **Description:** Retrieves all tasks.
*   **Authorization:** Bearer Token (User or Admin)
*   **Response:**
    ```json
    {
        "tasks": [
            {
                "_id": "651a8a2a7f8b9c0d1e2f3a4b",
                "title": "first task",
                "description": "sample task data",
                "status": "pending",
                "created_at": "2025-11-13T21:38:46.703768581+02:00",
                "due_date": "2025-11-13T22:38:46.703768455+02:00"
            }
        ]
    }
    ```

### Get Task by ID

*   **Method:** `GET`
*   **Path:** `/tasks/:id`
*   **Description:** Retrieves a single task by its ID.
*   **Authorization:** Bearer Token (User or Admin)
*   **Response:**
    ```json
    {
        "task": {
            "_id": "651a8a2a7f8b9c0d1e2f3a4c",
            "title": "Second task",
            "description": "sample second task data",
            "status": "Approved",
            "created_at": "2025-11-13T21:38:46.703768807+02:00",
            "due_date": "2025-11-13T22:38:46.703768455+02:00"
        }
    }
    ```

### Create Task (Admin Only)

*   **Method:** `POST`
*   **Path:** `/tasks`
*   **Description:** Creates a new task. Requires "admin" role.
*   **Authorization:** Bearer Token (Admin)
*   **Request Body:**
    ```json
    {
        "title": "My New Task",
        "description": "This is a description of my new task.",
        "status": "pending",
        "due_date": "2024-05-05T00:00:00Z"
    }
    ```
*   **Response:**
    ```json
    {
        "msg": "task created successfully"
    }
    ```

### Update Task (Admin Only)

*   **Method:** `PUT`
*   **Path:** `/tasks/:id`
*   **Description:** Updates a task. Requires "admin" role.
*   **Authorization:** Bearer Token (Admin)
*   **Request Body:**
    ```json
    {
        "title": "Updated Title"
    }
    ```
*   **Response:**
    ```json
    {
        "msg": "task updated successfully"
    }
    ```

### Delete Task (Admin Only)

*   **Method:** `DELETE`
*   **Path:** `/tasks/:id`
*   **Description:** Deletes a task. Requires "admin" role.
*   **Authorization:** Bearer Token (Admin)
*   **Response:**
    ```json
    {
        "msg": "task deleted successfully"
    }
    ```