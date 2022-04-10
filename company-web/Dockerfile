FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct


RUN echo "http://mirrors.ustc.edu.cn/alpine/v3.3/main/" > /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o ./server


FROM alpine
RUN echo "http://mirrors.ustc.edu.cn/alpine/v3.3/main/" > /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /build/server /app/server

CMD ["./server"]