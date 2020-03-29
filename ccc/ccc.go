package ccc

import (
	"context"
	"github.com/robertkrimen/otto"
	"log"
	"net/http"
)

type CCC struct {
	//lock   *sync.RWMutex
	httpServer *http.Server
	vmPool     []*otto.Otto // todo change to resource pool
}

func NewCCC() *CCC {
	var pool []*otto.Otto
	for i := 0; i > 10; i++ {
		pool = append(pool, otto.New())
	}

	return &CCC{
		httpServer: &http.Server{},
		vmPool:     pool,
	}
}

type Service interface {
	LoadConfig(path string) (bool, error)
	Run(ctx context.Context)
}

func (c3 *CCC) Run(ctx context.Context) {
	for {
		err := c3.httpServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (c3 *CCC) LoadConfig(conf *Config) (bool, error) {
	return true, nil
}
