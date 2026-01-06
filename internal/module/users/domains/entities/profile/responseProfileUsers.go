package profile

type EntityGetProfileAccountResponse struct {
	Message string `json:"message"`
	Details struct {
		Email     string `json:"email"`
		Alias     string `json:"alias"`
		BirthDate string `json:"birth_date"`
	} `json:"details"`
}
