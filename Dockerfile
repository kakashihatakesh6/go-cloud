#||* BUILD STAGE *||#
# Step 1: Use the official Golang image to build the Go application
FROM golang:1.23.1-alpine AS builder

# Step 2: Set the Current Working Directory inside the container
WORKDIR /app

# Step 3: Copy the source code into the container
COPY . .

# Step 4: Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Step 6: Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server cmd/main.go


#||* FINAL STAGE *||#
# Step 7: Use a smaller base image to run the compiled Go app
FROM alpine:latest

# Step 8: Set the Working Directory
WORKDIR /root/

# Step 9: Copy the compiled binary from the builder image
COPY --from=builder /app/server .

COPY .env .

# Step 11: Run the application
CMD ["./server"]