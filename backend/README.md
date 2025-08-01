# GeekHack Backend

This is the backend for the GeekHack application, a forum platform built with Go and Gin.

## Project Structure

```text
backend/
├───cmd/
│   └───main.go         # Entry point of the application
├───internal/
│   ├───database/       # Database connection and migration
│   │   └───database.go
│   ├───handlers/       # HTTP handlers for API endpoints
│   │   ├───auth_handler.go
│   │   ├───post_handler.go
│   │   └───thread_handler.go
│   ├───middleware/     # Gin middleware
│   │   └───require_auth.go
│   ├───models/         # GORM models for database tables
│   │   ├───post.go
│   │   ├───thread.go
│   │   └───user.go
│   └───utils/          # Utility functions
│       └───error_response.go
├───go.mod
├───go.sum
└───Dockerfile
```

## API Endpoints

### Auth

- **Register**

  - **Method:** `POST`
  - **Path:** `/register`
  - **Auth:** None
  - **Body:**

    ```json
    {
      "username": "string",
      "email": "string",
      "password": "string"
    }
    ```

- **Login**

  - **Method:** `POST`
  - **Path:** `/login`
  - **Auth:** None
  - **Body:**

    ```json
    {
      "email": "string",
      "password": "string"
    }
    ```

- **Validate**
  - **Method:** `GET`
  - **Path:** `/validate`
  - **Auth:** Required (JWT in `Authorization` header or cookie)

### Threads

- **Create Thread**

  - **Method:** `POST`
  - **Path:** `/api/threads`
  - **Auth:** Required
  - **Body:**

    ```json
    {
      "title": "string",
      "content": "string"
    }
    ```

- **Get Threads**

  - **Method:** `GET`
  - **Path:** `/api/threads`
  - **Auth:** None

- **Get Thread**

  - **Method:** `GET`
  - **Path:** `/api/threads/:id`
  - **Auth:** None

- **Update Thread**

  - **Method:** `PUT`
  - **Path:** `/api/threads/:id`
  - **Auth:** Required
  - **Body:**

    ```json
    {
      "title": "string",
      "content": "string"
    }
    ```

- **Delete Thread**
  - **Method:** `DELETE`
  - **Path:** `/api/threads/:id`
  - **Auth:** Required

### Posts

- **Create Post**

  - **Method:** `POST`
  - **Path:** `/api/threads/:id/posts`
  - **Auth:** Required
  - **Body:**

    ```json
    {
      "content": "string"
    }
    ```

- **Get Post**

  - **Method:** `GET`
  - **Path:** `/api/posts/:id`
  - **Auth:** None

- **Update Post**

  - **Method:** `PUT`
  - **Path:** `/api/posts/:id`
  - **Auth:** Required
  - **Body:**

    ```json
    {
      "content": "string"
    }
    ```

- **Delete Post**
  - **Method:** `DELETE`
  - **Path:** `/api/posts/:id`
  - **Auth:** Required
