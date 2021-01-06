package router

import (
	"Housekeeping/api"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("banner", api.GetBanners)
	}
	return r
}
