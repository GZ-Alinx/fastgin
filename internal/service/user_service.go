package service

import (
	"errors"
	"fastgin/internal/middleware"
	"fastgin/internal/model"
	"fastgin/internal/repository"

	"go.uber.org/zap"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func Login(req *LoginRequest) (*LoginResponse, error) {
	middleware.Logger.Info("用户尝试登录",
		zap.String("username", req.Username))

	var user model.User
	result := repository.DB.Where("username = ?", req.Username).First(&user)
	if result.Error != nil {
		middleware.Logger.Warn("用户不存在",
			zap.String("username", req.Username),
			zap.Error(result.Error))
		return nil, errors.New("用户不存在")
	}

	if !user.CheckPassword(req.Password) {
		middleware.Logger.Warn("密码错误",
			zap.String("username", req.Username))
		return nil, errors.New("密码错误")
	}

	// 生成 JWT token
	token, err := middleware.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		middleware.Logger.Error("生成token失败",
			zap.String("username", user.Username),
			zap.Error(err))
		return nil, err
	}

	middleware.Logger.Info("用户登录成功",
		zap.String("username", user.Username),
		zap.String("role", user.Role))

	return &LoginResponse{
		Token:    token,
		Username: user.Username,
		Role:     user.Role,
	}, nil
}

func CreateUser(req *CreateUserRequest) error {
	user := &model.User{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}

	if err := user.HashPassword(); err != nil {
		return err
	}

	return repository.DB.Create(user).Error
}

func UpdateUser(id uint, req *UpdateUserRequest) error {
	updates := make(map[string]interface{})

	if req.Username != "" {
		updates["username"] = req.Username
	}
	if req.Password != "" {
		hashedPassword, err := model.HashPassword(req.Password)
		if err != nil {
			return err
		}
		updates["password"] = hashedPassword
	}
	if req.Role != "" {
		updates["role"] = req.Role
	}

	return repository.DB.Model(&model.User{}).Where("id = ?", id).Updates(updates).Error
}

func DeleteUser(id uint) error {
	return repository.DB.Delete(&model.User{}, id).Error
}

func GetUser(id uint) (*model.User, error) {
	var user model.User
	err := repository.DB.First(&user, id).Error
	return &user, err
}

func ListUsers(page, pageSize int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	offset := (page - 1) * pageSize

	err := repository.DB.Model(&model.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = repository.DB.Offset(offset).Limit(pageSize).Find(&users).Error
	return users, total, err
}
