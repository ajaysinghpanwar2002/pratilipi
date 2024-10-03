# User Service

## Overview
The `user-service` manages user registration, authentication, and profile updates. Users can register, login, and update their profile details.

## Database Structure
The users are stored in the `users` table:

| Column     | Type    | Description                   |
|------------|---------|-------------------------------|
| id         | UUID    | Unique identifier for the user |
| username   | TEXT    | User's username                |
| email      | TEXT    | User's email address           |
| password   | TEXT    | User's password                |
| created_at | TIMESTAMP | User's creation timestamp      |
| updated_at | TIMESTAMP | Last profile update timestamp  |

## Event Consumption
- **`UserRegistered`**: Consumed when a new user registers, updating the `order-service` and other services.
- **`UserProfileUpdated`**: Consumed when a user updates their profile information.

Refer to the [Event Messages](../../docs/event-messages.md) for message formats.

## API Endpoints
For more details on the API, refer to the [API Endpoints](../../docs/api-endpoints.md).

---
