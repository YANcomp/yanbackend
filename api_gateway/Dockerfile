FROM golang:1.21.3-alpine AS builder

COPY . /github.com/YANcomp/yanbackend/api_gateway/source/
WORKDIR /github.com/YANcomp/yanbackend/api_gateway/source/

RUN go mod download
RUN go build -o ./bin/api_gateway cmd/app/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/YANcomp/yanbackend/api_gateway/source/bin/api_gateway .

CMD ["./api_gateway"]