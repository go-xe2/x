package xentity

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/utils/xvalid"
	"github.com/go-xe2/x/xf/ef/xq"
	"github.com/go-xe2/x/xf/ef/xq/sql"
	"github.com/go-xe2/x/xf/ef/xqi"
)

var _ xqi.EntityCUD = (*TEntity)(nil)

// 实现EntityEditor接口
func (ent *TEntity) FieldEdit(field xqi.EntField, value xqi.FieldValue) xqi.FieldValue {
	if value == nil {
		return nil
	}
	if ent.editFieldValues == nil {
		return value
	}
	if field.IsPrimary() {
		// 不能更新主键，只能作为更新条件表达式, 插入时，如果字段非autoincrement类型则允许插入主键
		ent.editKeyValue = value
		if ent.editState == entityUpdateState {
			return value
		}
	}
	switch ent.editState {
	case entityUpdateState:
		ent.editFieldValues = append(ent.editFieldValues, value)
		break
	case entityInsertState:
		ent.editFieldValues = append(ent.editFieldValues, value)
		break
	}
	return value
}

/*****************************/
/* 数据库实体插入api           */
/*****************************/
var _ xqi.EntityCUD = (*TEntity)(nil)

// 插入数据
func (ent *TEntity) Insert(values ...xqi.FieldValue) xqi.EntityInsert {
	return newEntityInsert(ent, values...)
}

func (ent *TEntity) InsertByMap(maps map[string]interface{}) (int, error) {
	if maps == nil {
		return 0, nil
	}
	fields := make([]xqi.FieldValue, 0)
	for k, v := range maps {
		if fd, ok := ent.fieldMaps[k]; ok {
			if tbFd, tbOk := fd.(xqi.SqlTableField); tbOk {
				fields = append(fields, tbFd.Set(v))
			}
		}
	}
	if len(fields) == 0 {
		return 0, nil
	}
	return ent.Insert(fields...).Execute()
}

func (ent *TEntity) BeginInsert() bool {
	if ent.editState != entityFreeState {
		return false
	}
	ent.editFieldValues = make([]xqi.FieldValue, 0)
	ent.editState = entityInsertState
	ent.editKeyValue = nil
	return true
}

// 提交插入
func (ent *TEntity) CommitInsert() (int, error) {
	if ent.editState != entityInsertState {
		return 0, nil
	}
	if len(ent.editFieldValues) == 0 {
		// 没有编译字段
		return 0, nil
	}
	n, err := ent.Insert(ent.editFieldValues...).Execute()
	ent.editFieldValues = nil
	ent.editState = entityFreeState
	return n, err
}

/*****************************/
/* 数据库实体更新api           */
/*****************************/
func (ent *TEntity) Update(values ...xqi.FieldValue) xqi.EntityUpdate {
	return newUpdate(ent, values...)
}

// 开始更新
func (ent *TEntity) BeginUpdate() bool {
	if ent.editState != entityFreeState {
		return false
	}
	ent.editFieldValues = make([]xqi.FieldValue, 0)
	ent.editState = entityUpdateState
	ent.editKeyValue = nil
	return true
}

// 提交更新
func (ent *TEntity) CommitUpdate(where ...xqi.SqlCondition) (int, error) {
	if ent.editState != entityUpdateState {
		return 0, nil
	}
	if len(ent.editFieldValues) == 0 {
		// 没有要更新的数据
		return 0, nil
	}
	var whereCond []xqi.SqlCondition
	if ent.editKeyValue != nil {
		whereCond = []xqi.SqlCondition{sql.Where(ent.editKeyValue.Field().Eq(ent.editKeyValue.Value()))}
	} else {
		whereCond = where
	}
	n, err := ent.Update(ent.editFieldValues...).Where(whereCond...).Execute()
	ent.editFieldValues = nil
	ent.editState = entityFreeState
	return n, err
}

// 插入或更新,如果设置的主键存在则更新否则插入数据
func (ent *TEntity) SaveMap(data map[string]interface{}) (int, error) {
	if len(data) == 0 {
		return 0, nil
	}
	fieldList := make([]xqi.FieldValue, 0)

	for k, v := range data {
		fd := ent.fieldMaps[k]
		if fd == nil {
			continue
		}
		if enFd, ok := fd.This().(xqi.EntField); ok {
			fieldList = append(fieldList, enFd.Set(v))
		}
	}
	if len(fieldList) == 0 {
		return 0, nil
	}
	return ent.Save(fieldList...)
}

// 插入或更新,如果设置了主键则更新等于主键值的行，否则插入
func (ent *TEntity) Save(fields ...xqi.FieldValue) (int, error) {
	fieldList := make([]xqi.FieldValue, 0)
	var keyFieldValue xqi.FieldValue = nil
	for _, field := range fields {
		if field.IsPrimaryKey() && keyFieldValue == nil {
			keyFieldValue = field
			continue
		}
		fieldList = append(fieldList, field)
	}
	if len(fieldList) == 0 {
		return 0, nil
	}
	if keyFieldValue != nil {
		// 检查数据是否存在，不存在时插入，否则更新数据
		b, err := ent.Exists(func(where xqi.SqlCondition, this interface{}) xqi.SqlCondition {
			return where.And(keyFieldValue.Field().Eq(keyFieldValue.Value()))
		})
		if err != nil {
			return 0, err
		}
		if b {
			// 存在该主键值，更新数据
			return ent.Update(fieldList...).Where(sql.Where(keyFieldValue.Field().Eq(keyFieldValue.Value()))).Execute()
		} else {
			return ent.Insert(append([]xqi.FieldValue{keyFieldValue}, fieldList...)...).Execute()
		}
	} else {
		// 增加
		keyField := ent.KeyField()
		if keyField != nil {
			autoIncrement := false
			// 处理主键问题
			if fieldDbAttribute := keyField.GetAnnotation(DbTypeAnnName).(EntityDbTypeAttribute); fieldDbAttribute != nil {
				autoIncrement = fieldDbAttribute.Increment()
			}
			if !autoIncrement {
				// 不是自动增长型，生成
				return 0, exception.Newf("不能向数据表%s添加数据，未传入主键值，并且主键不是自动增长型字段.", ent.TableName())
			}
		}
		return ent.Insert(fieldList...).Execute()
	}
}

func (ent *TEntity) ValidInsert(values []xqi.FieldValue, ruleCate ...string) (int, error) {
	fieldList := values
	// 数据插入前对数据有效性检查

	rules := ent.GetInsertRules(ruleCate...)
	if len(rules) > 0 {
		fieldValues := make(map[string]interface{})
		for _, value := range values {
			fieldValues[value.Field().FieldName()] = value.Value()
		}
		errs := xvalid.CheckMap(fieldValues, rules)
		if errs != nil {
			return 0, exception.NewText(errs.FirstString())
		}
	}

	if checker, existsCk := ent.This().(xqi.EntityInsertChecker); existsCk {
		var err error
		if fieldList, err = checker.CheckInsertValid(values); err != nil {
			return 0, err
		}
	}
	// 检查通过，开始插入数据
	return ent.Insert(fieldList...).Execute()
}

// 在更新数据之前，先检查数据有效后才更新，实体需要实现EntityUpdateChecker接口
func (ent *TEntity) ValidUpdate(values []xqi.FieldValue, where xqi.SqlCondition, validRule ...string) (int, error) {
	fieldList := values
	rules := ent.GetUpdateRules(validRule...)
	if len(rules) > 0 {
		fieldValues := make(map[string]interface{})
		for _, value := range values {
			fieldValues[value.Field().FieldName()] = value.Value()
		}
		errs := xvalid.CheckMap(fieldValues, rules)
		if errs != nil {
			return 0, exception.NewText(errs.FirstString())
		}
	}
	if checker, ok := ent.This().(xqi.EntityUpdateChecker); ok {
		var err error
		if fieldList, err = checker.CheckUpdateValid(values); err != nil {
			return 0, err
		}
	}
	// 检查通过，开始更新数据
	return ent.Update(fieldList...).Where(where).Execute()
}

// 使用参数数据源插入数据库，插入之前使用实体定义的插入规则过滤和检查数据有效性
func (ent *TEntity) InsertFromParams(params map[string]interface{}) (int, error) {
	fieldList, err := ent.GetInsertValuesFromParams(params, true)
	if err != nil {
		return 0, err
	}
	return ent.Insert(fieldList...).Execute()
}

// 使用参数数据源更新数据库，更新之前使用实体定义的更新规则过滤和检查数据有效性
func (ent *TEntity) UpdateFromParams(params map[string]interface{}, where ...xqi.SqlCondition) (int, error) {
	fieldList, err := ent.GetUpdateValuesFromParams(params, true)
	if err != nil {
		return 0, err
	}
	return ent.Update(fieldList...).Where(where...).Execute()
}

/*****************************/
/* 删除操作                   */
/****************************/

// 根据主键值删除
func (ent *TEntity) DeleteByKey(keyVal interface{}) (int, error) {
	if ent.KeyField() == nil {
		return 0, exception.NewText("实体未设置主键")
	}
	return ent.Delete(sql.Where(ent.KeyField().Eq(keyVal)))
}

// 根据条件删除
func (ent *TEntity) Delete(where ...xqi.SqlCondition) (int, error) {
	if observer, ok := ent.This().(xqi.EntityDeleteObserver); ok {
		var cond xqi.SqlCondition = nil
		if len(where) > 0 {
			cond = where[0]
		}
		if !observer.BeforeDelete(cond) {
			return 0, nil
		}
	}
	return xq.Delete(ent, ent.dbName...).Where(where...).Execute()
}
