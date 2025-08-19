FROM golang:1.21-alpine

WORKDIR /app

COPY . .

EXPOSE 25565
EXPOSE 8080

CMD ["go", "run", "src/main.go"]