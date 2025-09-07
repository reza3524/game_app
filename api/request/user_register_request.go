package request

type UserRegisterRequest struct {
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}
