<p>
<img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=86400" />
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-green.svg" />
  </a>
  <a href="https://twitter.com/xiaomoinfo" target="_blank">
    <img alt="Twitter: xiaomoinfo" src="https://img.shields.io/twitter/follow/xiaomoinfo.svg?style=social" />
  </a>
</p>

# 准备运行环境
```
go mod tidy 
cp config/config.yaml.example local/config.yaml
```

## 修改你的token
打开 [openai](https://beta.openai.com/account/api-keys) 并注册一个账号, 生成一个token并把token放到`local/config.yaml`的token下

```
chatgpt:
  keyword: 小莫
  token: sk-pKHZD1fLYqXDjjsdsdsdUvIODTT3ssjdfadsJC2gTuqqhTum
```

## 运行App
```
go run main.go
```

```
ain.go #gosetup
go: downloading github.com/eatmoreapple/openwechat v1.2.1
go: downloading github.com/sirupsen/logrus v1.6.0
go: downloading github.com/spf13/afero v1.9.2
go: downloading github.com/pelletier/go-toml/v2 v2.0.5
go: downloading golang.org/x/sys v0.0.0-20220908164124-27713097b956
/private/var/folders/8t/0nvj_2kn4dl517vhbc4rmb9h0000gn/T/GoLand/___go_build_main_go
访问下面网址扫描二维码登录
https://login.weixin.qq.com/qrcode/QedkOe1I4w==
```

会自动打开默认浏览器，如果没有打开也可以打动点击上面的链接打开二维码扫微信

```
2022/12/09 15:15:00 登录成功
2022/12/09 15:15:01 RetCode:0  Selector:2
2022/12/09 15:15:04 RetCode:0  Selector:2
INFO[0099] 0 <Friend:hxh,晓华>                            
INFO[0099] 1 <Friend:刘葵>                                
INFO[0099] 2 <Friend:吕>                                 
INFO[0099] 3 <Friend:wloscar>               
```
登陆成功后会拉取微信的好友和群组

## 如何使用
默认为`chatgpt`，如果想设置其他的触发方式可以修改`local/config.yaml`的keyword。此时，如果别人给你发消息带有关键字`chatgpt`，你的微信就会调用`chatGPT`AI自动回复你的好友。
当然，在群里也是可以的。

## 使用场景1
别人给你发消息时，如果消息中带有关键字，系统就会调用AI自动帮你回复此问题。    

<img src="screenshots/IMG_3837.png" alt="drawing" style="width:200px;display: inline"/>
<img src="screenshots/IMG_3840.png" alt="drawing" style="width:200px;display: inline"/>
<img src="screenshots/IMG_3850.png" alt="drawing" style="width:200px;display: inline"/>


## 使用场景2
别人在群里发消息时，如果消息中带有关键字，系统就会调用AI自动帮你回复此问题。   

<img src="screenshots/IMG_3845.png" alt="drawing" style="width:200px;display: inline"/>
<img src="screenshots/IMG_3847.png" alt="drawing" style="width:200px;display: inline"/>


## 使用场景3
自己给自己发消息时，如果消息中带有关键字，系统会也调用AI自动帮你回复此问题。   

<img src="screenshots/IMG_3843.png" alt="drawing" style="width:200px;display: inline"/>


## 总结
- 你可以把它当作你的智能助理，帮助你快速回复消息。   
- 你可以把它当作一个智能机器人，邀请在群里之后通过关键字帮助大家解答问题。   
- 你可以把它当作你的智多星，有什么问题不懂的时候随时问它。   


## 意外之喜
<img src="screenshots/IMG_3844.png" alt="drawing" style="width:200px;"/>   

这不比对象来的贴心？    

## 变爸爸事件
[用chatgpt写了个微信机器人结果变爸爸了](https://www.bilibili.com/video/BV1B24y1Q7us/)

##
如果大家有玩的时候有遇到一些奇怪的对话可以截图发PR分享给大家。另外对本项目有什么想法或者贡献的话欢迎提[issue](https://github.com/houko/wechatgpt/issues)或[pr](https://github.com/houko/wechatgpt/pulls)