package auth_service

import (
	"Chat/dto"
	"fmt"
)

func Register_Service(user dto.TRegister) dto.TRegisterResponse {
	fmt.Println(`
		Username: ` + user.Username + `
		Email: ` + user.Email + `
		Password: ` + user.Password + `

		Thank you for registering!
	`)

	newUser := dto.TRegisterResponse{
		Message: "Thank you for registering!",
		Status:  200,
		Data:    user,
	}

	return newUser
}
