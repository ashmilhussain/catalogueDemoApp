# Stage 1: Build the Go application

FROM golang:1.19 AS builder

WORKDIR /backend
# Copy the entire application code into the container

COPY . .
RUN go get

RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM golang:1.19-alpine as deployer
WORKDIR /backend
ENV PATH="/backend/:${PATH}"
COPY --from=builder /backend/app app
EXPOSE 8080
CMD ["./app"]