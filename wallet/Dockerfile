#Stage 1 - Build Stage
FROM golang:1.22.5-alpine AS build

#Install necessary packages
RUN apk add --no-cache git make protobuf

#Set destination for Copy
WORKDIR /app

#Copy go mod and download dependencies
COPY go.mod go.sum ./
RUN go mod download

#Copy resources code
COPY . .

#Install golangci-lint
# RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

#Run linter with increased time out
# RUN golangci-lint run --timeout 5m

#Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o ./go-crud-wallet ./wallet/main.go

#Stage 2 - Run stage
FROM alpine:edge

#Set the working directory
WORKDIR /app

#Copy binary from build stage
COPY --from=build /app/go-crud-wallet .

#Install CA certificates
RUN apk --no-cache add ca-certificates

#Run
CMD [ "./go-crud-wallet" ]