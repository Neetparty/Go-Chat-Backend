package service

import (
	"Chat/dto"
	"Chat/libs"
	model "Chat/modules/auth/model"
)

func Login_Service(user dto.TLogin) dto.TLoginResponse {

	usr := model.Login(user)

	if usr.Email == "" {
		return dto.TLoginResponse{
			Message: "User not found",
			Status:  404,
		}
	}

	if !libs.VerifyPassword(user.Password, usr.Password) {
		return dto.TLoginResponse{
			Message: "Wrong password",
			Status:  401,
		}
	}

	return dto.TLoginResponse{
		Message: "Login successful",
		Status:  200,
		Data:    usr,
	}
}
