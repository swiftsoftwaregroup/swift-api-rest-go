# Start from the latest golang base image
FROM golang:1.21-alpine

# will install the gcc, musl-dev and binutils packages (amongst others).
RUN apk add build-base

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Generate swagger documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --output ./swag

# Build the Go app
RUN CGO_ENABLED=1 go build -o main .

# Expose port 8001 to the outside world
EXPOSE 8001

# Command to run the executable
CMD ["./main"]