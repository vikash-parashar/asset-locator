# Use the official Golang image as the base image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files into the container
COPY . .

# Build the Golang application
RUN go build -o asset-locator .

# Expose the port that your application will run on
EXPOSE 8080

# Command to run the executable
CMD ["./asset-locator"]
