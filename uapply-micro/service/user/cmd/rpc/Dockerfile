FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY service/user/cmd/rpc/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/user service/user/cmd/rpc/user.go


FROM alpine:3.7

RUN echo -e https://mirrors.ustc.edu.cn/alpine/v3.7/main/ >> /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/user /app/user
COPY --from=builder /app/etc /app/etc

CMD ["./user", "-f", "etc/user.yaml"]
