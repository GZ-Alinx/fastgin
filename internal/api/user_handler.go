package api

import (
	"fastgin/internal/middleware"
	"fastgin/internal/service"
	"fastgin/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Login 保持原有的登录函数不变
// @Summary 用户登录
// @Description 处理用户登录请求，返回登录凭证
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body service.LoginRequests true "登录请求参数"
// @Success 200 {object} utils.Response{data=service.LoginResponses} "登录成功返回token信息"
// @Failure 400 {object} utils.Response{data=string} "无效的请求参数"
// @Failure 401 {object} utils.Response{data=string} "登录失败"
// @Router /login [post]
func Login(c *gin.Context) {
	var req service.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.Logger.Warn("无效的请求参数",
			zap.Error(err))
		utils.Error(c, 400, "无效的请求参数")
		return
	}

	resp, err := service.Login(&req)
	if err != nil {
		middleware.Logger.Error("登录失败",
			zap.String("username", req.Username),
			zap.Error(err))
		utils.Error(c, 401, err.Error())
		return
	}

	middleware.Logger.Info("登录成功",
		zap.String("username", req.Username))
	utils.Success(c, resp)
}

// CreateUser 创建用户
// @Summary 创建新用户
// @Description 创建一个新的用户账号
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body service.CreateUserRequests true "创建用户请求参数"
// @Success 200 {object} utils.Response{data=string} "创建用户成功"
// @Failure 400 {object} utils.Response{data=string} "无效的请求参数"
// @Failure 500 {object} utils.Response{data=string} "创建用户失败"
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var req service.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.Logger.Warn("创建用户：无效的请求参数", zap.Error(err))
		utils.Error(c, 400, "无效的请求参数")
		return
	}

	err := service.CreateUser(&req)
	if err != nil {
		middleware.Logger.Error("创建用户失败", zap.Error(err))
		utils.Error(c, 500, err.Error())
		return
	}

	utils.Success(c, "创建用户成功")
}

// UpdateUser 更新用户信息
// @Summary 更新用户信息
// @Description 更新指定用户的信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path uint true "用户ID"
// @Param request body service.UpdateUserRequests true "更新用户请求参数"
// @Success 200 {object} utils.Response{data=string} "更新用户成功"
// @Failure 400 {object} utils.Response{data=string} "无效的请求参数"
// @Failure 500 {object} utils.Response{data=string} "更新用户失败"
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req service.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.Logger.Warn("更新用户：无效的请求参数", zap.Error(err))
		utils.Error(c, 400, "无效的请求参数")
		return
	}

	err := service.UpdateUser(uint(userID), &req)
	if err != nil {
		middleware.Logger.Error("更新用户失败", zap.Error(err))
		utils.Error(c, 500, err.Error())
		return
	}

	utils.Success(c, "更新用户成功")
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 删除指定的用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path uint true "用户ID"
// @Success 200 {object} utils.Response{data=string} "删除用户成功"
// @Failure 500 {object} utils.Response{data=string} "删除用户失败"
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	err := service.DeleteUser(uint(userID))
	if err != nil {
		middleware.Logger.Error("删除用户失败", zap.Error(err))
		utils.Error(c, 500, err.Error())
		return
	}

	utils.Success(c, "删除用户成功")
}

// GetUser 获取用户信息
// @Summary 获取用户信息
// @Description 获取指定用户的详细信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path uint true "用户ID"
// @Success 200 {object} utils.Response{data=service.UserInfo} "获取用户信息成功"
// @Failure 500 {object} utils.Response{data=string} "获取用户信息失败"
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user, err := service.GetUser(uint(userID))
	if err != nil {
		middleware.Logger.Error("获取用户信息失败", zap.Error(err))
		utils.Error(c, 500, err.Error())
		return
	}

	utils.Success(c, user)
}

// ListUsers 获取用户列表
// @Summary 获取用户列表
// @Description 分页获取用户列表
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param page query int false "页码，默认为1"
// @Param page_size query int false "每页数量，默认为10"
// @Success 200 {object} utils.Response{data=[]service.UserInfo,total=int64} "获取用户列表成功"
// @Failure 500 {object} utils.Response{data=string} "获取用户列表失败"
// @Router /users [get]
func ListUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	users, total, err := service.ListUsers(page, pageSize)
	if err != nil {
		middleware.Logger.Error("获取用户列表失败", zap.Error(err))
		utils.Error(c, 500, err.Error())
		return
	}

	utils.SuccessWithPage(c, users, total, page, pageSize)
}
