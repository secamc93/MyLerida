package dtos

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponseDTO struct {
	Token        string
	UserId       uint
	UserName     string
	UserLastName string
}
