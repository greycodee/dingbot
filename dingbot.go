package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/greycodee/dingbot/message"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)


type DingBot struct{
	Secret 		string
	AccessToken string
}

type Response struct {
	Errcode int64 `json:"errcode"`
	Errmsg string `json:"errmsg"`
}

func (bot *DingBot) send(msg message.Message) (Response,error) {
	timestamp:=time.Now().UnixNano()/1e6
	var response Response
	signStr := sign(timestamp,bot.Secret)
	dingUrl := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%d&sign=%s",bot.AccessToken,timestamp,signStr)
	j,e:=json.Marshal(msg)
	if e!=nil {
		return response,e
	}
	resp,e:=http.Post(dingUrl,"application/json",strings.NewReader(string(j)))
	if e!=nil {
		return response,e
	}
	defer resp.Body.Close()
	
	body, _ := ioutil.ReadAll(resp.Body)
	e=json.Unmarshal(body,&response)
	if e!=nil {
		return response,e
	}
	return response,nil
}

func sign(t int64, secret string) string  {
	secStr := fmt.Sprintf("%d\n%s",t,secret)
	hmac256 := hmac.New(sha256.New,[]byte(secret))
	hmac256.Write([]byte(secStr))
	result := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(result)
}