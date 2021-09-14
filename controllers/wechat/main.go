package wechat

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/chatopera/chatopera-go-sdk"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
	"sort"
	"sync"
)

type MainController struct {
	BaseController
}

var chat *chatopera.Chatopera
var once sync.Once

/**
 * 验证 服务器
 */
func (c *MainController) CheckToken() {

	once.Do(func() {
		chat = chatopera.Chatbot(clientId,clientSecret)
	})

	signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	token := token
	echostr := c.GetString("echostr")

	var tempArray  = []string{token, timestamp, nonce}
	sort.Strings(tempArray)
	//将三个参数字符串拼接成一个字符串进行sha1加密
	var sha1String string = ""
	for _, v := range tempArray {
		sha1String += v
	}
	h := sha1.New()
	h.Write([]byte(sha1String))
	sha1String = hex.EncodeToString(h.Sum([]byte("")))
	//获得加密后的字符串可与signature对比
	if sha1String == signature {
		_, err := c.Ctx.ResponseWriter.Write([]byte(echostr))
		if err != nil {
			c.Data["json"] = err.Error()
			c.ServeJSON()
			c.StopRun()
		}
	} else {
		c.Data["json"] = "验证失败"
		c.ServeJSON()
		c.StopRun()
	}
}

func (c *MainController) Say() {
	once.Do(func() {
		chat = chatopera.Chatbot(clientId,clientSecret)
	})

	reply, err := chat.Faq("xiao", "傻瓜")

	response := make(map[string]interface{})
	if err != nil {
		log.Error(err)

		response["success"] = false
		response["err"] = err.Error()

	} else {
		log.Info("TestFaq reply:",reply)
		response["data"] = reply.Data
		response["success"] = true
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}



func (c *MainController) SayHello() {
	once.Do(func() {
		chat = chatopera.Chatbot(clientId,clientSecret)
	})
	uid := c.GetString("userId")
	msg := c.GetString("message")

	reply, err := chat.Faq(uid, msg)

	response := make(map[string]interface{})
	if err != nil {
		log.Error(err)

		response["success"] = false
		response["err"] = err.Error()

	} else {
		log.Info("TestFaq reply:",reply)
		response["data"] = reply.Data
		response["success"] = true
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}

func (c *MainController) Hello() {

	wc := wechat.NewWechat(config)

	log.Infoln(config)

	// 传入request和responseWriter
	server := wc.GetServer(c.Ctx.Request, c.Ctx.ResponseWriter)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		log.Infoln(msg)
		return ResponseMsgType(msg)
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}