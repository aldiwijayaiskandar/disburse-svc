# Disburse App

This is a example disburse microservice application. In this application all of the microservice is written in golang & communicate with RabbitMQ.

In this application the disburse the user's id is given by the request body, it's because to cut scope of authentication.

## Components

This application consist of serveral components:

1. Api Gateway, to help REST Api application to use the application
2. User service, provides with user data
3. Disburse service, provides business logic of the disburse process.
4. Wallet service, provides with user's wallet / balance data

<a href="https://ibb.co.com/cXk4y30"><img src="https://i.ibb.co.com/S75F6yb/service-diagram.png" alt="service-diagram" border="0"></a>

## Process Of The Disburse
<a href="https://ibb.co.com/fMQQ5bB"><img src="https://i.ibb.co.com/G7FFLXB/sequence.png" alt="sequence" border="0"></a>

## How to start

The easiest way is to use `docker-compose`:

```
docker-compose up --build
```
