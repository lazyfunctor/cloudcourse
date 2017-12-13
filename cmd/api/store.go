package main

import "github.com/lazyfunctor/cloudcourse/app"

func storeAPI(args []string) error {
	a := app.StoreApp{}
	addr := "0.0.0.0:4500"
	return a.Run(addr)
}
