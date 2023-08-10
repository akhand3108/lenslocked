# LensLocked - Web Gallery Application

![Project Status](https://img.shields.io/badge/status-active-brightgreen.svg)
![GitHub License](https://img.shields.io/badge/license-MIT-blue.svg)

LensLocked is a web application that allows users to create and manage their own gallery pages, upload photos, and share the gallery pages with others. It includes session management for user authentication and password reset functionality with email notifications.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Setup and Installation](#setup-and-installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Database](#database)
- [Contributing](#contributing)
- [License](#license)

## Features

- User Registration and Authentication
- Gallery Creation and Management
- Photo Upload to Galleries
- CSRF Protection
- Forgot Password and Password Reset (with Email Notifications)
- Sharing Gallery Pages
- Static Pages (Home, Contact, FAQ)
- Error Handling and Not Found Page

## Technologies Used

- Go (Golang)
- Chi Router
- Gorilla CSRF
- PostgreSQL
- HTML Templates
- SMTP for Email Notifications
- Godotenv for Configuration
- Tailwind CSS (for Styling)

## Setup and Installation

1. Clone the repository: `git clone https://github.com/theakhandpatel/lenslocked.git`
2. Install Go and PostgreSQL.
3. Set up your environment variables using the `.env` file (copy `.env.example`).
4. Run migrations: `go run main.go migrate`
5. Install dependencies: `go mod download`

## Usage

1. Start the server: `go run main.go`
2. Access the application in your browser at `http://localhost:8080`

## Endpoints

- `/` - Home page
- `/contact` - Contact page
- `/faq` - FAQ page
- `/signup` - User registration
- `/signin` - User login
- `/signout` - User logout
- `/forgot-pw` - Forgot password (with email verification)
- `/reset-pw` - Reset password
- `/galleries` - List all galleries
- `/galleries/new` - Create a new gallery
- `/galleries/{id}` - View a gallery
- `/galleries/{id}/edit` - Edit a gallery
- `/galleries/{id}/delete` - Delete a gallery
- `/galleries/{id}/images` - Upload images to a gallery

## Session Management

LensLocked uses session management to handle user authentication and authorization for various actions.

## Password Reset and Email Notifications

The application provides a password reset mechanism. In case a user forgets their password, they can request a password reset. An email containing a password reset link will be sent to the user's registered email address.

## Database

The application uses PostgreSQL as the database. Database settings can be configured in the `.env` file.

## Contributing

Contributions are welcome! Feel free to open issues and submit pull requests.

## License

This project is licensed under the [MIT License](LICENSE).
