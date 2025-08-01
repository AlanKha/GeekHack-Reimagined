# Testing Guide

This document provides instructions on how to run tests for the GeekHack-Reimagined project, covering both the backend (Go) and frontend (Next.js) components.

## 1. Backend Testing (Go)

The backend tests are written in Go and use the `testing` package along with `stretchr/testify` for assertions.

### Prerequisites

- Ensure you have Go installed (version 1.24.5 or higher is recommended).
- Ensure you have Node.js and npm (or yarn) installed.

### Running Tests

Navigate to the `backend/` directory in your terminal:

```bash
cd backend/
```

To run all tests:

```bash
go test ./...
```

To run tests and view code coverage:

```bash
go test -cover ./...
```

This command will output the coverage percentage for each package. For a more detailed HTML report, you can run:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

This will open a web browser displaying a detailed coverage report, highlighting covered and uncovered lines of code.

### Test Structure

- Test files are located in the same directory as the source code they test, with a `_test.go` suffix (e.g., `handler.go` has `handler_test.go`).
- Database-dependent tests utilize `internal/tests/test_utils.go` for setting up an in-memory SQLite database and tearing it down after tests complete.

## 2. Frontend Testing (Next.js / React)

The frontend uses ESLint for linting and code style checks. Currently, there are no dedicated unit or integration tests configured beyond linting.

### Installing Dependencies

Before running any frontend commands, install the necessary Node.js packages. Navigate to the `frontend/` directory:

```bash
cd frontend/
npm install
```

### Running Linting Tests

To run the linting checks and automatically fix some issues:

```bash
npm run lint
```

This command will report any linting errors or warnings and attempt to fix them if possible.

### Future Testing (Recommendations)

For more comprehensive frontend testing, consider implementing:

- **Unit Tests**: Using libraries like Jest and React Testing Library for individual component testing.
- **Integration Tests**: To test the interaction between multiple components or with API endpoints.
- **End-to-End (E2E) Tests**: Using tools like Cypress or Playwright to simulate user interactions across the entire application.
