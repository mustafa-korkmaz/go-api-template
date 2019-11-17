package main

import (
	"os"

	"github.com/mustafa-korkmaz/goapitemplate/pkg/api"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/utl/config"
)

func main() {

	cfg, err := config.Load(getConfigPath())

	checkErr(err)

	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func getConfigPath() string {

	dir, err := os.Getwd()

	if err != nil {
		return ""
	}

	path := dir + "./cmd/api/config/conf."
	args := os.Args

	if len(args) > 1 {
		if args[1] == "prod" {
			return path + "prod.yaml"
		}
	}

	return path + "dev.yaml"
}
