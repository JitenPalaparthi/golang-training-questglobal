# Go

- When you install go what happen?
    /usr/local/go

bin     -> go binary and other binaries stored
pkg     -> standard compiled packages exists
src     -> go source code. Go standard library

## To run go application

```go run hello.go```

- when you run --> it compiles -> links -> generates binary -> runs that binary

- where is my binary when run go applicatio

```go run --work hello.go```

- How to build go application

```go build hello.go```

- How to give different file name using build

```go build -o app hello.go```

- differnece between build and run

- build compiles, links, and generates the binary
- run compiles, links , generates the binary (in a temp location) and execute that binary

- cross compilation in go
- To find which all os and architecture that go can compile to

- to verify build steps

```go build -x hello.go```

- to get list of GOOS/GOARCH

```go tool dist list```

- list of go env

```go env```

- cross compile to windows

```GOOS=windows GOARCH=amd64 go build -o app.exe hello.go```

- To compile and link

```go tool compile hello.go```

```go tool link -v -o hello hello.o```

- to view plan9/go assembly language

```go build -gcflags="-S" hello.go```

- to get object dump

```go tool objdump -s main.main hello```

- to findout the size of the binary

```du -sh hello```

- to get metadata of a binary in unix based systems

```file hello```

- to strip off the size of the binary

```go build -ldflags="-s -w" hello.go```

- go install : It does everything that go build does and copies the binary to gobin directory

```go install hello.go```

