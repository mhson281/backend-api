# Backend API Documentation

## Overview

This API provides user registration, login, and a protected calculation endpoint. The API uses JWT (JSON Web Token) for authentication, allowing only registered and logged-in users to access the `/calculate` endpoint.

## Endpoints

- **POST** `/register` - Register a new user
- **POST** `/login` - Log in an existing user and receive a JWT token
- **POST** `/calculate` - Perform a calculation (protected by JWT authentication)

## Prerequisites

- Go (Golang) installed on your machine
- SQLite database (set up as part of the API)
- `curl` or any HTTP client (for testing API calls)

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/your-username/backend-api.git
cd backend-api
```
## Install Dependencies

- Ensure all dependencies are installed, including the sqlite3 package.

```bash
go mod tidy

```

## Set Up Environment Variables

- Create a .env file in the root of your project and set your JWT secret:

```bash
JWT_SECRET=your-very-secret-key
```

- Run the Server

- Start the server:

`go run cmd/api/main.go`

- The server will start at http://localhost:8080.

## Using the Endpoints

### 1. Register a New User

```plaintext
Endpoint

    URL: http://localhost:8080/register
    Method: POST
    Content-Type: application/json
```
Request Body

```json
{
  "username": "your_username",
  "password": "your_password"
}

```

Example curl Command

```bash
curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{
  "username": "exampleuser",
  "password": "securepassword"
}'

```

Response

```plaintext

Success:

Success (201 Created): Returns a confirmation message.


{
  "message": "User registered successfully"
}

Failure:

Error (409 Conflict): Username already exists.

{
  "error": "Username already exists"
}

```

### 2. Log in an Existing User

```plaintext
Endpoint

    URL: http://localhost:8080/login
    Method: POST
    Content-Type: application/json

```

Request Body

```json
{
  "username": "your_username",
  "password": "your_password"
}

```

Example curl Command

```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
  "username": "exampleuser",
  "password": "securepassword"
}'

```

Response

```plaintext
Success:

Success (200 OK): Returns a JWT token.
N

{
  "token": "your.jwt.token"
}

Failure:
Error (401 Unauthorized): Invalid credentials.

    {
      "error": "Invalid username or password"
    }

```


### 3. Perform a Calculation (Protected Endpoint)

This endpoint is protected by JWT authentication. You must include a valid token in the Authorization header.
Endpoint

```plaintext
    URL: http://localhost:8080/calculate
    Method: POST
    Content-Type: application/json
    Authorization: Bearer <your.jwt.token>

```

Request Body

```json
{
  "operation": "add", // Options: add, subtract, multiply, divide
  "operand1": 10,
  "operand2": 5
}

```

Example curl Command

- First, obtain the JWT token by logging in (as shown in the previous section).
- Then, use the token in the Authorization header for the /calculate request.

```bash
curl -X POST http://localhost:8080/calculate \
-H "Authorization: Bearer your.jwt.token" \
-H "Content-Type: application/json" \
-d '{
  "operation": "add",
  "operand1": 10,
  "operand2": 5
}'
```

*Note: Replace your.jwt.token with the actual token you received from the /login endpoint.
Response*

```plaintext

Success:

Success (200 OK): Returns the calculation result.
{
  "result": 15,
  "operation": "add",
  "operand1": 10,
  "operand2": 5
}

Failure:
Error (400 Bad Request): Invalid operation or division by zero.

{
  "error": "unsupported operation"
}

Error (401 Unauthorized): Missing or invalid JWT token.

    {
      "error": "Unauthorized"
    }
```



## Summary

### Testing the API with curl

- Register a new user with /register.
- Use curl to send a POST request with a JSON payload containing "username" and "password".
- Log in with the same credentials using /login.
- Copy the token value from the response.
- Use the Token to Access /calculate.
- Use the Authorization header with the value Bearer your.jwt.token.
- Make the API call with the JSON payload containing the operation and operands.

### Notes

 *JWT Expiration: Tokens expire after 24 hours. If you receive a 401 Unauthorized response, log in again to get a new token.*

Supported Operations: The /calculate endpoint supports the following operations:
  - add: Adds two numbers.
  - subtract: Subtracts the second number from the first.
  - multiply: Multiplies two numbers.
  - divide: Divides the first number by the second (returns an error if dividing by zero).
