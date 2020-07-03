package xdriveri

// 数据库类型
type DbDataType int

const (
	DbDataInt DbDataType = iota
	DbDataTinyint
	DbDataSmallint
	DbDataBigint

	DbDataFloat
	DbDataDouble
	DbDataDecimal

	DbDataDate
	DbDataTime
	DbDataDateTime
	DbDataTimestamp

	DbDataChar
	DbDataVarchar
	DbDataTinytext
	DbDataText
	DbDataLongText

	DbDataBlob
	DbDataTinyblob
	DbDataLongblob

	DbDataBinary
	DbDataVarbinary
)

var dbDataTypeNameMaps = map[string]DbDataType{
	"int": DbDataInt,
	"Int": DbDataInt,

	"tinyint": DbDataInt,
	"Tinyint": DbDataTinyint,

	"smallint": DbDataSmallint,
	"Smallint": DbDataSmallint,

	"bigint": DbDataBigint,
	"Bigint": DbDataBigint,

	"float": DbDataFloat,
	"Float": DbDataFloat,

	"double": DbDataDouble,
	"Double": DbDataDouble,

	"decimal": DbDataDecimal,
	"Decimal": DbDataDecimal,

	"date": DbDataDate,
	"Date": DbDataDate,

	"time": DbDataTime,
	"Time": DbDataTime,

	"datetime": DbDataDateTime,
	"Datetime": DbDataDateTime,

	"char": DbDataChar,
	"Char": DbDataChar,

	"varchar": DbDataVarchar,
	"Varchar": DbDataVarchar,

	"tinytext": DbDataTinytext,
	"Tinytext": DbDataTinytext,

	"text": DbDataText,
	"Text": DbDataText,

	"longtext": DbDataLongText,
	"Longtext": DbDataLongText,

	"blob": DbDataBlob,
	"Blob": DbDataBlob,

	"tinyblob": DbDataTinyblob,
	"Tinyblob": DbDataTinyblob,

	"longblob": DbDataLongblob,
	"Longblob": DbDataLongblob,

	"binary": DbDataBinary,
	"Binary": DbDataBinary,

	"varbinary": DbDataVarbinary,
	"Varbinary": DbDataVarbinary,
}

func (dt DbDataType) Name() string {
	switch dt {
	case DbDataInt:
		return "int"
	case DbDataTinyint:
		return "tinyint"
	case DbDataSmallint:
		return "smallint"
	case DbDataBigint:
		return "bigint"
	case DbDataFloat:
		return "float"
	case DbDataDouble:
		return "double"
	case DbDataDecimal:
		return "decimal"
	case DbDataDate:
		return "date"
	case DbDataTime:
		return "time"
	case DbDataDateTime:
		return "datetime"
	case DbDataTimestamp:
		return "timestamp"
	case DbDataChar:
		return "char"
	case DbDataVarchar:
		return "varchar"
	case DbDataTinytext:
		return "tinytext"
	case DbDataText:
		return "text"
	case DbDataLongText:
		return "longtext"
	case DbDataBlob:
		return "blob"
	case DbDataTinyblob:
		return "tinyblob"
	case DbDataLongblob:
		return "longblob"
	case DbDataBinary:
		return "binary"
	case DbDataVarbinary:
		return "varbinary"
	}
	return "unknown"
}

func (dt DbDataType) String() string {
	return dt.Name()
}

// 由字符串解析成DbDataType
func (dt *DbDataType) Parse(name string) DbDataType {
	if n, ok := dbDataTypeNameMaps[name]; ok {
		*dt = n
	} else {
		*dt = DbDataVarchar
	}
	return *dt
}
