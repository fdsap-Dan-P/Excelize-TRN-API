package request

type RegisteredRequest struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Retype_password string `json:"retype_password"`
}

type LogInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
