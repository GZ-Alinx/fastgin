1. 首先确保 GOPATH 已经正确设置并添加到 PATH 中：

```bash
echo 'export GOPATH=$HOME/go' >> ~/.zshrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.zshrc
source ~/.zshrc
```

2. 然后重新安装 swag：

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

3. 验证安装：

```bash
which swag
```

4. 现在直接运行（不需要 sudo）：

```bash
swag init -g cmd/server/main.go
```

解释：
1. 使用 `sudo` 运行 `swag` 时会使用 root 用户的 PATH，而不是你的用户 PATH，所以找不到命令
2. `swag` 应该安装在你的用户目录下的 `$GOPATH/bin` 中，不需要 sudo 权限
3. 确保将 `$GOPATH/bin` 添加到 PATH 中后，就可以直接使用 `swag` 命令了

如果执行完这些步骤后还是有问题，可以尝试：

```bash
go get -u github.com/swaggo/swag/cmd/swag
```

然后重新执行 `swag init` 命令。