FROM golang AS build-env

MAINTAINER yanorei32

COPY ./src /work/

RUN	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build \
		-ldflags "-s -w" \
		-a \
		-installsuffix cgo \
		-o /work/app \
		/work/main.go

FROM alpine

COPY --from=build-env /work/app /usr/local/bin/app

ENTRYPOINT ["/usr/local/bin/app"]

