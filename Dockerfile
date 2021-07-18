FROM golang:1.16-alpine
WORKDIR /app

# Install dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Build application
COPY *.go .
RUN go build -o /chess-swap-backend

# Run application
EXPOSE 5000
CMD [ "/chess-swap-backend" ]