# Project Overview for Gemini Agent

This document provides a concise overview of the GeekHack-Reimagined project, optimized for efficient understanding and operation by the Gemini agent.

## 1. Project Structure

The project is divided into two main components:

-   **`backend/`**: A Go-based API server.
-   **`frontend/`**: A Next.js (React) application.

## 2. Key Technologies

-   **Backend**: Go, Gin Gonic (web framework), GORM (ORM), PostgreSQL (via Supabase for production, SQLite for testing).
-   **Frontend**: Next.js, React, Tailwind CSS, ESLint, Prettier.

## 3. Test Execution Commands

### Backend (Go)

-   **Run all tests**: `go test ./...` (from `backend/` directory)
-   **Run tests with coverage**: `go test -cover ./...` (from `backend/` directory)
-   **Run specific test file**: `go test ./internal/middleware -cover` (example for middleware tests)

### Frontend (Next.js)

-   **Run linting/style checks**: `npm run lint` (from `frontend/` directory)
    -   *Note*: There is no dedicated `test` script in `package.json`. Linting serves as a primary quality gate.

## 4. Common Development Commands

-   **Install Frontend Dependencies**: `npm install` (from `frontend/` directory)
-   **Start Frontend Development Server**: `npm run dev` (from `frontend/` directory)
-   **Build Frontend**: `npm run build` (from `frontend/` directory)

## 5. Known Patterns and Considerations

-   **Go Type Assertions**: Be mindful of pointer vs. value type assertions (e.g., `user.(*models.User)` vs `user.(models.User)`). Panics related to `interface conversion` often indicate this.
-   **GORM Record Not Found**: When GORM's `First` or `Delete` methods don't find a record, they return an error and an empty struct (not `nil`). Assertions in tests should check for `err != nil` and `model.ID == 0` (or similar zero-value checks) rather than `model == nil`.
-   **Test Setup**: Use `tests.SetupTestDB(t)` for database-dependent tests in the backend. Remember to `defer teardown()` for proper cleanup.
-   **Environment Variables**: `JWT_SECRET` is crucial for authentication. Ensure it's set (e.g., via `os.Setenv` in tests or `.env` for local development).

## 6. Testing Strategy

-   **Unit Tests**: Located in `_test.go` files alongside the code they test (e.g., `handler_test.go` for `handler.go`).
-   **Coverage**: Aim to increase coverage for functional packages (`database`, `handlers`, `middleware`, `utils`). `cmd` (main entry), `models` (pure data structures), and `tests` (test utilities) typically have low or zero coverage and are not primary targets for unit testing.
-   **Adding New Tests**:
    1.  Create a `_test.go` file in the same package as the code to be tested.
    2.  Use `tests.SetupTestDB(t)` and `defer teardown()` for database interactions.
    3.  Use `httptest` and `gin.CreateTestContext` for testing HTTP handlers.
    4.  Use `assert` from `stretchr/testify` for clear assertions.