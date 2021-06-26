package utils

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	JwtSecret string
)

type App struct {
<<<<<<< HEAD
	PageSize          int
	JwtSecret         string
	JwtExpirationTime int
=======
	PageSize int
>>>>>>> parent of fee293e (登录生成token,后续添加验证)
}

var AppSetting = &App{}

func init() {
	var err error
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8080)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	AppSetting.PageSize = sec.Key("PAGE_SIZE").MustInt(10)
<<<<<<< HEAD
	AppSetting.JwtSecret = sec.Key("JWT_SECRET").MustString("23347$040412")
	AppSetting.JwtExpirationTime = sec.Key("JWT_EXPIRATION_TIME").MustInt(7)
=======
>>>>>>> parent of fee293e (登录生成token,后续添加验证)
}
