package wechat

import (
	"github.com/astaxie/beego"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/sirupsen/logrus"
)

var redisCache = cache.NewRedis(&cache.RedisOpts{
	Host: beego.AppConfig.String("redis::rHost")+":"+beego.AppConfig.String("redis::rPort"),
})
var token = beego.AppConfig.String("wechat::Token")

var config = &wechat.Config{
	AppID:          beego.AppConfig.String("wechat::AppID"),
	AppSecret:      beego.AppConfig.String("wechat::AppSecret"),
	Token:          token,
	EncodingAESKey: beego.AppConfig.String("wechat::EncodingAESKey"),
	Cache:			redisCache,
}

var clientId = beego.AppConfig.String("chatopera::clientId")
var clientSecret = beego.AppConfig.String("chatopera::clientSecret")

var limit = 0.6

var log = logrus.New()
