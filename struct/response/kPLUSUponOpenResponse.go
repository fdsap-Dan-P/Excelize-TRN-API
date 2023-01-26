package response

// swagger:response DashboardMenuResponse
type SplashScreenResponse struct {
	Action            string `json:"action"`
	Title             string `json:"title"`
	Message           string `json:"message"`
	Sub_message       string `json:"sub_message"`
	Image_url         string `json:"image_url"`
	Show              string `json:"show"`
	Created_date      string `json:"created_date"`
	Id                int    `json:"id"`
	Last_updated_by   int    `json:"last_updated_by"`
	Last_updated_date string `json:"last_updated_date"`
	Redirect_link     string `json:"redirect_link"`
}

type InstiParam struct {
	Insti_code       int    `json:"insti_code"`
	Institution      string `json:"institution"`
	Address          string `json:"address"`
	Telephone_number string `json:"telephone_number"`
	Email            string `json:"email"`
}

type GetParamResponse struct {
	Param_id    int    `json:"param_id"`
	Param_name  string `json:"param_name"`
	Param_value string `json:"param_value"`
}

type ParamResponse struct {
	Param_id    int                    `json:"param_id"`
	Param_name  string                 `json:"param_name"`
	Param_value map[string]interface{} `json:"param_value"`
}

type GetParamWebToolResponse struct {
	Param_id          string `json:"param_id"`
	Param_name        string `json:"param_name"`
	Param_value       string `json:"param_value"`
	Param_desc        string `json:"param_desc"`
	App_type          string `json:"app_type"`
	Created_by        string `json:"created_by"`
	Created_date      string `json:"created_date"`
	Last_updated_by   string `json:"last_updated_by"`
	Last_updated_date string `json:"last_updated_date"`
	Param_status      string `json:"param_status"`
	Process_id        string `json:"process_id"`
	Value_type        string `json:"value_type"`
	Value_lookup      string `json:"value_lookup"`
}
