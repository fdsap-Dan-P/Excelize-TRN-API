package controller

import (
	"APITransactionGenerator/struct/request"
	"APITransactionGenerator/struct/response"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Developer			Roldan
// @summary 	  		CREDENTIAL Base64
// @Description	  		Encoding/Decoding Credentials
// @Tags		  		JANUS REPORT GENERATION
// @Produce		  		json
// @Success		  		200 {object} response.TransactionResponse
// @Failure		  		400 {object} response.ResponseModel
// @Router				/public/v1/credentials/register_sign_up [get]
func Registered(c *fiber.Ctx) error {
	UserInfo := request.RegisteredRequest{}

	if parsErr := c.BodyParser(&UserInfo); parsErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad request",
			Data:    parsErr.Error(),
		})
	}
	message := "Hello"
	base64Text := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println("base64: ", base64Text)

	originalText, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("origial text: ", string(originalText))

	data := []byte("hello")
	byte16 := md5.Sum(data)
	pass := hex.EncodeToString(byte16[:])
	fmt.Println("HASH: ", string(pass))

	// if insertErr := database.DBConn.Raw("INSERT INTO trn_gen.admin (admin_id, username)")

	return c.JSON(UserInfo)
}

// 5d41402abc4b2a76b9719d911017c592
