package e2e

//go:generate protoc -I . -I /usr/local/include -I ../options --go_out=. service.proto
//go:generate protoc -I . -I /usr/local/include -I ../options --graphql_out=. service.proto
//go:generate protoc -I . -I /usr/local/include -I ../options --graphql_out=gqlgen=false,dest=withoutgqlgen:. service.proto
