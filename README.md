# Employee Management System

A full-stack application for managing employee records. The backend is built with **Golang**, and the frontend uses **Angular**. This project provides CRUD operations to manage employee data and supports running the application via Docker or traditional methods.

---

## Features
- **View employees** in a paginated list.
- **Add new employees** via a form with validation.
- **Edit employee details** seamlessly.
- **Delete employees** with confirmation alerts.

---

## Project Structure
```
Employee Management System
├── Backend/         # Golang-based backend API
├── Frontend/        # Angular-based frontend application
├── docker-compose.yml # For running the app using Docker
```

---

## Installation

### Running Locally (Without Docker)
#### Backend
1. Ensure **Go** is installed on your system.
2. Navigate to the `Backend` directory:
   ```bash
   cd Backend
   ```
3. Install dependencies and run the server:
   ```bash
   go mod tidy
   go run main.go
   ```
4. The backend will be available at `http://localhost:8080`.

#### Frontend
1. Ensure **Node.js** and **Angular CLI** are installed.
2. Navigate to the `Frontend` directory:
   ```bash
   cd Frontend
   ```
3. Install dependencies and start the development server:
   ```bash
   npm install
   ng serve
   ```
4. The frontend will be available at `http://localhost:4200`.

---

### Running via Docker
1. Ensure **Docker** and **Docker Compose** are installed on your system.
2. From the project root directory, build and run the containers:
   ```bash
   docker-compose up --build
   ```
3. Access the application:
   - Backend: `http://localhost:8080`
   - Frontend: `http://localhost:4200`

---

## API Endpoints
The backend provides the following RESTful API endpoints:

- `GET /employees` : Fetch all employees.
- `GET /employees/{id}` : Retrieve an employee by ID.
- `POST /employees` : Add a new employee.
- `PUT /employees/{id}` : Update an existing employee.
- `DELETE /employees/{id}` : Remove an employee.

---

### Technologies Used
#### Backend:
- **Golang**
- **MongoDB**
- **Gin Framework**
#### Frontend:
- **Angular**
- **Bootstrap**