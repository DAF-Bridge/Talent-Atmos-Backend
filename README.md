# Talent-Atmos-Backend

## Generate Swagger API Document
```
swag init -g .\main.go -o ./docs --parseDependency --parseInternal
```

## Build the project
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
