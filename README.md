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
      - "80:80"
    command: ./main -url=postgresql://kullop:TwvU5euWR0DoffTvpCKO@sms-task.ct4mxefnosgh.us-east-1.rds.amazonaws.com:5432/sms_task?sslmode=disable -port=0.0.0.0:80
```

- Run the app.

```
docker compose up
```

### Give it a try

- You can check out the endpoints with inputs via PostMan by opening [this](SMS-Task.postman_collection.json)

[Back To The Top](#sms-task)

---

## Author Info

- Twitter - [@dev_bck](https://twitter.com/dev_bck)

[Back To The Top](#sms-task)