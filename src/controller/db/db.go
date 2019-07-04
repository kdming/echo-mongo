// 数据库增删改查模块
package db

import (
	db "web_model/src/common/mongo"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"strconv"
	"fmt"
)

// 创建数据
func Create(collection string, item interface{}) error {
	s, c := db.GetSession(collection)
	defer s.Close()
	err := c.Insert(item)
	return err
}

// 获取一条数据
func FindOne(collection string, object, query interface{}, extraQuery func(*mgo.Query) *mgo.Query) error {
	s, c := db.GetSession(collection)
	defer s.Close()
	q := c.Find(query)
	if extraQuery != nil {
		q = extraQuery(q)
	}
	err := q.One(object)
	return err
}

// 获取全部数据
func Find(collection string, objects, query interface{}, extraQuery func(*mgo.Query) *mgo.Query) error {
	s, c := db.GetSession(collection)
	defer s.Close()
	q := c.Find(query)
	if extraQuery != nil {
		q = extraQuery(q)
	}
	err := q.All(objects)
	return err
}

// 更新数据
func Update(collection string, query, updatedItem interface{}, UpdateAll bool) error {
	s, c := db.GetSession(collection)
	defer s.Close()
	var err error
	if UpdateAll {
		_, err = c.UpdateAll(query, bson.M{"$set": updatedItem})
	} else {
		err = c.Update(query, updatedItem)
	}
	return err
}

// 删除数据
func Delete(collection string, query interface{}, RemoveAll bool) error {
	s, c := db.GetSession(collection)
	defer s.Close()
	var err error
	if RemoveAll {
		_, err = c.RemoveAll(query)
		return err
	}
	err = c.Remove(query)
	return err
}

// 聚合查询一条数据
func AggGet(collection string, object interface{}, queries ...bson.M) error {
	s, c := db.GetSession(collection)
	defer s.Close()
	err := c.Pipe(queries).One(object)
	return err
}

// 聚合查询全部数据
func AggGetAll(collection string, objects interface{}, queries ...bson.M) error {
	s, c := db.GetSession(collection)
	defer s.Close()
	err := c.Pipe(queries).All(objects)
	return err
}

// 分页查询
func FindByPage(collection string, query bson.M, page string, limit string) []bson.M {
	// string to int
	iPage, _ := strconv.Atoi(page)
	iLimit, _ := strconv.Atoi(limit)

	var skip int
	if iPage == 1 {
		skip = 0
	} else {
		skip = (iPage - 1) * iLimit
	}
	// query
	s, c := db.GetSession(collection)
	defer s.Close()
	result := []bson.M{}
	err := c.Find(query).Limit(iLimit).Skip(skip).All(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}