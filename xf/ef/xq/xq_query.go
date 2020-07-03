package xq

import (
	xqcomm2 "github.com/go-xe2/x/xf/ef/xq/xqcomm"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type SqlQuery = xqi.SqlQuery

type (
	SqlField         = xqi.SqlField
	SqlTable         = xqi.SqlTable
	SqlCondition     = xqi.SqlCondition
	SqlConditionItem = xqi.SqlConditionItem
	SqlLogic         = xqi.SqlConditionLogic
	SqlTableField    = xqi.SqlTable
	SqlFun           = xqi.SqlFun
	SqlCase          = xqi.SqlCase
	SqlCaseThenElse  = xqi.SqlCaseThenElse

	SqlBaseTable = *xqcomm2.TSqlTable
	SqlTables    = xqi.SqlTables

	DbType = xqi.DbType
)

// 数据库数据类型
var (
	// 数据库int类型
	DTInt = xqcomm2.DTInt
	// 数据库tinyint类型
	DTTinyint = xqcomm2.DTTinyint
	// 数据库smallint类型
	DTSmallint = xqcomm2.DTSmallint
	// 数据库bigint类型
	DTBigint = xqcomm2.DTBigint

	// 数据库float类型
	DTFloat = xqcomm2.DTFloat
	// 数据库double类型
	DTDouble = xqcomm2.DTDouble
	// 数据库decimal类型
	DTDecimal = xqcomm2.DTDecimal

	// 数据库date类型
	DTDate = xqcomm2.DTDate
	// 数据库time类型
	DTTime = xqcomm2.DTTime
	// 数据库datetime类型
	DTDatetime = xqcomm2.DTDatetime
	// 数据库timestamp类型
	DTTimestamp = xqcomm2.DTTimestamp

	// 数据库char类型
	DTChar = xqcomm2.DTChar
	// 数据库varchar类型
	DTVarchar = xqcomm2.DTVarchar
	// 数据库tinytext类型
	DTTinytext = xqcomm2.DTTinytext
	// 数据库text类型
	DTText = xqcomm2.DTText
	// 数据库longtext类型
	DTLongtext = xqcomm2.DTLongtext

	// 数据库blob类型
	DTBlob = xqcomm2.DTBlob
	// 数据库tinyblob类型
	DTTinyblob = xqcomm2.DTTinyblob
	// 数据库longblob类型
	DTLongblob = xqcomm2.DTLongblob

	// 数据库binary类型
	DTBinary = xqcomm2.DTBinary
	// 数据库varbinary类型
	DTVarbinary = xqcomm2.DTVarbinary
)

// sql 常用函数方法
var (
	SFCase       = xqcomm2.SqlFunCase
	SFSubstr     = xqcomm2.SqlFunSubstring
	SFConcat     = xqcomm2.SqlFunStrConcat
	SFDateFormat = xqcomm2.SqlFunDateFormat
	SFDateDiff   = xqcomm2.SqlFunDateFormat
	SFDateAdd    = xqcomm2.SqlFunDateAdd
	SFDateSub    = xqcomm2.SqlFunDateSub
	SFDateToUnix = xqcomm2.SqlFunDateToUnix
	SFUnixToDate = xqcomm2.SqlFunUnixToDate
	SFIfNull     = xqcomm2.SqlFunIfNull
	SFIsNull     = xqcomm2.SqlFunIsNull
	SFIf         = xqcomm2.SqlFunIf
	SFCast       = xqcomm2.SqlFunCast
)

// 创建方法
var (
	NewBaseTable = xqcomm2.NewSqlTable
	NewBaseField = xqcomm2.NewSqlTableField
	NewField     = xqcomm2.NewSqlField
	RealValue    = func(val interface{}, varName ...string) xqi.SqlVarExpr {
		return xqcomm2.NewSqlVar(nil, val, varName...)
	}
	NewSqlContext = xqcomm2.NewSqlCompileContext
)
