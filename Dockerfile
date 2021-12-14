FROM golang:1.17.1-alpine
WORKDIR /app
COPY . .
RUN go mod download
EXPOSE 8080
RUN go build -o main .
CMD ["/app/main"]
