# eCommerce API

A work-in-progress e-commerce REST API built with **Go** and **PostgreSQL**.

The project is structured around controllers, services, repositories, and middleware, with JWT authentication and PostgreSQL for data persistence.

## Current Features

* JWT-based authentication
* Product listing
* Product creation
* Product stock management
* Checkout / order creation
* Order history
* PostgreSQL database integration
* Transaction-based checkout
* Inventory locking during checkout to help prevent overselling

## API Endpoints

| Method | Endpoint    | Description             |
| ------ | ----------- | ----------------------- |
| `POST` | `/login`    | Login and receive a JWT |
| `GET`  | `/products` | Get available products  |
| `POST` | `/products` | Create a product        |
| `POST` | `/checkout` | Create an order         |
| `GET`  | `/orders`   | Get orders              |

## Project Structure

```text
eCommerceAPI/
├── controllers/
├── middleware/
├── models/
├── repository/
├── services/
├── main.go
├── go.mod
└── go.sum
```

The project follows a simple layered architecture:

```text
Controller
    ↓
Service
    ↓
Repository
    ↓
PostgreSQL
```

## Tech Stack

* Go
* PostgreSQL
* JWT
* `database/sql`
* `lib/pq`

## Running Locally

Clone the repository:

```bash
git clone https://github.com/chnert/eCommerceAPI.git
cd eCommerceAPI
```

Configure the required environment variables:

```env
DB_CONN_STR=your_postgresql_connection_string
JWT_SECRET=your_jwt_secret
PORT=:8080
```

Run the API:

```bash
go run .
```

The server runs on `http://localhost:8080` by default.

## Planned Features

The API is still under development. Planned additions include:

* Proper user registration and authentication
* Password hashing
* User roles and authorization
* Product update and deletion
* Shopping cart functionality
* Improved order management
* Payment integration
* Order status management
* Database migrations
* API documentation / Swagger
* Tests and additional validation

More features and improvements will be added as development continues.

## Status

**Work in Progress**

The API is currently focused on establishing the core e-commerce functionality and backend architecture.
