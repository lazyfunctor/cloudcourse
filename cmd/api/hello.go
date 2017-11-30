package main

import "github.com/lazyfunctor/cloudcourse/app"

func helloApp(args []string) error {
	a := app.HelloApp{}
	addr := "0.0.0.0:4500"
	return a.Run(addr)
}
