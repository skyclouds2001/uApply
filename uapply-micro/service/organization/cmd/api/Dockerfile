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
COPY service/organization/cmd/api/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/org service/organization/cmd/api/org.go


FROM alpine:3.7

RUN echo -e https://mirrors.ustc.edu.cn/alpine/v3.7/main/ >> /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/org /app/org
COPY --from=builder /app/etc /app/etc

CMD ["./org", "-f", "etc/org-api.yaml"]
