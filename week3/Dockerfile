FROM golang:1.18.2 AS builder

WORKDIR /build

ENV GOPROXY https://mirrors.aliyun.com/goproxy/
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o gin-server .

FROM alpine:3.16.0 AS final

WORKDIR /app
COPY --from=builder /build/gin-server /app/

ENTRYPOINT ["/app/gin-server"]