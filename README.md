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
2. Install dependencies:
    ```bash
    npm install
    ```

3. Run the application:
    ```bash
    npm start
    ```

## Database Setup (MySQL)

This project requires a MySQL database. You can download the initial database file from this link: [MySQL Database File](https://drive.google.com/file/d/1F5AvJBU_wA1YPcZsLUq_Nm8QVpQJJMnM/view?usp=sharing).

### Steps to Set Up

1. **Create Database**: 
   - Create a MySQL database named `fiber_libraries_app`:
   ```sql
   CREATE DATABASE fiber_libraries_app;

## Testing with Postman

You can access the Postman collection for testing the API endpoints here: [Postman Collection](https://orange-trinity-586014.postman.co/workspace/ba2fd21e-faab-475c-ba96-402a4b6ca449/folder/9072736-c9c8d5df-ae0b-4d8f-8112-f1742058a7f9).

Alternatively, you can download the collection file and import it into Postman using the "Import" feature. [Download the Postman collection](https://drive.google.com/file/d/1xtSbEKyLTYlTULRE1VxPeiIc-PaMiaO9/view?usp=sharing).

## API Documentation

The complete API documentation for this project is available online. You can access it at the link below:

[View API Documentation on Postman](https://documenter.getpostman.com/view/9072736/2sAYBRGuM9)

## Documentation
