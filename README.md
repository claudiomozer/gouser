# User CRUD

This application contains 4 basic operations with a simple user structure and a database.

## Stack

* Go
* Gin
* Tern (migrations)
* PostgreSQL

## Requirements

To run this application are necessaries the technologies:
* Go > 1.21
* PostgreSQL or Docker with docker-compose

## Running the application

First of all it's necessary to install the go libraries to run the project, for this on your terminal run:

```shell
    make setup
```

After that rename the `.env.example` file to `.env` with your values. So, you can export the values to the OS typing:
```shell
    export $(< .env)
```

If you don't have postgres installed in your machine you can use the `docker-compose.yml` on the project root. To start your postgres execute:
```shell
   docker compose up -d 
```

Now, run the migrations with the command:
```shell
   make migrations 
```

Then run the application:
```shell
   make run 
```

# Endpoints

## Create User

### POST localhost:`{APP_PORT}`/v1/user

Body:
```json
{
	"name": "Valid Name",
	"email": "valid@email.com",
	"role": "admin" // could be modifier or watcher
}
```

## Get User

### GET localhost:`{APP_PORT}`/v1/user

Body:
```json
{
    "id": "valid-uuid",
	"name": "Valid Name",
	"email": "valid@email.com",
	"role": "admin" // could be modifier or watcher
}
```

## Update User Role

### PATCH localhost:`{APP_PORT}`/v1/user

Body:
```json
{
    "id": "valid-uuid",
	"role": "admin" // could be modifier or watcher
}
```

## Delete User 

### DELETE localhost:`{APP_PORT}`/v1/user/`:id`

