package xqcomm

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/xf/ef/xdriveri"
	"github.com/go-xe2/x/xf/ef/xqi"
)

type tSqlInsertExp struct {
	// 要插入数据的表
	insertTable xqi.SqlTable
	// 要插入的字段列表
	fieldExps []xqi.FieldValue
	// 从其他表或查询结果插入时的源表
	fromTable xqi.SqlTable
}

var _ xqi.SqlInsertExp = (*tSqlInsertExp)(nil)

func NewSqlInsertExp(insertTable xqi.SqlTable) xqi.SqlInsertExp {
	inst := &tSqlInsertExp{}
	return inst.Table(insertTable)
}

func (exp *tSqlInsertExp) TokenType() xqi.SqlTokenType {
	return xqi.SqlInsertTokenType
}

func (exp *tSqlInsertExp) This() interface{} {
	return exp
}

// 编译表达式
func (exp *tSqlInsertExp) Compile(builder xdriveri.DbDriverSqlBuilder, cxt xqi.SqlCompileContext, unPrepare ...bool) xqi.SqlToken {
	if exp.insertTable == nil {
		panic(exception.NewText("没有选择要插入的表"))
	}
	if exp.fieldExps == nil {
		panic(exception.NewText("没有设置要插入数据的字段"))
	}

	result := NewSqlToken("", exp.TokenType())

	// 要插入数据的表名
	szTable := ""
	// 要插入数据的字段列表
	szFieldList := ""
	// 要插入的值列表
	szValueList := ""
	// insert from 表达式中的数据源
	szFromExp := ""

	isPrepare := true
	if len(unPrepare) > 0 {
		isPrepare = !unPrepare[0]
	}
	cxt.PushState(xqi.SCPBuildInsertTableState)
	if tk := exp.insertTable.Compile(builder, cxt, unPrepare...); tk != nil && tk.TType() != xqi.SqlEmptyTokenType {
		szTable = tk.Val()
	} else {
		return EmptySqlToken
	}
	cxt.PopState()

	// 用于接收关联到的表

	// 添加待插入数据的表到当前上下文
	cxt.Clear()
	cxt.Tables().Add(exp.insertTable)

	// 生成插入字段列表
	fdCount := len(exp.fieldExps)
	for i := 0; i < fdCount; i++ {
		fdExp := exp.fieldExps[i]
		if fdExp == nil {
			continue
		}
		field := fdExp.Field()
		fieldValue := fdExp.Value()
		if field == nil {
			continue
		}
		cxt.PushState(xqi.SCPBuildInsertFieldNameState)
		if ftk := field.Compile(builder, cxt, unPrepare...); ftk == nil || ftk.TType() == xqi.SqlEmptyTokenType {
			panic(exception.Newf("insert表达式错误，插入表%s的字段%s编译错误", exp.insertTable.TableName(), field.FieldName()))
		} else {
			szFieldList += "," + ftk.Val()
		}
		cxt.PopState()

		// 编译value部份
		szValue := ""
		if c, ok := fieldValue.(xqi.SqlCompiler); ok {
			cxt.PushState(xqi.SCPBuildInsertFieldValueState)
			if tk := c.Compile(builder, cxt, unPrepare...); tk == nil || tk.TType() == xqi.SqlEmptyTokenType {
				panic(exception.Newf("insert表达式错误，插入表%s的字段%s编译错误", exp.insertTable.TableName(), field.FieldName()))
			} else {
				szValue = tk.Val()
			}
			cxt.PopState()
		} else {
			if isPrepare {
				vn := cxt.MakeParamId()
				result.AddParam(vn, fieldValue)
				cxt.AddParam(vn, fieldValue)
				szValue = builder.PlaceHolder(vn)
			} else {
				szValue = builder.MakeRealValue(fieldValue)
			}
		}
		szValueList += "," + szValue
	}

	// 收集表及参数
	tables := cxt.Tables().All()

	for _, table := range tables {
		if table != exp.fromTable && table != exp.insertTable {
			szMsg := ""
			if exp.fromTable != nil {
				szMsg = "及数据源表" + exp.fromTable.TableName()
			}
			panic(exception.Newf("表达式中表[%s %s]无效，表达式只用到表:%s %s", table.TableName(), table.TableAlias(), exp.insertTable.TableName(), szMsg))
		}
		cxt.Tables().Add(table)
	}

	// 暂时不需要当前表达式返回的参数
	//params := cxt.Params()
	//for _, param := range params {
	//	cxt.Add(param)
	//	result.AddParam(param.Name(), param.Val())
	//}

	isFromInsertExpr := cxt.Tables().Count() > 1
	if isFromInsertExpr && exp.fromTable == nil {
		panic(exception.Newf("缺少数据来源表from表达式f"))
	}

	if isFromInsertExpr {
		cxt.PushState(xqi.SCPBuildInsertFromState)
		if tk := exp.fromTable.Compile(builder, cxt, unPrepare...); tk == nil || tk.TType() == xqi.SqlEmptyTokenType {
			panic(exception.Newf("from表达式错误"))
		} else {
			szFromExp = tk.Val()
		}
		cxt.PopState()
	}
	szExp := builder.BuildInsert(szTable, szFieldList[1:], szValueList[1:], szFromExp)
	return result.SetVal(szExp)
}
