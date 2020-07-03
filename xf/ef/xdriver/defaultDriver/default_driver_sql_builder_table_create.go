package defaultDriver

import (
	"bytes"
	"fmt"
	"github.com/go-xe2/x/type/t"
	. "github.com/go-xe2/x/xf/ef/xdriveri"
)

func (dr *TDbDefaultDriver) buildDbTypeString(ddType DbDataType, size int, decimal int) string {
	switch ddType {
	case DbDataInt:
		return "int"
	case DbDataTinyint:
		if size > 0 {
			return fmt.Sprintf("tinyint(%d)", size)
		}
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
		s := "decimal(18,2)"
		if size > 0 && decimal == 0 {
			s = fmt.Sprintf("decimal(%d,2)", size)
		} else if size == 0 && decimal > 0 {
			s = fmt.Sprintf("decimal(%d,%d)", size+decimal, decimal)
		}
		return s
	case DbDataDate:
		return "date"
	case DbDataTime:
		return "time"
	case DbDataDateTime:
		return "datetime"
	case DbDataChar:
		if size == 0 {
			return "char(30)"
		}
		return fmt.Sprintf("char(%d)", size)
	case DbDataVarchar:
		if size == 0 {
			return "varchar(60)"
		}
		return fmt.Sprintf("varchar(%d)", size)
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
		if size > 0 {
			return fmt.Sprintf("binary(%d)", size)
		}
		return fmt.Sprintf("binary(30)")
	case DbDataVarbinary:
		if size > 0 {
			return fmt.Sprintf("varbinary(%d)", size)
		}
		return fmt.Sprint("varbinary(30)")
	default:
		return "varchar(60)"
	}
}

func (dr *TDbDefaultDriver) buildTableFieldCreate(fields []DBFieldDefine) string {
	buf := bytes.NewBufferString("")
	for _, field := range fields {
		buf.WriteString(",\n")
		buf.WriteString(field.FieldName())
		buf.WriteString(dr.buildDbTypeString(field.Type(), field.Size(), field.Decimal()))
		if field.AutoIncrement() {
			buf.WriteString(" auto_increment")
		}
		defaultV := field.Default()
		if defaultV != nil {
			if defaultV == "" {
				defaultV = "''"
			}
			buf.WriteString(" default ")
			buf.WriteString(t.String(field.Default()))
		}

		if field.AllowNull() {
			buf.WriteString(" null")
		} else {
			buf.WriteString(" not null")
		}
		if field.IsPrimary() {
			buf.WriteString(" primary key")
		}
	}
	return buf.String()[1:]
}

// 生成创建数据库表脚本
func (dr *TDbDefaultDriver) BuildTableCreate(tableDefine DbTableDefine) string {
	if tableDefine == nil {
		return ""
	}
	fields := dr.buildTableFieldCreate(tableDefine.Fields())
	s := fmt.Sprintf("create table if not exists %s(\n%s\n);", tableDefine.TableName(), fields)
	return s
}
