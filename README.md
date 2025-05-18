# HMS Backend - Healthcare Management System

A comprehensive Healthcare Management System built with Go, designed to manage hospital operations including patient records, appointments, prescriptions, and medical records.

## Overview

HMS Backend is a robust healthcare management platform built with a clean architecture pattern. The system provides a secure API for healthcare operations with role-based access control for administrators, doctors, and patients.

**API Documentation**: [APIdog Documentation](https://apidog.com/apidoc/shared-ef4c599e-8373-4716-ba6f-6e3ab9a18a01)

## Features

- **User Management**: Registration and authentication with JWT tokens
- **Role-Based Access Control**: Different permissions for admin, doctor, and patient roles
- **Patient Management**: Store and manage patient details
- **Doctor Management**: Maintain doctor profiles with specialties
- **Appointment Scheduling**: Create, update, and manage appointments
- **Medical Records**: Document patient diagnoses and treatments
- **Prescriptions**: Issue and track medication prescriptions
- **Audit Logging**: Track system activities for security and compliance

## Tech Stack

- **Backend**: Go (Golang)
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **Containerization**: Docker and Docker Compose

## Project Structure

The project follows a clean architecture approach with the following main components:

- **API**: Contains HTTP handlers, routes, and middleware
- **Bootstrap**: Application initialization and configuration
- **Domain**: Core business entities and interfaces
- **Repository**: Database interaction implementations
- **Usecase**: Business logic implementation
- **Internal**: Utilities and helper functions

## Prerequisites

- Go 1.24+
- PostgreSQL
- Docker and Docker Compose (for containerized deployment)

## Database Setup

Before running the application, you need to set up the PostgreSQL database. Use the following SQL queries to create the required tables:

```sql
CREATE TYPE user_role AS ENUM ('admin', 'doctor', 'patient');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role user_role NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE patients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE REFERENCES users(id),
    cpf VARCHAR(14) UNIQUE NOT NULL,
    date_birth DATE NOT NULL,
    phone VARCHAR(20),
    address TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE doctors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE REFERENCES users(id),
    crm VARCHAR(20) UNIQUE NOT NULL,
    specialty VARCHAR(100),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE appointments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    doctor_id UUID NOT NULL REFERENCES doctors(id),
    appointment_date TIMESTAMPTZ NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('scheduled', 'completed', 'canceled')),
    notes TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE medical_records (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    doctor_id UUID NOT NULL REFERENCES doctors(id),
    diagnosis TEXT NOT NULL,
    treatment TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE prescriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    doctor_id UUID NOT NULL REFERENCES doctors(id),
    medical_record_id UUID REFERENCES medical_records(id),
    medication_details TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE audit_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    action TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
```

## Configuration

Create a `.env` file in the root directory with the following variables:

```
# Server Configuration
SERVER_ADDRESS=:8080
CONTEXT_TIMEOUT=10

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=yourpassword
DB_NAME=hms_backend

# JWT Configuration
ACCESS_TOKEN_EXPIRY_HOUR=24
REFRESH_TOKEN_EXPIRY_HOUR=168
ACCESS_TOKEN_SECRET=your_access_token_secret
REFRESH_TOKEN_SECRET=your_refresh_token_secret
```

## Installation and Setup

### Local Development

1. Clone the repository:

   ```bash
   git clone https://github.com/DevShiba/hms-backend
   cd hms-backend
   ```

2. Set up the database (see Database Setup section)

3. Create and configure the `.env` file

4. Install dependencies:

   ```bash
   go mod download
   ```

5. Run the application:
   ```bash
   go run cmd/main.go
   ```

### Using Docker

1. Clone the repository:

   ```bash
   git clone https://github.com/DevShiba/hms-backend
   cd hms-backend
   ```

2. Create and configure the `.env` file

3. Build and run with Docker Compose:
   ```bash
   docker-compose up -d
   ```

This will start both the application and the PostgreSQL database in containers.

## API Endpoints

The HMS Backend API provides the following endpoint groups:

### Authentication

- **POST /register**: Register a new user
- **POST /login**: Authenticate a user and get tokens
- **POST /refresh**: Refresh the access token

### Patients

- **POST /patients**: Create a new patient
- **GET /patients**: List all patients
- **GET /patients/:id**: Get a specific patient
- **GET /patients/doctor/:doctor_id**: Get patients for a specific doctor
- **PATCH /patients/:id**: Update a patient
- **DELETE /patients/:id**: Delete a patient

### Doctors

- **POST /doctors**: Create a new doctor
- **GET /doctors**: List all doctors
- **GET /doctors/:id**: Get a specific doctor
- **PATCH /doctors/:id**: Update a doctor
- **DELETE /doctors/:id**: Delete a doctor

### Appointments

- **POST /appointments**: Create a new appointment
- **GET /appointments**: List all appointments
- **GET /appointments/:id**: Get a specific appointment
- **GET /appointments/patient/:patient_id**: Get appointments for a specific patient
- **GET /appointments/doctor/:doctor_id**: Get appointments for a specific doctor
- **PATCH /appointments/:id**: Update an appointment
- **DELETE /appointments/:id**: Delete an appointment

### Medical Records

- **POST /medical_records**: Create a new medical record
- **GET /medical_records**: List all medical records
- **GET /medical_records/:id**: Get a specific medical record
- **GET /medical_records/doctor/:doctor_id**: Get medical records for a specific doctor
- **PATCH /medical_records/:id**: Update a medical record
- **DELETE /medical_records/:id**: Delete a medical record

### Prescriptions

- **POST /prescriptions**: Create a new prescription
- **GET /prescriptions**: List all prescriptions
- **GET /prescriptions/:id**: Get a specific prescription
- **GET /prescriptions/patient/:patient_id**: Get prescriptions for a specific patient
- **GET /prescriptions/doctor/:doctor_id**: Get prescriptions for a specific doctor
- **PATCH /prescriptions/:id**: Update a prescription
- **DELETE /prescriptions/:id**: Delete a prescription

### Audit Logs

- **POST /audit_logs**: Create a new audit log
- **GET /audit_logs**: List all audit logs
- **GET /audit_logs/:id**: Get a specific audit log
- **PATCH /audit_logs/:id**: Update an audit log
- **DELETE /audit_logs/:id**: Delete an audit log

## Role-Based Access Control

The system implements role-based access control with three roles:

1. **Admin**: Has full access to all parts of the system
2. **Doctor**: Can manage appointments, medical records, and prescriptions related to their patients
3. **Patient**: Can view their own data, appointments, and medical records

## Architecture

The application follows clean architecture principles with the following layers:

1.  **Domain Layer**: Contains business entities, interfaces, and use case interfaces.

    - Located in the `domain/` directory.
    - Defines the core data structures (e.g., `User`, `Patient`, `Appointment`) and business rules, independent of other layers.
    - Includes definitions for request/response objects, error responses, and custom JWT claims.

2.  **Usecase Layer**: Implements application-specific business logic and orchestrates operations between the domain and repository layers.

    - Located in the `usecase/` directory.
    - Contains the core logic for each feature (e.g., `LoginUsecase`, `AppointmentUsecase`).
    - Depends on interfaces defined in the Domain Layer and implemented by the Repository Layer.

3.  **Repository Layer**: Handles data persistence and retrieval, abstracting the database interactions.

    - Located in the `repository/` directory.
    - Implements data access interfaces defined in the Domain Layer (e.g., `UserRepository`, `PatientRepository`).
    - Manages database queries and translates between database models and domain entities.

4.  **Delivery/API Layer**: Handles HTTP requests, responses, and routing. This is the entry point for external interactions.

    - Located in the `api/` directory.
    - `api/controller/`: Contains controllers that process incoming requests, call usecases, and formulate HTTP responses.
    - `api/middleware/`: Includes middleware for tasks like JWT authentication (`jwt_auth_middleware.go`) and Role-Based Access Control (`rbac_middleware.go`).
    - `api/route/`: Defines the API routes and maps them to their respective controllers.

5.  **Infrastructure/Bootstrap Layer**: Manages application startup, configuration, and external dependencies like the database connection.

    - `bootstrap/`: Contains files for application setup (`app.go`), database connection (`database.go`), and environment variable management (`env.go`).
    - `cmd/`: Contains the main application entry point (`main.go`) which initializes and starts the application.

6.  **Internal Layer**: Houses shared utilities, internal services, and components not meant for direct external use or import by higher layers like `api` or `usecase` directly (though services might be injected).
    - Located in the `internal/` directory.
    - `internal/auditservice/`: Provides a dedicated service for audit logging.
    - `internal/tokenutil/`: Contains utility functions for JWT token generation and validation.

This layered approach promotes separation of concerns, testability, and maintainability.

## Security Features

- **JWT Authentication**: Secure authentication using access and refresh tokens
- **Password Hashing**: All passwords are securely hashed
- **Role-Based Access Control**: Granular permissions based on user roles
- **Audit Logging**: Tracking all significant system actions

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature-name`
3. Commit your changes: `git commit -m 'Add some feature'`
4. Push to the branch: `git push origin feature/your-feature-name`
5. Submit a pull request

## License

This project is licensed under the MIT License.
