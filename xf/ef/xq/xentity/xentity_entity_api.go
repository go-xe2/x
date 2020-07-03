package xentity

import (
	"fmt"
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/type/slice/xstrSlice"
	"github.com/go-xe2/x/type/xstring"
	"github.com/go-xe2/x/xf/ef/xdriver"
	"github.com/go-xe2/x/xf/ef/xq"
	"github.com/go-xe2/x/xf/ef/xq/xqcomm"
	"github.com/go-xe2/x/xf/ef/xqi"
)

// 创建实体之间的关系
func (ent *TEntity) buildRelation() xqi.SqlTables {
	if ent.isBuildRelation() {
		return ent.foreignTables
	}
	cxt := xqcomm.NewSqlCompileContext()
	builder := xdriver.GetSqlBuilderByName(xq.Database().Driver())

	ent.joinConditions = make(map[string]*foreignItem)

	// 字段与关联到的表实体
	fieldTables := make(map[xq.SqlField][]xqi.SqlTable)

	// 收集所用的所有实体,关联

	for _, field := range ent.foreignFields {
		cxt.Tables().Clear()
		item := &foreignItem{}
		item.table = field.JoinTable()
		item.joinType = field.JoinType()
		item.relationTables = make([]xqi.SqlTable, 0)

		condition := field.On()(xqcomm.NewSqlCondition(), item.table, ent.foreignTables)
		// 预编译join条件表达式，以便获取字段所关联到的表实体
		condition.Compile(builder, cxt, true)

		item.onCondition = (func(condition xqi.SqlCondition) func(joinTable xqi.SqlTable, otherTables xqi.SqlTables, on xqi.SqlCondition) xqi.SqlCondition {
			return func(joinTable xqi.SqlTable, otherTables xqi.SqlTables, on xqi.SqlCondition) xqi.SqlCondition {
				return condition
			}
		})(condition)
		// 合并字段所使用到的表实体
		tables := cxt.Tables().All()
		// 关联到的表
		item.relationTables = append(item.relationTables, tables...)
		tables = append(tables, item.table)
		for _, table := range tables {
			if !ent.foreignTables.HasTable(table.TableName()) {
				panic(exception.Newf("外联字段%s使用到的实体%s未定义关联关系", field.AliasName(), table.TableName()))
			}
		}
		tbName := item.table.TableName()
		if item.table.TableAlias() != "" {
			tbName = item.table.TableAlias()
		}
		if _, ok := ent.joinConditions[tbName]; !ok {
			ent.joinConditions[tbName] = item
		}
		fieldTables[field] = tables
	}
	// 建立字段与实体关系
	ent.fieldConditions = make(map[xqi.SqlField][]*foreignItem)
	for field, tables := range fieldTables {
		fieldItems := make([]*foreignItem, 0)
		for _, table := range tables {
			if table == ent.This() {
				// 不需要添加表实体本身
				continue
			}
			tbName := table.TableName()
			if table.TableAlias() != "" {
				tbName = table.TableAlias()
			}
			condition := ent.joinConditions[tbName]
			fieldItems = append(fieldItems, condition)
		}
		ent.fieldConditions[field] = fieldItems
	}
	return ent.foreignTables
}

type fieldForeignItem struct {
	item  *foreignItem
	field xqi.SqlField
}

func (ent *TEntity) buildJoinCondition(fields []xqi.SqlField) (result []*foreignItem, nameMaps map[string]string) {
	ent.buildRelation()
	tmp := make(map[xqi.SqlTable]*fieldForeignItem)
	// 字段定义与查询结果名称映射
	nameMaps = make(map[string]string)

	for _, field := range fields {
		inst := field.This()
		ef, ok := inst.(xqi.EntField)
		if !ok {
			panic(exception.Newf("字段%s非实体字段", field.AliasName()))
		}
		if ef.Entity() != ent.This() {
			panic(exception.Newf("字段%s非实体%s定义的字段", field.AliasName(), ent.TableName()))
		}
		if _, ok := inst.(xqi.EFForeign); ok {
			if fieldItems, ok := ent.fieldConditions[field]; ok {
				for _, item := range fieldItems {
					if _, ok := tmp[item.table]; !ok {
						tmp[item.table] = &fieldForeignItem{item: item, field: field}
					}
				}
			}
		}
		fdName := ef.FieldName()
		if ef.AliasName() != "" {
			fdName = ef.AliasName()
		}
		nameMaps[ef.DefineName()] = fdName
	}

	// 根据引用关系排序，与实体本身有关系的项排在最前，实关联到的实体项应该在该项之前定义
	nLen := len(tmp)
	result = make([]*foreignItem, 0)
	var isDefine = func(table xqi.SqlTable) bool {
		if table == ent.This() {
			return true
		}
		for _, tb := range result {
			if tb.table == table {
				return true
			}
		}
		return false
	}

	//k := 0
search:
	for k, item := range tmp {
		for _, relationItem := range item.item.relationTables {
			if relationItem == item.item.table {
				continue
			}
			if isDefine(relationItem) {
				// 关联的关系已经定义，直接加入
				goto isDefineProcess
			}
		}
		continue
	isDefineProcess:
		result = append(result, item.item)
		delete(tmp, k)
		goto search
	}
	if len(result) != nLen {
		szMsg := ""
		for _, item := range tmp {
			szMsg += fmt.Sprintf("字段%s关联表%s未能关联到表%s\n", item.field.AliasName(), item.item.table.TableName(), ent.TableName())
		}
		panic(exception.Newf("实体关系定义不完整:%s", szMsg))
	}
	return
}

func (ent *TEntity) LastSql() string {
	return ent.lastSql
}

func (ent *TEntity) KeyField() xqi.EntField {
	if ent.keyField == nil {
		for _, f := range ent.fieldMaps {
			if efd, ok := f.This().(xqi.EntField); ok {
				if efd.IsPrimary() {
					ent.keyField = efd
					break
				}
			}
		}
	}
	return ent.keyField
}

// 获取指定规则的字段列表
func (ent *TEntity) getFieldListByRule(rule string) []xqi.SqlField {
	fieldList := ent.fields
	if rule != "" {
		fieldList = make([]xqi.SqlField, 0)
		for i := 0; i < ent.FieldCount(); i++ {
			field, ok := ent.Field(i).(xqi.EntField)
			if !ok {
				continue
			}
			szRule := field.Rule()
			if szRule == "" {
				// 没有设置规则，默认为公共字段
				fieldList = append(fieldList, field)
				continue
			}
			ruleItems := xstring.Split(szRule, "|")
			needRules := xstring.Splits(rule, "|", ",", " ")
			for _, r := range needRules {
				if xstrSlice.Contain(ruleItems, r) {
					fieldList = append(fieldList, field)
					break
				}
			}
		}
	}
	return fieldList
}
