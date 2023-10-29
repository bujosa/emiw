# Emiw
This is a simple project in go for testing a simple functions. 

In this project we can find unit testing. Also CI/CD with github actions for testing the project.

## How to use

### Install the project
```bash
git clone https://github.com/bujosa/emiw
cd emiw
```

### Run the project
```bash
$ go run main.go
```

### Run the tests
```bash
$ go test -v
```

### Run the tests with coverage
```bash
$ go test -coverprofile=coverage ./...
```

### Run the tests with coverage and html
```bash
$ go test -coverprofile=coverage ./... 
$ go tool cover -html=coverage
```


