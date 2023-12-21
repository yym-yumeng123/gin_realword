package response

type UserRegistrationResponse struct {
	User UserRegistrationBody `json:"user"`
}

type UserRegistrationBody struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}
