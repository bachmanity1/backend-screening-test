FROM golang:latest
WORKDIR /terra
ADD . /terra
RUN make build
ENTRYPOINT ["bin/terra"]
