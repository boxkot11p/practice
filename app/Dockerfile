FROM golang:1.20

WORKDIR /workspace
COPY . ./
RUN go mod download

EXPOSE 9000
CMD go run cmd/main.go