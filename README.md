# Talent-Atmos-Backend

## Installation Dependencies
Ensure you have [Go](https://go.dev/doc/install) installed.
After you cloning the project you have to install the package/library that will be using in the project from this command.
```
go mod tidy
```

## Environment Variables
>> Create a .env file in the root directory and configure it based on .env.example

## Running the project
```
go run main.go
```
### OR
If you want hot reload you can install air from this package: 
```
go install github.com/air-verse/air@latest
```
Then run this command for building the project
```
air
```

## Generate Swagger API Document
```
swag init -g .\main.go -o ./docs --parseDependency --parseInternal
```

## Docker
Building the docker image
```
docker build -t DOCKER_USERNAME/IMAGE_NAME .
```

Running the container
```
docker run --env-file .<path to your env file> -p "<exposed port>:8080"  IMAGE_NAME:TAG
```
