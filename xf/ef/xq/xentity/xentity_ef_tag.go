package xentity

type TEFTag struct {
	// 数据表字段名称
	Name string
	// 是否主键
	PK bool
	// 是否外键
	FK bool
	// 数据库类型名称,支持int,tinyint,smallint,bigint,float,double,decimal,date,time,datetime,timestamp,char,varchar,tinytext,text,longtext,blob,tinyblob,longblob,binary,varbinary类型
	DbTypeName string
	// 字段大小
	Size int
	// 小数点位数，只对decimal数据类型有效
	Decimal int
	// 是否允许为空
	AllowNull bool
	// 是否自动增长，对int,bigint类型有效
	AutoIncrement bool
	// 字段
	LookEntity string
}
