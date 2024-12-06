package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	Name         string  `gorm:"size:255;not null" json:"name"`
	Address      string  `gorm:"type:text" json:"address"`
	PhoneNumber  string  `gorm:"size:20" json:"phone_number"`
	Email        string  `gorm:"size:100" json:"email"`
	OpeningHours string  `gorm:"size:100" json:"opening_hours"`
	CuisineType  string  `gorm:"size:100" json:"cuisine_type"`
	AverageCost  float64 `gorm:"type:decimal(10,2)" json:"average_cost"`
}

type Transportation struct {
	gorm.Model
	Name         string  `gorm:"size:255;not null" json:"name"`
	Description  string  `gorm:"type:text" json:"description"`
	Capacity     int     `gorm:"not null" json:"capacity"`
	PricePerKm   float64 `gorm:"type:decimal(10,2)" json:"price_per_km"`
	Availability bool    `gorm:"default:true" json:"availability"`
}

type Service struct {
	gorm.Model
	Name         string  `gorm:"size:255;not null" json:"name"`
	Description  string  `gorm:"type:text" json:"description"`
	Price        float64 `gorm:"type:decimal(10,2)" json:"price"`
	ContactPhone string  `gorm:"size:20" json:"contact_phone"`
	ContactEmail string  `gorm:"size:100" json:"contact_email"`
	Available    bool    `gorm:"default:true" json:"available"`
}

type Tourism struct {
	gorm.Model
	Name         string  `gorm:"size:255;not null" json:"name"`
	Location     string  `gorm:"size:255" json:"location"`
	Description  string  `gorm:"type:text" json:"description"`
	EntryFee     float64 `gorm:"type:decimal(10,2)" json:"entry_fee"`
	OpeningHours string  `gorm:"size:100" json:"opening_hours"`
	ContactPhone string  `gorm:"size:20" json:"contact_phone"`
	Website      string  `gorm:"size:255" json:"website"`
}
