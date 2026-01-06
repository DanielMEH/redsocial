package register

import (
	"fmt"
	"regexp"
	"time"

	"github.com/redsocial/internal/module/users/domains/command"
)

var regValidationEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type EntityRegisterAccountRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	DisplayName     string `json:"display_name"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	BirthDate       string `json:"birth_date"`
}

func NewEntityRegisterAccount(data command.EntityRegisterAccountCommand) (*EntityRegisterAccountRequest, error) {
	if !regValidationEmail.MatchString(data.Email) {
		return nil, fmt.Errorf(`"El correo ingresado no es valido"`)
	}

	if data.Password != data.ConfirmPassword {
		return nil, fmt.Errorf(`"Las contraseñas tienen que ser iguales"`)

	}
	if len(data.FirstName) < 2 {
		return nil, fmt.Errorf("el nombre es obligatorio y debe tener al menos 2 caracteres")
	}
	if len(data.LastName) < 2 {
		return nil, fmt.Errorf("el apellido es obligatorio y debe tener al menos 2 caracteres")
	}

	_, err := time.Parse("2006-01-02", data.BirthDate)
	if err != nil {
		return nil, fmt.Errorf("formato de fecha de nacimiento inválido (debe ser YYYY-MM-DD)")
	}
	if data.DisplayName == "" {
		return nil, fmt.Errorf("el alias o nombre de usuario es obligatorio")
	}

	if len(data.Password) < 5 {
		return nil, fmt.Errorf("la contraseña debe tener al menos 8 caracteres")
	}

	return &EntityRegisterAccountRequest{
		Email:           data.Email,
		Password:        data.Password,
		ConfirmPassword: data.ConfirmPassword,
		DisplayName:     data.DisplayName,
		FirstName:       data.FirstName,
		LastName:        data.LastName,
		BirthDate:       data.BirthDate,
	}, nil

}
