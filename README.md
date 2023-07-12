# Goqr - QRCode Generator

Goqr is a web application built with Svelte and Go that provides QRCode generation functionality. It allows users to create, update, and delete QRCodes. The application uses Fiber as the web server and PostgreSQL as the database.

## Features

- Create QRCodes to represent URLs.
- Update existing QRCodes to redirect to new URLs.
- Delete QRCodes.
- Easy deployment with Docker for the database.

## Technologies Used

- Svelte: A modern JavaScript framework for building user interfaces.
- Go: A statically typed, compiled programming language.
- Fiber: A fast, flexible, and easy-to-use web framework for Go.
- PostgreSQL: A powerful, open-source relational database management system.
- Docker: A platform that allows you to automate the deployment and scaling of applications using containerization.
- Yarn: A fast, reliable, and secure package manager for JavaScript.

## Prerequisites

To run Goqr locally, ensure that you have the following installed:

- Docker: [Installation Guide](https://docs.docker.com/get-docker/)
- Yarn: [Installation Guide](https://yarnpkg.com/getting-started/install)
- Go: [Installation Guide](https://golang.org/doc/install)
- PostgreSQL: [Installation Guide](https://www.postgresql.org/download/)

## Getting Started locally

Follow the steps below to get Goqr up and running on your local machine:

1. Clone the repository:

   ```bash
   git clone https://github.com/RNAV2019/goqr.git
   ```

2. Change into the project directory.

   ```bash
   cd goqr
   ```

3. Start the PostgreSQL database using Docker:

   ```bash
   docker run --name goqr -d -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=test postgres
   ```

4. Install the dependencies for Goqr using Yarn:

   ```bash
   cd frontend
   yarn install
   ```

5. Start the Svelte development server (hosted on http://localhost:5173):

   ```bash
   yarn dev
   ```

6. Change into the backend directory:

   ```bash
   cd ../backend
   ```

7. Change directory to the model:

   ```bash
   cd model
   ```
8. Edit model.go:
   ```go
   // Uncomment the getenv dsn string
   dsn := os.Getenv("DSN")

   // And comment out the other dsn string
   // dsn := "postgres://admin:test@" + os.Getenv("DB_HOST") + ":5432/admin?sslmode=disable"

   ```
9. Change directory back to the backend folder
   ```bash
   cd ..
   ```

8. Start the Go backend server:
   
   ```bash
   go run main.go
   ```

Ensure you create your own .env file with your connection string.
The API endpoints will be available at http://localhost:3000.

## Getting Started using Docker
1. Clone the repository:

   ```bash
   git clone https://github.com/RNAV2019/goqr.git
   ```

2. Change into the project directory.

   ```bash
   cd goqr
   ```
3. Build the docker-compose file (make sure docker daemon is running):

   ```bash
   docker-compose build
   ```
4. Run the docker-compose file in detached mode:

   ```bash
   docker-compose up -d
   ```

The Svelte Website will be available at http://localhost:8080.
The API endpoints will be available at http://localhost:3000.

## Configuration

Goqr uses environment variables for configuration. By default, the application uses the following values:

### Frontend
- `PUBLIC_DB=http://localhost:3000/`
- `PUBLIC_WS=ws://localhost:3000/ws`

### Backend
- `DSN=postgres://user:password@domain:5432/database?sslmode=disable`

You can modify these values using the example in the `.env.example` file located in the corresponding directory.