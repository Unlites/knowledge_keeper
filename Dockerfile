FROM golang:1.19.5-buster as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /knowledge_keeper ./cmd/knowledge_keeper/main.go

FROM scratch

COPY --from=builder knowledge_keeper /bin/knowledge_keeper
COPY --from=builder /app/migrations /migrations

CMD ["/bin/knowledge_keeper"]
