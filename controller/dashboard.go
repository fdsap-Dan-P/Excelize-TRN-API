package controller

import (
	"APITransactionGenerator/middleware/go-utils/database"
	"APITransactionGenerator/struct/request"
	"APITransactionGenerator/struct/response"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Developer		Roldan
// @summary 	  	DASHBOARD-MENU
// @Description	  	Provide the menu to show in kPLUS
// @Tags		  	DASHBOARD
// @Accept		  	json
// @Produce		  	json
// @Param       	getInstiCode body request.InstiCodeRequest true  "Insti Code"
// @Success		  	200 {object} response.DashboardMenuResponse
// @Failure		  	400 {object} response.ResponseModel
// @Router			/public/v1/dashboard/dashboardMenu [post]
func DashboardMenu(c *fiber.Ctx) error {
	getInstiCode := request.InstiCodeRequest{}

	fmt.Println(getInstiCode.Insti_code)
	if parsErr := c.BodyParser(&getInstiCode); parsErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad request",
			Data:    parsErr.Error(),
		})
	}
	dashResponse := response.DashboardMenuResponse{}

	database.DBConn.Debug().Table("c_menu").Where("insti_code = ?", getInstiCode.Insti_code).Find(&dashResponse)

	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    dashResponse,
	})
}

// // Developer			Roldan
// // Get_DashboardMenu	godoc
// // @summary 	  		GET DASHBOARD-MENU
// // @Description	  		Provide the menu to show in kPLUS
// // @Tags		  		DASHBOARD
// // @Produce		  		json
// // @Success		  		200 {object} response.DashboardMenuResponse
// // @Failure		  		400 {object} response.ResponseModel
// // @Router				/public/v1/dashboard/dashboardMenu [get]
// func Get_DashboardMenu(c *fiber.Ctx) error {

// 	dashResponse := []response.DashboardMenuResponse{}

// 	if fetchErr := database.DBConn.Debug().Table("c_menu").Find(&dashResponse).Error; fetchErr != nil {
// 		return c.JSON(response.ResponseModel{
// 			RetCode: "400",
// 			Message: "can't fetch data",
// 			Data:    fetchErr,
// 		})
// 	}
// 	return c.JSON(dashResponse)

// }
