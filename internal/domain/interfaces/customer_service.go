package interfaces

import "github.com/Bigtalljosh/gostrife/internal/entity"

type CustomerService interface {
	GetAll() ([]entity.Customer, error)
	GetByID(id int) (entity.Customer, error)
	Create(customer entity.Customer) (entity.Customer, error)
	Update(customer entity.Customer) (entity.Customer, error)
	Delete(id int) error
}