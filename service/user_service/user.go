package user_service

import (
	"fmt"
	"time"
	"web_model/dao"
	"web_model/models"
	"web_model/pkg/encrypt"

	"gopkg.in/mgo.v2/bson"
)

type UserSvc struct {
}

func (u *UserSvc) Add(user *models.User) error {
	user.Salt = time.Now().String()
	user.Password = encrypt.Md5WithSalt(user.Password, user.Salt)
	if err := dao.Create("user", *user); err != nil {
		return err
	}
	return nil
}

func (u *UserSvc) GetById(id bson.ObjectId) error {
	return nil
}

func (u *UserSvc) GetByName(name string) (*models.User, error) {
	user := &models.User{}
	if err := dao.FindOne("user", &user, &bson.M{"name": name}, nil); err != nil {
		fmt.Println(err.Error(), "err")
		return nil, err
	}
	return user, nil
}
