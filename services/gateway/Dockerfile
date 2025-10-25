FROM golang:1.25

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENBABLED=0 GOOS=linux go build -o /gateway ./cmd/main.go

EXPOSE 8080

CMD ["/gateway"]