package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/golang/api/handler"
	"github.com/gorilla/mux"
)

//App função servidor
type App struct {
	Router *mux.Router
	Db     *sql.DB
}

//StartServer função servidor
func (a *App) StartServer() {
	a.Router = mux.NewRouter()

	s := a.Router.PathPrefix("/api/v1").Subrouter()

	s.HandleFunc("/health", handler.HealthChecker).Methods(http.MethodGet)

	s.HandleFunc("/clientes", handler.GravaCliente).Methods(http.MethodPost)
	s.HandleFunc("/update", handler.UpdateCliente).Methods(http.MethodPost)
	s.HandleFunc("/update/{id}", handler.UpdateCliente).Methods(http.MethodPut)
	s.HandleFunc("/clientes", handler.GetClientes).Methods(http.MethodGet)
	s.HandleFunc("/clientes/{id}", handler.DeleteCliente).Methods(http.MethodDelete)
	s.HandleFunc("/clientes/{id}", handler.GetCliente).Methods(http.MethodGet)
	s.HandleFunc("/clientes", handler.GetClientes).Methods(http.MethodGet)

	a.Router.Handle("/api/v1/{_:.*}", a.Router)
	port := 8091
	log.Println("Starting server on port : ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router))
}
