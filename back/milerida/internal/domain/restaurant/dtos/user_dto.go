package dtos

type RestaurantDTO struct {
	ID           uint
	Name         string
	Address      string
	PhoneNumber  string
	Email        string
	OpeningHours string
	CuisineType  string
	AverageCost  float64
}
