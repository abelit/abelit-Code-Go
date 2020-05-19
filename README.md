# abelit-Go
Learn Go Programming language.

# Git&go get使用Shadowsocks代理

* Linux/Mac

```bash
export http_proxy="http://127.0.0.1:1080"
export https_proxy="http://127.0.0.1:1080"

# Socks5
export http_proxy="socks5://127.0.0.1:1080"
export https_proxy="socks5://127.0.0.1:1080"
```

* Windows

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

# go get使用goproxy

* Linux/Mac
```bash
export GOPROXY="https://goproxy.io"
```

* Windows powershell
```powershell
$env:GOPROXY = "https://goproxy.io"
```

# go module
* Linux/Mac
```bash
set GO111MODULE=on
```

* Windows
```powershell
export GO111MODULE=on
```

# 安装VSCode的Go开发依赖包
```
go get -u -v github.com/mdempsky/gocode
go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -u -v github.com/ramya-rao-a/go-outline
go get -u -v github.com/acroca/go-symbols
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v github.com/sqs/goreturns
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v github.com/go-delve/delve/cmd/dlv
go get -u -v github.com/stamblerre/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v github.com/sqs/goreturns
go get -u -v golang.org/x/lint/golint
```

```
go install github.com/mdempsky/gocode
go install github.com/uudashr/gopkgs/cmd/gopkgs
go install github.com/ramya-rao-a/go-outline
go install github.com/acroca/go-symbols
go install golang.org/x/tools/cmd/guru
go install github.com/sqs/goreturns
go install golang.org/x/tools/cmd/gorename
go install github.com/go-delve/delve/cmd/dlv
go install github.com/stamblerre/gocode
go install github.com/rogpeppe/godef
go install github.com/sqs/goreturns
go install golang.org/x/lint/golint
```

# 安装beego框架和bee工具
```powershell
go get -u github.com/astaxie/beego
go get -u github.com/beego/bee
```

```powershell
go install github.com/astaxie/beego
go install github.com/beego/bee
```

# Q&A
Q1. git或go get的其他包源码无法同步到自建的github repository？

> A1. 删除下载的repository下的".git"文件夹。