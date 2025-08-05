# User Handlers

This package handles API endpoints related to user profiles.

## Endpoints

### `GET /api/users/:id`

Retrieves a user's public profile information.

- **URL Parameters**:
  - `id` (uint): The ID of the user.
- **Responses**:
  - `200 OK`: Successfully returns the user's profile.
  - `400 Bad Request`: Invalid user ID.
  - `404 Not Found`: User with the specified ID does not exist.

### `PUT /api/users/:id`

Updates a user's profile (e.g., avatar, signature). Requires authentication and ownership of the profile.

- **URL Parameters**:
  - `id` (uint): The ID of the user.
- **Request Body**:

  ```json
  {
    "avatar_url": "string",
    "signature": "string"
  }
  ```

- **Responses**:
  - `200 OK`: User profile successfully updated.
  - `400 Bad Request`: Invalid user ID or request body.
  - `401 Unauthorized`: User is not authorized to update this profile.
  - `404 Not Found`: User with the specified ID does not exist.
  - `500 Internal Server Error`: Failed to update the user profile.
