package defaultDriver

import (
	. "github.com/go-xe2/x/xf/ef/xdriveri"
)

type TDbDefaultDriver struct {
}

var _ DbDriver = &TDbDefaultDriver{}

func (dr *TDbDefaultDriver) This() interface{} {
	return dr
}

func (dr *TDbDefaultDriver) DriverType() DbDriverType {
	return MySqlDriver
}

func (dr *TDbDefaultDriver) SqlBuilder() DbDriverSqlBuilder {
	return dr
}
