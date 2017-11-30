package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// HelloApp struct
type HelloApp struct {
	Router     *mux.Router
	CostFactor int
}

func (a *HelloApp) initializeRoutes() {
	a.Router.HandleFunc("/", a.healthCheck).Methods("GET")
}

func (a *HelloApp) healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Test</h1><div>Hello world!!!</div>")

}

// Run function is the entry point
func (a *HelloApp) Run(addr string) error {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
	return http.ListenAndServe(addr, a.Router)
}
