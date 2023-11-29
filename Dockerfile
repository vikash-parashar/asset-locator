# Stage 1: Build the Go app
FROM golang:1.21 AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Create the final image
FROM postgres:latest

# Set the working directory to /app
WORKDIR /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/main .

# Copy the .env file to set environment variables for PostgreSQL
COPY .env .

# Expose port 8080 to the outside world (for the Go app)
EXPOSE 8080

# Command to run the Go executable
CMD ["./main"]
