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

	/*
		prefix: /api/
		    method:
		      - GET
		      - HEAD
		      - OPTION
		    cache:
		      name: "api_cache"   # spec is snake_case and ascii
		      enable: yes         # no
		      condition: time     # or compute
		      langage: javascript # or lua
		      src: ../script/api.js
	*/
}

type Condition string

type Language string

type Cache struct {
	Name      string
	Enable    bool
	Condition *Condition
	Language  *Language
	SRC       string
}

func NewConfig(fp string) (*Config, error) {
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal([]byte(data), c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
