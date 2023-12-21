package request

type UserRegistrationRequest struct {
	User UserRegistrationBody `json:"user"`
}

type UserRegistrationBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
