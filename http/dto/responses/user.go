package responses

type UserLoginResponse struct {
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	AccessToken string `json:"access_token"`
}
