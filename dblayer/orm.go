package dblayer

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rossi1/musicstore/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname string) (*DBORM, error) {
	db, err := gorm.Open(mysql.Open(dbname), &gorm.Config{})
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

func (db *DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

func (db *DBORM) GetCustomerByName(firstname string, lastname string) (customer models.Customer, err error) {
	return customer, db.Where(&models.Customer{FirstName: firstname, LastName: lastname}).Find(&customer).Error
}

func (db *DBORM) GetCustomerByID(id int) (customer models.Customer, err error) {
	return customer, db.First(&customer, id).Error
}

func (db *DBORM) GetProduct(id int) (product models.Product, error error) {
	return product, db.First(&product, id).Error
}

func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error) {

	hashPassword(&customer.Pass)
	customer.LoggedIn = true
	err := db.Create(&customer).Error
	customer.Pass = ""
	return customer, err
}

func (db *DBORM) SignInUser(email, pass string) (customer models.Customer, err error) {
	result := db.Table("customers").Where(&models.Customer{Email: email})
	err = result.First(&customer).Error
	if err != nil {
		return customer, err

	}

	if !checkPassword(customer.Pass, pass) {
		return customer, ErrINVALIDPASSWORD{err: errors.New("errors")}
	}
	customer.Pass = ""
	err = result.Update("loggedin", 1).Error

	if err != nil {
		return customer, err

	}
	return customer, result.Find(&customer).Error
}
func (db *DBORM) SignOutUserById(id int) error {
	//Create a customer Go struct with the provided if
	customer := models.Customer{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	//Update the customer row to reflect the fact that the customer is not logged in
	return db.Table("Customers").Where(&customer).Update("loggedin", 0).Error
}

func (db *DBORM) GetCustomerOrdersByID(id int) (orders []models.Order, err error) {
	return orders, db.Table("orders").Select("*").Joins("join customers on customers.id = customer_id").Joins("join products on products.id = product_id").Where("customer_id=?", id).Scan(&orders).Error

}

func hashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}

	//convert password string to byte slice so that we can use it with the bcrypt package
	sBytes := []byte(*s)

	//Obtain hashed password via the GenerateFromPassword() method
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	//update password string with the hashed version
	*s = string(hashedBytes[:])
	return nil
}

func checkPassword(existingHash, incomingPass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(incomingPass)) == nil
}

//Add the order to the orders table
func (db *DBORM) AddOrder(order models.Order) error {
	return db.Create(&order).Error
}

//Get the id representing the credit card from the database
func (db *DBORM) GetCreditCardCID(id int) (string, error) {
	cusomterWithCCID := struct {
		models.Customer
		CCID string `gorm:"column:cc_customerid"`
	}{}
	return cusomterWithCCID.CCID, db.First(&cusomterWithCCID, id).Error
}

//Save the credit card information for the customer
func (db *DBORM) SaveCreditCardForCustomer(id int, ccid string) error {
	result := db.Table("customers").Where("id=?", id)
	return result.Update("cc_customerid", ccid).Error
}
