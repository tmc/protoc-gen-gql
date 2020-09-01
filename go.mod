module github.com/tmc/protoc-gen-graphql

go 1.14

require (
	github.com/99designs/gqlgen v0.12.2
	github.com/agnivade/levenshtein v1.1.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/golang/protobuf v1.4.1
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/lyft/protoc-gen-star v0.4.14
	github.com/matryer/moq v0.0.0-20200816112511-720d53e65d2f // indirect
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/stretchr/testify v1.5.1
	github.com/twitchtv/twirp v5.10.1+incompatible
	github.com/urfave/cli/v2 v2.2.0 // indirect
	github.com/vektah/dataloaden v0.3.0 // indirect
	github.com/vektah/gqlparser/v2 v2.0.1
	golang.org/x/tools v0.0.0-20200831203904-5a2aa26beb65 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/99designs/gqlgen => github.com/tmc/gqlgen v0.0.0-20200901050952-6383e6ad1368
