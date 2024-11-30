package service

import (
	"loyalty_central/internal/dto"
	"loyalty_central/internal/models"
	"loyalty_central/internal/repository"
	"loyalty_central/internal/utils"
)

type AdminService interface {
	GetAllAdmins() ([]*dto.AdminDTO, error)
	GetAdminByUsername(username string) (*dto.AdminDTO, error)
	CreateAdmin(admin *models.Admin) (*models.Admin, error)
	UpdateAdmin(admin *models.Admin, adminID uint) (*models.Admin, error)
	DeleteAdmin(adminID uint) error
	UndeleteAdminByID(adminID uint) error
	GetDeletedAdmins() ([]*dto.AdminDTO, error)
}

type adminService struct {
	repository repository.AdminRepository
}

func NewAdminService(repository repository.AdminRepository) AdminService {
	return &adminService{repository}
}

func (s *adminService) GetAllAdmins() ([]*dto.AdminDTO, error) {
	admins, err := s.repository.GetAllAdmins()
	if err != nil {
		return nil, err
	}

	var dtos []*dto.AdminDTO
	for _, admin := range admins {
		dtos = append(dtos, &dto.AdminDTO{
			ID:        admin.ID,
			Username:  admin.Username,
			Name:      admin.Name,
			Role:      admin.Role,
			Phone:     admin.PhoneNumber,
			CreatedAt: admin.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return dtos, nil
}

func (s *adminService) GetAdminByUsername(username string) (*dto.AdminDTO, error) {
	admin, err := s.repository.GetAdminByUsername(username)
	if err != nil {
		return nil, err
	}

	return &dto.AdminDTO{
		ID:        admin.ID,
		Username:  admin.Username,
		Name:      admin.Name,
		Role:      admin.Role,
		Phone:     admin.PhoneNumber,
		CreatedAt: admin.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *adminService) CreateAdmin(admin *models.Admin) (*models.Admin, error) {
	hashedPassword, err := utils.HashPassword(admin.Password)
	if err != nil {
		return nil, err
	}
	admin.Password = string(hashedPassword)
	return s.repository.CreateAdmin(admin)
}

func (s *adminService) UpdateAdmin(admin *models.Admin, adminID uint) (*models.Admin, error) {
	hashedPassword, err := utils.HashPassword(admin.Password)
	if err != nil {
		return nil, err
	}
	admin.Password = string(hashedPassword)
	return s.repository.UpdateAdmin(admin, adminID)
}

func (s *adminService) DeleteAdmin(adminID uint) error {
	return s.repository.DeleteAdmin(adminID)
}

func (s *adminService) UndeleteAdminByID(adminID uint) error {
	return s.repository.UndeleteAdminByID(adminID)
}

func (s *adminService) GetDeletedAdmins() ([]*dto.AdminDTO, error) {
	admins, err := s.repository.GetDeletedAdmins()
	if err != nil {
		return nil, err
	}

	var dtos []*dto.AdminDTO
	for _, admin := range admins {
		dtos = append(dtos, &dto.AdminDTO{
			ID:        admin.ID,
			Username:  admin.Username,
			Name:      admin.Name,
			Role:      admin.Role,
			Phone:     admin.PhoneNumber,
			CreatedAt: admin.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return dtos, nil
}
