# API Endpoints Documentation

This documentation provides details on the available API endpoints for various services, including User, Product, and Order management. Each service is defined using GraphQL or RESTful methods, supporting operations such as querying data and making mutations.

## Base URLs

- **GraphQL Gateway:** `http://localhost:8080/query`
- **User Service:** `http://localhost:8081`
- **Product Service:** `http://localhost:8082`
- **Order Service:** `http://localhost:8083`

---

## GraphQL Gateway

### 1. Get User by ID
- **Method:** POST
- **Endpoint:** `/query`
- **Description:** Retrieve a user's details by their ID.
- **Query:**
  ```graphql
  query {
    user(id: "b091999c-319d-4e88-b9d7-d98505064a56") {
      id
      username
      email
    }
  }
  ```
  
### 2. Get All Users
- **Method:** POST
- **Endpoint:** `/query`
- **Description:** Retrieve a list of all users.
- **Query:**
  ```graphql
  query {
    users {
      id
      username
      email
    }
  }
  ```

### 3. Get All Products
- **Method:** POST
- **Endpoint:** `/query`
- **Description:** Retrieve a list of all products.
- **Query:**
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

### 4. Get Product by ID
- **Method:** POST
- **Endpoint:** `/query`
- **Description:** Retrieve a product's details by its ID.
- **Query:**
  ```graphql
  query {
    product(id: "39d55f47-7620-4505-8f39-0e5714d73d87") {
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

### 5. Get All Orders
- **Method:** POST
- **Endpoint:** `/query`
- **Description:** Retrieve a list of all orders.
- **Query:**
  ```graphql
  query {
    orders {
      id
      user_id
      product_id
      quantity
      total_price
      status
    }
  }
  ```

### 6. Get Order by ID
- **Method:** POST
- **Endpoint:** `/query`
- **Description:** Retrieve an order's details by its ID.
- **Query:**
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
- **Variables:**
  ```json
  {
    "id": "98ff0676-2b67-407f-9d85-0b3fa791f121"
  }
  ```

### 7. Register User

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

### 8. Create product

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

### 9. Place an Order

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

---

## User Service

### 1. Register User
- **Method:** POST
- **Endpoint:** `/register`
- **Description:** Register a new user.
- **Request Body:**
  ```json
  {
    "username": "ayushsingh",
    "email": "ayushsingh@gmail.com",
    "password": "ayushsingh2002pass"
  }
  ```

### 2. Get All Users
- **Method:** GET
- **Endpoint:** `/users`
- **Description:** Retrieve a list of all registered users.

### 3. Get User by ID
- **Method:** GET
- **Endpoint:** `/users/{userId}`
- **Description:** Retrieve a user's details by their ID.

### 4. User Login
- **Method:** POST
- **Endpoint:** `/login`
- **Description:** Authenticate a user by username and password.
- **Request Body:**
  ```json
  {
    "username": "ayushsingh",
    "password": "ayushsingh2002pass"
  }
  ```

### 5. Update Profile
- **Method:** PUT
- **Endpoint:** `/profile`
- **Description:** Update the user's profile.
- **Headers:**
  - `Authorization: <JWT_TOKEN>`
- **Request Body:**
  ```json
  {
    "email": "heyajaysingh123@gmail.com"
  }
  ```

---

## Product Service

### 1. Create Product
- **Method:** POST
- **Endpoint:** `/products`
- **Description:** Create a new product.
- **Request Body:**
  ```json
  {
    "name": "bullet classic 350",
    "description": "elegant and masterpiece",
    "price": 100012,
    "stock": 300
  }
  ```

### 2. Get Product by ID
- **Method:** GET
- **Endpoint:** `/products/{productId}`
- **Description:** Retrieve product details by product ID.

### 3. Get All Products
- **Method:** GET
- **Endpoint:** `/products`
- **Description:** Retrieve a list of all available products.

### 4. Update Product
- **Method:** PUT
- **Endpoint:** `/products/{productId}`
- **Description:** Update product details.
- **Request Body:**
  ```json
  {
    "stock": 205
  }
  ```

### 5. Delete Product
- **Method:** DELETE
- **Endpoint:** `/products/{productId}`
- **Description:** Delete a product by its ID.

---

## Order Service

### 1. Place Order
- **Method:** POST
- **Endpoint:** `/orders`
- **Description:** Place an order for a product.
- **Request Body:**
  ```json
  {
    "user_id": "44775972-75b1-4f3f-90c2-17f9e2776068",
    "product_id": "37cab164-7a59-4a58-a3de-042280fc8d7e",
    "quantity": 10
  }
  ```

### 2. Get All Orders
- **Method:** GET
- **Endpoint:** `/orders`
- **Description:** Retrieve a list of all orders.

### 3. Get Order by ID
- **Method:** GET
- **Endpoint:** `/orders/{orderId}`
- **Description:** Retrieve order details by order ID.

---