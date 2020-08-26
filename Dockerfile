FROM golang:1.14.6


WORKDIR /app
COPY . .

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

RUN go build github.com/ggdream/ngo

EXPOSE 80 8888
CMD ["./ngo", ":80", "./static"]