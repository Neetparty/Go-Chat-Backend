package model

import (
	"Chat/configs"
	"Chat/dto"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func Login(user dto.TLogin) dto.TRegister {
	client := configs.Mongo
	usr_coll := client.Database("go-chat").Collection("user")

	filter := bson.D{{Key: "email", Value: user.Email}}

	var result dto.TRegister

	get_err := usr_coll.FindOne(context.TODO(), filter).Decode(&result)

	if get_err != nil {
		panic(get_err)
	}

	return result
}
