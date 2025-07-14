# 🚀 GeekHack, Reimagined

[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://github.com/your-username/geekhack-reimagined)
[![Go Version](https://img.shields.io/badge/go-1.18+-blue)](https://golang.org/)
[![Next.js Version](https://img.shields.io/badge/next.js-13+-blue)](https://nextjs.org/)

> A complete overhaul of a legacy forum, architected from the ground up as a modern, high-performance web application. This project uses a modular Golang backend and a server-rendered React frontend to deliver a fast, scalable, and developer-friendly experience.

## ✨ Key Features

- **🎨 Modern, Responsive UI**: Built with Tailwind CSS for a utility-first design that is fast and looks great on all devices
- **⚡ Optimized for Performance & SEO**: Leverages Next.js for Server-Side Rendering (SSR), ensuring fast initial page loads and excellent search engine visibility
- **🛡️ Robust Backend**: Engineered with a modular, RESTful API in Golang using the Gin framework for high-performance routing
- **🔒 Type-Safe Database Interaction**: Utilizes GORM to interact with a PostgreSQL database, providing type safety and preventing common SQL injection vulnerabilities
- **🔐 Secure Authentication**: Implements a secure user registration and login system using JWT (JSON Web Tokens)

## 🛠️ Tech Stack

| Area | Technology |
|------|------------|
| **Frontend** | React, Next.js, Tailwind CSS |
| **Backend** | Golang, Gin, GORM |
| **Database** | PostgreSQL |
| **Dev & Deploy** | Docker, Docker Compose |

## 📁 Project Structure

This project is structured as a monorepo to keep the frontend and backend code in a single, easy-to-manage repository.

```
geekhack-reimagined/
├── .github/              # GitHub Actions workflows
├── backend/              # Golang REST API
│   ├── cmd/              # Main application entry point
│   ├── internal/         # Private application logic (handlers, services, repos)
│   ├── pkg/              # Public library code
│   ├── Dockerfile        # Backend Docker image definition
│   └── go.mod            # Go module dependencies
├── frontend/             # Next.js Web Application
│   ├── app/              # Next.js App Router structure
│   ├── components/       # Reusable React components
│   ├── lib/              # Helper functions, API clients
│   ├── public/           # Static assets
│   └── Dockerfile        # Frontend Docker image definition
├── docker-compose.yml    # Local development environment setup
├── .env.example          # Example environment variables
└── README.md             # This file
```

## 🚀 Getting Started

Follow these instructions to get the project up and running on your local machine for development and testing purposes.

### Prerequisites

You must have the following software installed on your machine:

- **Go** (v1.18+)
- **Node.js** (v18+)
- **Docker & Docker Compose**

### Installation

1. **Clone the repository:**

```bash
git clone https://github.com/your-username/geekhack-reimagined.git
cd geekhack-reimagined
```

2. **Configure Environment Variables:**

Create a `.env` file in the project root by copying the example file. This file will be used by Docker Compose to configure all services.

```bash
cp .env.example .env
```

Now, open the `.env` file and fill in the values. The defaults are suitable for local development.

3. **Build and Run with Docker Compose:**

This single command will build the Docker images for the frontend and backend, start the PostgreSQL database container, and run the entire application.

```bash
docker-compose up --build
```

You should now be able to access:

- **Frontend Application**: http://localhost:3000
- **Backend API**: http://localhost:8080

## 📋 To-Do List & Project Roadmap

This is the development plan. Check off items as they are completed.

### Phase 0: Project Setup & Foundation

- [X] Initialize Git repository
- [X] Set up monorepo folder structure (frontend, backend)
- [X] Create initial docker-compose.yml for local development
- [X] Create .env.example with necessary variables (DB credentials, ports, JWT secret)
- [X] Write initial README.md file

### Phase 1: Backend API (Golang)

- [X] Initialize Go module in `/backend`
- [X] Integrate Gin for routing
- [X] Set up GORM and establish a database connection module
- [X] **Database Design**: Define GORM models for User, Thread, and Post
- [X] Implement GORM AutoMigrate to create tables on startup

#### User Authentication:
- [X] `/register` endpoint (hash password before saving)
- [X] `/login` endpoint (verify credentials, issue JWT)
- [X] JWT authentication middleware to protect routes

#### Thread Endpoints (CRUD):
- [X] `POST /api/threads` (Create, protected)
- [X] `GET /api/threads` (Read all)
- [X] `GET /api/threads/:id` (Read one)
- [X] `PUT /api/threads/:id` (Update, protected)
- [X] `DELETE /api/threads/:id` (Delete, protected)

#### Post Endpoints (CRUD):
- [X] `POST /api/threads/:id/posts` (Create, protected)
- [X] `GET /api/posts/:id` (Read one, likely not needed)
- [X] `PUT /api/posts/:id` (Update, protected)
- [X] `DELETE /api/posts/:id` (Delete, protected)

- [X] Implement robust error handling and standardized JSON responses
- [X] Set up basic unit tests for service logic

### Phase 2: Frontend UI (Next.js)

- [ ] Initialize Next.js project in `/frontend` with TypeScript and Tailwind CSS

#### Component Library:
- [ ] Build reusable UI components (Button, Input, Card, Navbar)
- [ ] **API Client**: Create a library in `/lib/api.ts` to handle all fetch requests to the backend

#### Static Pages & Layout:
- [ ] Create main application layout with Navbar and Footer
- [ ] Build Homepage (`/`) to fetch and display all threads
- [ ] Build `/login` and `/register` pages with forms

#### Dynamic Pages:
- [ ] Build Thread Detail page (`/thread/[id]`) to fetch and display a single thread and its posts

#### Authentication Flow:
- [ ] Create a React Context for managing authentication state
- [ ] On login, store JWT securely (HttpOnly cookie)
- [ ] Create protected routes/components that are only visible to logged-in users

#### User Interaction:
- [ ] Build form for creating new threads
- [ ] Build form for creating new posts/replies within a thread
- [ ] Implement logic for updating/deleting threads and posts

### Phase 3: Deployment

- [ ] Create a production-ready Dockerfile for the Go backend (multi-stage build)
- [ ] Create a production-ready Dockerfile for the Next.js frontend
- [ ] Provision a managed PostgreSQL instance (e.g., AWS RDS, Google Cloud SQL)
- [ ] Set up a container registry (e.g., Docker Hub, Google Artifact Registry)
- [ ] Write deployment scripts/configuration for a cloud provider (e.g., cloudbuild.yaml for Google Cloud Run)
- [ ] Deploy backend and frontend services
- [ ] Configure DNS and secure with SSL/TLS

### Phase 4: Future Features (Backlog)

- [ ] User Profiles
- [ ] Markdown support for posts with preview
- [ ] Pagination for threads and posts
- [ ] Real-time notifications (e.g., using WebSockets)
- [ ] Full-text search functionality
- [ ] Admin roles and moderation tools

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with ❤️ using modern web technologies
- Inspired by the original GeekHack forum community