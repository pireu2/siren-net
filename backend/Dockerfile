# Use the official Go image as the base image
FROM golang:1.24

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the backend code
COPY . .

# Copy the .env file
COPY .env .env

# Build the application
RUN CGO_ENABLED=1 go build -o main .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]