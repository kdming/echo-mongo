package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

// 加密字符串(加盐生成md5)
func EncryptStr(salt, pwd string) string {
	m5 := md5.New()
	m5.Write([]byte(pwd))
	m5.Write([]byte(string(salt)))
	st := m5.Sum(nil)
	return hex.EncodeToString(st)
}

// 创建嵌套文件夹
func MkdirAll(filePath string) string {
	exists := pathExists(filePath)
	if exists == true {
		return filePath
	}
	// 递归创建目录
	err := os.MkdirAll(filePath, 0755)
	if err != nil {
		fmt.Println(err)
	}
	return filePath
}

// 判断所给路径文件/文件夹是否存在
func pathExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 公用返回数据方法
func ReturnBody(code int, data interface{}, msg string) map[string]interface{} {
	response := make(map[string]interface{})
	response["code"] = code
	response["message"] = msg
	response["data"] = data
	return response
}
