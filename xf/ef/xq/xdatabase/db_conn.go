package xdatabase

import (
	"database/sql"
	"fmt"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/xf/ef/xdbUtil"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

type cluster struct {
	master     []*sql.DB
	masterSize int
	slave      []*sql.DB
	slaveSize  int
}

type TDbConn struct {
	config DbConfigCluster
	driver string
	prefix string
	dbs    *cluster
	logger DbLogger
}

var _ DbConn = (*TDbConn)(nil)

// NewDbConn : 初始化 TDbConn 结构体对象指针
func NewDbConn(conf ...interface{}) (conn *TDbConn, err exception.IException) {
	dbCon := new(TDbConn)
	if len(conf) == 0 {
		return
	}

	// 使用默认的log, 如果自定义了logger, 则只需要调用 Use() 方法即可覆盖
	dbCon.Use(DefaultLogger())

	switch v := conf[0].(type) {
	// 传入的是单个配置
	case DbConfig:
		err = dbCon.bootSingle(v)
	// 传入的是集群配置
	case DbConfigCluster:
		dbCon.config = v
		err = dbCon.bootCluster()
	default:
		panic(fmt.Sprint("DbConn创建参数只支持DbConfig及DbConfigCluster接口的配置,当前传入:", conf))
	}
	return dbCon, err
}

// 使用插件
func (c *TDbConn) Use(closers ...func(e DbConn)) {
	for _, closer := range closers {
		closer(c)
	}
}

// Ping
func (c *TDbConn) Ping() exception.IException {
	if err := c.GetQueryDB().Ping(); err != nil {
		return exception.Wrap(err, "连接数据库失败")
	}
	return nil
}

// SetPrefix 设置表前缀
func (c *TDbConn) SetPrefix(pre string) {
	c.prefix = pre
}

// GetPrefix 获取前缀
func (c *TDbConn) GetPrefix() string {
	return c.prefix
}

func (c *TDbConn) GetDriver() string {
	return c.driver
}

func (c *TDbConn) SetLogger(logger DbLogger) {
	c.logger = logger
}

// GetQueryDB : get a slave db for using query operation
// GetQueryDB : 获取一个从库用来做查询操作
func (c *TDbConn) GetQueryDB() *sql.DB {
	if c.dbs.slaveSize == 0 {
		return c.GetExecuteDB()
	}
	var randInt = xdbUtil.MakeRandomInt(c.dbs.slaveSize)
	return c.dbs.slave[randInt]
}

// GetExecuteDB : 获取一个主库用来做查询之外的操作
func (c *TDbConn) GetExecuteDB() *sql.DB {
	if c.dbs.masterSize == 0 {
		return nil
	}
	var randInt = xdbUtil.MakeRandomInt(c.dbs.masterSize)
	return c.dbs.master[randInt]
}

func (c *TDbConn) GetLogger() DbLogger {
	return c.logger
}

func (c *TDbConn) This() interface{} {
	return c
}

func (c *TDbConn) bootSingle(conf DbConfig) exception.IException {
	// 如果传入的是单一配置, 则转换成集群配置, 方便统一管理
	var cc = NewDbConfigCluster()
	cc.AddMaster(conf)
	c.config = cc
	return c.bootCluster()
}

func (c *TDbConn) bootCluster() exception.IException {
	//fmt.Println(len(c.config.Slave))
	slaves := c.config.Slave()
	masters := c.config.Master()
	if len(slaves) > 0 {
		for _, item := range slaves {
			db, err := c.bootReal(item)
			if err != nil {
				return err
			}
			if c.dbs == nil {
				c.dbs = new(cluster)
			}
			c.dbs.slave = append(c.dbs.slave, db)
			c.dbs.slaveSize++
			c.driver = item.Driver()
		}
	}
	var pre, dr string
	if len(masters) > 0 {
		for _, item := range masters {
			db, err := c.bootReal(item)

			if err != nil {
				return err
			}
			if c.dbs == nil {
				c.dbs = new(cluster)
			}
			c.dbs.master = append(c.dbs.master, db)
			c.dbs.masterSize = c.dbs.masterSize + 1
			c.driver = item.Driver()
			//fmt.Println(c.dbs.masterSize)
			if item.Prefix() != "" {
				pre = item.Prefix()
			}
			if item.Driver() != "" {
				dr = item.Driver()
			}
		}
	}
	// 如果config没有设置prefix,且configcluster设置了prefix,则使用cluster的prefix
	if pre != "" && c.prefix == "" {
		c.prefix = pre
	}
	// 如果config没有设置driver,且configcluster设置了driver,则使用cluster的driver
	if dr != "" && c.driver == "" {
		c.driver = dr
	}
	return nil
}

// 生成数据库连接
func (c *TDbConn) bootReal(dbConf DbConfig) (db *sql.DB, err exception.IException) {
	//db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/test?charset=utf8")
	// 开始驱动
	var e error
	db, e = sql.Open(dbConf.Driver(), dbConf.Dsn())
	if e != nil {
		return nil, exception.Wrap(e, "打开数据失败")
	}

	// 检查是否可以ping通
	e = db.Ping()
	if e != nil {
		return nil, exception.Wrap(e, "连接数据失败")
	}
	// 连接池设置
	db.SetMaxOpenConns(dbConf.MaxOpenCons())
	db.SetMaxIdleConns(dbConf.MaxIdleCons())
	return
}
