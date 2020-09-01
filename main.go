package main

import (
	"io/ioutil"
	"log"
	"os"

	pgs "github.com/lyft/protoc-gen-star"
	"github.com/tmc/protoc-gen-graphql/gengraphql"
	"github.com/tmc/protoc-gen-graphql/internal/gocopy/modfile"
)

func main() {
	modname := getImportPath()
	log.SetOutput(ioutil.Discard)
	pgs.Init(pgs.DebugEnv("DEBUG")).
		RegisterModule(gengraphql.New(modname)).
		Render()
}

func getImportPath() string {
	bts, err := ioutil.ReadFile("go.mod")
	if os.IsNotExist(err) {
		return ""
	} else {
		must(err)
	}
	modf, err := modfile.Parse("go.mod", bts, nil)
	must(err)
	return modf.Module.Mod.Path
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
