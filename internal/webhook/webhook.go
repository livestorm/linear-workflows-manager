package webhook

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type apiRouter struct {
	Mux *mux.Router
}

type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

// New returns an API structure.
func New() (*apiRouter, error) {
	router := &apiRouter{}
	router.initRoutes()
	return router, nil
}

func (a *apiRouter) WriteJson(rw http.ResponseWriter, code int, value interface{}) {
	if value == nil {
		value = map[string]string{}
	}
	b, err := json.Marshal(value)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(code)
	if _, err := rw.Write(b); err != nil {
		if err != http.ErrHandlerTimeout {
			log.Printf("Write HTTP response failed: %v", err)
		}
	}
}

func (a *apiRouter) addRoutes(mux *mux.Router, routes []Route) {
	for _, route := range routes {
		mux.Path(route.Path).Methods(route.Method).HandlerFunc(route.HandlerFunc)
	}
}

func (a *apiRouter) initRoutes() {
	mux := mux.NewRouter()
	a.addRoutes(mux, a.linearRoutes())
	a.Mux = mux
}
