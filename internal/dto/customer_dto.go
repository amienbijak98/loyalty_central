package dto

type CustomerDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	TotalPoints int    `json:"total_points"`
}
