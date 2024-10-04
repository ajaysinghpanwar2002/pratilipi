# Project Overview

## Purpose
This project consists of three microservices (`user-service`, `product-service`, and `order-service`) that manage user registration, product inventory, and order placement. The services communicate asynchronously using RabbitMQ for event-driven architecture and use PostgreSQL for data persistence.

## Key Features
- **Microservices**: Independent services for users, products, and orders.
- **Asynchronous Communication**: RabbitMQ for event-driven message passing.
- **Data Persistence**: PostgreSQL with migrations handled via Golang's `sqlx`.
- **Caching**: Redis is used for caching to improve read performance.
- **Authentication**: JWT-based authentication in the `user-service`.

## Technologies
- Golang (Fiber Framework)
- PostgreSQL
- RabbitMQ
- Redis
- Docker (for containerization)
- gqlgen (for the GraphQL API gateway)

---

## Documentation

- [User Service Documentation](./cmd/user_service/readme.md)
- [Product Service Documentation](./cmd/product_service/readme.md)
- [Order Service Documentation](./cmd/order_service/readme.md)
- [API Endpoints Documentation](./docs/api-endpoints.md)
- [Event Communication Documentation](./docs/event-messages.md)

---

## Local Setup

Follow these steps to run the project locally:

1. **Clone the repository:**

   ```bash
   git clone https://github.com/ajaysinghpanwar2002/pratilipi.git
   ```

2. **Navigate into the repository:**

   ```bash
   cd pratilipi
   ```

3. **Build the services:**

   ```bash
   make build
   ```

4. **Start the services using Docker:**

   ```bash
   make start
   ```

5. **Accessing the services:**
   - GraphQL Gateway: `http://localhost:8080`
   - User Service: `http://localhost:8081`
   - Product Service: `http://localhost:8082`
   - Order Service: `http://localhost:8083`

---

## Postman Collection

I am providing a Postman collection to test the API endpoints. You can download it using the link: [Postman Collection](./pratilipi.postman_collection.json)
