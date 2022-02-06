# 原神祈愿记录

## 功能

- [X] 将爬取结果保存到 `<uid>.json` 文件中（原神官方祈愿记录只保留六个月，获取到新的祈愿记录时，会与本地数据合并后保存）
- [X] 以 HTML 形式粗略展示祈愿记录
- [X] 输入 url 后自动获取祈愿记录（Android、iOS）
- [X] 通过加载本地记录自动获取祈愿记录（PC）
- [ ] 选择是否输出4星祈愿记录

## 食用方式

首先需要拥有 Golang 环境或通过 Release 下载已编译版本。

编译：

```shell
go build .
```

### PC

在本地运行原神，打开到祈愿的历史记录页面即可。

![pc-history](https://user-images.githubusercontent.com/41776735/152675521-e104c9aa-f54b-46b8-8b66-b17fbea5a160.png)

接着直接运行编译后的二进制文件即可，程序会自动打开 output.html。（如果它不自动，自己打开也可以）

```shell
.\gwr.exe
```

### Android

同样的也是打开到这个页面，然后断网，点击右上角的刷新按钮。

![mobile-history](https://user-images.githubusercontent.com/41776735/152675661-89ecf91d-c4e4-4658-9f1d-c3c81522a66a.jpg)

将其中的 url 复制，注意该 url 是以 `https://` 开头，`#log` 结尾。

![mobile-url](https://user-images.githubusercontent.com/41776735/152675686-9e29306e-a0f5-454d-89f4-4aafc9189d2f.jpg)

在运行时使用 -i 参数输入到程序中。（注意双引号）

```shell
.\gwr.exe -i "https://xxxxxxxxxxxxxxxxxx#log"
```

### iOS

> 还没开始写

## 预览

UI 够用就行

![display](https://user-images.githubusercontent.com/41776735/152642642-25a3c387-b44e-46d8-9a4d-6d308e07c374.png)