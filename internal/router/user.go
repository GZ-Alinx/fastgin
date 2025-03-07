package router

import (
	"fastgin/internal/api"
	"fastgin/internal/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine, Enforcer *casbin.Enforcer) {
	// 需要认证的路由组
	authorized := r.Group("/api")
	authorized.Use(middleware.JWTAuth())
	authorized.Use(middleware.Authorize(Enforcer))
	{
		// 用户管理接口
		users := authorized.Group("/users")
		{
			users.POST("", api.CreateUser)
			users.PUT("/:id", api.UpdateUser)
			users.DELETE("/:id", api.DeleteUser)
			users.GET("/:id", api.GetUser)
			users.GET("", api.ListUsers)
		}
	}
}
