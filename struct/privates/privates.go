package privates

import (
	"github.com/gofiber/fiber/v2"

	fiberUtils "APITransactionGenerator/middleware/go-utils/fiber"
	"APITransactionGenerator/struct/errors"
	"APITransactionGenerator/struct/response"
)

func checkHealth() response.ResponseModel {
	return response.ResponseModel{
		RetCode: "100",
		Message: "Request success!",
		Data: errors.ErrorModel{
			Message:   "Service is available!",
			IsSuccess: true,
			Error:     nil,
		},
	}
}

func CheckServiceHealth(c *fiber.Ctx) error {
	// jwt := jwt.GetUserFromJWTClaim(c)
	fiberUtils.Ctx.New(c)
	health := checkHealth()
	response := errors.ErrorModel{}
	response = health.Data.(errors.ErrorModel)
	if !response.IsSuccess {
		return response.Error
	}
	return c.JSON(health)
}
