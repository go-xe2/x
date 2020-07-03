package xqcomm

import (
	. "github.com/go-xe2/x/xf/ef/xdriveri"
	. "github.com/go-xe2/x/xf/ef/xqi"
)

var (
	DTInt      = DatabaseDataType(DbDataInt)
	DTTinyint  = DatabaseDataTypeS(DbDataTinyint, 0)
	DTSmallint = DatabaseDataTypeS(DbDataSmallint, 0)
	DTBigint   = DatabaseDataType(DbDataBigint)

	DTFloat   = DatabaseDataType(DbDataFloat)
	DTDouble  = DatabaseDataType(DbDataDouble)
	DTDecimal = DatabaseDataTypeSD(DbDataDecimal, 0, 0)

	DTDate      = DatabaseDataType(DbDataDate)
	DTTime      = DatabaseDataType(DbDataTime)
	DTDatetime  = DatabaseDataType(DbDataDateTime)
	DTTimestamp = DatabaseDataType(DbDataTimestamp)

	DTChar     = DatabaseDataTypeS(DbDataChar, 0)
	DTVarchar  = DatabaseDataTypeS(DbDataVarchar, 0)
	DTTinytext = DatabaseDataType(DbDataTinytext)
	DTText     = DatabaseDataType(DbDataText)
	DTLongtext = DatabaseDataType(DbDataLongText)

	DTBlob     = DatabaseDataType(DbDataBlob)
	DTTinyblob = DatabaseDataType(DbDataTinyblob)
	DTLongblob = DatabaseDataType(DbDataLongblob)

	DTBinary    = DatabaseDataTypeS(DbDataBinary, 0)
	DTVarbinary = DatabaseDataTypeS(DbDataVarbinary, 0)
)

type tDataType struct {
	dbType  DbDataType
	size    int
	decimal []int
}

var _ DbType = &tDataType{}

func DatabaseDataType(typ DbDataType) DbType {
	return &tDataType{
		dbType:  typ,
		size:    0,
		decimal: nil,
	}
}

func DatabaseDataTypeS(typ DbDataType, size int) DbTypeS {
	return &tDataType{
		dbType:  typ,
		size:    size,
		decimal: nil,
	}
}

func DatabaseDataTypeSD(typ DbDataType, size int, decimal ...int) DbTypeSD {
	return &tDataType{
		dbType:  typ,
		size:    size,
		decimal: decimal,
	}
}

func (dt *tDataType) GetType() DbDataType {
	return dt.dbType
}

func (dt *tDataType) GetSize() int {
	return dt.size
}

func (dt *tDataType) GetDecimal() []int {
	return dt.decimal
}

func (dt *tDataType) Size(size int) DbType {
	return &tDataType{
		dbType:  dt.dbType,
		size:    size,
		decimal: dt.decimal,
	}
}

func (dt *tDataType) SizeDecimal(size int, decimal ...int) DbType {
	return &tDataType{
		dbType:  dt.dbType,
		size:    size,
		decimal: decimal,
	}
}
