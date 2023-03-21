package controller

import (
	"APITransactionGenerator/middleware/go-utils/database"
	"APITransactionGenerator/struct/request"
	"APITransactionGenerator/struct/response"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Developer			Roldan
// @summary 	  		CREDENTIAL Base64/md5 hashing
// @Description	  		Encoding/Decoding/Hashing Credentials
// @Tags		  		JANUS REPORT GENERATION
// @Produce		  		json
// @Success		  		200 {object} response.RegisteredRequest
// @Failure		  		400 {object} response.ResponseModel
// @Router				/public/v1/credentials/register_sign_up [post]
func Registered(c *fiber.Ctx) error {
	UserCredentials := request.RegisteredRequest{}
	NewRegistered := request.RegisteredRequest{}

	if parsErr := c.BodyParser(&UserCredentials); parsErr != nil {
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

	if UserCredentials.Password == UserCredentials.Retype_password {
		usernameBase64 := base64.StdEncoding.EncodeToString([]byte(UserCredentials.Username))
		passwordByte16 := md5.Sum([]byte(UserCredentials.Password))
		passwordHashString := hex.EncodeToString(passwordByte16[:])

		fmt.Println(usernameBase64)
		fmt.Println(passwordHashString)

		if insertErr := database.DBConn.Debug().Raw("INSERT INTO public.admin_accounts (username, password, retype_password) VALUES(?, ?, ?)", usernameBase64, passwordHashString, passwordHashString).Scan(&NewRegistered).Error; insertErr != nil {
			return c.JSON(response.ResponseModel{
				RetCode: "400",
				Message: insertErr.Error(),
				Data:    insertErr,
			})
		}
	} else {
		return c.JSON("Passwords Do Not Match")
	}

	return c.JSON(NewRegistered)
}

// 5d41402abc4b2a76b9719d911017c592
