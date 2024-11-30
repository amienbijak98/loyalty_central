package repository

import (
	"errors"

	"gorm.io/gorm"

	"loyalty_central/internal/models"
)

type AdminRepository interface {
	GetAllAdmins() ([]*models.Admin, error)
	GetAdminByUsername(username string) (*models.Admin, error)
	CreateAdmin(admin *models.Admin) (*models.Admin, error)
	UpdateAdmin(admin *models.Admin, adminID uint) (*models.Admin, error)
	DeleteAdmin(adminID uint) error
	UndeleteAdminByID(adminID uint) error
	GetDeletedAdmins() ([]*models.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) GetAllAdmins() ([]*models.Admin, error) {
	var admins []*models.Admin
	err := r.db.Find(&admins).Error
	if err != nil {
		return nil, err
	}

	return admins, nil
}

func (r *adminRepository) GetAdminByUsername(username string) (*models.Admin, error) {
	var admin models.Admin
	err := r.db.Where("username = ?", username).First(&admin).Error
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

func (r *adminRepository) CreateAdmin(admin *models.Admin) (*models.Admin, error) {
	err := r.db.Create(admin).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *adminRepository) UpdateAdmin(admin *models.Admin, adminID uint) (*models.Admin, error) {
	err := r.db.Model(&admin).
		Where("id = ?", adminID).
		Updates(map[string]interface{}{
			"username":     admin.Username,
			"password":     admin.Password,
			"name":         admin.Name,
			"role":         admin.Role,
			"phone_number": admin.PhoneNumber,
		}).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *adminRepository) DeleteAdmin(adminID uint) error {
	err := r.db.Delete(&models.Admin{}, adminID).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *adminRepository) UndeleteAdminByID(adminID uint) error {
	err := r.db.Unscoped().Model(&models.Admin{}).Where("id = ?", adminID).Update("deleted_at", nil).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *adminRepository) GetDeletedAdmins() ([]*models.Admin, error) {
	var admins []*models.Admin
	err := r.db.Unscoped().Where("deleted_at IS NOT NULL").Find(&admins).Error
	if err != nil {
		return nil, err
	}

	if len(admins) == 0 {
		return nil, errors.New("no deleted admin found")
	}

	return admins, nil
}
