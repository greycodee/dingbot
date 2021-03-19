# dingbot

钉钉机器人 go库

[钉钉官方文档](https://developers.dingtalk.com/document/app/custom-robot-access)

## 快速开始
```shell script
    go get github.com/greycodee/dingbot
```
示例程序：
```go
package main

import (
	"fmt"
	"github.com/greycodee/dingbot"
	"github.com/greycodee/dingbot/message"
	"time"
)


func main() {
	bot:= dingbot.DingBot{
		Secret:      "你的加签秘钥",
		AccessToken: "你的AccessToken【从钉钉机器人的url上获取】",
	}
	msg := message.Message{
		MsgType: message.TextStr,
		Text:    message.Text_{
			Content: "go-钉钉机器人测试",
		},
	}
	bot.Send(msg)

}
```

## 消息支持
- [x] text类型
- [x] link类型
- [x] markdown类型
- [x] 整体跳转ActionCard类型
- [x] 独立跳转ActionCard类型
- [x] FeedCard类型

## 使用

[查看使用文档](./README-use.md)