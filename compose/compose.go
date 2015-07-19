// Package compose reads docker-compose yaml files and generates architecture json output
package compose

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// Compose Attribute maps to attributes of a microservice
type ComposeAttributes struct {
	Build string `yaml:"build,omitempty"`
	Image string `yaml:"image,omitempty"`
	Links []string `yaml:"links,omitempty"`
}

// Compose type to extract interesting data from compose yaml
type ComposeYaml map[string]ComposeAttributes

// ReadCompose
func ReadCompose(compose string) ComposeYaml {
	fn := "compose_yaml/" + compose + ".yml"
	log.Println("Loading compose yaml from " + fn)
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	var c ComposeYaml
	e := yaml.Unmarshal(data, &c)
	if e == nil {
		return c
	} else {
		log.Fatal(e)
		return nil
	}
}