# ✈️ Backend - Voucher Seat Assignment API

## Requirements

- **Golang**: Make sure you have Go installed. [Download Go](https://golang.org/dl/)
- **Mockgen**: Used for generating mocks. Install it using:
  ```bash
  go install github.com/golang/mock/mockgen@v1.6.0
  ```
- **Gin**: A web framework for Go. Install it using:
  ```bash
  go get -u github.com/gin-gonic/gin
  ```
- **SQLite**: Used as the database for storing vouchers. Install the driver via:
  ```bash
  go get -u gorm.io/driver/sqlite
  go get -u gorm.io/gorm
  ```

## Database Migration

- Use the `Makefile` to handle database migrations. Run the following command:
  ```bash
  make migrate
  ```

## Installation

1. **Install Golang**:
   - Follow the instructions on the [official Go website](https://golang.org/doc/install).

2. **Setup Configuration**:
   - Copy the sample environment file:
     ```bash
     cp sample.env .env
     ```
   - Update the `.env` file with your configuration.
   - Install dependencies:
    ```bash
    go mod tidy
    ```

3. **Run the Application**:
   - Run your application:
     ```bash
     go run main.go
     ```
     you can access the API with localhost:8080

## Generating Mocks

- To generate mocks for a service, use the following command:
  ```bash
  mockgen -source=service/voucher_service.go -destination=tests/mock/service/voucher_service.go
  ```
- Replace `service/voucher_service.go` with the path to the file you want to mock.

## Unit Test

- To run the unit test and check the coverage, use the following command:
  ```bash
  make unit-test
  ```