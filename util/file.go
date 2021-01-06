/*
	文件操作辅助类，主要功能如下：
	- 文件是否存在
	- 删除文件
*/
package util

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"time"
)

//判断文件是否存在
func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

//创建文件上传目录
func MkdirFileDirectory(path string) string {
	date := time.Now().Format("2006-01-02")
	if ok := IsFileExist(path + "/" + date); !ok {
		_ = os.MkdirAll(path+"/"+date, 0755)
	}
	return path + "/" + date + "/"
}

//删除文件
func DeleteFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println(filePath+"文件删除失败，原因是", err)
	}
	return
}

//读取文件内容
func ReadFile(filePath string) string {
	if filePath == "" {
		return ""
	}
	fileSuffix := path.Ext(filePath) //获取文件后缀
	if fileSuffix == ".zip" {
		return ""
	}
	file, err := os.Open("." + filePath)
	fmt.Println(err)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(bytes)

}
