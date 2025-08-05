# Thread Handlers

This package provides API endpoints for managing discussion threads.

## Endpoints

### `POST /api/threads`

Creates a new thread. Requires authentication.

- **Request Body**:

  ```json
  {
    "title": "string",
    "content": "string",
    "category_id": "uint"
  }
  ```

- **Responses**:
  - `201 Created`: Thread successfully created.
  - `400 Bad Request`: Invalid request body.
  - `500 Internal Server Error`: Failed to create the thread.

### `GET /api/threads`

Retrieves a list of all threads. Can be filtered by category.

- **Query Parameters**:
  - `category_id` (uint, optional): Filter threads by category ID.
- **Responses**:
  - `200 OK`: Successfully returns a list of threads.
  - `500 Internal Server Error`: Failed to retrieve threads.

### `GET /api/threads/:id`

Retrieves a single thread by its ID.

- **URL Parameters**:
  - `id` (uint): The ID of the thread.
- **Responses**:
  - `200 OK`: Successfully returns the thread and its posts.
  - `400 Bad Request`: Invalid thread ID.
  - `404 Not Found`: Thread with the specified ID does not exist.

### `PUT /api/threads/:id`

Updates a thread's title and content. Requires authentication and ownership of the thread.

- **URL Parameters**:
  - `id` (uint): The ID of the thread.
- **Request Body**:

  ```json
  {
    "title": "string",
    "content": "string"
  }
  ```

- **Responses**:
  - `200 OK`: Thread successfully updated.
  - `400 Bad Request`: Invalid thread ID or request body.
  - `401 Unauthorized`: User does not own the thread.
  - `404 Not Found`: Thread with the specified ID does not exist.
  - `500 Internal Server Error`: Failed to update the thread.

### `DELETE /api/threads/:id`

Deletes a thread. Requires authentication and ownership of the thread.

- **URL Parameters**:
  - `id` (uint): The ID of the thread.
- **Responses**:
  - `200 OK`: Thread successfully deleted.
  - `400 Bad Request`: Invalid thread ID.
  - `401 Unauthorized`: User does not own the thread.
  - `404 Not Found`: Thread with the specified ID does not exist.
  - `500 Internal Server Error`: Failed to delete the thread.
