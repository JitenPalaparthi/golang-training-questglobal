FROM golang:alpine  AS build

WORKDIR /app

COPY . .

RUN go mod tidy && \
go build -o app main.go

FROM scratch

COPY --from=build app .

EXPOSE 8080

ENTRYPOINT [ "./app" ]

