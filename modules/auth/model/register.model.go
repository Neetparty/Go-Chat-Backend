package auth_model

import (
	"Chat/configs"
	"Chat/dto"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user dto.TRegister) dto.TRegisterResponse {
	client := configs.Mongo
	usr_coll := client.Database("go-chat").Collection("user")
	room_col := client.Database("go-chat").Collection("room")

	objectID, err := usr_coll.InsertOne(context.TODO(), user)

	if err != nil {
		panic(err)
	}

	_, room_err := room_col.InsertOne(context.TODO(), bson.D{{Key: "user_id", Value: objectID.InsertedID}})

	if room_err != nil {
		panic(room_err)
	}

	filter := bson.D{{Key: "_id", Value: objectID.InsertedID}}

	get_err := usr_coll.FindOne(context.TODO(), filter).Decode(&user)

	if get_err != nil {
		usr_coll.DeleteOne(context.TODO(), filter)
		panic(get_err)
	}

	newUser := dto.TRegisterResponse{
		Message: "Thank you for registering!",
		Status:  200,
		Data:    user,
	}

	return newUser
}
