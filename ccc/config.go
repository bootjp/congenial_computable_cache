package ccc

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Specs []*Spec
}

type Spec struct {
	prefix string
	method []string // src/net/http/method.go
	Cache  *Cache
}

type Cache struct {
	Name      string
	Enable    bool
	Condition *Condition
	Language  *Language
	SRC       string
}

type Condition string

type Language string

func NewConfig(fp string) (*Config, error) {
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
