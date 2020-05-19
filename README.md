# abelit-Go

Learn Go Programming language.

# Git&go get 使用 Shadowsocks 代理

- Linux/Mac

```bash
export http_proxy="http://127.0.0.1:1080"
export https_proxy="http://127.0.0.1:1080"

# Socks5
export http_proxy="socks5://127.0.0.1:1080"
export https_proxy="socks5://127.0.0.1:1080"
```

- Windows

```powershell
set http_proxy="http://127.0.0.1:1080"
set https_proxy="http://127.0.0.1:1080"

# Socks5
set http_proxy="socks5://127.0.0.1:1080"
set https_proxy="socks5://127.0.0.1:1080"

# 或
$env:http_proxy="http://127.0.0.1:1080"
$env:https_proxy="http://127.0.0.1:1080"

# Socks5
$env:http_proxy="socks5://127.0.0.1:1080"
$env:https_proxy="socks5://127.0.0.1:1080"
```

# go get 使用 goproxy

- Go 1.13 及以上（推荐）

```go
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
```

- Linux/Mac (如果您使用的 Go 版本是 1.12 及以下)

```bash
$ export GO111MODULE=on
$ export GOPROXY=https://goproxy.cn

或

$ echo "export GO111MODULE=on" >> ~/.profile
$ echo "export GOPROXY=https://goproxy.cn" >> ~/.profile
$ source ~/.profile
```

- Windows (如果您使用的 Go 版本是 1.12 及以下)

```powershell
C:\> $env:GO111MODULE = "on"
C:\> $env:GOPROXY = "https://goproxy.cn"

或

1. 打开“开始”并搜索“env”
2. 选择“编辑系统环境变量”
3. 点击“环境变量…”按钮
4. 在“<你的用户名> 的用户变量”章节下（上半部分）
5. 点击“新建…”按钮
6. 选择“变量名”输入框并输入“GO111MODULE”
7. 选择“变量值”输入框并输入“on”
8. 点击“确定”按钮
9. 点击“新建…”按钮
10. 选择“变量名”输入框并输入“GOPROXY”
11. 选择“变量值”输入框并输入“https://goproxy.cn”
12. 点击“确定”按钮
```

- 自托管 Go 模块代理

你的代码永远只属于你自己，因此我们向你提供目前世界上最炫酷的自托管 Go 模块代理搭建方案。通过使用 Goproxy 这个极简主义项目，你可以在现有的任意 Web 服务中轻松地加入 Go 模块代理支持，要知道 goproxy.cn 就是基于它搭建的。

创建一个名为 goproxy.go 的文件

```go
package main

import (
	"net/http"
	"os"

	"github.com/goproxy/goproxy"
)

func main() {
	g := goproxy.New()
	g.GoBinEnv = append(
		os.Environ(),
		"GOPROXY=https://goproxy.cn,direct", // 使用 goproxy.cn 作为上游代理
		"GOPRIVATE=git.example.com",         // 解决私有模块的拉取问题（比如可以配置成公司内部的代码源）
	)
	http.ListenAndServe("localhost:8080", g)
}
```

并且运行它

```
$ go run goproxy.go
```

然后通过把 GOPROXY 设置为 http://localhost:8080 来试用它。另外，我们也建议你把 GO111MODULE 设置为 on。

就这么简单，一个功能完备的 Go 模块代理就搭建成功了。事实上，你可以将 Goproxy 结合着你钟爱的 Web 框架一起使用，比如 Gin 和 Echo，你所需要做的只是多添加一条路由而已。更高级的用法请查看文档。

- 常用 GOPROXY

```
https://goproxy.baidu.com
https://goproxy.cn
https://goproxy.io
https://mirrors.aliyun.com/goproxy/
https://gonexus.dev
https://athens.azurefd.net
https://gocenter.io
https://proxy.golang.org
```

# go module

- Linux/Mac

```bash
export GO111MODULE=on
```

- Windows

```powershell
set GO111MODULE=on
```

# 安装 VSCode 的 Go 开发依赖包

```bash
go get -u -v github.com/bytbox/golint
go get -u -v github.com/golang/tools
go get -u -v github.com/lukehoban/go-outline
go get -u -v github.com/newhook/go-symbols
go get -u -v github.com/mdempsky/gocode
go get -u -v github.com/uudashr/gopkgs/v2/cmd/gopkgs
go get -u -v github.com/ramya-rao-a/go-outline
go get -u -v github.com/acroca/go-symbols
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v github.com/cweill/gotests
go get -u -v github.com/fatih/gomodifytags
go get -u -v github.com/josharian/impl
go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct
go get -u -v github.com/haya14busa/goplay/cmd/goplay
go get -u -v github.com/godoctor/godoctor
go get -u -v github.com/go-delve/delve/cmd/dlv
go get -u -v github.com/stamblerre/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v github.com/sqs/goreturns
go get -u -v golang.org/x/lint/golint
go get -u -v github.com/gin-gonic/gin
go get -u -v golang.org/x/tools/gopls
```

```bash
go install github.com/bytbox/golint
go install github.com/golang/tools
go install github.com/lukehoban/go-outline
go install github.com/newhook/go-symbols
go install github.com/mdempsky/gocode
go install github.com/uudashr/gopkgs/v2/cmd/gopkgs
go install github.com/ramya-rao-a/go-outline
go install github.com/acroca/go-symbols
go install golang.org/x/tools/cmd/guru
go install golang.org/x/tools/cmd/gorename
go install github.com/cweill/gotests
go install github.com/fatih/gomodifytags
go install github.com/josharian/impl
go install github.com/davidrjenni/reftools/cmd/fillstruct
go install github.com/haya14busa/goplay/cmd/goplay
go install github.com/godoctor/godoctor
go install github.com/go-delve/delve/cmd/dlv
go install github.com/stamblerre/gocode
go install github.com/rogpeppe/godef
go install github.com/sqs/goreturns
go install golang.org/x/lint/golint
go install github.com/gin-gonic/gin
go install golang.org/x/tools/gopls
```

# 安装 beego 框架和 bee 工具

```powershell
go get -u github.com/astaxie/beego
go get -u github.com/beego/bee
```

```powershell
go install github.com/astaxie/beego
go install github.com/beego/bee
```

# Q&A

Q1. git 或 go get 的其他包源码无法同步到自建的 github repository？

> A1. 删除下载的 repository 下的".git"文件夹。
