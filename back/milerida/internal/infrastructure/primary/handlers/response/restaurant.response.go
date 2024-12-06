package response

type RestaurantResponse struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Address      string  `json:"address"`
	PhoneNumber  string  `json:"phone_number"`
	Email        string  `json:"email"`
	OpeningHours string  `json:"opening_hours"`
	CuisineType  string  `json:"cuisine_type"`
	AverageCost  float64 `json:"average_cost"`
}
