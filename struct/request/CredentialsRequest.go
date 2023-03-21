package request

type RegisteredRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	Retype_password string `json:"retype_password"`
}
