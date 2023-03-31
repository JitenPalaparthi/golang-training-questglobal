# Swagger how to

- install swag tool

```go install github.com/swaggo/swag/cmd/swag@latest```

- to run swggerui in docker

```docker run -p 80:8080 -e SWAGGER_JSON=/foo/swagger.json -v /home/jiten/workspace/personal/training/stackroute/golang-training-questglobal/67-swagger/docs:/foo swaggerapi/swagger-ui```