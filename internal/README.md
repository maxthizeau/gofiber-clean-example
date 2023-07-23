./internal

# entity

Struct that can be used anywhere in the internal folder
They are are representation of the database, but also the main objets of our application.

# configuration (Todo: merge/refactor config/configuration folders)

config file needed to operate the server/database

# model (Todo: refactor + implement converters)

Represent JSON/Data that the user can send to the API, or receive from the API.
Incoming Data is suffixed with "Input" (ex : CreateUserInput, UpdateUserInput)
Outgoing Data has converters to transform it to entities, or to transform entities to model.

# controller

Controller define routes and handlers. Handlers are used by one route and can use . A Route calls a Handler (return a JSON response) that calls a Service (that return an entity)
The Handler is responsible to decode & validate user input, and transform result from the service to JSON response (using a model)
Similar to : https://github.com/Creatly/creatly-backend/tree/main/internal/delivery/http/v1

# service

🪄 Where the magic happens : Apply the business rules here.

It has a file "service.go" that define all service interface, and import all deps. It also create all services by passing dependend services + deps.

# repository

A repository is responsible to getting, and setting data from/to the database. Functions receive an entity, and should return an entity if neeeded (and error if needed)

# server (Todo)

Create the server(s), apply middleware, ...

# app (Todo)

Connect to the database, setup repositories, services, controllers, start the server(s), the logger, the cache, ... This is what is called by main.go
