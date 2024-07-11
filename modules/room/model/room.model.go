package model

import (
	"Chat/configs"
	"Chat/dto"
	"Chat/libs"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateRoom(room dto.TRoom) dto.TRoomResponse {
	client := configs.Mongo
	room_coll := client.Database("go-chat").Collection("room")

	dv := libs.CreateDefaultValue()

	newRoom := dto.TRoom{
		Users:     room.Users,
		CreatedAt: dv.CreatedAt,
		UpdatedAt: dv.UpdatedAt,
		IsDeleted: dv.IsDeleted,
	}

	objectID, err := room_coll.InsertOne(context.TODO(), newRoom)

	if err != nil {
		panic(err)
	}

	filter := bson.D{{Key: "_id", Value: objectID.InsertedID}}

	get_err := room_coll.FindOne(context.TODO(), filter).Decode(&newRoom)

	if get_err != nil {
		panic(get_err)
	}

	return dto.TRoomResponse{
		Message: "Room created",
		Status:  200,
		Data:    newRoom,
	}
}
