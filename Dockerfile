FROM golang:1.23
WORKDIR /app
COPY . .
RUN go build -o risk-service
CMD ["./risk-service"]
EXPOSE 8080
