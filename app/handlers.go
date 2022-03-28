package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/daisuke8000/go-rest/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"xcity"`
	Zipcode string `json:"zip_code" xml:"zip"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{Name: "User1", City: "New York", Zipcode: "110075"},
	//	{Name: "User2", City: "San Francisco", Zipcode: "110088"},
	//}

	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "app/xml" {
		w.Header().Add("Content-Type", "app/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "app/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "request post render")
}

//// currentTimeAPI
//func getCurrentTime(w http.ResponseWriter, r *http.Request) {
//	n := time.Now()
//	// request json
//	curtime := struct {
//		RequestTime string `json:"current_time"`
//	}{
//		RequestTime: n.Format(time.RFC3339),
//	}
//	w.Header().Add("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(curtime)
//	// plan
//	//fmt.Fprint(w, n.Format(time.RFC3339))
//}
