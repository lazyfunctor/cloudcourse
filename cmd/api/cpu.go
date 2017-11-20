package main

import (
	"strconv"

	"github.com/lazyfunctor/cloudcourse/app"
)

func cpuloadAPI(args []string) error {
	a := app.App{}
	costFactor, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	a.CostFactor = costFactor
	addr := "0.0.0.0:4500"
	return a.Run(addr)
}
