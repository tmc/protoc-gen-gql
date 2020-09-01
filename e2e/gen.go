package e2e

//go:generate protoc -I . -I /usr/local/include -I ../options --go_out=. service.proto
//go:generate protoc -I . -I /usr/local/include -I ../options --gql_out=. service.proto
//go:generate protoc -I . -I /usr/local/include -I ../options --gql_out=gqlgen=false,dest=withoutgqlgen:. service.proto
