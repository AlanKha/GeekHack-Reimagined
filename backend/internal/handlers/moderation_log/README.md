# Moderation Log Handlers

This package provides endpoints for viewing and managing moderation logs, which record actions taken by moderators.

## Endpoints

### `GET /api/moderation-logs`

Retrieves a list of all moderation log entries. Requires admin privileges.

- **Responses**:
  - `200 OK`: Successfully returns a list of moderation logs.
  - `500 Internal Server Error`: Failed to retrieve moderation logs.

### `POST /api/moderation-logs`

Creates a new moderation log entry. This is typically called internally by other handlers when a moderation action is performed.

- **Request Body**:

  ```json
  {
    "user_id": "uint",
    "action": "string",
    "reason": "string",
    "thread_id": "uint",
    "post_id": "uint"
  }
  ```

- **Responses**:
  - `201 Created`: Moderation log successfully created.
  - `400 Bad Request`: Invalid request body.
  - `500 Internal Server Error`: Failed to create the moderation log.
