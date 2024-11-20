# Go - Library App
# PT. Orde Digital Intelektual - Backend Developer

Dokumen ini berisi penjelasan terkait studi kasus sistem perpustakaan yang digunakan dalam proses seleksi backend developer di PT. Orde Digital Intelektual. Studi kasus ini dirancang untuk menguji kemampuan peserta dalam mengembangkan solusi backend yang efisien, terstruktur, dan sesuai dengan kebutuhan fungsionalitas sistem perpustakaan.

## Features

### 1. Database Design
- A relational database schema tailored for a library system, including tables for books, users, loans, etc.
- ER diagram included in the `/docs` folder.

### 2. REST API Endpoints
- Implements at least **6 endpoints**, including:
  * User login
  * Add, update, delete, and list books
  * Book Borrowing Transaction
- Includes **authentication and authorization** for secured endpoints.

### 3. Configuration File
### 4. Built with Go
- Uses Golang for implementation with optional frameworks like : fiber

## Installation and Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/sidikimamsetiyawan/gp-project-library-app.git
    ```
2. Navigate to the Project Directory:
    ```bash
    cd go-project-library-app
    cd server
    ```
3. Install Dependencies:
   ```bash
    go mod tidy
    ```
5. Configure the Environment:
   Check for a configuration file like .env or config.json. If it exists. Open the file and set necessary values like database credentials, ports, etc.
   Example .env:
   ```bash
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_username
    DB_PASSWORD=your_password
    DB_NAME=library_app
    APP_PORT=8080
    ```
7. Set Up the Database (Optional)
8. Run the Project
   ```bash
    go run server.go
    ```
10. Test the Application
    ```bash
    http://localhost:8080
    ```

## Database Setup (MySQL)

This project requires a MySQL database. You can download the initial database file from this link: [MySQL Database File](https://drive.google.com/file/d/1Snp6ngPTu1Kcg5XDgg8Zzu7Dtm7bLD0u/view?usp=sharing).

### Steps to Set Up

1. **Create Database**: 
   - Create a MySQL database named `fiber_libraries_app`:
   ```sql
   CREATE DATABASE fiber_libraries_app;

## Testing with Postman

You can access the Postman collection for testing the API endpoints here: [Postman Collection](https://orange-trinity-586014.postman.co/workspace/PT.-Orde-Digital-Intelektual~34d83240-57f7-4e38-9fb4-8d8d91802b39/collection/9072736-8f6733dd-05f0-4b9f-917d-59de8deb7076?action=share&creator=9072736).

## API Documentation

The complete API documentation for this project is available online. You can access it at the link below:

[View API Documentation on Postman](https://documenter.getpostman.com/view/9072736/2sAYBRGuM9)

## Documentation
