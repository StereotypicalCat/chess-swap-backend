FROM golang:1.16-alpine
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go .
RUN go build -o /chess-swap-backend

EXPOSE 5000
CMD [ "/chess-swap-backend" ]