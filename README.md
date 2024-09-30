# Microservices Delivery-management Project

This project is a microservices-based delivery management system built with Go, utilizing Gin framework, Kafka for event streaming, and PostgreSQL for data storage.

## Services

1. Gateway Service
2. Order Service
3. Delivery Service
4. User Service
5. Notification Service

## Prerequisites

- Go 1.23.1 or later
- Docker and Docker Compose
- PostgreSQL
- Kafka

## Getting Started

1. Clone the repository:
   ```
   git clone https://github.com/Sotatek-GiapHoang/delivery-management.git
   cd delivery-management
   ```

2. Set up environment variables:
   Copy the `.env.example` file to `.env` in each service directory and update the values as needed.

3. Start the services:
   ```
   docker-compose up -d
   ```

4. API Gateway will run at `http://localhost:8080`

## API Documentation

The documentation is available at `http://localhost:8080/swagger/index.html` when the server is running.

## Authentication

API use authentication with Bearer Token. To authenticate a request, add the `Authorization` header with the value as the JWT token:

```
Authorization: Bearer <token>
```
## Endpoints

- `/api/v1/users/*`: User service endpoints
- `/api/v1/orders/*`: Order service endpoints
- `/api/v1/deliveries/*`: Delivery service endpoints

The detail of each endpoint is in the swagger documentation.

## Development

To run a specific service locally:

1. Navigate to the service directory:
   ```
   cd order-service
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Run the service:
   ```
   go run cmd/server/main.go
   ```


## Additional Information

Ensure you have all the required dependencies installed and your environment is properly set up before running the commands above. If you encounter any issues, please refer to the documentation of the respective tools or seek help from the project maintainers.