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