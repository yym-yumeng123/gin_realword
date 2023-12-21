package request

type UserRegistrationRequest struct {
	User UserRegistrationBody `json:"user"`
}

type UserLoginRequest struct {
	User UserLoginBody
}

type UserRegistrationBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
