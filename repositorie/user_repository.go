package repositorie

import (
	"context"
	"cycling-tracker-server/models"
	"cycling-tracker-server/mongo_db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type UserRepository struct {
}

func (r *UserRepository) LoginUser(login models.Login) (user models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userNotChecked models.User
	err = mongo_db.GetCollection("users").FindOne(ctx, bson.D{{"username", login.Username}}).Decode(&userNotChecked)
	if err != nil || userNotChecked.Password != login.Password {
		return
	}

	return userNotChecked, nil
}
