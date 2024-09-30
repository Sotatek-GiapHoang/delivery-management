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

4. The services should now be running:
   - Gateway Service: http://localhost:8080
   - Order Service: http://localhost:8081
   - Delivery Service: http://localhost:8082
   - User Service: http://localhost:8083
   - Notification Service: http://localhost:8084

## API Documentation

(Add information about your API endpoints, request/response formats, etc.)

## Development

To run a specific service locally:

1. Navigate to the service directory:
   ```
   cd order-service
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Run the service:
   ```
   go run cmd/server/main.go
   ```


## Additional Information

Ensure you have all the required dependencies installed and your environment is properly set up before running the commands above. If you encounter any issues, please refer to the documentation of the respective tools or seek help from the project maintainers.