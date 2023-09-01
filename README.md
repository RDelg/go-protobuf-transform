# Description
Go script that given a protobuf definition and a list of paths inside the definition transforms matching paths given a transformation function.

# Usage
```go
protoc A.proto --go_out=.
go run main.go
```