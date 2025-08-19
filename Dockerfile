FROM golang:1.21-alpine

WORKDIR /app

COPY . .

EXPOSE 25565

CMD ["go", "run", "src/main.go"]