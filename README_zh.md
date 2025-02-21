<div align="center">
<a href="https://github.com/yangyezhuang/git-proxy"><img src="build/appicon.png" width="120"/></a>
</div>
<h1 align="center">GitProxy</h1>
<h4 align="center"><strong><a href="/">English</a></strong> | 简体中文</h4>
<div align="center">

<strong>一个现代化轻量级的跨平台Git代理切换桌面客户端，支持Mac、Windows和Linux</strong>
</div>

<picture>
 <source media="(prefers-color-scheme: dark)" srcset="docs/screenshots/light_en1.png">
 <source media="(prefers-color-scheme: light)" srcset="docs/screenshots/light_en1.png">
 <img alt="screenshot" src="docs/screenshots/light_en1.png">
</picture>

<picture>
 <source media="(prefers-color-scheme: dark)" srcset="docs/screenshots/light_en2.png">
 <source media="(prefers-color-scheme: light)" srcset="docs/screenshots/light_en2.png">
 <img alt="screenshot" src="docs/screenshots/light_en2.png">
</picture>

## 功能特性

* 极度轻量，基于Webview2，无内嵌浏览器
* 界面精美易用，提供浅色/深色主题
* 多国语言支持：英文/中文
* 快速的连接管理：支持默认模式/HTTP代理/SOCKS5代理


## 安装

提供Mac、Windows和Linux安装包，可[免费下载](https://github.com/yangyezhuang/git-proxy/releases)。


## 构建项目

### 运行环境要求

* Go（最新版本）
* Node.js >= 16
* NPM >= 9

### 安装wails

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### 拉取代码

```bash
git clone https://github.com/yangyezhuang/git-proxy --depth=1
```

### 构建前端代码

```bash
npm install --prefix ./frontend
```

或者

```bash
cd frontend
npm install
```

### 编译运行开发版本

```bash
wails dev
```

## 关于

如果你也同为独立开发者（团队），喜欢开源，可以关注一下，探讨心得，反馈意见，交个朋友。


### 赞助

该项目完全为爱发电，如果对你有所帮助，可以请作者喝杯咖啡 ☕️

* 微信赞赏

<img src="docs/images/wechat_sponsor.png" alt="wechat" width="200" />