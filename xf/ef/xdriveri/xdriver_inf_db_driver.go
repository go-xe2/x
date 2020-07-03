package xdriveri

type DbDriver interface {
	This() interface{}
	DriverType() DbDriverType
	SqlBuilder() DbDriverSqlBuilder
}
