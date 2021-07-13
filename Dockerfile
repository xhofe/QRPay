FROM alpine:latest as builder
LABEL stage=go-builder
WORKDIR /app/
COPY ./ ./
RUN apk add --no-cache go gcc musl-dev; \
      go build -o ./bin/app -ldflags '-s -w' ./cmd/main.go ./cmd/init.go

FROM alpine:latest
LABEL MAINTAINER="i@nn.ci"
WORKDIR /opt/app/
COPY --from=builder /app/bin/app ./
EXPOSE 5244
CMD [ "./app" ]