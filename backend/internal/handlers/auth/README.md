# Auth Handlers

This package manages user authentication, including registration, login, and session validation.

## Endpoints

### `POST /register`

Registers a new user in the system.

- **Request Body**:

  ```json
  {
    "username": "string",
    "email": "string",
    "password": "string"
  }
  ```

- **Responses**:
  - `200 OK`: User successfully created.
  - `400 Bad Request`: Invalid request body or username/email already exists.
  - `500 Internal Server Error`: Failed to hash the password.

### `POST /login`

Authenticates a user and returns a JWT.

- **Request Body**:

  ```json
  {
    "email": "string",
    "password": "string"
  }
  ```

- **Responses**:
  - `200 OK`: Login successful. Returns a JWT token and sets an `Authorization` cookie.
  - `400 Bad Request`: Invalid request body or incorrect email/password.
  - `500 Internal Server Error`: Failed to generate the JWT.

### `GET /validate`

Validates the current user's session using the JWT from the `Authorization` cookie.

- **Responses**:
  - `200 OK`: Session is valid. Returns user information.
  - `401 Unauthorized`: Invalid or expired token.
