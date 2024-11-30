package dto

import (
	"loyalty_central/internal/models"

	"gorm.io/gorm"
)

type AdminDTO struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
}

func (a *AdminDTO) ToAdmin() *models.Admin {
	return &models.Admin{
		Model:       gorm.Model{},
		Username:    a.Username,
		Password:    "",
		Name:        a.Name,
		PhoneNumber: a.Phone,
		Role:        a.Role,
	}
}
