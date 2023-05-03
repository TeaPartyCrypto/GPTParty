# Use the official Golang image as the base image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files into the container
COPY go.mod go.sum ./

# Download and cache the dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the program
RUN cd cmd && go build -o ../main

# Expose the necessary port for the Discord bot (optional)
EXPOSE 8080

# Run the program
CMD ["./main"]
