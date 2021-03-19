## 发送Text消息

```go
func send() {
	bot:= dingbot.DingBot{
		Secret:      "你的加签秘钥",
		AccessToken: "你的AccessToken【从钉钉机器人的url上获取】",
	}
	msg := message.Message{
        MsgType: message.TextStr,
        Text:    message.Text_{
            Content: "go-钉钉机器人测试",
            At: message.At_{
                AtMobiles: []string{"188xxxxxxxx"},
                IsAtAll:   false,
            },
        },
    }
	bot.Send(msg)

}
```

## 发送link类型消息

```go
func send() {
	bot:= dingbot.DingBot{
		Secret:      "你的加签秘钥",
		AccessToken: "你的AccessToken【从钉钉机器人的url上获取】",
	}
	msg := message.Message{
        MsgType: message.LinkStr,
        Link: message.Link_{
            Text:       "link测试123123",
            Title:      "go钉钉机器人",
            PicUrl:     "",
            MessageUrl: "https://developers.dingtalk.com/document/app/custom-robot-access/title-72m-8ag-pqw",
        },
    		
    }
	bot.Send(msg)

}
```

## 发送markdown类型消息

```go
func send() {
	bot:= dingbot.DingBot{
		Secret:      "你的加签秘钥",
		AccessToken: "你的AccessToken【从钉钉机器人的url上获取】",
	}
	msg := message.Message{
        MsgType: message.MarkdownStr,
        Markdown: message.Markdown_{
            Title: "go钉钉",
            Text:  "## go钉钉测试 @188xxxxxxxx \n>文本123",
            At:    message.At_{
                AtMobiles: []string{"188xxxxxxxx"},
                IsAtAll:   false,
            },
        },
    		
    }
	bot.Send(msg)

}
```

## 发送整体跳转ActionCard类型消息
   
```go
func send() {
	bot:= dingbot.DingBot{
		Secret:      "你的加签秘钥",
		AccessToken: "你的AccessToken【从钉钉机器人的url上获取】",
	}
	msg := message.Message{
        MsgType: message.ActionCardStr,
        ActionCard: message.ActionCard_{
            Title:          "ActionCard整体跳转11",
            Text:           "ActionCardt整体跳转1223",
            SingleTitle:    "阅读全文",
            SingleURL:      "https://developers.dingtalk.com/document/app/custom-robot-access/title-72m-8ag-pqw",
        },
    }
	bot.Send(msg)

}
```  

## 发送独立跳转ActionCard类型消息

```go
func send() {
    bot:= dingbot.DingBot{
        Secret:      "你的加签秘钥",
        AccessToken: "你的AccessToken【从钉钉机器人的url上获取】",
    }
    msg := message.Message{
        MsgType: message.ActionCardStr,
        ActionCard: message.ActionCard_{
            Title:          "ActionCard跳转11",
            Text:           "ActionCardt跳转1223",
            BtnOrientation: "1",
            HideAvatar:     "0",
            BtnS:           []message.Btn_{
                {
                    Title:     "按钮1",
                    ActionURL: "https://developers.dingtalk.com/",
                },
                {
                    Title:     "按钮2",
                    ActionURL: "https://developers.dingtalk.com/",
                },
            },
        },
    }
    bot.Send(msg)
}
```

## 发送FeedCard类型消息

```go
func send() {
    bot:= dingbot.DingBot{
        Secret:      "你的加签秘钥",
        AccessToken: "你的AccessToken【从钉钉机器人的url上获取】",
    }
    msg := message.Message{
        MsgType:  message.FeedCardStr,
        FeedCard: message.FeedCard_{[]message.Link_{
            {
                Title:      "标题1",
                PicUrl:     "",
                MessageUrl: "https://developers.dingtalk.com/",
            },
            {
                Title:      "标题2",
                PicUrl:     "",
                MessageUrl: "https://developers.dingtalk.com/",
            },
        }},
    }
    bot.Send(msg)
}
```