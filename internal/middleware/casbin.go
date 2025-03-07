package middleware

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewCasbinMiddleware(db *gorm.DB) (*casbin.Enforcer, error) {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}

	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		return nil, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, err
	}

	return enforcer, nil
}

func Authorize(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		if role == "" {
			Logger.Warn("未找到用户角色")
			c.AbortWithStatusJSON(401, gin.H{"code": 401, "message": "未授权"})
			return
		}

		obj := c.Request.URL.Path
		act := c.Request.Method

		Logger.Debug("权限检查",
			zap.String("role", role),
			zap.String("path", obj),
			zap.String("method", act))

		ok, err := e.Enforce(role, obj, act)
		if err != nil {
			Logger.Error("权限检查错误",
				zap.Error(err),
				zap.String("role", role),
				zap.String("path", obj),
				zap.String("method", act))
			c.AbortWithStatusJSON(500, gin.H{"code": 500, "message": "权限检查错误"})
			return
		}

		if !ok {
			Logger.Warn("权限不足",
				zap.String("role", role),
				zap.String("path", obj),
				zap.String("method", act))
			c.AbortWithStatusJSON(403, gin.H{"code": 403, "message": "没有权限"})
			return
		}

		Logger.Debug("权限检查通过",
			zap.String("role", role),
			zap.String("path", obj),
			zap.String("method", act))
		c.Next()
	}
}
