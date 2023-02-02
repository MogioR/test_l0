package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	PSQLPort    int32  `yaml:"psql_port"`
	PSQLHost    string `yaml:"psql_host"`
	PSQLUser    string `yaml:"psql_user"`
	PSQLPass    string `yaml:"psql_pass"`
	PSQLName    string `yaml:"psql_name"`
	NutsCluster string `yaml:"nuts_cluster"`
	NutsClient  string `yaml:"nuts_client"`
	NutsChenel  string `yaml:"nuts_chanel"`
	NutsHost    string `yaml:"nuts_host"`
}

func (c *Config) LoadConfig(source string) (err error) {
	yamlFile, err := ioutil.ReadFile(source)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)

	return
}
