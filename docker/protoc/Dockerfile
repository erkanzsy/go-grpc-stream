FROM golang:1.21-alpine3.18

RUN apk add protobuf
RUN apk add --no-cache protobuf protobuf-dev
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3

WORKDIR /application

CMD ["tail", "-f", "/dev/null"]