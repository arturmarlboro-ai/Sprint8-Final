FROM golang:1.21
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -t parcels:v1.0.0 .
CMD ["./parcels:v1.0.0"]