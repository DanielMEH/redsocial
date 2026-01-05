package config

import "net/http"

// AppError representa el error personalizado similar al enum de Rust
type AppError struct {
	Message      string `json:"message"`
	Code         int    `json:"code"`
	InternalCode string `json:"internal_code"`
	Status       int    `json:"status"`
}

// Implementa la interfaz 'error' de Go
func (e *AppError) Error() string {
	return e.Message
}

// ResponseSystem es el envoltorio estándar para todas las respuestas (éxito y error)
type ResponseSystem struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data"`
	Message      string      `json:"message"`
	Code         int         `json:"code"`
	InternalCode string      `json:"internal_code"`
	Status       int         `json:"status"`
}

/*

Codigos de errores ejemplo 401(1001,1003) con procesos diferente pero conectado al mismo codigo de error
*/
const (
	ErrCodeAuthInvalidCredentials = 1001
)

// NewAuthInvalidCredentials imita a Self::AuthInvalidCredentials(String)
func NewAuthInvalidCredentials(detail string) *AppError {
	return &AppError{
		Message:      "Credenciales inválidas: " + detail,
		Code:         ErrCodeAuthInvalidCredentials,
		InternalCode: "AUTH_INVALID_CREDENTIALS",
		Status:       http.StatusUnauthorized,
	}
}
