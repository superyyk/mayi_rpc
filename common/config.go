package common

var (
	OneDayOfHours = 60 * 60 * 24
	JWTSecret     = "yyk*2012"
)

var errmsg struct {
	ERROR   string
	SUCCESS string
}

var Base struct {
	BaseUrl string
	IP      string
	Port    int
}

func init() {
	// errmsg.ERROR = "error"
	// errmsg.SUCCESS = "success"
	// cfg, err := ini.Load("yunboss/conf/config.ini")
	// if err != nil {
	// 	log.Fatal("读取配置文件失败: ", err)
	// }
	// Base.BaseUrl = cfg.Section("Base").Key("server_base_url").String()
	// Base.Port, _ = cfg.Section("Base").Key("port").Int()
	Base.BaseUrl = "https://www.kissnet.cn:"
	Base.IP = "127.0.0.1:" //47.106.160.38
	Base.Port = 10030
}
