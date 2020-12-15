package xqconf

import (
	"errors"
	"github.com/go-xe2/x/xf/ef/xq"
	"github.com/go-xe2/x/xf/ef/xq/xdatabase"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type MysqlSingleConfig struct {
	Charset     string `json:"charset"`
	ParseTime   bool   `json:"parse_time"`
	Protocol    string `json:"protocol"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Prefix      string `json:"prefix"`
	DB          string `json:"db"`
	MaxOpenCons int    `json:"max_open_cons"`
	MaxIdleCons int    `json:"max_idle_cons"`
}

type MysqlClusterConfig struct {
	MysqlSingleConfig
	Master []MysqlSingleConfig `json:"master"`
	Slave  []MysqlSingleConfig `json:"slave"`
}

// mysql 数据库配置
type MysqlConfig struct {
	Cluster bool               `json:"cluster"`
	Options MysqlClusterConfig `json:"options"`
}

// InitMySqlDatabase 初始化mysql数据库数据库连接
func InitMySqlDatabase(cfg *MysqlConfig, dbInstance ...string) {
	var conn interface{}
	var err error
	if cfg.Cluster {
		conn, err = initDbWithCluster(&cfg.Options)
		if err != nil {
			panic(err)
		}
	} else {
		conn, err = initDbSingle(&MysqlSingleConfig{
			Charset:     cfg.Options.Charset,
			ParseTime:   cfg.Options.ParseTime,
			Protocol:    cfg.Options.Protocol,
			Host:        cfg.Options.Host,
			Port:        cfg.Options.Port,
			User:        cfg.Options.User,
			Password:    cfg.Options.Password,
			Prefix:      cfg.Options.Prefix,
			DB:          cfg.Options.DB,
			MaxOpenCons: cfg.Options.MaxOpenCons,
			MaxIdleCons: cfg.Options.MaxIdleCons,
		})
		if err != nil {
			panic(err)
		}
	}
	db := xq.Database(dbInstance...)
	if e := db.Connection(conn); e != nil {
		panic(e)
	}
}

// 初始化集群方式的数据库连接
func initDbWithCluster(config *MysqlClusterConfig) (xqi.DbConfigCluster, error) {
	prefix := config.Prefix
	if len(config.Master) == 0 {
		return nil, errors.New("至少应配置一个主数据库连接参数")
	}
	if len(config.Slave) == 0 {
		return nil, errors.New("至少应配置一个从数据库参数")
	}
	connCluster := xdatabase.NewDbConfigCluster()
	masterCount := 0
	for _, v := range config.Master {
		conn := xdatabase.NewMysqlConfig()
		conn.SetCharset(v.Charset)
		conn.SetHost(v.Host)
		conn.SetPort(v.Port)
		conn.SetDatabase(v.DB)
		conn.SetPrefix(v.Prefix)
		conn.SetProtocol(v.Protocol)
		conn.SetMaxIdleCons(v.MaxIdleCons)
		conn.SetMaxOpenCons(v.MaxOpenCons)
		conn.SetUser(v.User)
		conn.SetPassword(v.Password)
		conn.SetParseTime(v.ParseTime)
		connCluster.AddMaster(conn)
		masterCount++
	}
	if masterCount == 0 {
		return nil, errors.New("至少应配置一个主数据库连接参数")
	}
	slaveCount := 0
	for _, v := range config.Slave {
		conn := xdatabase.NewMysqlConfig()
		conn.SetCharset(v.Charset)
		conn.SetHost(v.Host)
		conn.SetPort(v.Port)
		conn.SetDatabase(v.DB)
		conn.SetPrefix(v.Prefix)
		conn.SetProtocol(v.Protocol)
		conn.SetMaxIdleCons(v.MaxIdleCons)
		conn.SetMaxOpenCons(v.MaxOpenCons)
		conn.SetUser(v.User)
		conn.SetPassword(v.Password)
		conn.SetParseTime(v.ParseTime)
		connCluster.AddSlave(conn)
		slaveCount++
	}
	if slaveCount == 0 {
		return nil, errors.New("至少应配置一个从数据库参数")
	}
	connCluster.SetPrefix(prefix)
	connCluster.SetDriver("mysql")
	return connCluster, nil
}

// 初始化单节点数据库连接
func initDbSingle(config *MysqlSingleConfig) (xqi.DbConfig, error) {
	conn := xdatabase.NewMysqlConfig()
	conn.SetCharset(config.Charset)
	conn.SetHost(config.Host)
	conn.SetPort(config.Port)
	conn.SetDatabase(config.DB)
	conn.SetPrefix(config.Prefix)
	conn.SetProtocol(config.Protocol)
	conn.SetMaxIdleCons(config.MaxIdleCons)
	conn.SetMaxOpenCons(config.MaxOpenCons)
	conn.SetUser(config.User)
	conn.SetPassword(config.Password)
	conn.SetParseTime(config.ParseTime)
	return conn, nil
}
