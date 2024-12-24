# Event System Backend

Event System Backend is a RESTful API built with Golang and Gin framework to manage events, users, vendors, and authentication for an event management platform.

## Table of Contents

- [About](#about)
- [Features](#features)
- [Technologies](#technologies)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Database Migration](#database-migration)
  - [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)

## About

The Event System Backend is designed to provide a reliable and scalable solution for managing events, vendors, and users. The backend handles user authentication, event creation, list events and status updates.

## Features

- User authentication (JWT with Bearer token)
- Event management (create, detail and list events)
- Vendor management (view all vendors, filter by name)
- Event approval system (approve or reject events)
- Pagination and filtering of event listings
- Secure API with Bearer token authentication

## Technologies

- **Golang** - Programming language
- **Gin** - Web framework
- **JWT** - Authentication with JSON Web Tokens
- **Bcrypt** - Encryption
- **GORM** - ORM for database interaction
- **MySQL** - Database
- **Swagger** - API documentation and testing

## Getting Started

### Prerequisites

Before begin, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (Version 1.18+)
- [MySQL](https://dev.mysql.com/downloads/) (Version 5.7 or higher recommended)
- [Swag](https://github.com/swaggo/swag) (for generating Swagger documentation)
- [Migrate](https://github.com/golang-migrate/migrate) (for database migrations)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/vendor-event-management/event-system-backend.git
   cd event-system-backend
   ```

2. Install Go dependencies:

   ```bash
   go mod tidy
   ```

3. Set up the MySQL database:

   - Create a MySQL database.

4. Create a `.env` file in the root directory and add necessary environment variables. Example:

   ```env
   DB_DSN=mysql://username:password@tcp(localhost:3306)/dbname?charset=utf8&parseTime=True&loc=Local
   JWT_SECRET=your_jwt_secret
   PORT=5000
   ENVIRONMENT=local
   JWT_SECRET_KEY=123
   CORS_ALLOW_ORIGINS=http://example.com,http://localhost:3000
   CORS_ALLOW_METHOD=GET,POST,PUT,DELETE
   ```

### Database Migration

To handle database migrations, we use the [migrate](https://github.com/golang-migrate/migrate) tool. Follow these steps to create and run migrations:

1. **Install the migrate tool**:

   ```bash
   go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```

2. **Create a migration file**:

   To create a new migration file, run the following command:

   ```bash
   migrate create -ext sql -dir migrations create_table_events
   ```

   - This command generates two SQL files: one for applying the migration (`.up.sql`) and one for rolling back the migration (`.down.sql`).

3. **Write migration scripts**:

   After running the above command, two files will be created under the `migrations` directory:

   - `migrations/xxxxxx_create_table_events.up.sql` - For creating the events table.
   - `migrations/xxxxxx_create_table_events.down.sql` - For rolling back the events table creation.

   Example migration `up.sql` to create the events table:

   ```sql
   CREATE TABLE events (
       id VARCHAR(255) AUTO_INCREMENT PRIMARY KEY,
       name VARCHAR(255) NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
   );
   ```

   Example migration `down.sql` to drop the events table:

   ```sql
   DROP TABLE IF EXISTS events;
   ```

4. **Run the migration**:

   To apply the migration to your MySQL database, run:

   ```bash
   migrate -path migrations -database "mysql://username:password@tcp(localhost:3306)/dbname" up
   ```

   Replace `username`, `password`, and `dbname` with your MySQL credentials.

5. **Running Migrations Automatically**:

   When you run the application with the following command:

   ```bash
   go run cmd/app/main.go
   ```

   The migration will be automatically applied if there are any pending migrations. This ensures that your database schema is always up-to-date with the latest changes without needing to manually run the migration tool.

6. **Rollback a migration**:

   To roll back the last migration, run:

   ```bash
   migrate -path migrations -database "mysql://username:password@tcp(localhost:3306)/dbname" down
   ```

### Running the Application

Once the application is installed, you can start the server using the following command:

```bash
go run main.go
```

### API Documentation

The API is documented using Swagger. To access the Swagger UI:

1. Run the application as described above.
2. Open your browser and go to: http://localhost:5000/swagger/index.html#/.
