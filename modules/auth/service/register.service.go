package service

import (
	"Chat/dto"
	"Chat/libs"
	"Chat/modules/auth/model"
	room_service "Chat/modules/room/service"
)

func Register_Service(user dto.TRegister) dto.TRegisterResponse {

	pwd, err := libs.HashPassword(user.Password)

	if err != nil {
		panic(err)
	}

	user.Password = pwd

	newUser := model.CreateUser(user)

	room := dto.TRoom{
		Users: []string{user.Username},
	}

	room_service.CreateRoom_Service(room)

	return newUser
}
