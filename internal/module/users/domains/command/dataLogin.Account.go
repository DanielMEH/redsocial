package command

type EntityLoginAccountCommand struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
