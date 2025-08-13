package request

type RegisterUserRequest struct {
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
}
