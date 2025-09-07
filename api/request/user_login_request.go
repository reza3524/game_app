package request

type UserLoginRequest struct {
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}
