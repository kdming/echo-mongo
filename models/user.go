package models

import (
	"gopkg.in/mgo.v2/bson"
)

// 用户信息
type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`     // 用户名
	Password string        `bson:"password"` // 用户密码
	Salt     string        `bson:"salt"`     // 存储盐值
}
