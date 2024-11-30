package repository

import (
	"errors"

	"gorm.io/gorm"

	"loyalty_central/internal/models"
)

type CustomerRepository interface {
	GetAllCustomers() ([]*models.Customer, error)
	GetCustomerByID(customerID uint) (*models.Customer, error)
	CreateCustomer(customer *models.Customer) (*models.Customer, error)
	GetCustomerByPhoneNumber(phoneNumber string) (*models.Customer, error)
	UpdateCustomer(customer *models.Customer, customerID uint) (*models.Customer, error)
	DeleteCustomer(customerID uint) error
	GetDeletedCustomers() ([]*models.Customer, error)
	UndeleteCustomerByID(customerID uint) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) CreateCustomer(customer *models.Customer) (*models.Customer, error) {
	err := r.db.Create(customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *customerRepository) GetCustomerByPhoneNumber(phoneNumber string) (*models.Customer, error) {
	customer := new(models.Customer)

	err := r.db.Where("phone_number = ?", phoneNumber).First(&customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *customerRepository) UpdateCustomer(customer *models.Customer, customerID uint) (*models.Customer, error) {
	err := r.db.Model(&customer).
		Where("id = ?", customerID).
		Updates(map[string]interface{}{
			"name":         customer.Name,
			"dob":          customer.DOB,
			"phone_number": customer.PhoneNumber,
			"total_points": customer.TotalPoints,
		}).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *customerRepository) DeleteCustomer(customerID uint) error {
	err := r.db.Delete(&models.Customer{}, customerID).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *customerRepository) GetAllCustomers() ([]*models.Customer, error) {
	customers := make([]*models.Customer, 0)

	err := r.db.Find(&customers).Error
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *customerRepository) GetCustomerByID(customerID uint) (*models.Customer, error) {
	customer := new(models.Customer)

	err := r.db.Where("id = ?", customerID).First(&customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *customerRepository) GetDeletedCustomers() ([]*models.Customer, error) {
	customers := make([]*models.Customer, 0)

	err := r.db.Unscoped().Where("deleted_at IS NOT NULL").Find(&customers).Error
	if err != nil {
		return nil, err
	}

	if len(customers) == 0 {
		return nil, errors.New("no deleted customer found")
	}

	return customers, nil
}

func (r *customerRepository) UndeleteCustomerByID(customerID uint) error {
	err := r.db.Unscoped().Model(&models.Customer{}).Where("id = ?", customerID).Update("deleted_at", nil).Error
	if err != nil {
		return err
	}

	return nil
}
