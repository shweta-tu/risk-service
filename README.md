# Risk Management API

## Overview

This project is a simple API for managing risks, built using the Gin framework. It allows users to create and retrieve  risks through RESTful endpoints.

## Features

- **CRUD Operations**: Create and retrieve risks.
- **JSON-based**: All data transfer is done in JSON format.

## Prerequisites

- **Go 1.20+**: Ensure you have Go installed. Download it from the [official site](https://golang.org/dl/).
- **Docker** (Optional): For containerization and easier deployment.
- **Postman** (Optional): For testing API endpoints.

## Setup Instructions

1. **Clone the repository**:

   ```bash
   git clone https://github.com/shweta-tu/risk-service.git
   cd risk-service

2. **Install Dependencies**:
Ensure all dependencies are installed:

    ```bash
    go mod tidy

3. **Running the Application**:
Run the application locally:

    ```bash
    go run main.go

The server will start on <http://localhost:8080>.

4. **Running Tests**:
Run the tests to ensure everything is working correctly:

    ```bash
    go test ./... -cover

This command will execute all tests and provide a coverage report.

5. **Docker Setup (Optional)**:
   For containerized deployment:

   - **Build the Docker Image**:

     ```bash
     docker build -t risk-service .

   - **Run the Docker Container**:

     ```bash
     docker run -p 8080:8080 risk-service

The application will be accessible at <http://localhost:8080>.

6. **Rate Limiting**
Rate limiting is not implemented. We can use Gin Middleware for Rate Limiting or implement rate limiting at the cloud level using services like AWS API Gateway

## API Documentation

### Risk Object

The Risk object represents a risk in the system. It has the following structure:

```
{
  "id": "string (UUID)",
  "state": "enum (State)",
  "title": "string",
  "description": "string"
}
```

- **id**: The unique identifier for the risk. This is a UUID and is auto-generated.
- **state**: The state of the risk. Must be one of the following:
  - `open`
  - `closed`
  - `accepted`
  - `investigating`
- **title**: The title of the risk.
- **description**: A detailed description of the risk.

### Endpoints

### GET /v1/risks

Retrieve all risks.

Response:

- 200 OK: Returns an array of Risk objects.

Example:

  ```
  curl -X GET http://localhost:8080/v1/risks -H "Content-Type: application/json"
  ```

Response:

  ```

  [
    {
        "id": "9bde444b-25a1-40d7-9e14-c3c4c9b060cd",
        "state": "open",
        "title": "Risk 1",
        "description": "Description of Risk 1"
    },
    {
        "id": "de305d54-75b4-431b-adb2-eb6b9e546014",
        "state": "closed",
        "title": "Risk 2",
        "description": "Description of Risk 2"
    }
]
```

### GET /v1/risks/:id

Retrieve a risk by its UUID.

URL Parameters:
id: (string) The UUID of the risk to retrieve.

Response:

- 200 OK: Returns the requested Risk object.
- 400 Bad Request: Returns if the risk ID is invalid.
- 404 Not Found: Returns if the risk with the given ID is not found.

Example:

  ```
  curl -X GET http://localhost:8080/v1/risks/e51c2f96-43d6-4e2b-a48e-5c109ad9b4f1 -H "Content-Type: application/json"
  ```

Response:

  ```
  {
      "id": "e51c2f96-43d6-4e2b-a48e-5c109ad9b4f1",
      "state": "open",
      "title": "New Risk",
      "description": "Description of the new risk"
  }
  ```

### POST /v1/risks

Create a new risk.

### Request Body

A Risk object without the id field. The id will be auto-generated.

### Response

- 201 Created: Returns the created Risk object.
- 400 Bad Request: Returns if the payload is invalid or the state value is not allowed.

Example:

  ```
  curl -X POST http://localhost:8080/v1/risks -H "Content-Type: application/json" -d '{
    "state": "open",
    "title": "New Risk",
    "description": "Description of the new risk"
  }'
  ```

Response:

  ```
  {
      "id": "e51c2f96-43d6-4e2b-a48e-5c109ad9b4f1",
      "state": "open",
      "title": "New Risk",
      "description": "Description of the new risk"
  }
  ```

### Error Responses

The API uses standard HTTP status codes to indicate the success or failure of an API request. Some common response codes are:

- **200 OK**: The request was successful.
- **201 Created**: The resource was successfully created.
- **400 Bad Request**: The request was invalid or cannot be served.
- **404 Not Found**: The requested resource could not be found.
- **500 Internal Server Error**: An error occurred on the server.
