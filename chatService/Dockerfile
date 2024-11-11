FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Update and add system packages && software dependencies
RUN apt-get update && \
    apt-get -y upgrade && \
    apt-get -y install curl make && \
    curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash && \
    apt-get update && \
    apt-get install migrate && \
    apt-get install -y protobuf-compiler && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
    go get ./... && \
    go mod tidy && \
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

RUN export PATH="$PATH:$(go env GOPATH)/bin"

EXPOSE 7000 50051

## Runing migrations
#RUN make migratedown && \
#    make migrateup

# Running application
CMD ["go", "run", "cmd/chatservice/main.go"]
