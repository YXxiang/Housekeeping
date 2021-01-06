/**
 * @Author: LOFTER
 * @Description:
 * @File:  main
 * @Date: 2021/1/5 3:23 下午
 */
package main

import (
	"Housekeeping/cache"
	"Housekeeping/model"
	"Housekeeping/router"
	"Housekeeping/util"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
	model.Initialized()
	cache.Redis()
	util.InitLogger()
}

func main() {

	// 装载路由
	r := router.NewRouter()
	_ = r.Run(":8080")

}
