# Use a lightweight base image with the desired Go version
FROM golang:1.20-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Copy the application files
COPY main.go .
COPY internal /app/internal

# Build the GoLang app
RUN go build -o main .

# Use a scratch base image for the final image
FROM scratch

# Copy the built executable from the builder stage
COPY --from=builder /app/main /bin/app

# Command to run the application
CMD ["/bin/app"]
