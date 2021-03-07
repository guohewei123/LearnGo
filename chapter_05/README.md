## Go语言的依赖管理
### 依赖管理的三个阶段 GOPATH, GOVENDOR, go mod
### 1. GOPATH 依赖管理
- GOPATH 依赖管理是通过设置依赖查找路径，来实现依赖导入。
- 缺点：1. 所有代码的依赖都放置到同一个地方，当不同项目对相同依赖的版本不一致时，将无法处理; 2. 所有的项目代码需要放置到 GOPATH 目录下的 src 目录中，只有这样代码运行时，才能找到 GOPATH 下的依赖。
```shell script
# 配置成GOPATH模式
go env  # 查看配置项
go env -w GO111MODULE=off   # 关闭go mod 模式，临时设置 export GO111MODULE=off
go env -w GOPATH=/root/go   # 临时设置 export GOPATH=/root/go 
# 需要在GOPATH目录下创建 src 文件夹，用于放置自己编写的代码
mkdir /root/go/src

# 安装测试依赖
go get -u go.uber.org/zap 
```

### 2. GOVENDOR 依赖管理
- GOVENDOR 依赖管理是通过在当前项目目录下新建 vendor 文件夹，放置本项目使用依赖库，
代码运行时，优先从 vendor 目录下查找依赖，如果没有找到，就会到 GOPATH 目录下查找依赖库。 
- 缺点：需要用户手动拷贝依赖库到 vender 目录下，自动化操作时需要使用第三方依赖库（如：glide, dep, go dep, ...）

### 3. go mod 依赖管理
- 配置系统为 go mod 模式
    ```shell script
    go env -w GO111MODULE=on  # 配置 go mode 模式
    go env -w GOPROXY=https://goproxy.cn,direct  # 使用国内源
    ```
- 由go命令统一管理，用户不必关心项目所在目录和vendor目录
- 初始化 go mod: `go mode init [name], eg: go mode init gomodtest`， 执行命令后将生成 go.mod 文件
- 安装依赖，在 go.mod 文件所在目录下执行 
    ```shell script
    go get -u go.uber.org/zap   # 安装最新版本
    go get go.uber.org/zap@1.11  # 安装指定版本
    ```
- 更新依赖后通过整理命令清理旧的依赖记录
    ```shell script
    go get go.uber.org/zap@1.11   # 安装zap@1.11
    go get go.uber.org/zap@1.12   # 更新到zap@1.12
    go mod tidy    # 重新整理依赖记录
    ```
- 将旧项目迁移到go mod:
    ```shell script
    go mod init [mod_name]  # 初始化 go.mod 文件
    go build ./...          # 将所有项目依赖记录到go.mod中
    ```

