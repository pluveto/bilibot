package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

// BotConfig 机器人配置
type BotConfig struct {
	Jct     string `json:"jct"`
	UID     int32  `json:"uid"`
	Session string `json:"sess"`
	SendAPI string `json:"sendAPI"`
}

// MessageService 消息类型
type MessageService struct {
	Config BotConfig
}

// Receiver 接收者
type Receiver struct {
	ID   int
	Type int // 1 私信，2 群聊
}

// SendText 发送私信
func (t MessageService) SendText(to Receiver, content string) error {
	var cookies []*http.Cookie
	cookies = append(cookies, &http.Cookie{Name: "SESSDATA", Value: t.Config.Session})
	jar, _ := cookiejar.New(nil)
	u, _ := url.Parse(t.Config.SendAPI)
	jar.SetCookies(u, cookies)
	client := &http.Client{Jar: jar}
	resp, err := client.PostForm(t.Config.SendAPI, url.Values{
		"msg[receiver_type]": {fmt.Sprintf("%d", to.Type)},
		"msg[receiver_id]":   {fmt.Sprintf("%d", to.ID)},
		"msg[msg_type]":      {fmt.Sprintf("%d", 1)},
		"msg[content]":       {fmt.Sprintf("{\"content\":\"%s\"}", jsonEscape(content))},
		"msg[timestamp]":     {fmt.Sprintf("%d", time.Now().Unix())},
		"msg[sender_uid]":    {fmt.Sprintf("%d", t.Config.UID)},
		"csrf_token":         {t.Config.Jct},
	})

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	fmt.Printf("%s\n", string(body))

	if err != nil {
		return err
	}

	return nil
}

func jsonEscape(i string) string {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	// Trim the beginning and trailing " character
	return string(b[1 : len(b)-1])
}
