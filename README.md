[![pipeline status](https://api.travis-ci.org/assetcloud/assetchain.svg?branch=master)](https://travis-ci.org/assetcloud/assetchain/)
[![Go Report Card](https://goreportcard.com/badge/github.com/assetcloud/assetchain)](https://goreportcard.com/report/github.com/assetcloud/assetchain)
[![Windows Build Status](https://ci.appveyor.com/api/projects/status/github/assetcloud/assetchain?svg=true&branch=master&passingText=Windows%20-%20OK&failingText=Windows%20-%20failed&pendingText=Windows%20-%20pending)](https://ci.appveyor.com/project/assetcloud/assetchain)
[![Macos Build Status](https://github.com/assetcloud/assetchain/actions/workflows/MacOS.yml/badge.svg)](https://github.com/assetcloud/assetchain/actions/workflows/MacOS.yml)

# 基于 chain 区块链开发 框架 开发的 assetchain 系统（v6.8.16）

```
注意: master 分支不是 发布版本, 不要用于生产环境
```

# bug 奖励

我们会对 bug 评价 4 个等级(不会奖励人民币，等值虚拟资产)。
只有影响现有在线运行系统的，并且会产生严重分叉等行为的，才会评价为 L3

```
L0 1000
L1 3000
L2 10000
L3 20000
```

## 安装

#### golang 1.19+

#### 支持 make file 的平台

```
git clone https://github.com/assetcloud/assetchain $GOPATH/src/github.com/assetcloud/assetchain

//开启mod功能
export GO111MODULE=on

//国内用户需要导入阿里云代理，用于下载依赖包
export GOPROXY=https://mirrors.aliyun.com/goproxy

cd $GOPATH/src/github.com/assetcloud/assetchain

make
```

就可以完成编译安装

### 更新 go.mod

```
make update
```
