package xdatabase

import (
	"fmt"
	"github.com/go-xe2/x/type/t"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type TDbBaseConfig struct {
	driver string // 驱动: mysql/sqlite3/oracle/mssql/postgres/clickhouse, 如果集群配置了驱动, 这里可以省略
	// mysql 示例:
	// root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=true
	dsn         string // 数据库链接
	maxOpenCons int    // (连接池)最大打开的连接数，默认值为0表示不限制
	maxIdleCons int    // (连接池)闲置的连接数, 默认0
	prefix      string // 表前缀, 如果集群配置了前缀, 这里可以省略
	this        interface{}
}

var _ DbConfig = (*TDbBaseConfig)(nil)

func NewBaseConfig(driver, dsn string, maxOpenCons, maxIdleCons int, prefix string, inherited ...interface{}) *TDbBaseConfig {
	inst := &TDbBaseConfig{
		driver:      driver,
		dsn:         dsn,
		maxOpenCons: maxOpenCons,
		maxIdleCons: maxIdleCons,
		prefix:      prefix,
	}
	inst.this = inst
	if len(inherited) > 0 && inherited[0] != nil {
		if _, ok := inherited[0].(DbConfig); ok {
			inst.this = inherited[0]
		}
	}
	return inst
}

func (cfg *TDbBaseConfig) LoadFromMap(c map[string]interface{}) {
	cfg.driver = t.String(c["driver"], "mysql")
	cfg.maxOpenCons = t.Int(c["maxOpenCons"], 0)
	cfg.maxIdleCons = t.Int(c["maxIdleCons"], 0)
	cfg.prefix = t.String(c["prefix"], "")
}

func (cfg *TDbBaseConfig) String() string {
	return fmt.Sprintf("dsn:%s,driver:%s,maxOpenConns:%d,maxIdleConns:%d,prefix:%s", cfg.dsn, cfg.driver, cfg.maxOpenCons, cfg.maxIdleCons, cfg.prefix)
}

func (cfg *TDbBaseConfig) Driver() string {
	return cfg.driver
}

func (cfg *TDbBaseConfig) Dsn() string {
	return cfg.dsn
}

func (cfg *TDbBaseConfig) MaxOpenCons() int {
	return cfg.maxOpenCons
}

func (cfg *TDbBaseConfig) MaxIdleCons() int {
	return cfg.maxIdleCons
}

func (cfg *TDbBaseConfig) Prefix() string {
	return cfg.prefix
}

func (cfg *TDbBaseConfig) This() interface{} {
	return cfg.this
}
