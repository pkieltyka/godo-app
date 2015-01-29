Godo app Server - Example Go REST API
=====================================

# Curl commands

## Sign up a new user

```
curl -d 'username=peter&password=shhhh' http://localhost:3333/signup
```

## Login

```
curl -d 'username=peter&password=shhhh' http://localhost:3333/login
```

## Get Todos list

```
curl -H "Authorization: BEARER jwt-token-here" http://localhost:3333/todos
```

## Create a new Todo

```
curl -H "Authorization: BEARER jwt-token-here" -d '{"title":"clean my dishes"}' http://localhost:3333/todos
```

## Fetch a Todo

```
curl .........
```

## Update a Todo

```
curl -H "Authorization: BEARER jwt-token-here" -X "PUT" -d '{"user_id":"abc", "title":"mmmmhmmm"}' http://localhost:3333/todos/:id
```

## Deleting a Todo

```
...
-X "DELETE" http://localhost:3333/todos/:id
```

## Benchmarks with wrk

```
wrk -c 50 -t 5 -H "Authorization: BEARER jwt-token-here" http://localhost:3333/todos
```
