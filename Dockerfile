# syntax=docker/dockerfile:1

FROM golang:1.22.3

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gopostgres

EXPOSE 8080

# Run
CMD ["/docker-gopostgres"]