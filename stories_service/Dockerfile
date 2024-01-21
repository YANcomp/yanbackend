FROM golang:1.21.3-alpine AS builder

COPY . /github.com/YANcomp/yanbackend/stories_service/source/
WORKDIR /github.com/YANcomp/yanbackend/stories_service/source/

RUN go mod download
RUN go build -o ./bin/stories_service cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/YANcomp/yanbackend/stories_service/source/prod.env .
COPY --from=builder /github.com/YANcomp/yanbackend/stories_service/source/bin/stories_service .

CMD ["./stories_service"]