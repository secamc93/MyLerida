package response

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	UserId       uint   `json:"user_id"`
	UserName     string `json:"user_name"`
	UserLastName string `json:"user_last_name"`
}
