FROM golang:1.20 AS builder

RUN cat /etc/issue
COPY . /opt/

WORKDIR /opt
RUN export GO111MODULE=on
RUN export GOPROXY=https://goproxy.io
RUN make

FROM ubuntu:latest

COPY --from=builder /opt/x-project /bin/

RUN mkdir /etc/x-project/
COPY config/config.yaml /etc/x-project/config.yaml
WORKDIR /

ENTRYPOINT ["/bin/x-project"]
CMD ["-c", "/etc/x-project/config.yaml", "start"]
