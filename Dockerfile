FROM golang:1.15-alpine AS build

WORKDIR /go/src/github.com/aeekayy/go-api-base

ENV GOPATH=/go

RUN apk update && apk add -U build-base git curl libstdc++ ca-certificates nodejs npm python3 fftw-dev gcc g++ make libc6-compat make pkgconfig glib-dev poppler-dev python2 \
	&& apk add vips vips-dev --repository=http://dl-cdn.alpinelinux.org/alpine/edge/community

ADD . /go/src/github.com/aeekayy/go-api-base

RUN make build-docker-binary

FROM alpine:latest
RUN apk update && apk add curl bash
WORKDIR /app
COPY --from=build /go/src/github.com/aeekayy/go-api-base/bin/go-api-base /app/
COPY --from=build /go/src/github.com/aeekayy/go-api-base/build-manifest.yml /app/

EXPOSE 8080/tcp

# Adding nsswitch.conf
# https://github.com/golang/go/issues/35305
RUN echo "hosts:          files dns" > /etc/nsswitch.conf

ENV PATH="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
WORKDIR /app

ENV BUILD_MANIFEST /app/build-manifest.yml

ENTRYPOINT ["/app/go-api-base"]
