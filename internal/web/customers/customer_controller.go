package customers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bigtalljosh/gostrife/internal/domain/interfaces"
	"github.com/Bigtalljosh/gostrife/internal/entity"
	"github.com/gorilla/mux"
)

type CustomerController struct {
	service interfaces.CustomerService
}

func NewCustomerController(service interfaces.CustomerService) *CustomerController {
	return &CustomerController{service}
}

func (c *CustomerController) RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/api/customers", c.GetCustomers).Methods("GET")
    router.HandleFunc("/api/customers/{id}", c.GetCustomer).Methods("GET")
    router.HandleFunc("/api/customers", c.CreateCustomer).Methods("POST")
    router.HandleFunc("/api/customers/{id}", c.UpdateCustomer).Methods("PUT")
    router.HandleFunc("/api/customers/{id}", c.DeleteCustomer).Methods("DELETE")
}

func (c *CustomerController) CreateCustomer(w http.ResponseWriter, r *http.Request) {
    var customer entity.Customer
    _ = json.NewDecoder(r.Body).Decode(&customer)
    createdCustomer, err := c.service.Create(customer)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(createdCustomer)
}

func (c *CustomerController) GetCustomers(w http.ResponseWriter, r *http.Request) {
    customers, err := c.service.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(customers)
}

func (c *CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    customer, err := c.service.GetByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(customer)
}

func (c *CustomerController) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
    var customer entity.Customer
    _ = json.NewDecoder(r.Body).Decode(&customer)
    updatedCustomer, err := c.service.Update(customer)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(updatedCustomer)
}

func (c *CustomerController) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    err := c.service.Delete(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
