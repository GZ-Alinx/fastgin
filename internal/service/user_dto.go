package service

// LoginRequests 登录请求参数
type LoginRequests struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"123456"`
}

// LoginResponse 登录响应
type LoginResponses struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIs..."`
}

// CreateUserRequest 创建用户请求
type CreateUserRequests struct {
	Username string `json:"username" binding:"required" example:"newuser"`
	Password string `json:"password" binding:"required" example:"123456"`
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Nickname string `json:"nickname" example:"新用户"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequests struct {
	Email    string `json:"email" binding:"omitempty,email" example:"user@example.com"`
	Nickname string `json:"nickname" example:"用户昵称"`
	Password string `json:"password" example:"123456"`
}

// UserInfo 用户信息
type UserInfo struct {
	ID        uint   `json:"id" example:"1"`
	Username  string `json:"username" example:"admin"`
	Email     string `json:"email" example:"admin@example.com"`
	Nickname  string `json:"nickname" example:"管理员"`
	CreatedAt string `json:"created_at" example:"2023-01-01 12:00:00"`
	UpdatedAt string `json:"updated_at" example:"2023-01-01 12:00:00"`
}

// PageResponse 分页响应
type PageResponse struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}
