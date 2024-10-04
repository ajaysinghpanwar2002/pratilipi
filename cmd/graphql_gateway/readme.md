# GraphQL Gateway for Microservices

## Overview

This project is a GraphQL gateway that consolidates multiple microservices into a single GraphQL API. It is designed to provide a unified interface for managing users, products, and orders, making it easier to interact with different services in one place.

### Features

- **Users Service**: Manage user registration, login, and profile updates.
- **Products Service**: Manage product creation, updates, and retrieval.
- **Orders Service**: Manage order placement and retrieval.
- **GraphQL API**: Unified schema for interacting with the microservices.
- **JWT Authentication**: Secured with JWT middleware for mutation requests.
- **RabbitMQ Integration**: Event-based communication between services.
  
## Architecture

This gateway interacts with three microservices:
1. **User Service** (Running on port 8081)
2. **Product Service** (Running on port 8082)
3. **Order Service** (Running on port 8083)

The GraphQL API aggregates these services and provides a single endpoint for client interaction. It communicates with these services using HTTP calls and RabbitMQ events.

## Endpoints

- **GraphQL Playground**: [http://localhost:8080/](http://localhost:8080/)
- **GraphQL Query Endpoint**: `/query`

### Example Queries

#### 1. Get All Users

```graphql
query {
  users {
    id
    username
    email
  }
}

```

#### 2. Get User by ID

```graphql
query {
  user(id: "user-id") {
    id
    username
    email
  }
}
```

#### 3. Get All Products

```graphql
query {
  products {
    id
    name
    price
    stock
    createdAt
    updatedAt
  }
}
```

#### 4. Get Product by ID

```graphql
query {
  product(id: "productId") {
    id
    name
    description
    price
    stock
    createdAt
    updatedAt
  }
}
```

#### 5. Get All orders

```graphql
query {
  product(id: "productId") {
    id
    name
    description
    price
    stock
    createdAt
    updatedAt
  }
}

```

#### 6. Get Order by ID

```graphql
query getOrder($id: ID!) {
  order(id: $id) {
    id
    user_id
    product_id
    quantity
    total_price
    status
  }
}

```

***variables***

```json
{
  "id": "98ff0676-2b67-407f-9d85-0b3fa791f121"
}
```


#### 7. Register User

```graphql
mutation RegisterUser($input: RegisterInput!) {
  registerUser(input: $input) {
    id
    username
    email
    createdAt
    updatedAt
  }
}
```

***Variables:***

```json
{
  "input": {
    "username": "ajay",
    "email": "ajay@gmail.com",
    "password": "12345"
  }
}

```

#### 8. Create product

```graphql
mutation CreateProduct($input: ProductInput!) {
  createProduct(input: $input) {
    id
    name
    description
    price
    stock
    createdAt
    updatedAt
  }
}
```

***variables***
```json
{
  "input": {
    "name": "Test product",
    "description": "A test product",
    "price": 1913,
    "stock": 1100
  }
}
```

#### 9. Place an Order

```graphql
mutation PlaceOrder($input: OrderInput!) {
  placeOrder(input: $input) {
    id
    user_id
    product_id
    quantity
    total_price
    status
    createdAt
    updatedAt
  }
}
```

***Variables:***

```json
{
  "input": {
    "user_id": "b091999c-319d-4e88-b9d7-d98505064a56",
    "product_id": "39d55f47-7620-4505-8f39-0e5714d73d87",
    "quantity": 20
  }
}
```

### Running the Project

1. **Clone the Repository**

   ```bash
   git clone https://github.com/ajaysinghpanwar2002/pratilipi.git
   cd pratilipi
   ```

2. **Build the Docker Containers**

   You can spin up all services using Docker:

   ```bash
   make build
   make start
   ```

3. **Access the GraphQL Playground**

   Open [http://localhost:8080/](http://localhost:8080/) to access the GraphQL playground for running queries and mutations.

4. **Test using Postman**

   I am providing a Postman collection to test the API endpoints. You can download it using the link below:

   [Download Postman Collection](./pratilipi.postman_collection.json)

### Project Structure

```
pratilipi/
├── cmd/
│   ├── graphql_gateway/
│   ├── user-service/
│   ├── product-service/
│   ├── order-service/
├── pkg/
│   ├── db/
│   ├── rabbitmq/
├── docker-compose.yml
├── Makefile
├── go.mod
├── go.sum
└── pratilipi.postman_collection.json
```