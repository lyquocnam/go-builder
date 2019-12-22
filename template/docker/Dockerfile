FROM golang:1.13.5 as builder

ARG GITHUB_TOKEN

RUN git config --global url."https://${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"

WORKDIR /build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

RUN go build -o app


FROM alpine:latest as alpine
RUN apk --no-cache add tzdata zip ca-certificates
WORKDIR /usr/share/zoneinfo
RUN zip -r -0 /zoneinfo.zip .

FROM scratch
WORKDIR /app

ENV ZONEINFO=/zoneinfo.zip \
    TZ=Asia/Saigon

COPY --from=builder /build /app/
COPY --from=alpine /zoneinfo.zip /
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./app", "serve"]