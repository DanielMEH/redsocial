package profile

import "time"

type EntityGetProfileAccountResponse struct {
	Message string `json:"message"`
	Details struct {
		Email     string    `json:"email"`
		Alias     string    `json:"alias"`
		BirthDate time.Time `json:"bith_date"`
	} `json:"details"`
}
