package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

//package setting
//
//import (
//	"log"
//	"time"
//
//	"github.com/go-ini/ini"
//)
//
//type App struct {
//	JwtSecret       string
//	PageSize        int
//	RuntimeRootPath string
//
//	ImagePrefixUrl string
//	ImageSavePath  string
//	ImageMaxSize   int
//	ImageAllowExts []string
//
//	LogSavePath string
//	LogSaveName string
//	LogFileExt  string
//	TimeFormat  string
//}
//
//var AppSetting = &App{}
//
//type Server struct {
//	RunMode      string
//	HttpPort     int
//	ReadTimeout  time.Duration
//	WriteTimeout time.Duration
//}
//
//var ServerSetting = &Server{}
//
//type Database struct {
//	Type        string
//	User        string
//	Password    string
//	Host        string
//	Name        string
//	TablePrefix string
//}
//
//var DatabaseSetting = &Database{}
//
//type Redis struct {
//	Host        string
//	Password    string
//	MaxIdle     int
//	MaxActive   int
//	IdleTimeout time.Duration
//}
//
//var RedisSetting = &Redis{}
//
//func Setup() {
//	Cfg, err := ini.Load("conf/app.ini")
//	if err != nil {
//		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
//	}
//
//	err = Cfg.Section("app").MapTo(AppSetting)
//	if err != nil {
//		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
//	}
//
//	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
//
//	err = Cfg.Section("server").MapTo(ServerSetting)
//	if err != nil {
//		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
//	}
//
//	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
//	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
//
//	err = Cfg.Section("database").MapTo(DatabaseSetting)
//	if err != nil {
//		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
//	}
//}