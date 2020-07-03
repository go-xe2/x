package xdriver

import . "github.com/go-xe2/x/xf/ef/xdriveri"

var Engine = newEngin()

func RegisterDriver(drType DbDriverType, driver DbDriver) {
	Engine.RegisterBuilder(drType, driver)
}

func GetDriver(drType DbDriverType) DbDriver {
	return GetDriver(drType)
}

func GetSqlBuilder(drType DbDriverType) DbDriverSqlBuilder {
	return Engine.GetBuilder(drType)
}

func GetDriverByName(driverName string) DbDriver {
	drType := DbDriverTypeByName(driverName)
	return GetDriver(drType)
}

func GetSqlBuilderByName(driverName string) DbDriverSqlBuilder {
	drType := DbDriverTypeByName(driverName)
	return Engine.GetBuilder(drType)
}

func init() {
	RegisterDriver(MySqlDriver, DbDefaultDriver)
}
