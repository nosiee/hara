FROM golang:1.18-alpine

COPY . .
WORKDIR server/

RUN apk add --no-cache musl-dev gcc make linux-headers pkgconfig imagemagick imagemagick-dev ffmpeg bash file
RUN go get ./...

RUN make build
