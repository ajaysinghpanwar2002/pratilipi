# Order Service

## Overview
The `order-service` is responsible for order placement. It validates the user and product availability before placing an order.

## Database Structure
The orders are stored in the `orders` table:

| Column     | Type     | Description                     |
|------------|----------|---------------------------------|
| id         | UUID     | Unique identifier for the order |
| user_id    | UUID     | The ID of the user placing the order |
| product_id | UUID     | The ID of the product being ordered |
| quantity   | INTEGER  | Quantity of product ordered     |
| Status     | TEXT     | Available or Booked             |
| total_price| FLOAT    | Total price for the order       |
| created_at | TIMESTAMP | Order creation timestamp       |
| updated_at | TIMESTAMP | Order updation timestamp       |

## Event Emission
- **`OrderPlaced`**: Emitted after successfully placing an order, consumed by other services.

## API Endpoints
For more details on the API, refer to the [API Endpoints](api-endpoints.md).

---
