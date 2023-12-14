# Library-API

API for virtual library management, where possible:

- Perform CRUD operations in books
- Upload the book to a Google Cloud Platform bucket
- Delete book from Google Cloud bucket
- Download books

## API Documentation

See [swagger.json](/docs/swagger.json)

## How to run

### Download project dependencies

- Inside lib-api directory, run:

```go
go mod vendor
```

### Setup Database

- Install docker

Docker is a software platform that allows developers to create, deploy, and run applications in containers. See the official documentation for install the version compatible with your OS: [Install Docker Engine](https://docs.docker.com/engine/install/)

- Run Docker Compose

```bash
docker compose -f config/docker/docker-compose.yml up -d --build
```