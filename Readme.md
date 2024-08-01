# Disburse App

This is a example disburse microservice application. In this application all of the microservice is written in golang & communicate with RabbitMQ.

In this application the disburse the user's id is given by the request body, it's because to cut scope of authentication.

## Components

This application consist of serveral components:

1. Api Gateway, to help REST Api application to use the application
2. User service, provides with user data
3. Disburse service, provides business logic of the disburse process.
4. Wallet service, provides with user's wallet / balance data

## Process Of The Disburse

## How to start

The easiest way is to use `docker-compose`:

```
docker-compose up --build
```
