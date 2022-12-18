# Cake Store API

## Run locally

- Copy env-example to .env
- Prepare environment, fill DB parameters:

```bash
export DB_NAME=cake_store
export DB_PASSWORD=root
export DB_USER=root

export TEST_DB_NAME=cake_store_test
export TEST_DB_PASSWORD=root
export TEST_DB_USER=root

export DB_CONTAINER_NAME=cake-store-db

```

- Build and run image of docker:

```bash
docker-compose build
docker-compose up
```

Server is listening on localhost:8080

Before running Migration Commands first create database cake-store

#### Migration Commands

| Command             | Desc                            |
| ------------------- | ------------------------------- |
| `make migrate-up`   | runs migration up command       |
| `make migrate-down` | runs migration down command     |
| `make drop`         | Drop everything inside database |

## Test

## License

Copyright (c) 2022 Rendered Text

Distributed under the MIT License. See the file LICENSE.
