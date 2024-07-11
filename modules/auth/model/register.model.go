package model

import (
	"Chat/configs"
	"Chat/dto"
	"Chat/libs"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user dto.TRegister) dto.TRegisterResponse {
	client := configs.Mongo
	usr_coll := client.Database("go-chat").Collection("user")

	dv := libs.CreateDefaultValue()

	nUser := dto.TRegister{
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Image:     user.Image,
		CreatedAt: dv.CreatedAt,
		UpdatedAt: dv.UpdatedAt,
		IsDeleted: dv.IsDeleted,
	}

	objectID, err := usr_coll.InsertOne(context.TODO(), nUser)

	if err != nil {
		panic(err)
	}

	filter := bson.D{{Key: "_id", Value: objectID.InsertedID}}

	get_err := usr_coll.FindOne(context.TODO(), filter).Decode(&nUser)

	if get_err != nil {
		usr_coll.DeleteOne(context.TODO(), filter)
		panic(get_err)
	}

	newUser := dto.TRegisterResponse{
		Message: "Thank you for registering!",
		Status:  200,
		Data:    nUser,
	}

	return newUser
}
