package router

import (
	"fastgin/docs" // 这个包会在运行 swag init 后自动生成
	"fastgin/internal/api"
	"fastgin/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// 注册中间件（调整顺序）
	r.Use(middleware.Loggers())  // 日志中间件放在最前面
	r.Use(middleware.Recovery()) // 添加恢复中间件
	r.Use(middleware.Cors())     // CORS中间件放在后面

	// 注册 Swagger 路由
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 初始化 Casbin
	Enforcer, err := middleware.NewCasbinMiddleware(db)
	if err != nil {
		panic(err)
	}

	// 添加公开路由组
	public := r.Group("/api")
	{
		public.POST("/login", api.Login) // 登录接口
	}

	// 用户模块路由
	UserRouter(r, Enforcer)

	return r
}
