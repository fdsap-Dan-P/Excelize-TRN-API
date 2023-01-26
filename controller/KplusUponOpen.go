package controller

import (
	"APITransactionGenerator/middleware/go-utils/database"
	"APITransactionGenerator/struct/response"
	"encoding/json"

	//"encoding/json"
	//"fmt"

	"github.com/gofiber/fiber/v2"
)

// Developer			Roldan
// UponOpen				godoc
// @summary 	  		kPLUS UPON OPEN
// @Description	  		Provide the data that will be used by kPLUS upon opening the application
// @Tags		  		kPLUS UPON OPEN
// @Produce		  		json
// @Success		  		200 {object} response.GetParamResponse
// @Failure		  		400 {object} response.ResponseModel
// @Router				/public/v1/kplus/kplus_upon_open [get]
func KplusUponOpen(c *fiber.Ctx) error {
	//--------------------------------------------
	// G E T   I N S T I - P A R A M
	//--------------------------------------------
	instiParamResponse := []response.InstiParam{}

	if fetchErr := database.DBConn.Debug().Table("c_get_insti_param").Find(&instiParamResponse).Error; fetchErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "can't update table",
			Data:    fetchErr,
		})
	}

	//---------------------------------------------------------
	//  G E T   P A R A M   A N D   S P L A S H   S C R E E N
	//---------------------------------------------------------
	paramResponseList := []response.GetParamResponse{}
	paramResponse := response.GetParamResponse{}
	SplashScreenResponse := []response.SplashScreenResponse{}
	paramValueHolder := make(map[string]interface{})

	if fetchErr := database.DBConn.Debug().Table("c_splash_screen").Where("show = 'true' AND action = 'upon_login'").Find(&SplashScreenResponse).Error; fetchErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "can't fetch splash screen table",
			Data:    fetchErr,
		})
	}

	if fetchErr := database.DBConn.Debug().Table("c_get_param").Where("param_id != 214").Find(&paramResponseList).Error; fetchErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "can't fetch get param list table",
			Data:    fetchErr,
		})
	}

	if fetchErr := database.DBConn.Debug().Table("c_get_param").Where("param_id = 214").Update("param_value", SplashScreenResponse).Error; fetchErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "can't fetch param [214] table",
			Data:    fetchErr,
		})
	}

	if fetchErr := database.DBConn.Debug().Table("c_get_param").Where("param_id = 214").Find(&paramResponse).Error; fetchErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "can't fetch get param table",
			Data:    fetchErr,
		})
	}

	json.Unmarshal(json.RawMessage(paramResponse.Param_value), &paramValueHolder)

	GetSplashScreen := response.ParamResponse{
		Param_id:    paramResponse.Param_id,
		Param_name:  paramResponse.Param_name,
		Param_value: paramValueHolder,
	}

	param_splash_screen := append([]interface{}{paramResponseList}, []interface{}{GetSplashScreen})

	//-------------------------------------
	// G E T   P A R A M  W E B T O O L
	//-------------------------------------
	GetParamWebTool := []response.GetParamWebToolResponse{}

	if fetchErr := database.DBConn.Debug().Table("m_param_config").Find(&GetParamWebTool).Error; fetchErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "can't fetch get param list table",
			Data:    fetchErr,
		})
	}

	//---------------------------------------------
	//      DISPLAY ALL UPON OPENING KPLUS
	//---------------------------------------------
	result := append([]interface{}{instiParamResponse}, []interface{}{param_splash_screen})
	result1 := append(result, []interface{}{GetParamWebTool})
	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    result1,
	})
}

// Developer			Roldan
// UponOpen				godoc
// @summary 	  		API FOR SPLASH SCREEN
// @Description	  		Provide the data FROM c_splash_screen TABLE that will be used by kPLUS upon opening the application
// @Tags		  		kPLUS UPON OPEN
// @Produce		  		json
// @Success		  		200 {object} response.SplashScreenResponse
// @Failure		  		400 {object} response.ResponseModel
// @Router				/public/v1/kplus/splash_screen [get]
func SplashScreen(c *fiber.Ctx) error {
	splashScreenResponse := []response.SplashScreenResponse{}

	if fetchErr := database.DBConn.Debug().Table("c_splash_screen").Find(&splashScreenResponse).Error; fetchErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "can't update table",
			Data:    fetchErr,
		})
	}

	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    splashScreenResponse,
	})
}

// Developer			Roldan
// UponOpen				godoc
// @summary 	  		API FOR INSTI PARAM
// @Description	  		Provide the data FROM insti_param TABLE that will be used by kPLUS upon opening the application
// @Tags		  		kPLUS UPON OPEN
// @Produce		  		json
// @Success		  		200 {object} response.InstiParam
// @Failure		  		400 {object} response.ResponseModel
// @Router				/public/v1/kplus/insti_param [get]
func InstiParam(c *fiber.Ctx) error {
	instiParamResponse := []response.InstiParam{}

	if fetchErr := database.DBConn.Debug().Table("c_get_insti_param").Find(&instiParamResponse).Error; fetchErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "can't update table",
			Data:    fetchErr,
		})
	}

	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    instiParamResponse,
	})
}

// Developer			Roldan
// UponOpen				godoc
// @summary 	  		API FOR GET PARAM
// @Description	  		Provide the data FROM getParam TABLE that will be used by kPLUS upon opening the application
// @Tags		  		kPLUS UPON OPEN
// @Produce		  		json
// @Success		  		200 {object} response.InstiParam
// @Failure		  		400 {object} response.ResponseModel
// @Router				/public/v1/kplus/get_param [get]
func GetParam(c *fiber.Ctx) error {
	paramResponseList := []response.GetParamResponse{}
	paramResponse := response.GetParamResponse{}
	SplashScreenResponse := []response.SplashScreenResponse{}
	paramValueHolder := make(map[string]interface{})

	if fetchErr := database.DBConn.Debug().Table("c_splash_screen").Where("show = 'true' AND action = 'upon_login'").Find(&SplashScreenResponse).Error; fetchErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "can't fetch splash screen table",
			Data:    fetchErr,
		})
	}

	if fetchErr := database.DBConn.Debug().Table("c_get_param").Where("param_id != 214").Find(&paramResponseList).Error; fetchErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "can't fetch get param list table",
			Data:    fetchErr,
		})
	}

	if fetchErr := database.DBConn.Debug().Table("c_get_param").Where("param_id = 214").Update("param_value", SplashScreenResponse).Error; fetchErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "can't fetch param [214] table",
			Data:    fetchErr,
		})
	}

	if fetchErr := database.DBConn.Debug().Table("c_get_param").Where("param_id = 214").Find(&paramResponse).Error; fetchErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "can't fetch get param table",
			Data:    fetchErr,
		})
	}

	json.Unmarshal(json.RawMessage(paramResponse.Param_value), &paramValueHolder)

	GetSplashScreen := response.ParamResponse{
		Param_id:    paramResponse.Param_id,
		Param_name:  paramResponse.Param_name,
		Param_value: paramValueHolder,
	}

	result := append([]interface{}{paramResponseList}, []interface{}{GetSplashScreen})

	return c.JSON(response.ResponseModel{
		RetCode: "200",
		Message: "success",
		Data:    result,
	})
}
