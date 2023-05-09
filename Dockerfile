FROM golang:1.19

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY pkg /app/pkg
COPY cmd /app/cmd
COPY config /app/config

RUN go build -o main /app/cmd/chooser_bot/

EXPOSE 8080

CMD ["/app/main"]


