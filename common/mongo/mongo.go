package mongo

import (
	config "web_model/common/config"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
)

// 全局Session管理，其他模块都从这里拿session
var GlobalMgoSession *mgo.Session
// 数据库名称
var DataBase string

// 连接数据库
func Connect () bool {
	// 读取配置文件，初始化mongo链接配置信息
	conf := config.GetConfig()
	DataBase = conf.DB_NAME
	info := &mgo.DialInfo{
		Addrs:    []string{conf.DB_HOST},
		Timeout:  60 * time.Second,
		Database: conf.DB_NAME,
		Username: conf.DB_USER,
		Password: conf.DB_PWD,
	}
	// 链接数据库
	globalMgoSession, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
		return false
	}
	GlobalMgoSession = globalMgoSession
	GlobalMgoSession.SetMode(mgo.Monotonic, true)
	//default is 4096
	GlobalMgoSession.SetPoolLimit(500) // 设置session连接池最大值
	fmt.Println("数据库连接成功")
	return true
}

// 获取session
func GetSession (collection string) (*mgo.Session, *mgo.Collection) {
	s := GlobalMgoSession.Copy()
	c := s.DB(DataBase).C(collection)
	return s, c
}