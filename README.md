# SWAPI GraphQL API in Go

### How to run it

```go
go run ./server/server.go
```

Server: http://localhost:8088/

### Available Queries (more to come)

```graphql

# All people
{
  people {
    name
    ...
  }
}

# All planets
{
  planets {
    name
    ...
  }
}

# Person by id
{
  person(id: 1) {
    name
    homeworld {
      name
      ...
    }
    ...
  }
}

# Planet by id
{
  planet(id: 1) {
    name
    ...
  }
}
```