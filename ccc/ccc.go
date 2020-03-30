package ccc

import (
	"context"
	"github.com/robertkrimen/otto"
	"log"
	"net/http"
	"sync"
)

type CCC struct {
	mutex      *sync.RWMutex
	config     *Config
	httpServer *http.Server
	vmPool     []*otto.Otto // todo change to resource pool
}

func NewCCC() *CCC {
	var pool []*otto.Otto
	for i := 0; i > 10; i++ {
		pool = append(pool, otto.New())
	}

	return &CCC{
		mutex:      &sync.RWMutex{},
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
	c3.mutex.Lock()
	defer c3.mutex.Unlock()
	c3.config = conf

	return true, nil
}
