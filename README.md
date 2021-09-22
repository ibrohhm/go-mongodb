# Go Mongo

## Description
personal repo about go using mongo as database and include with middleware, localization and unit test

## Setup
- Install mongodb on local env and create `go_mongo_learn` database, in this repo no need username and password
- Running the project
  ```
  make run
  ```
- Unit testing
  ```
  make test
  ```

## Endpoints
- `get: /healtzh` - for sanity check
- `get: /products` - retrieve all products
- `get: /products/:id` - retrieve specific product given id
- `post: /products` - create new products
- `put: /products/:id` - update specific products given id
- `delete: /products/:id` - delete specific product given id