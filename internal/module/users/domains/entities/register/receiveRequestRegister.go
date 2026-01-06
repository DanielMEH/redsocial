package register

import (
	"fmt"
	"regexp"

	"github.com/redsocial/internal/module/users/domains/command"
)

var regValidationEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type EntityRegisterAccountRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	DisplayName     string `json:"display_name"`
}

func NewEntityRegisterAccount(data command.EntityRegisterAccountCommand) (*EntityRegisterAccountRequest, error) {
	if !regValidationEmail.MatchString(data.Email) {
		return nil, fmt.Errorf(`"El correo ingresado no es valido"`)
	}

	if data.Password != data.ConfirmPassword {
		return nil, fmt.Errorf(`"Las contraseñas tienen que ser iguales"`)

	}

	if len(data.Password) < 5 {
		return nil, fmt.Errorf("la contraseña debe tener al menos 8 caracteres")
	}

	return &EntityRegisterAccountRequest{
		Email:           data.Email,
		Password:        data.Password,
		ConfirmPassword: data.ConfirmPassword,
		DisplayName:     data.DisplayName,
	}, nil

}
