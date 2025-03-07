package middleware

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

const (
	JWTSecret = "your-secret-key" // 实际应用中应该从配置文件读取
	Bearer    = "Bearer "
)

func GenerateToken(userID uint, username, role string) (string, error) {
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWTSecret))

	if err != nil {
		Logger.Error("生成token失败",
			zap.Error(err),
			zap.Uint("userID", userID),
			zap.String("username", username))
		return "", err
	}

	Logger.Info("生成token成功",
		zap.Uint("userID", userID),
		zap.String("username", username))
	return tokenString, nil
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			Logger.Warn("未提供token")
			c.AbortWithStatusJSON(401, gin.H{"code": 401, "message": "未授权"})
			return
		}

		if !strings.HasPrefix(auth, Bearer) {
			Logger.Warn("无效的token格式", zap.String("token", auth))
			c.AbortWithStatusJSON(401, gin.H{"code": 401, "message": "无效的token格式"})
			return
		}

		tokenString := auth[len(Bearer):]
		token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				Logger.Warn("无效的签名方法",
					zap.String("method", token.Method.Alg()))
				return nil, errors.New("无效的签名方法")
			}
			return []byte(JWTSecret), nil
		})

		if err != nil {
			Logger.Error("token解析失败",
				zap.Error(err),
				zap.String("token", tokenString))
			c.AbortWithStatusJSON(401, gin.H{"code": 401, "message": "无效的token"})
			return
		}

		if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
			Logger.Debug("token验证成功",
				zap.Uint("userID", claims.UserID),
				zap.String("username", claims.Username),
				zap.String("role", claims.Role))

			c.Set("userID", claims.UserID)
			c.Set("username", claims.Username)
			c.Set("role", claims.Role)
			c.Next()
		} else {
			Logger.Warn("无效的token声明")
			c.AbortWithStatusJSON(401, gin.H{"code": 401, "message": "无效的token声明"})
			return
		}
	}
}
