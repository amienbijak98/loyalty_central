package service

import (
	"loyalty_central/internal/dto"
	"loyalty_central/internal/models"
	"loyalty_central/internal/repository"
)

type CustomerService interface {
	GetAllCustomers() ([]*models.Customer, error)
	GetCustomerByID(customerID uint) (*dto.CustomerDTO, error)
	CreateCustomer(customer *models.Customer) (*models.Customer, error)
	GetCustomerByPhoneNumber(phoneNumber string) (*models.Customer, error)
	UpdateCustomer(customer *models.Customer, customerID uint) (*models.Customer, error)
	DeleteCustomer(customerID uint) error
	GetDeletedCustomers() ([]*models.Customer, error)
	UndeleteCustomerByID(customerID uint) error
}

type customerService struct {
	repository repository.CustomerRepository
}

func NewCustomerService(repository repository.CustomerRepository) CustomerService {
	return &customerService{repository}
}

func (s *customerService) CreateCustomer(customer *models.Customer) (*models.Customer, error) {
	return s.repository.CreateCustomer(customer)
}

func (s *customerService) GetCustomerByPhoneNumber(phoneNumber string) (*models.Customer, error) {
	return s.repository.GetCustomerByPhoneNumber(phoneNumber)
}

func (s *customerService) UpdateCustomer(customer *models.Customer, customerID uint) (*models.Customer, error) {
	return s.repository.UpdateCustomer(customer, customerID)
}

func (s *customerService) DeleteCustomer(customerID uint) error {
	return s.repository.DeleteCustomer(customerID)
}

func (s *customerService) GetAllCustomers() ([]*models.Customer, error) {
	return s.repository.GetAllCustomers()
}

func (s *customerService) GetCustomerByID(customerID uint) (*dto.CustomerDTO, error) {
	customer, err := s.repository.GetCustomerByID(customerID)
	if err != nil {
		return nil, err
	}

	// Mapping data dari models.Customer ke dto.CustomerDTO
	customerDTO := dto.CustomerDTO{
		ID:          customer.ID,
		Name:        customer.Name,
		PhoneNumber: customer.PhoneNumber,
		TotalPoints: customer.TotalPoints,
	}

	return &customerDTO, nil
}

func (s *customerService) GetDeletedCustomers() ([]*models.Customer, error) {
	return s.repository.GetDeletedCustomers()
}

func (s *customerService) UndeleteCustomerByID(customerID uint) error {
	return s.repository.UndeleteCustomerByID(customerID)
}
