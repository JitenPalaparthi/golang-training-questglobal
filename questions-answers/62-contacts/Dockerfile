FROM golang:alpine AS builder
#golang and alpine based image

ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/app
COPY . .
#apk is simialr to apt or apt-get in ubuntu
# in fedora it is called yum
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev git
RUN go mod tidy
RUN go build -ldflags="-s -w" -o app main.go

FROM scratch

WORKDIR /app

COPY --from=builder /go/src/app/app .

ENV PORT=8080
ARG type=dev
EXPOSE ${PORT}

ENTRYPOINT [ "./app" ]