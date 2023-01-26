package response

// swagger:response DashboardMenuResponse
type DashboardMenuResponse struct {
	Dashboard   *string `json:"dashboard"`
	Menu_Item   *string `json:"menu_item"`
	Enable      *int    `json:"enable"`
	Description *string `json:"description"`
	Soon        *int    `json:"soon"`
	Insti_code  *int    `json:"insti_code"`
	Icon        *int    `json:"icon"`
	Order       *int    `json:"order"`
}
