package requests

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
	Avatar   string `json:"avatar"`
}


type LoginInput struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}