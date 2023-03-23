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
// @Success		  		200 {object} request.RegisteredRequest
// @Failure		  		400 {object} response.ResponseModel
// @Router				/public/v1/credentials/log_in [post]
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
	message := "1030108389"
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

	//check if passwords are similar before it saves
	if UserCredentials.Password == UserCredentials.Retype_password {
		//encoding username before it saves
		usernameBase64 := base64.StdEncoding.EncodeToString([]byte(UserCredentials.Username))
		//hashing password before it saves
		passwordByte16 := md5.Sum([]byte(UserCredentials.Password))
		passwordHashString := hex.EncodeToString(passwordByte16[:])

		if insertErr := database.DBConn.Debug().Raw("INSERT INTO public.admin_accounts (name, username, password, retype_password) VALUES(?, ?, ?, ?)", UserCredentials.Name, usernameBase64, passwordHashString, passwordHashString).Scan(NewRegistered).Error; insertErr != nil {
			return c.JSON(response.ResponseModel{
				RetCode: "400",
				Message: insertErr.Error(),
				Data:    insertErr,
			})
		}
	} else {
		return c.JSON("Passwords Do Not Match")
	}

	return c.JSON(UserCredentials)
}

// Developer			Roldan
// @summary 	  		CREDENTIAL Base64/md5 hashing
// @Description	  		Encoding/Decoding/Hashing Credentials
// @Tags		  		JANUS REPORT GENERATION
// @Produce		  		json
// @Success		  		200 {object} response.LogInResponse
// @Failure		  		400 {object} response.ResponseModel
// @Router				/public/v1/credentials/register_sign_up [post]
func Log_in(c *fiber.Ctx) error {
	userCredentials := request.LogInRequest{}
	ClientInfo := response.LogInResponse{}

	if parsErr := c.BodyParser(&userCredentials); parsErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad request",
			Data:    parsErr.Error(),
		})
	}
	//encoding username
	usernameBase64 := base64.StdEncoding.EncodeToString([]byte(userCredentials.Username))
	//hashing password
	passwordByte16 := md5.Sum([]byte(userCredentials.Password))
	passwordHashString := hex.EncodeToString(passwordByte16[:])

	if fetchErr := database.DBConn.Debug().Raw("SELECT name, username FROM admin_accounts WHERE username = ? AND password = ?", usernameBase64, passwordHashString).Scan(&ClientInfo).Error; fetchErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad request",
			Data:    fetchErr.Error(),
		})
	}

	//decoding username to display
	originalClientUsername, decodErr := base64.StdEncoding.DecodeString(ClientInfo.Username)
	if decodErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad request",
			Data:    decodErr.Error(),
		})
	}
	ClientInfo.Username = string(originalClientUsername)

	return c.JSON(ClientInfo)

}
