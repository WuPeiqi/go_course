package main

import (
	"fmt"
)

type IMessage interface {
	send() bool
}

type Email struct {
	email   string
	content string
}

func (p *Email) send() bool {
	fmt.Println("发送邮件提醒", p.email, p.content)
	return true
}

type Wechat struct {
	wid     string
	content string
}

func (p *Wechat) send() bool {
	fmt.Println("发送微信提醒", p.wid, p.content)
	return true
}

func something(objectList []IMessage) {
	for _, item := range objectList {
		result := item.send()
		fmt.Println(result)
	}
}

func main() {
	// 用户注册...
	messageObjectList := []IMessage{
		&Email{"wupeiqi@live.com", "注册成功"},
		&Wechat{"wupeiqi@live.com", "注册成功"},
	}
	something(messageObjectList)
}