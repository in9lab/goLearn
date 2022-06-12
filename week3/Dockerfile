FROM golang:1.18.2 AS builder

WORKDIR /build
RUN adduser -u 10001 -D app-runner

ENV GOPROXY https://goproxy.cn
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o gin-server .

FROM alpine:3.16.0 AS final

WORKDIR /app
COPY --from=builder /build/gin-server /app/

USER app-runner
ENTRYPOINT ["/app/gin-server"]