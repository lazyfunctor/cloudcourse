package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// App struct
type App struct {
	Router     *mux.Router
	CostFactor int
}

//Event struct
type Event struct {
	Data string `json:"data"`
}

//Response struct
type Response struct {
	Value        string `json:"value"`
	ErrorMessage string `json:"error"`
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/event/", a.createEvent).Methods("POST")
	a.Router.HandleFunc("/", a.healthCheck).Methods("GET")
}

func (a *App) createEvent(w http.ResponseWriter, r *http.Request) {
	var ev Event
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ev); err != nil {
		err = errors.Wrap(err, "JSON unmarshal error")
		fmt.Printf("%+v\n", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(ev.Data), a.CostFactor)
	if err != nil {
		fmt.Printf("%+v\n", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	resp := &Response{string(hash), ""}
	respondWithJSON(w, 200, resp)

}

func (a *App) healthCheck(w http.ResponseWriter, r *http.Request) {
	resp := &Response{"", ""}
	respondWithJSON(w, 200, resp)
}

// Run function is the entry point
func (a *App) Run(addr string) error {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
	return http.ListenAndServe(addr, a.Router)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, &Response{"", message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.WriteHeader(code)
	w.Write(response)
}
