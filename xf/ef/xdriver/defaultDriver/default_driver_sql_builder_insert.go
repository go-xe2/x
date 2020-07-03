package defaultDriver

import (
	. "github.com/go-xe2/x/xf/ef/xdriveri"
)

func (dr *TDbDefaultDriver) joinFields(fields []string, sep string) string {
	result := ""
	for _, s := range fields {
		result += dr.QuotesName(s) + ","
	}
	return result[0 : len(result)-2]
}

func (dr *TDbDefaultDriver) InsertExecute(ssl *TExprInsertSqlSSL) error {
	//if ssl == nil {
	//	return exception.NewText("中间脚本无效，插入数据失败.")
	//}
	//if ssl.Execute == nil {
	//	return exception.NewText("未绑定执行方法，插入数据失败")
	//}
	//if len(ssl.Values) == 0 {
	//	return exception.NewText("没有要插入的数据.")
	//}
	//if ssl.ValueMode == SqlInsertStaticValueType && ssl.Values == nil {
	//	return exception.NewText("没有要插入的数据")
	//}
	//if ssl.ValueMode == SqlInsertFromScriptType && ssl.QuerySql == "" {
	//	return exception.NewText("没有设置数据查询语句")
	//}
	//fields := dr.joinFields(ssl.Fields, ",")
	//switch ssl.ValueMode {
	//case SqlInsertStaticValueType:
	//	rowCount := len(ssl.Values)
	//	values := xdbUtil.MakeStrAndFill(len(ssl.Fields), dr.PlaceHolder(), ",")
	//	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES ", dr.QuotesName(ssl.TableName), fields)
	//	if rowCount == 1 {
	//		// 只有一行数据时
	//		sql += values
	//		ssl.Execute(sql, ssl.Values[0])
	//	} else {
	//		isPrepare := true
	//		sql += values
	//		if ssl.ExecutePrepare != nil {
	//			isPrepare = ssl.ExecutePrepare(sql)
	//		}
	//		if isPrepare {
	//			// 执行插入
	//			for _, vars := range ssl.Values {
	//				ssl.Execute(sql, vars)
	//			}
	//		}
	//	}
	//	break
	//case SqlInsertFromScriptType:
	//	// 使用sql查询结果插入
	//	sql := fmt.Sprintf("INSERT INTO %s (%s) \n%s", dr.QuotesName(ssl.TableName), fields, ssl.QuerySql)
	//	vars := make([]interface{}, 0)
	//	if len(ssl.Values) > 0 {
	//		vars = ssl.Values[0]
	//	}
	//	ssl.Execute(sql, vars)
	//	break
	//}
	//return nil
	return nil
}
