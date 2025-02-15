<div align="center">
<a href="https://github.com/yangyezhuang/git-proxy"><img src="build/appicon.png" width="120"/></a>
</div>
<h1 align="center">GitProxy</h1>
<h4 align="center"><strong>English</strong> | <a href="https://github.com/yangyezhuang/git-proxy/blob/main/README_zh.md">
简体中文</a></h4>



<strong>GitProxy is a modern and lightweight cross platform Git proxy switching desktop client that supports Mac, Windows, and Linux.</strong>
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

## Feature

* Super lightweight, built on Webview2, without embedded browsers (Thanks
  to [Wails](https://github.com/wailsapp/wails)).
* Provides visually and user-friendly UI, light and dark themes (Thanks to [Naive UI](https://github.com/tusen-ai/naive-ui)
  and [IconPark](https://iconpark.oceanengine.com)).
* Multi-language support ([Need more languages ? Click here to contribute](.github/CONTRIBUTING.md)).
* Better connection management: supports SSH Tunnel/SSL/Sentinel Mode/Cluster Mode/HTTP proxy/SOCKS5 proxy.


## Installation

Available to download for free from [here](https://github.com/yangyezhuang/git-proxy/releases).


## Build Guidelines

### Prerequisites

* Go (latest version)
* Node.js >= 16
* NPM >= 9

### Install Wails

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### Pull the Code

```bash
git clone https://github.com/yangyezhuang/git-proxy --depth=1
```

### Build Frontend

```bash
npm install --prefix ./frontend
```

or

```bash
cd frontend
npm install
```

### Compile and Run

```bash
wails dev
```
## About

If you are also an independent developer (team) and enjoy open source, you can follow us to discuss your experiences, provide feedback, and make friends.

### Sponsor

If this project helpful for you, feel free to buy me a cup of coffee ☕️.

* Wechat Sponsor

<img src="docs/images/wechat_sponsor.png" alt="wechat" width="200" />

