# Library-API

API for virtual library management, where possible:

- Perform CRUD operations in books
- Upload the book to a Google Cloud Platform bucket
- Delete book from Google Cloud bucket
- Download books

## API Documentation

See [swagger.json](/docs/swagger.json)

## How to run

### Configure Bucket in GCP

- Create the bucket
- Create a service account in the following directory:

```bash
/config/credentials
```

### Setup Database

- Install docker

Docker is a software platform that allows developers to create, deploy, and run applications in containers. See the official documentation for install the version compatible with your OS: [Install Docker Engine](https://docs.docker.com/engine/install/)

- Run Docker Compose

```bash
docker compose -f ./config/docker/docker-compose.yml up -d --build
```

### Download project dependencies

- Inside lib-api directory, run:

```go
go mod vendor
```

### Configure the project

1. Setup the environment variables listed below in directory:

```bash
/config/env/envs.yaml
```

2. Inside lib-api directory, run:

```go
go run cmd/server/main.go
```

## Test
1. Inside lib-api directory, run:
```go
go test ./... -v
```

You should see messages like:

![test output](/docs/test-output.png)

## **Environment Variables**

| Name | Description | Default |
| --- | --- | --- |
| DATABASE.USER | Username to access database | root |
| DATABASE.PASSWORD | Password to access database | 12345678 |
| GCP.BUCKET_NAME | Name of the bucket that will manage the book files on Google Cloud Platform | library |
| GCP.CREDENTIALSPATH | Path of the service account that will give free access to the bucket | config/credentials/service_account.json |