package xdriver

import . "github.com/go-xe2/x/xf/ef/xdriveri"

type TDriverEngin struct {
	drivers map[DbDriverType]DbDriver
}

func newEngin() *TDriverEngin {
	return &TDriverEngin{
		drivers: make(map[DbDriverType]DbDriver),
	}
}

func (se *TDriverEngin) RegisterBuilder(drType DbDriverType, driver DbDriver) {
	se.drivers[drType] = driver
}

func (se *TDriverEngin) GetDriver(drType DbDriverType) DbDriver {
	if v, ok := se.drivers[drType]; ok {
		return v
	}
	return nil
}

func (se *TDriverEngin) GetBuilder(drType DbDriverType) DbDriverSqlBuilder {
	if b, ok := se.drivers[drType]; ok {
		return b.SqlBuilder()
	}
	return nil
}
