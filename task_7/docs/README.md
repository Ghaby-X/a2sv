# Task Manager API

This document provides an overview of the Task Manager API and its available endpoints.

## Postman Collection

A Postman collection is available in this directory to test the API. You can import the `TaskManager.postman_collection.json` file into Postman to get started.

## API Endpoints

The following endpoints are available:

### Create Task

*   **Method:** POST
*   **Path:** `/tasks`
*   **Description:** Creates a new task.
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

### Get All Tasks

*   **Method:** GET
*   **Path:** `/tasks`
*   **Description:** Retrieves all tasks.
*   **Response:**
    ```json
    {
        "tasks": [
            {
                "id": 1,
                "title": "first task",
                "description": "sample task data",
                "status": "Pending",
                "created_at": "2025-11-13T21:38:46.703768581+02:00",
                "due_date": "2025-11-13T22:38:46.703768455+02:00"
            },
            {
                "id": 2,
                "title": "Second task",
                "description": "sample second task data",
                "status": "Approved",
                "created_at": "2025-11-13T21:38:46.703768807+02:00",
                "due_date": "2025-11-13T22:38:46.703768455+02:00"
            }
        ]
    }
    ```

### Get Task by ID

*   **Method:** GET
*   **Path:** `/tasks/:id`
*   **Description:** Retrieves a single task by its ID.
*   **Response:**
    ```json
    {
        "task": {
            "id": 2,
            "title": "Second task",
            "description": "sample second task data",
            "status": "Approved",
            "created_at": "2025-11-13T21:38:46.703768807+02:00",
            "due_date": "2025-11-13T22:38:46.703768455+02:00"
        }
    }
    ```

### Update Task

*   **Method:** PUT
*   **Path:** `/tasks/:id`
*   **Description:** Updates a task.
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

### Delete Task

*   **Method:** DELETE
*   **Path:** `/tasks/:id`
*   **Description:** Deletes a task.
*   **Response:**
    ```json
    {
        "msg": "task deleted successfully"
    }
    ```
