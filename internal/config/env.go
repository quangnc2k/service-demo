package config

import (
	"log"

	"git.cyradar.com/phinc/my-awesome-project/pkg/util"
)

var Env = &env

var env Environment

type Environment struct {
	MongoURL string `envconfig:"mongo_url"`
}

func Init() *Environment{
	var environment Environment
	err := util.ReadEnvVars(&environment)
	if err != nil {
		log.Fatalln("Unable to read env variables: ", err)
	}

	log.Println("Env read", environment)

	return &environment
}
