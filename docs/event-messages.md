## Event-Driven Communication

### Overview
The microservices in the system communicate asynchronously through RabbitMQ message queues for better state management and decoupled interaction. Each microservice emits and listens to specific events to ensure consistent state updates across the system.

### Queues

- `user_events`: Handles user-related events.
- `product_events`: Handles product-related events.
- `order_events`: Handles order-related events.

---

### 1. **User Registered**

- **Queue**: `user_events`
- **Event Type**: `UserRegistered`
- **Producer**: User Service
- **Consumers**: Order Service
- **Description**: This event is emitted when a new user is registered.

  **Event Payload**:
  ```json
  {
    "user_id": "b091999c-319d-4e88-b9d7-d98505064a56",
    "username": "ajaysinghpanwar2002",
    "email": "ajaysinghpanwar2002@gmail.com"
  }
  ```

  **Consumer Actions**:
  - **Order Service**: Updates user references in existing or future orders.
---

### 2. **User Profile Updated**

- **Queue**: `user_events`
- **Event Type**: `UserProfileUpdated`
- **Producer**: User Service
- **Consumers**: Order Service
- **Description**: This event is emitted when a user's profile is updated.

  **Event Payload**:
  ```json
  {
    "id": "b091999c-319d-4e88-b9d7-d98505064a56",
    "username": "updated_username",
    "email": "updated_email@gmail.com"
  }
  ```

  **Consumer Actions**:
  - **Order Service**: Updates user information in existing orders.

---

### 3. **Product Created**

- **Queue**: `product_events`
- **Event Type**: `ProductCreated`
- **Producer**: Product Service
- **Consumers**: Order Service
- **Description**: This event is emitted when a new product is created in the system.

  **Event Payload**:
  ```json
  {
    "product_id": "795a75e9-6d06-47c6-b868-4e2ac104adb6",
    "name": "classic 350",
    "price": 13200,
    "stock": 150
  }
  ```

  **Consumer Actions**:
  - **Order Service**: Updates the catalog of available products for future orders.

---

### 4. **Product Updated**

- **Queue**: `product_events`
- **Event Type**: `ProductUpdated`
- **Producer**: Product Service
- **Consumers**: Order Service
- **Description**: This event is emitted when a product’s details (such as stock or price) are updated.

  **Event Payload**:
  ```json
  {
    "product_id": "795a75e9-6d06-47c6-b868-4e2ac104adb6",
    "name": "classic 350",
    "price": 14000,
    "stock": 100
  }
  ```

  **Consumer Actions**:
  - **Order Service**: Updates product details in ongoing or future orders.

---

### 5. **Order Placed**

- **Queue**: `order_events`
- **Event Type**: `OrderPlaced`
- **Producer**: Order Service
- **Consumers**: Product Service
- **Description**: This event is emitted when a new order is placed.

  **Event Payload**:
  ```json
  {
    "order_id": "98ff0676-2b67-407f-9d85-0b3fa791f121",
    "user_id": "b091999c-319d-4e88-b9d7-d98505064a56",
    "product_id": "39d55f47-7620-4505-8f39-0e5714d73d87",
    "quantity": 10,
    "status": "Placed"
  }
  ```

  **Consumer Actions**:
  - **Product Service**: Updates the inventory to reflect the ordered quantity.

---

### 6. **Inventory Updated**

- **Queue**: `product_events`
- **Event Type**: `InventoryUpdated`
- **Producer**: Product Service
- **Consumers**: Order Service (optional)
- **Description**: This event is emitted when the inventory of a product is updated (e.g., due to a purchase or restocking).

  **Event Payload**:
  ```json
  {
    "product_id": "39d55f47-7620-4505-8f39-0e5714d73d87",
    "stock": 500
  }
  ```

  **Consumer Actions**:
  - **Order Service**: Updates stock availability for future orders.

---

### Summary

- **Event-driven architecture** allows each service to remain loosely coupled, reacting to state changes without direct calls.
- **Queues** like `user_events`, `product_events`, and `order_events` manage different service-specific events, ensuring state consistency across the system.

--- 