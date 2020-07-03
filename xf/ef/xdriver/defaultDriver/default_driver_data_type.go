package defaultDriver

import (
	"fmt"
	. "github.com/go-xe2/x/xf/ef/xdriveri"
)

type DDTType = DbDataType

var DbDataTypeSizes = map[DDTType][]int{
	DbDataInt:      {0, 4},
	DbDataTinyint:  {0, 1},
	DbDataSmallint: {6, 6},
	DbDataBigint:   {0, 8},

	DbDataFloat:   {0, 4},
	DbDataDouble:  {0, 8},
	DbDataDecimal: {0, 18},

	DbDataDate:      {0, 3},
	DbDataTime:      {0, 3},
	DbDataDateTime:  {0, 8},
	DbDataTimestamp: {0, 4},

	DbDataChar:     {128, 255},
	DbDataVarchar:  {128, 65535},
	DbDataTinytext: {0, 255},
	DbDataText:     {0, 65535},
	DbDataLongText: {0, 4294967295},

	DbDataBlob:     {0, 65535},
	DbDataTinyblob: {0, 255},
	DbDataLongblob: {0, 4294967295},

	DbDataBinary:    {1, 65535},
	DbDataVarbinary: {128, 65535},
}

var DbDataTypeDefineFunMaps = map[DDTType]func(size int, decimal ...int) string{
	DbDataInt: func(size int, decimal ...int) string {
		return "int"
	},
	DbDataTinyint: func(size int, decimal ...int) string {
		if size == 0 {
			return fmt.Sprintf("tinyint")
		}
		return fmt.Sprintf("tinyint(%d)", size)
	},
	DbDataSmallint: func(size int, decimal ...int) string {
		if size == 0 {
			size = DbDataTypeSizes[DbDataSmallint][0]
		}
		return fmt.Sprintf("smallint(%d)", size)
	},
	DbDataBigint: func(size int, decimal ...int) string {
		return "bigint"
	},
	DbDataFloat: func(size int, decimal ...int) string {
		return "float"
	},
	DbDataDouble: func(size int, decimal ...int) string {
		return "double"
	},
	DbDataDecimal: func(size int, decimal ...int) string {
		if size == 0 && len(decimal) > 0 {
			size = decimal[0] + 2
			return fmt.Sprintf("decimal(%d,%d)", size, decimal[0])
		} else if size == 0 && len(decimal) == 0 {
			return "decimal"
		} else if size > 0 && len(decimal) > 0 {
			if size < decimal[0] {
				size = decimal[0] + 2
			}
			return fmt.Sprintf("decimal(%d,%d)", size, decimal[0])
		} else if size > 0 && len(decimal) == 0 {
			return fmt.Sprintf("decimal(%d, 0)", size)
		} else {
			return "decimal"
		}
	},
	DbDataDate: func(size int, decimal ...int) string {
		return "date"
	},
	DbDataTime: func(size int, decimal ...int) string {
		return "time"
	},
	DbDataDateTime: func(size int, decimal ...int) string {
		return "datetime"
	},
	DbDataTimestamp: func(size int, decimal ...int) string {
		return "timestamp"
	},
	DbDataChar: func(size int, decimal ...int) string {
		if size == 0 {
			size = DbDataTypeSizes[DbDataChar][0]
		}
		return fmt.Sprintf("char(%d)", size)
	},
	DbDataVarchar: func(size int, decimal ...int) string {
		if size == 0 {
			size = DbDataTypeSizes[DbDataVarchar][0]
		}
		return fmt.Sprintf("varchar(%d)", size)
	},
	DbDataTinytext: func(size int, decimal ...int) string {
		return "tinytext"
	},
	DbDataText: func(size int, decimal ...int) string {
		return "text"
	},
	DbDataLongText: func(size int, decimal ...int) string {
		return "longtext"
	},
	DbDataTinyblob: func(size int, decimal ...int) string {
		return "tinyblob"
	},
	DbDataBlob: func(size int, decimal ...int) string {
		return "blob"
	},
	DbDataLongblob: func(size int, decimal ...int) string {
		return "longblob"
	},
	DbDataBinary: func(size int, decimal ...int) string {
		if size == 0 {
			size = DbDataTypeSizes[DbDataBinary][0]
		}
		return fmt.Sprintf("binary(%d)", size)
	},
	DbDataVarbinary: func(size int, decimal ...int) string {
		if size == 0 {
			size = DbDataTypeSizes[DbDataVarbinary][0]
		}
		return fmt.Sprintf("varbinary(%d)", size)
	},
}
