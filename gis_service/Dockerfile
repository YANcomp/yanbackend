FROM golang:1.21.3-alpine AS builder

COPY . /github.com/YANcomp/yanbackend/gis_service/source/
WORKDIR /github.com/YANcomp/yanbackend/gis_service/source/

RUN go mod download
RUN go build -o ./bin/gis_service cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/YANcomp/yanbackend/gis_service/source/bin/gis_service .

CMD ["./gis_service"]