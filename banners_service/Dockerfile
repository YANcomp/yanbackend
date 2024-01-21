FROM golang:1.21.3-alpine AS builder

COPY . /github.com/YANcomp/yanbackend/banners_service/source/
WORKDIR /github.com/YANcomp/yanbackend/banners_service/source/

RUN go mod download
RUN go build -o ./bin/banners_service cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/YANcomp/yanbackend/banners_service/source/bin/banners_service .

CMD ["./banners_service"]