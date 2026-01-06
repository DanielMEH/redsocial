package register

type EntityRegisterAccountResponse struct {
	Message string `json:"message"`
	Details struct {
		AccountToken     string `json:"account_token"`
		AccountSessionId string `json:"account_session_id"`
	} `json:"details"`
}
