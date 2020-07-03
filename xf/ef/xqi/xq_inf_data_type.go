package xqi

import (
	"github.com/go-xe2/x/xf/ef/xdriveri"
)

type DbType interface {
	GetType() xdriveri.DbDataType
	GetSize() int
	GetDecimal() []int
}

type DbTypeS interface {
	DbType
	Size(size int) DbType
}

type DbTypeSD interface {
	DbType
	SizeDecimal(size int, decimal ...int) DbType
}
