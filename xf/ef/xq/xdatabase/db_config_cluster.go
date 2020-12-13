package xdatabase

import (
	"fmt"
	. "github.com/go-xe2/x/xf/ef/xqi"
	"strings"
)

type TDbConfigCluster struct {
	master []DbConfig
	slave  []DbConfig
	driver string
	prefix string
}

var _ DbConfigCluster = &TDbConfigCluster{}

func NewDbConfigCluster() *TDbConfigCluster {
	return &TDbConfigCluster{
		master: make([]DbConfig, 0),
		slave:  make([]DbConfig, 0),
		driver: "",
		prefix: "",
	}
}

func (dcc *TDbConfigCluster) AddMaster(config DbConfig) DbConfigCluster {
	dcc.master = append(dcc.master, config)
	return dcc
}

func (dcc *TDbConfigCluster) AddSlave(config DbConfig) DbConfigCluster {
	dcc.slave = append(dcc.slave, config)
	return dcc
}

func (dcc *TDbConfigCluster) Master() []DbConfig {
	return dcc.master
}

func (dcc *TDbConfigCluster) Slave() []DbConfig {
	return dcc.slave
}

func (dcc *TDbConfigCluster) Driver() string {
	return dcc.driver
}

func (dcc *TDbConfigCluster) SetDriver(driver string) {
	dcc.driver = driver
}

func (dcc *TDbConfigCluster) Prefix() string {
	return dcc.prefix
}

func (dcc *TDbConfigCluster) SetPrefix(prefix string) {
	dcc.prefix = prefix
}

func (dcc *TDbConfigCluster) This() interface{} {
	return dcc
}

func (dcc *TDbConfigCluster) String() string {
	s1 := make([]string, len(dcc.master))
	for i := 0; i < len(dcc.master); i++ {
		s1[i] = dcc.master[i].String()
	}
	s2 := make([]string, len(dcc.slave))
	for i := 0; i < len(dcc.slave); i++ {
		s2[i] = dcc.slave[i].String()
	}
	return fmt.Sprintf("prefix:%s,driver:%s, master:[%s], slave:[%s]", dcc.prefix, dcc.driver, strings.Join(s1, ","), strings.Join(s2, ","))
}
