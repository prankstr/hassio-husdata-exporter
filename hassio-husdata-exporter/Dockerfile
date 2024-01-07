ARG BUILD_FROM
FROM docker.io/golang:1.21-rc-alpine AS build

WORKDIR /app
COPY . /app

RUN GO111MODULE=on CGO_ENABLED=0 go build -o husdata_exporter

##
FROM $BUILD_FROM

COPY --from=build /app/husdata_exporter .
COPY run.sh /
RUN chmod a+x /run.sh

CMD [ "/run.sh" ]
