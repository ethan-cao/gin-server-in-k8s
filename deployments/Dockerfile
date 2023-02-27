# Set base image
FROM golang:1.20-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

# Download dependencies. Dependencies are cached if the go.mod and go.sum files remain unchanged
RUN go mod download

# Copy source from the current directory to the Working Directory inside the container
COPY . .
COPY *.go ./

# Build the Go app
RUN go build -o /main

# Expose port 8080 to the outside world
EXPOSE 8080

# run the executable
CMD [ "/main" ]