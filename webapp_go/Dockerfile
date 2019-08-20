## Build ##

FROM golang:alpine3.10 AS build
WORKDIR /home/app
COPY . .
ENV GO111MODULE=on
RUN apk --no-cache add git make build-base \
	&& make \
	&& go get github.com/labstack/echo/v4

## App ##

FROM alpine:3.9.4
WORKDIR /home/app
RUN apk --no-cache add tzdata && \
	cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
COPY --from=build /home/app/main /home/app
EXPOSE 80
ENTRYPOINT ["./main"]
CMD []
