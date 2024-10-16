package login

import "github.com/gin-gonic/gin"

// RegisterRoutes 注册 User 模块的路由
func RegisterRoutes(r *gin.RouterGroup) {
	r.POST("login", LoginController)
}
