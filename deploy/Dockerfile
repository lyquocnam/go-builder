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

# it will take the flags from the environment
RUN go build -o myapp

FROM alpine:latest as alpine
RUN apk --no-cache add tzdata zip ca-certificates
WORKDIR /usr/share/zoneinfo
# -0 means no compression.  Needed because go's
# tz loader doesn't handle compressed data.
RUN zip -r -0 /zoneinfo.zip .

FROM scratch
WORKDIR /app
# the test program:
COPY --from=builder /build /app/
# the timezone data:
ENV ZONEINFO=/zoneinfo.zip \
    TZ=Asia/Saigon

COPY --from=alpine /zoneinfo.zip /
# the tls certificates:
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["./myapp", "serve"]

ENTRYPOINT ["./myapp", "serve"]