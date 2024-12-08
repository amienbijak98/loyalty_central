package dto

import (
	"time"
)

// MenuDTO represents a menu dto
type MenuDTO struct {
	ID           uint       `json:"id"`
	Name         string     `json:"name"`
	Descriptions string     `json:"descriptions"`
	ImgPath      string     `json:"img_path"`
	Type         string     `json:"type"`
	PriceMoney   int        `json:"price_money"`
	PricePoints  int        `json:"price_points"`
	Points       int        `json:"points"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}
