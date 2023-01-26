package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	PSQLPort int32  `yaml:"PSQL.psql_port"`
	PSQLHost string `yaml:"PSQL.psql_host"`
	PSQLUser string `yaml:"PSQL.psql_user"`
	PSQLPass string `yaml:"PSQL.psql_pass"`
	PSQLName string `yaml:"PSQL.psql_name"`
}

func (c *Config) LoadConfig(source string) (err error) {
	yamlFile, err := ioutil.ReadFile(source)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)

	return
}
