# Category Handlers

This package manages the API endpoints for forum categories.

## Endpoints

### `POST /api/categories`

Creates a new category. Requires admin privileges.

- **Request Body**:

  ```json
  {
    "name": "string",
    "description": "string"
  }
  ```

- **Responses**:
  - `201 Created`: Category successfully created.
  - `400 Bad Request`: Invalid request body.
  - `500 Internal Server Error`: Failed to create the category.

### `GET /api/categories`

Retrieves a list of all categories.

- **Responses**:
  - `200 OK`: Successfully returns a list of categories.
  - `500 Internal Server Error`: Failed to retrieve categories.

### `GET /api/categories/:id`

Retrieves a single category by its ID.

- **URL Parameters**:
  - `id` (uint): The ID of the category.
- **Responses**:
  - `200 OK`: Successfully returns the category.
  - `400 Bad Request`: Invalid category ID.
  - `404 Not Found`: Category with the specified ID does not exist.

### `PUT /api/categories/:id`

Updates a category's details. Requires admin privileges.

- **URL Parameters**:
  - `id` (uint): The ID of the category.
- **Request Body**:

  ```json
  {
    "name": "string",
    "description": "string"
  }
  ```

- **Responses**:
  - `200 OK`: Category successfully updated.
  - `400 Bad Request`: Invalid request body or category ID.
  - `404 Not Found`: Category with the specified ID does not exist.
  - `500 Internal Server Error`: Failed to update the category.

### `DELETE /api/categories/:id`

Deletes a category. Requires admin privileges.

- **URL Parameters**:
  - `id` (uint): The ID of the category.
- **Responses**:
  - `200 OK`: Category successfully deleted.
  - `400 Bad Request`: Invalid category ID.
  - `404 Not Found`: Category with the specified ID does not exist.
  - `500 Internal Server Error`: Failed to delete the category.
