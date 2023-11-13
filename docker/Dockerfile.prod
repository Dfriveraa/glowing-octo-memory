FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN swag init -g ./app/cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o gateway ./app/cmd/main.go

FROM alpine

WORKDIR /app

RUN apk add --no-cache ca-certificates
COPY --from=build /app/gateway .


EXPOSE 8000