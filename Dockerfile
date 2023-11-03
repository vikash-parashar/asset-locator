# Use an official Golang runtime as a parent image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy your application code into the container
COPY . .

# Build your application
RUN go build -o main .

# Expose the port your application will listen on
EXPOSE 8080

# Define the command to run your application
CMD ["./main"]
