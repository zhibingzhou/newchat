package setting

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/go-ini/ini"
)

type commonConf struct {
	HttpPort  string
	RPCPort   string
	Cluster   bool
	CryptoKey string
	Weburl    string
}

type redis struct {
	Host   string
	Port   string
	Pwd    string
	DBName int
}

type rabbitMq struct {
	Admin   string `mapstructure:"admin" json:"admin" yaml:"admin"`
	Pwd     string `mapstructure:"pwd" json:"pwd" yaml:"pwd"`
	Port    int    `mapstructure:"port" json:"port" yaml:"port"`
	Ip      string `mapstructure:"ip" json:"ip" yaml:"ip"`
	Verhost string `mapstructure:"verhost" json:"verhost" yaml:"verhost"`
}

type tool struct {
	ToolName string
}

var CommonSetting = &commonConf{}

var CommonRedis = &redis{}

var CommonRabbitMq = &rabbitMq{}

var CommonTool = &tool{}

type etcdConf struct {
	Endpoints []string
}

var EtcdSetting = &etcdConf{}

type global struct {
	LocalHost      string //本机内网IP
	ServerList     map[string]string
	ServerListLock sync.RWMutex
}

var GlobalSetting = &global{}

var cfg *ini.File

func Setup() {
	configFile := flag.String("c", "conf/app.ini", "-c conf/app.ini")

	var err error
	cfg, err = ini.Load(*configFile)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("common", CommonSetting)
	mapTo("etcd", EtcdSetting)
	mapTo("redis", CommonRedis)
	mapTo("rabbitMq", CommonRabbitMq)
	mapTo("tool", CommonTool)
	fmt.Println(CommonRedis)
	GlobalSetting = &global{
		LocalHost:  getIntranetIp(),
		ServerList: make(map[string]string),
	}
}

func Default() {
	CommonSetting = &commonConf{
		HttpPort:  "6000",
		RPCPort:   "7000",
		Cluster:   false,
		CryptoKey: "Adba723b7fe06819",
	}

	GlobalSetting = &global{
		LocalHost:  getIntranetIp(),
		ServerList: make(map[string]string),
	}
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

//获取本机内网IP
func getIntranetIp() string {
	addrs, _ := net.InterfaceAddrs()

	for _, addr := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}

	return ""
}
