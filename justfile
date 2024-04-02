set positional-arguments

fmt:
  go fmt ./...

run:
  go run main.go

@test run:
  go test ./... -v -run $1
