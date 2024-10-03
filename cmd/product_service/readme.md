# Product Service

## Overview
The `product-service` handles product management, including creating, retrieving, updating, and deleting products. Inventory stock levels are also managed here.

## Database Structure
The products are stored in the `products` table:

| Column     | Type     | Description                    |
|------------|----------|--------------------------------|
| id         | UUID     | Unique identifier for the product |
| name       | TEXT     | Product name                   |
| Description| TEXT     | Product Description            |
| price      | FLOAT    | Product price                  |
| stock      | INTEGER  | Available stock for the product |
| created_at | TIMESTAMP | Product creation timestamp     |
| updated_at | TIMESTAMP | Last product update timestamp  |

## Event Consumption
- **`ProductCreated`**: Consumed when a new product is added.
- **`InventoryUpdated`**: Consumed when product stock is updated.

## API Endpoints
For more details on the API, refer to the [API Endpoints](api-endpoints.md).

---
