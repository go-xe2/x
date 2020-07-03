package xentity

import (
	"github.com/go-xe2/x/xf/ef/xq/xbinder"
	"github.com/go-xe2/x/xf/ef/xqi"
)

func (eqs *tEntitySelect) setLastSql(sql string) {
	eqs.entity.lastSql = sql
}

func (eqs *tEntitySelect) getFieldConvert() map[string]func(old interface{}) interface{} {
	if eqs.fieldNameMaps == nil {
		return nil
	}
	// 处理从字段定义名 -> 别名的对应数据格式化方法
	result := make(map[string]func(old interface{}) interface{})
	formatters := eqs.entity.fieldFormatter
	for k, v := range eqs.fieldNameMaps {
		if fn, ok := formatters[k]; ok {
			result[v] = fn
		}
	}
	return result
}

func (eqs *tEntitySelect) initBinder(binder xqi.DbQueryBinder) xqi.DbQueryBinder {
	convert := eqs.getFieldConvert()
	if binder == nil {
		if convert == nil {
			return xbinder.GetQueryBinder(xqi.MapBinder)
		} else {
			return xbinder.NewQryMapBinder(nil, convert)
		}
	}
	if convert == nil {
		return binder
	}
	return binder.NewInstance(map[string]interface{}{
		"converts": convert,
	})
}
