package dblayer

import (
	"github.com/rossi1/musicstore/models"
	model "github.com/rossi1/musicstore/models"
)

type DBLayer interface {
	GetAllProducts() ([]model.Product, error)
	GetPromos() ([]model.Product, error)
	GetCustomerByName(string, string) (model.Customer, error)
	GetCustomerByID(int) (model.Customer, error)
	GetProduct(int) (model.Product, error)
	AddUser(model.Customer) (model.Customer, error)
	SignInUser(username, password string) (model.Customer, error)
	SignOutUserById(int) error
	GetCustomerOrdersByID(int) ([]models.Order, error)
	AddOrder(model.Order) error
	GetCreditCardCID(int) (string, error)
	SaveCreditCardForCustomer(int, string) error
}

type ErrINVALIDPASSWORD struct {
	err error
}

func (e *ErrINVALIDPASSWORD) Error() string { return e.Error() }
