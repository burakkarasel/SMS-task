# _SMS Task_


---

### Table of Contents

- [Description](#description)
- [How To Use](#how-to-use)
- [Author Info](#author-info)

---

## Description

SMS Task is a Student Management System tool. 

## Technologies

### Main Technologies

- [Go](https://go.dev/)
- [Gin Framework](https://github.com/gin-gonic/gin)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)

### Libraries

- [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
- [lib/pq](https://github.com/lib/pq)
- [gin-gonic/gin](https://github.com/gin-gonic/gin)

[Back To The Top](#sms-task)

---

## How To Use

### Tools

- [Go](https://go.dev/dl/)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [TablePlus](https://tableplus.com/download)

### Setup Database

- Create a database in PostgreSQL and add its URL as a flag while running the app with ```-url``` flag.

### Run tests

- Create a test.sh file in root directory. Add executable with ```chmod +x test.sh```.
- Then put this piece of code in it.

```
#!/bin/bash

go test -cover -url="postgresql://{user}:{password}@{db url}/{db name}?sslmode=disable"
```

- Then you can run make test command.

```
make test
```

### Start App

- Create a Dockerfile.

```
# Build Stage
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/web/main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY /internal/db/migration ./internal/db/migration

EXPOSE 8080
CMD ./main -url=postgresql://postgresql://{user}:{password}@{db url}/{db name}?sslmode=disable -port=0.0.0.0:8080
```

- Create a docker-compose.yml file.

```
services:
  api:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    command: ./main -url=postgresql://kullop:TwvU5euWR0DoffTvpCKO@sms-task.ct4mxefnosgh.us-east-1.rds.amazonaws.com:5432/sms_task?sslmode=disable -port=0.0.0.0:8080
```

- Run the app.

```
docker compose up
```

### Give it a try

#### Routes

- Base URL: _http://localhost:8080/api/v1/_

##### Endpoint: /students

##### Endpoint: /classes

Creates a new class with the provided name and professor.

Request Body
The request body must be a JSON object containing the following fields:

```
name	string	The name of the class. Required, minimum length of 2.
professor	string	The name of the professor of the class. Required, minimum length of 4.
```

Example:

```
{
"name": "Math 101",
"professor": "John Doe"
}
```

Success Response
If the class is created successfully, the server will respond with HTTP status code 201 Created and a JSON object representing the created class.

```
{
"success": true,
"code": 201,
"message": "",
"data": {
"id": 1,
"name": "Math 101",
"professor": "John Doe",
"createdAt": "2023-03-14T10:30:00Z",
"updatedAt": "2023-03-14T10:30:00Z"
}
}
```
Error Response
If there is an error creating the class, the server will respond with an HTTP status code indicating the type of error that occurred and a JSON object containing an error message.

```
Invalid Request Body	400 Bad Request	{"success":false,"code":400,"message":"<error message>","data":null}
Internal Server Error	500 Internal Server Error	{"success":false,"code":500,"message":"<error message>","data":null}
```
[Back To The Top](#sms-task)

---

## Author Info

- Twitter - [@dev_bck](https://twitter.com/dev_bck)

[Back To The Top](#sms-task)