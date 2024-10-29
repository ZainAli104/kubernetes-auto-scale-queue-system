FROM golang:1.23-alpine

WORKDIR .

COPY . .

ENV GOPATH=/

RUN go get

RUN go build -o bin/fs

CMD [ "./bin/fs" ]