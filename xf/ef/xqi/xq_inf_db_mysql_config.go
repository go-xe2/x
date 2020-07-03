package xqi

type DbMysqlConfig interface {
	DbConfig
	// 连接协议tcp,udp
	Protocol() string
	// 数据库服务器地址
	Host() string
	// 数据库服务器端口号
	Port() int
	// 数据库名称
	DbName() string
	// 账户名
	User() string
	// 密码
	Password() string
	// 字符编码，默认utf-8
	Charset() string
	// 是否解析日期, 默认true
	ParseTime() bool
	// 其他连接参数
	Params() map[string]string

	SetOption(host string, port int, user string, password string, db string, protocol ...string) DbMysqlConfig
	SetDatabase(db string) DbMysqlConfig
	SetCharset(charset string) DbMysqlConfig
	SetParseTime(b bool) DbMysqlConfig
	AddParam(name string, value string) DbMysqlConfig

	DbConfig() DbConfig
}
