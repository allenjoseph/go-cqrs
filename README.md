# CQRS Pattern

Sample of CQRS Pattern following this [article](https://outcrawl.com/go-microservices-cqrs-docker)

## How to use

1. [Install Docker and Docker Compose](https://docs.docker.com/compose/install)
2. In a terminal:
```
cd go-cqrs
docker-compose up
```
### Command request
```
curl --header "Content-Type: application/json" --request POST --data '{"message":"woof woof!"}' http://localhost:8080/woofs
```

### Query request
```
curl --header "Content-Type: application/json" --request GET http://localhost:9090/woofs?offset=0&limit=10
```