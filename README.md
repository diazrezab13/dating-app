# Dating App Backend with Gin Gonic and PostgreSQL

This is a backend service for a simple dating app built with the Gin Gonic framework and PostgreSQL as the database. The app provides functionalities such as user registration, login, swiping on dating profiles, and purchasing premium packages.

## Features

- **User Authentication**: Sign up and login functionality for users.
- **Daily Swipe Limit**: Users can swipe on a maximum of 10 profiles (like or pass) per day.
- **Premium Features**:
  - Users can purchase premium packages to unlock features such as unlimited swipes.
- **Persistent Data**: User data and swipe records are stored in a PostgreSQL database.

## Technologies Used

- **Golang**: The programming language used for building the backend service.
- **Gin Gonic**: A web framework for building APIs in Go.
- **PostgreSQL**: The database for storing user and swipe data.
- **GORM**: An ORM library for Golang used for database operations.
- **Docker & Docker Compose**: For containerization and easy setup.

## Installation and Setup

### Prerequisites

- Docker and Docker Compose installed on your system.
- Go (Golang) installed on your system.

### Setup Instructions

1. **Clone the Repository**:
   git clone https://github.com/yourusername/your-dating-app-repo.git
   cd your-dating-app-repo

2. **Create a .env File: Create a .env file in the root directory with the following content**:

DB_HOST=localhost
DB_PORT=5432
DB_USER=yourusername
DB_PASSWORD=yourpassword
DB_NAME=yourdatabasename
JWT_SECRET=yourjwtsecret

3. **Build and Run the App Using Docker Compose**:

docker-compose up --build -d

4. **Access the API**: 

The API will be accessible at http://localhost:8080.

# API Endpoints
- **Auth Endpoints**
POST /signup: Create a new user
POST /login: Log in and get a JWT token

- **Profile Endpoints**
GET /profile: Get List of profile to swipe
POST /profile/swipe: Swipe left (pass) or swipe right (like) on a profile
POST /profile/upgrade: Upgrade user account to premium

# Running Tests 
**To run unit tests for the controllers**:

go test ./... -v

**To check code coverage**:

go test ./... -cover