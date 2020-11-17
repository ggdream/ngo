FROM golang:1.15.5


WORKDIR /app
COPY . .

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

RUN go install

EXPOSE 54639
CMD ["ngo"]