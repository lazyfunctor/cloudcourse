package main

import (
	"github.com/lazyfunctor/cloudcourse/app"
)

func cpuloadAPI(args []string) error {
	a := app.App{}
	a.CostFactor = 11
	addr := "0.0.0.0:4500"
	return a.Run(addr)
}
