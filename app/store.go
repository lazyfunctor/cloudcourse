package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// StoreApp struct
type StoreApp struct {
	Router     *mux.Router
	CostFactor int
}

func (a *StoreApp) initializeRoutes() {
	a.Router.HandleFunc("/state/", a.saveStateOptions).Methods("OPTIONS")
	a.Router.HandleFunc("/state/", a.saveState).Methods("POST")
	a.Router.HandleFunc("/state/", a.getState).Methods("GET")
}

func (a *StoreApp) saveStateOptions(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.WriteHeader(http.StatusOK)
}

func (a *StoreApp) saveState(w http.ResponseWriter, r *http.Request) {
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			fmt.Println(fmt.Sprintf("%v: %v", name, h))
		}
	}
	b, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(b))
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = ioutil.WriteFile("state.json", b, 0644)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.WriteHeader(200)
}

func (a *StoreApp) getState(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("state.json")
	fmt.Println(string(data))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// w.WriteHeader(200)
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Write(data)
}

// Run function is the entry point
func (a *StoreApp) Run(addr string) error {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
	handler := cors.AllowAll().Handler(a.Router)
	return http.ListenAndServe(addr, handler)
}
