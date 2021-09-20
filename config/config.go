package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/smallfish/simpleyaml"
)

type Config struct {
	PathPrefix string
}

var config *Config

// GetLocalization with params
// lang: language (but only english in this repo)
// pathFormat: path in yaml file
// values: arguments for fmt.Sprintf
func (c *Config) GetLocalization(lang string, pathFormat string, values ...interface{}) string {
	b, err := c.ReadFile(lang)
	if err != nil {
		log.Fatal(err)
	}

	y, err := simpleyaml.NewYaml(b)
	if err != nil {
		log.Fatal(err)
	}

	pathArray := strings.Split(pathFormat, ".")
	format, err := c.Get(y, pathArray)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf(format, values...)
}

// ReadFile with dynamis language
func (c *Config) ReadFile(lang string) ([]byte, error) {
	return ioutil.ReadFile(c.PathPrefix + "config/localization/" + lang + ".yaml")
}

// Get string format from yaml
func (c *Config) Get(y *simpleyaml.Yaml, pathArray []string) (string, error) {
	yaml := y.Get(pathArray[0])
	remPathArray := pathArray[1:]

	if len(remPathArray) != 0 {
		return c.Get(yaml, remPathArray)
	}

	return yaml.String()
}

// GetLocalization with params
// lang: language (but only english in this repo)
// pathFormat: path in yaml file
// values: arguments for fmt.Sprintf
func GetLocalization(lang string, pathFormat string, values ...interface{}) string {
	return config.GetLocalization(lang, pathFormat, values...)
}

// Initialize method really helpful for unit testing
func Initialize(pathPrefix string) *Config {
	config = &Config{
		PathPrefix: pathPrefix,
	}

	return config
}
