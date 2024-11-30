package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username    string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	Name        string `gorm:"not null"`
	PhoneNumber string `gorm:"not null" json:"phone_number"`
	Role        string `gorm:"default:casheer"`
}

type Customer struct {
	gorm.Model
	Name        string    `gorm:"not null"`
	DOB         time.Time `gorm:"not null"`
	PhoneNumber string    `gorm:"unique;not null" json:"phone_number"`
	TotalPoints int
}

type Menu struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Descriptions string
	ImgPath      string
	Type         string
	PriceMoney   int
	PricePoints  int
	Points       int
}

type Purchase struct {
	gorm.Model
	CustomerID          int `gorm:"index"`
	AdminID             int `gorm:"index"`
	PaymentMethod       string
	TotalPurchaseMoney  int
	TotalPurchasePoints int
	TotalPointsEarned   int
	Customer            Customer         `gorm:"foreignKey:CustomerID"`
	Admin               Admin            `gorm:"foreignKey:AdminID"`
	PurchaseDetails     []PurchaseDetail `gorm:"foreignKey:PurchaseID"`
}

type PurchaseDetail struct {
	gorm.Model
	PurchaseID     int `gorm:"index"`
	MenuID         int `gorm:"index"`
	MenuName       string
	PriceMoney     int
	PricePoints    int
	Qty            int
	SubTotalMoney  int
	SubTotalPoints int
	PointsEarned   int
	Purchase       Purchase `gorm:"foreignKey:PurchaseID"`
	Menu           Menu     `gorm:"foreignKey:MenuID"`
}
