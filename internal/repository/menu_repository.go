package repository

import (
	"loyalty_central/internal/dto"
	"loyalty_central/internal/models"

	"gorm.io/gorm"
)

// MenuRepository provides menu repository
type MenuRepository interface {
	GetAllMenus() ([]*models.Menu, error)
	GetMenuByID(menuID uint) (*models.Menu, error)
	CreateMenu(menu *models.Menu) (*models.Menu, error)
	UpdateMenu(menu *dto.MenuDTO, menuID uint) (*models.Menu, error)
	DeleteMenu(menuID uint) error
}

type menuRepository struct {
	db *gorm.DB
}

// NewMenuRepository creates an instance of MenuRepository
func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) GetAllMenus() ([]*models.Menu, error) {
	var menus []*models.Menu
	err := r.db.Find(&menus).Error
	if err != nil {
		return nil, err
	}

	return menus, nil
}

func (r *menuRepository) GetMenuByID(menuID uint) (*models.Menu, error) {
	menu := &models.Menu{}
	err := r.db.First(menu, menuID).Error
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (r *menuRepository) CreateMenu(menu *models.Menu) (*models.Menu, error) {
	err := r.db.Create(menu).Error
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (r *menuRepository) UpdateMenu(menu *dto.MenuDTO, menuID uint) (*models.Menu, error) {
	menuModel := &models.Menu{}
	err := r.db.Model(menuModel).Where("id = ?", menuID).Updates(map[string]interface{}{
		"Name":         menu.Name,
		"Descriptions": menu.Descriptions,
		"ImgPath":      menu.ImgPath,
		"Type":         menu.Type,
		"PriceMoney":   menu.PriceMoney,
		"PricePoints":  menu.PricePoints,
		"Points":       menu.Points,
	}).Error
	if err != nil {
		return nil, err
	}

	err = r.db.First(menuModel, menuID).Error
	if err != nil {
		return nil, err
	}

	return menuModel, nil
}

func (r *menuRepository) DeleteMenu(menuID uint) error {
	menu := &models.Menu{}
	err := r.db.Delete(menu, menuID).Error
	if err != nil {
		return err
	}

	return nil
}
