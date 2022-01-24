# minimalist-web-notepad

一款带API的开源轻量级简洁在线笔记本。

## demo

[剪切板](https://clip.ncc.cx)

## 项目地址

* [minimalist-web-notepad](https://github.com/fghwett/minimalist-web-notepad)

## 部署

### 1. 源码部署

需要GO环境才能运行( [安装Go环境](https://go.dev/doc/install) )

```shell
git clone https://github.com/fghwett/minimalist-web-notepad.git

cd minimalist-web-notepad

go run ./
```

> 可采用 `nohup` 或 `screen` 后台运行程序

### 2. docker-compose部署

需要主机上已经有docker或docker-compose

```shell
wget -O https://raw.githubusercontent.com/fghwett/minimalist-web-notepad/main/docker-compose.yml

docker-compose up -d
```

> 目前docker镜像只支持X64平台

## API

1. 修改剪切板内容

    GET/POST https://example.com/clipname?text=your content

2. 获取剪切板内容

    GET https://example.com/clipname?raw

## 相关链接

* [minimalist-web-notepad](https://github.com/pereorga/minimalist-web-notepad)
* [Minimalist-Web-Notepad-API](https://github.com/Xiaobin2333/Minimalist-Web-Notepad-API)