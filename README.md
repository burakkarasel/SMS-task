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

- Create a Dockerfiles.

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