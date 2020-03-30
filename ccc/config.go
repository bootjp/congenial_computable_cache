package ccc

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config []Spec

type Spec struct {
	Prefix string   `yaml:"prefix"`
	Method []string `yaml:"method"` // src/net/http/method.go
	Cache  Cache    `yaml:"cache"`
}

type Cache struct {
	Name      string     `yaml:"name"`
	Enable    bool       `yaml:"enable"`
	Condition *Condition `yaml:"condition"`
	Language  *Language  `yaml:"language"`
	SRC       string     `yaml:"src"`
}

type Condition string

type Language string

func NewConfig(fp string) (*Config, error) {
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	c := Config{}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
