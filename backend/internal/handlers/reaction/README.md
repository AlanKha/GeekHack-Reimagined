# Reaction Handlers

This package handles API endpoints for post reactions.

## Endpoints

### `POST /api/posts/:id/reactions`

Adds a reaction to a specific post. Requires authentication.

- **URL Parameters**:
  - `id` (uint): The ID of the post to react to.
- **Request Body**:

  ```json
  {
    "reaction_type": "string" // e.g., "like", "love", "angry"
  }
  ```

- **Responses**:
  - `201 Created`: Reaction successfully created.
  - `400 Bad Request`: Invalid post ID or request body.
  - `500 Internal Server Error`: Failed to create the reaction (e.g., user already reacted).

### `GET /api/posts/:id/reactions`

Retrieves all reactions for a specific post.

- **URL Parameters**:
  - `id` (uint): The ID of the post.
- **Responses**:
  - `200 OK`: Successfully returns a list of reactions.
  - `400 Bad Request`: Invalid post ID.
  - `500 Internal Server Error`: Failed to retrieve reactions.
