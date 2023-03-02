- to run all test 

```go test ./...```


- to run a test with in given time

```go test -timeout=5s```

- to run a test based on pattern

```go test -timeout 30s -run ^TestArea$ shapes/shape/rect```

- to run all tests in a file 

```go test -timeout 30s -run ^(TestArea|TestPerimeter)$ shapes/shape/rect```

```go test shapes/shape/rect -v```

- to check test result along with coverage

```go test -coverprofile=coverage.out ./...```



