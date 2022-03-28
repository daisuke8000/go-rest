package app

import (
	"github.com/daisuke8000/go-rest/domain"
	"github.com/daisuke8000/go-rest/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	// define route
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	//router.HandleFunc("/time", getCurrentTime).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
