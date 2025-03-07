# FastGin 服务框架

FastGin 是一个基于 Gin 框架的 Go 语言服务端脚手架，集成了常用的中间件和功能组件。

## 功能特性

- JWT 认证
- Casbin 权限管理
- Swagger API 文档
- Zap 日志系统
- 配置热重载
- 优雅关闭
- CORS 支持
- 数据库集成 (MySQL)

## 环境要求

- Go 1.16 或更高版本
- MySQL 5.7 或更高版本

## 安装部署

### MacOS 环境

1. 安装 Go
```bash
brew install go
```

2. 安装 MySQL
```bash
brew install mysql
```

3. 启动 MySQL
```bash
brew services start mysql
```

4. 克隆项目
```bash
git clone <项目地址>
cd fastgin
```

5. 安装依赖
```bash
go mod tidy
```

6. 安装 Swagger
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

7. 生成 API 文档
```bash
swag init -g cmd/server/main.go
```

8. 创建数据库
```bash
mysql -u root -p
CREATE DATABASE fastgin;
```

9. 修改配置文件
```bash
cp config/config.yaml.example config/config.yaml
# 编辑 config.yaml 修改数据库配置
```

10. 运行项目
```bash
go run cmd/server/main.go
```

### Windows 环境

1. 安装 Go
   - 下载 Go 安装包：https://golang.org/dl/
   - 运行安装程序
   - 添加环境变量：
     - GOPATH: `C:\Users\{用户名}\go`
     - Path: 添加 `%GOPATH%\bin`

2. 安装 MySQL
   - 下载 MySQL 安装包：https://dev.mysql.com/downloads/installer/
   - 运行安装程序
   - 记住设置的 root 密码

3. 克隆项目
```bash
git clone <项目地址>
cd fastgin
```

4. 安装依赖
```bash
go mod tidy
```

5. 安装 Swagger
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

6. 生成 API 文档
```bash
swag init -g cmd/server/main.go
```

7. 创建数据库
```sql
mysql -u root -p
CREATE DATABASE fastgin;
```

8. 修改配置文件
```bash
copy config/config.yaml.example config/config.yaml
# 编辑 config.yaml 修改数据库配置
```

9. 运行项目
```bash
go run cmd/server/main.go
```

## 配置说明

配置文件位于 `config/config.yaml`：

```yaml
server:
  port: "8080"     # 服务端口
  mode: "debug"    # 运行模式

database:
  driver: "mysql"  # 数据库类型
  host: "localhost"
  port: "3306"
  username: "root"
  password: "your-password"
  dbname: "fastgin"

jwt:
  secret: "your-secret"  # JWT 密钥
  expire: 7200          # Token 过期时间（秒）

log:
  level: "debug"       # 日志级别
  filename: "logs/app.log"
  maxSize: 128        # 单个日志文件大小限制（MB）
  maxBackups: 30      # 最大日志文件保留数
  maxAge: 7          # 日志文件保留天数
  compress: true     # 是否压缩
```

## 访问服务

- API 服务：http://localhost:8080
- Swagger 文档：http://localhost:8080/swagger/index.html

## 常见问题

1. 如果遇到权限问题：
   - MacOS：确保 MySQL 用户有适当权限
   - Windows：以管理员身份运行命令行

2. 如果 `swag` 命令未找到：
   - 确保 `GOPATH/bin` 在系统 PATH 中
   - 重新打开终端

3. 如果数据库连接失败：
   - 检查 MySQL 服务是否启动
   - 验证数据库配置信息是否正确

## 开发指南

详细的开发文档请参考 [开发指南](docs/development.md)

## 许可证

MIT License
