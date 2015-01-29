Godo app Server - Example Go REST API
=====================================

# Usage

1. Make sure you have Go (1.3+) installed and $GOPATH setup
2. Install dependencies: `$ make tools && make deps`
3. Copy/edit sample config file from godo.conf.sample to godo.conf
4. Ensure Mongodb is running in the background
5. Compile the Godo API Server: `$ make build`
6. Boot it up! `$ ./bin/godo-app-server -config=godo.conf`
7. Create yourself a new user by POSTing a form to /signup:
    * `$ curl -d 'username=peter&password=shhhh' http://localhost:3333/signup`
8. Login:
    * Open your browser to `http://localhost:3333/login?username=peter&password=shhhh`
9. Use the app!
    * Point browser to: `http://localhost:3333/`

## Curl commands

### Sign up a new user

```
curl -d 'username=peter&password=shhhh' http://localhost:3333/signup
```

### Login

```
curl -d 'username=peter&password=shhhh' http://localhost:3333/login
```

### Get Todos list

```
curl -H "Authorization: BEARER jwt-token-here" http://localhost:3333/todos
```

### Create a new Todo

```
curl -H "Authorization: BEARER jwt-token-here" -d '{"title":"clean my dishes"}' http://localhost:3333/todos
```

### Fetch a Todo

```
curl -H "Authorization: BEARER jwt-token-here" http://localhost:3333/todos/:id
```

### Update a Todo

```
curl -H "Authorization: BEARER jwt-token-here" -X "PUT" -d '{"title":"mmmmhmmm"}' http://localhost:3333/todos/:id
```

### Deleting a Todo

```
curl -H "Authorization: BEARER jwt-token-here" -X "DELETE" http://localhost:3333/todos/:id
```

### Benchmarks with wrk

```
wrk -c 50 -t 5 -H "Authorization: BEARER jwt-token-here" http://localhost:3333/todos
```
