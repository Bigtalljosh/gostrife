package services

import (
	"fmt"

	"github.com/Bigtalljosh/gostrife/internal/entity"
)

// For now I have no Db set up so I'm using a collection in memory
type InMemoryCustomerService struct {
    customers []entity.Customer
}

func (db *InMemoryCustomerService) GetAll() ([]entity.Customer, error) {
	return db.customers, nil
}

func (db *InMemoryCustomerService) GetByID(id int) (entity.Customer, error) {
    for _, customer := range db.customers {
        if customer.ID == id {
            return customer, nil
        }
    }
    return entity.Customer{}, fmt.Errorf("customer not found")
}

func (db *InMemoryCustomerService) Create(customer entity.Customer) (entity.Customer, error) {
    customer.ID = len(db.customers) + 1
    db.customers = append(db.customers, customer)
    return customer, nil
}

func (db *InMemoryCustomerService) Update(customer entity.Customer) (entity.Customer, error) {
    for i, existingCustomer := range db.customers {
        if existingCustomer.ID == customer.ID {
            db.customers[i] = customer
            return customer, nil
        }
    }
    return entity.Customer{}, fmt.Errorf("customer not found")
}

func (db *InMemoryCustomerService) Delete(id int) error {
    for i, customer := range db.customers {
        if customer.ID == id {
            db.customers = append(db.customers[:i], db.customers[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("customer not found")
}