package auth_service

import (
	"Chat/dto"
	auth_model "Chat/modules/auth/model"
)

func Register_Service(user dto.TRegister) dto.TRegisterResponse {

	newUser := auth_model.CreateUser(user)

	return newUser
}
