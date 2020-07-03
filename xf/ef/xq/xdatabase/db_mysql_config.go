package xdatabase

import (
	"fmt"
	"github.com/go-xe2/x/type/t"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TDbMysqlConfig struct {
	*TDbBaseConfig
	// 数据库服务器地址
	dbName string
	// 连接协议,默认tcp
	protocol string
	host     string
	// 数据库端口号
	port int
	// 用户名
	user string
	// 密码
	password string
	// 字符编码类型
	charset string
	// 是否解析日期
	parseTime bool
	// 其他连接参数
	params map[string]string
}

var _ DbMysqlConfig = (*TDbMysqlConfig)(nil)

// 从配置文件创建数据库配置
func NewMysqlConfigFromMap(c map[string]interface{}) *TDbMysqlConfig {
	inst := NewMysqlConfig()
	inst.LoadFromMap(c)
	return inst
}

func NewMysqlConfig() *TDbMysqlConfig {
	return NewMysqlConfig6("127.0.0.1", 3306, "root", "", "", "tcp")
}

// 创建数据库配置
func NewMysqlConfig6(host string, port int, user string, password string, db string, protocol ...string) *TDbMysqlConfig {
	pro := "tcp"
	if len(protocol) > 0 {
		pro = protocol[0]
	}
	inst := &TDbMysqlConfig{
		host:      host,
		port:      port,
		user:      user,
		password:  password,
		protocol:  pro,
		charset:   "utf8",
		parseTime: true,
		params:    make(map[string]string),
		dbName:    db,
	}
	inst.TDbBaseConfig = NewBaseConfig("mysql", "", 0, 0, "", inst)
	return inst
}

// mysql 示例:
// root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=true
func (cfg *TDbMysqlConfig) Dsn() string {
	userPwd := cfg.user
	if cfg.password != "" {
		userPwd += ":" + cfg.password
	}
	s := fmt.Sprintf("%s@%s(%s:%d)/%s?charset=%s&parseTime=%v", userPwd, cfg.protocol, cfg.host, cfg.port, cfg.dbName, cfg.charset, cfg.parseTime)
	szParams := ""
	if cfg.params != nil {
		for k, v := range cfg.params {
			if szParams != "" {
				szParams += "&"
			}
			szParams += k + "=" + v
		}
	}
	if szParams != "" {
		s += "&" + szParams
	}
	return s
}

func (cfg *TDbMysqlConfig) Protocol() string {
	return cfg.protocol
}

// 数据库服务器地址
func (cfg *TDbMysqlConfig) Host() string {
	return cfg.host
}

// 数据库服务器端口号
func (cfg *TDbMysqlConfig) Port() int {
	return cfg.port
}

// 数据库名称
func (cfg *TDbMysqlConfig) DbName() string {
	return cfg.dbName
}

// 账户名
func (cfg *TDbMysqlConfig) User() string {
	return cfg.user
}

// 密码
func (cfg *TDbMysqlConfig) Password() string {
	return cfg.password
}

// 字符编码，默认utf-8
func (cfg *TDbMysqlConfig) Charset() string {
	return cfg.charset
}

// 是否解析日期, 默认true
func (cfg *TDbMysqlConfig) ParseTime() bool {
	return cfg.parseTime
}

// 其他连接参数
func (cfg *TDbMysqlConfig) Params() map[string]string {
	return cfg.params
}

func (cfg *TDbMysqlConfig) SetOption(host string, port int, user string, password string, db string, protocol ...string) DbMysqlConfig {
	cfg.host = host
	cfg.user = user
	cfg.password = password
	cfg.port = port
	cfg.dbName = db
	if len(protocol) > 0 {
		cfg.protocol = protocol[0]
	}
	return cfg
}

func (cfg *TDbMysqlConfig) SetDatabase(db string) DbMysqlConfig {
	cfg.dbName = db
	return cfg
}

func (cfg *TDbMysqlConfig) SetCharset(charset string) DbMysqlConfig {
	cfg.charset = charset
	return cfg
}

func (cfg *TDbMysqlConfig) SetParseTime(b bool) DbMysqlConfig {
	cfg.parseTime = b
	return cfg
}

func (cfg *TDbMysqlConfig) AddParam(name string, value string) DbMysqlConfig {
	cfg.params[name] = value
	return cfg
}

func (cfg *TDbMysqlConfig) LoadFromMap(c map[string]interface{}) {
	// 基类加载配置
	cfg.TDbBaseConfig.LoadFromMap(c)
	cfg.protocol = t.String(c["protocol"], "tcp")
	cfg.host = t.String(c["host"], "127.0.0.1")
	cfg.port = t.Int(c["port"], 3306)
	cfg.user = t.String(c["user"], "root")
	cfg.password = t.String(c["password"], "")
	cfg.dbName = t.String(c["db"], "")
	cfg.charset = t.String(c["charset"], "utf8")
	cfg.parseTime = t.Bool(c["parseTime"], true)
	if tmp := c["params"]; tmp != nil {
		if pms, ok := tmp.(map[string]interface{}); ok {
			for k, v := range pms {
				cfg.params[k] = t.String(v)
			}
		}
	}
}

func (cfg *TDbMysqlConfig) DbConfig() DbConfig {
	return cfg
}
