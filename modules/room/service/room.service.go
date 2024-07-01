package service

import (
	"Chat/dto"
	"Chat/modules/room/model"
)

func CreateRoom_Service(room dto.TRoom) dto.TRoomResponse {
	return model.CreateRoom(room)
}
