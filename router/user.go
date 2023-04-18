package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jimmyann2023/Gin/api"
	"net/http"
)

func InitUserRoutes() {
	RegisterRouter(func(PublicRouter *gin.RouterGroup, AuthRouter *gin.RouterGroup) {

		// api 实体处理对象
		userApi := api.NewUserApi()

		// 创建 User 的公开路由
		PublicUser := PublicRouter.Group("user")
		{
			PublicUser.POST("/login", userApi.Login)
		}

		// 创建 User的 鉴权路由组
		AuthUser := AuthRouter.Group("user")
		{
			AuthUser.GET("", func(ctx *gin.Context) {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
					"data": []map[string]any{
						{"id": 1, "name": "张三"},
						{"id": 2, "name": "李四"},
					},
				})
			})
			AuthUser.GET("/:id", func(ctx *gin.Context) {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
					"id":   1,
					"name": "张三",
				})
			})
		}

	})
}
