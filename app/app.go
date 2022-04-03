package app

import (
	"log"
	"net/http"

	"github.com/daisuke8000/go-rest/domain"
	"github.com/daisuke8000/go-rest/service"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	//wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	// define route
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	//router.HandleFunc("/time", getCurrentTime).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
