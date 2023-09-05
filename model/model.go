package model

type SignupResponse struct {
	Message string `json:"message"`
	User    struct {
		UserID   string `json:"user_id"`
		Nickname string `json:"nickname"`
	} `json:"user"`
}

type Res struct {
	Message map[string]interface{}
	ResUser map[string]interface{}
}
