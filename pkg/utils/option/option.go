package option

import (
	"flag"
	"github.com/vharitonsky/iniflags"
)

type Option struct {
	Address string // 服务地址
	Port int // 服务端口
	Db *DB // 数据库链接信息
	CSRFName string //csrf名称
	CSRFToken string //csrf值
	SessionName string //session的名称
	LogFile string //log file
}

type DB struct {
	Host string
	Port int
	Database string
	User string
	Passwd string
}


var GlobalOption = &Option{Db: &DB{}}

func (o *Option) InitOptions() {
	flag.StringVar(&o.Address, "a", "0.0.0.0", "bind address")
	flag.IntVar(&o.Port, "p", 8085, "bind port")
	flag.StringVar(&o.Db.Host, "dbhost", "127.0.0.1", "database host")
	flag.IntVar(&o.Db.Port, "dbport", 3306, "database port")
	flag.StringVar(&o.Db.User, "dbuser", "root", "database user")
	flag.StringVar(&o.Db.Passwd, "dbpasswd", "", "database password")
	flag.StringVar(&o.Db.Database, "dbname", "gomanager", "database name")
	flag.StringVar(&o.CSRFName, "csrf_name", "gomanagercsrf", "csrf name")
	flag.StringVar(&o.CSRFToken, "csrf_value", "*****", "csrf value")
	flag.StringVar(&o.SessionName, "session_name", "***", "session name")
	flag.StringVar(&o.LogFile, "log_file", "", "log file path")
	iniflags.Parse()
}