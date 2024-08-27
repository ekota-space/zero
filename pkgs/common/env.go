package common

import (
	"log"

	"github.com/Netflix/go-env"
)

type environment struct {
	Port   int `env:"PORT,default=8080"`
	Extras env.EnvSet
}

var Env environment

func SetupEnvironmentVars() {
	es, error := env.UnmarshalFromEnviron(&Env)
	if error != nil {
		log.Fatal(error)
	}
	Env.Extras = es
}
