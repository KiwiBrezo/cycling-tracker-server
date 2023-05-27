package Repositories

import (
	"context"
	"cycling-tracker-server/Models"
	"cycling-tracker-server/Mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type UserRepository struct {
}

func (r *UserRepository) LoginUser(login Models.Login) (user Models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userNotChecked Models.User
	err = Mongodb.GetCollection("users").FindOne(ctx, bson.D{{"username", login.Username}}).Decode(&userNotChecked)
	if err != nil || userNotChecked.Password != login.Password {
		return
	}

	return userNotChecked, nil
}
