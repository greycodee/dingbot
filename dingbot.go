package main

import (
	"encoding/json"
	"fmt"
	"github.com/greycodee/dingbot/message"
)
func main() {
	var text = message.Text_{
		Content: "22",
		At: message.At_{
			AtMobiles: []string{"1111"},
			IsAtAll:   false,
		} ,
	}

	var msg = message.Message{
		MsgType: message.TextStr,
		Text: text,
	}
	j,e:=json.Marshal(msg)
	if e!=nil {
		fmt.Println(e)
	}
	fmt.Println(string(j))
}
