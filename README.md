# Go Rest API Boilerplate

This project is a boilerplate for creating a REST API using the Go (Golang) programming language.

## Usage

**1. Install Dependencies**

Ensure that you have Go installed on your system. Then, run the following command to install project dependencies:
```
go mod tidy
```

**2. Run the Application**

```
go run main.go
```
The application will run on http://localhost:8080 by default.



## Endpoint

**Product Endpoint**

| Method | Endpoint              | Handler Function     | Description                 |
|--------|-----------------------|----------------------|-----------------------------|
| GET    | /v1/product           | productHandler.ShowProducts   | Display all products      |
| GET    | /v1/product/:id       | productHandler.ShowProduct    | Display details of a product |
| POST   | /v1/product           | productHandler.CreateProduct  | Create a new product       |
| PUT    | /v1/product/:id       | productHandler.UpdateProduct  | Update a product           |
| DELETE | /v1/product/:id       | productHandler.DestroyProduct | Delete a product           |

**Unit Endpoint**

| Method | Endpoint              | Handler Function     | Description                 |
|--------|-----------------------|----------------------|-----------------------------|
| GET    | /v1/unit           | unitHandler.ShowUnits   | Display all units      |
| GET    | /v1/unit/:id       | unitHandler.ShowUnit    | Display details of a unit |
| POST   | /v1/unit           | unitHandler.CreateUnit  | Create a new unit       |
| PUT    | /v1/unit/:id       | unitHandler.UpdateUnit  | Update a unit           |
| DELETE | /v1/unit/:id       | unitHandler.DestroyUnit | Delete a unit           |


Regards.
