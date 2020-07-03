package xqi

// 数据库连接配置

type DbConfig interface {
	Driver() string // 驱动: mysql/sqlite3/oracle/mssql/postgres/clickhouse, 如果集群配置了驱动, 这里可以省略
	// mysql 示例:
	// root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=true
	Dsn() string      // 数据库链接
	MaxOpenCons() int // (连接池)最大打开的连接数，默认值为0表示不限制
	MaxIdleCons() int // (连接池)闲置的连接数, 默认0
	Prefix() string   // 表前缀, 如果集群配置了前缀, 这里可以省略
	This() interface{}
	String() string
	LoadFromMap(cfg map[string]interface{})
}

type DbConfigCluster interface {
	Master() []DbConfig // 主
	Slave() []DbConfig  // 从
	Driver() string     // 驱动
	Prefix() string     // 前缀
	This() interface{}
	String() string
	// 添加主数据库连接
	AddMaster(config DbConfig) DbConfigCluster
	// 添加从数据库连接
	AddSlave(config DbConfig) DbConfigCluster
}
