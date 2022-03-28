package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "User3", City: "NewYork", Zipcode: "999001", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "1002", Name: "User2", City: "San Francisco", Zipcode: "777002", DateofBirth: "2000-02-01", Status: "2"},
	}
	return CustomerRepositoryStub{customers}
}
