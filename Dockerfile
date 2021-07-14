FROM node:lts-alpine as node-builder

WORKDIR /build

COPY . /build

RUN cd /build/assets/frontend \
    && yarn && yarn build

FROM golang:alpine as go-builder

WORKDIR /build

COPY --from=node-builder /build /build

RUN go build -o /build/golist

FROM alpine:latest

COPY --from=go-builder /build/golist /app/golist

VOLUME [ "/data" ]

ENTRYPOINT [ "/app/golist" ]

CMD ["-p","8080","-s","/data"]