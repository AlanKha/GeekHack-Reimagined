# Post Handlers

This package contains handlers for managing posts within threads.

## Endpoints

### `POST /api/threads/:id/posts`

Creates a new post in a specific thread. Requires authentication.

- **URL Parameters**:
  - `id` (uint): The ID of the thread to post in.
- **Request Body**:

  ```json
  {
    "content": "string"
  }
  ```

- **Responses**:
  - `201 Created`: Post successfully created.
  - `400 Bad Request`: Invalid thread ID or request body.
  - `500 Internal Server Error`: Failed to create the post.

### `GET /api/posts/:id`

Retrieves a single post by its ID.

- **URL Parameters**:
  - `id` (uint): The ID of the post.
- **Responses**:
  - `200 OK`: Successfully returns the post.
  - `400 Bad Request`: Invalid post ID.
  - `404 Not Found`: Post with the specified ID does not exist.

### `PUT /api/posts/:id`

Updates a post's content. Requires authentication and ownership of the post.

- **URL Parameters**:
  - `id` (uint): The ID of the post.
- **Request Body**:

  ```json
  {
    "content": "string"
  }
  ```

- **Responses**:
  - `200 OK`: Post successfully updated.
  - `400 Bad Request`: Invalid post ID or request body.
  - `401 Unauthorized`: User does not own the post.
  - `404 Not Found`: Post with the specified ID does not exist.
  - `500 Internal Server Error`: Failed to update the post.

### `DELETE /api/posts/:id`

Deletes a post. Requires authentication and ownership of the post.

- **URL Parameters**:
  - `id` (uint): The ID of the post.
- **Responses**:
  - `200 OK`: Post successfully deleted.
  - `400 Bad Request`: Invalid post ID.
  - `401 Unauthorized`: User does not own the post.
  - `404 Not Found`: Post with the specified ID does not exist.
  - `500 Internal Server Error`: Failed to delete the post.
