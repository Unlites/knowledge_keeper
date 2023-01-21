FROM golang:1.19.5-buster

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /bin/knowledge_keeper ./cmd/knowledge_keeper/main.go

CMD ["/bin/knowledge_keeper"]

# CMD ["go", "run", "./cmd/knowledge_keeper/main.go"]