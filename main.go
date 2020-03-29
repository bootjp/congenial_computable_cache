package main

import (
	"context"
	"github.com/bootjp/congenial_computable_cache/ccc"
)

func main() {
	ctx := context.TODO()
	ccc.NewCCC().Run(ctx)
}
