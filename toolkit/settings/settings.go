package settings

import (
	"log"
	"time"
	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
	JwtTokenValid time.Duration

	PageSize int

	ImageSavePath    string
	ImageMaxSize     int
	ImageAllowExists []string
}

var AppSettings = &App{}

type Server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

var ServerSettings = Server{}

type Mysql struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var MysqlSettings = Mysql{}

type Redis struct {
	Host        string
	Password    string
	MinIdleConn int
	PoolSize    int
	MaxActive   time.Duration
	IdleTimeout time.Duration
	DB          int
}

var RedisSettings = Redis{}

type Log struct {
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var LogSettings = Log{}

var cfg *ini.File
func Setup()  {
	var err error
	cfg, err= ini.Load("conf/app.ini")
	if err != nil{
		log.Fatalf("setting.Setup, fail to parse 'conf/api.ini': %v", err)
	}
	mapTo("app", AppSettings)
	mapTo("server", ServerSettings)
	mapTo("mysql", MysqlSettings)
	mapTo("redis", RedisSettings)
	mapTo("log", LogSettings)

	AppSettings.JwtTokenValid = AppSettings.JwtTokenValid * time.Hour // jwt valid 1 h
	AppSettings.ImageMaxSize = AppSettings.ImageMaxSize *  1024 * 1024 // MB
	ServerSettings.ReadTimeout = ServerSettings.ReadTimeout * time.Second
	ServerSettings.WriteTimeout = ServerSettings.WriteTimeout * time.Second
	RedisSettings.IdleTimeout = RedisSettings.IdleTimeout * time.Second
	RedisSettings.MaxActive = RedisSettings.MaxActive * time.Second
}

func mapTo(section string, v interface{})  {
	err := cfg.Section(section).MapTo(v)
	if err != nil{
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}