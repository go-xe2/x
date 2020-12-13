package xentity

import (
	"github.com/go-xe2/x/core/exception"
	"github.com/go-xe2/x/sync/xsafeMap"
	"github.com/go-xe2/x/type/xstring"
	"github.com/go-xe2/x/xf/anno"
	"github.com/go-xe2/x/xf/ef/xqi"
	"reflect"
)

var entityClasses = xsafeMap.NewAnyAnyMap()

func registerEntityClasses(typ reflect.Type, class xqi.EntityClass) {
	if entityClasses.Contains(typ) {
		panic(exception.Newf("实体类型%s已经存在", typ.Name()))
	}
	entityClasses.Set(typ, class)
}

func getEntityClass(typ reflect.Type) xqi.EntityClass {
	if v := entityClasses.Get(typ); v != nil {
		return v.(xqi.EntityClass)
	}
	return nil
}

// 获取系统定义的所有实例
func GetAllEntities() []xqi.EntityClass {
	result := make([]xqi.EntityClass, 0)
	entityClasses.Foreach(func(k interface{}, v interface{}) bool {
		result = append(result, v.(xqi.EntityClass))
		return true
	})
	return result
}

func GetEntityAnnotations(typ reflect.Type) map[string]map[string]anno.AnnotationContainer {
	cls := getEntityClass(typ)
	if cls != nil {
		return cls.Annotations()
	}
	return nil
}

type entityCreate struct {
	tableName  string
	tableAlias string
	// 当前实例
	instance xqi.Entity
	// 被继承的根基类
	rootParent *TEntity
	// 实例关联的元注解，饱含继承的父类
	annotations       map[string]map[string]interface{}
	fields            map[string]xqi.EntField
	fieldConstructors map[interface{}]map[string]xqi.EntityFieldClass
	updateValidItems  map[string]xqi.FieldValid
	insertValidItems  map[string]xqi.FieldValid
}

func loopNewEntity(cls xqi.EntityClass, createInfo *entityCreate) xqi.Entity {

	inst := cls.Constructor()()
	// 收集字段类类型
	if createInfo.instance == nil {
		// 当前实例
		createInfo.instance = inst
	}
	// 创建实体字段
	fieldClsMap := cls.Fields()
	instV := reflect.ValueOf(inst)
	for instV.Kind() == reflect.Ptr {
		instV = instV.Elem()
	}

	createInfo.fieldConstructors[instV] = fieldClsMap

	// 收集字段数据检查规则
	for k, v := range cls.InsertValidItems() {
		if _, ok := createInfo.insertValidItems[k]; !ok {
			createInfo.insertValidItems[k] = v
		}
	}
	for k, v := range cls.UpdateValidItems() {
		if _, ok := createInfo.updateValidItems[k]; !ok {
			createInfo.updateValidItems[k] = v
		}
	}

	// 获取元注解
	for fdName, items := range cls.Annotations() {
		fieldAnnotations := make(map[string]interface{})
		for annK, annC := range items {
			fieldAnnotations[annK] = annC.Create(createInfo.instance)
		}
		if _, ok := createInfo.annotations[fdName]; !ok {
			createInfo.annotations[fdName] = fieldAnnotations
		}
	}

	// 父类
	parent := inst.Supper()
	if parent != nil {
		// 继承类
		var implement interface{}
		implementT := reflect.TypeOf(parent)
		for implementT.Kind() == reflect.Ptr {
			implementT = implementT.Elem()
		}
		if implementT == entityType {
			createInfo.rootParent = newEntity(createInfo.instance, createInfo.tableName, createInfo.tableAlias)
			implement = createInfo.rootParent
		} else {
			cls := getEntityClass(implementT)
			if cls != nil {
				implement = loopNewEntity(cls, createInfo)
			}
		}
		if implement != nil {
			inst.Implement(implement)
		}
	}
	return inst
}

func NewEntity(entityCls xqi.EntityClass) interface{} {
	if entityCls == nil {
		return nil
	}
	tableName := entityCls.TableName()
	tableAlias := entityCls.TableAlias()

	createInfo := &entityCreate{
		tableName:         tableName,
		tableAlias:        tableAlias,
		instance:          nil,
		rootParent:        nil,
		insertValidItems:  make(map[string]xqi.FieldValid),
		updateValidItems:  make(map[string]xqi.FieldValid),
		annotations:       make(map[string]map[string]interface{}),
		fields:            make(map[string]xqi.EntField),
		fieldConstructors: make(map[interface{}]map[string]xqi.EntityFieldClass),
	}

	inst := loopNewEntity(entityCls, createInfo)
	if createInfo.rootParent != nil {
		entity := createInfo.rootParent
		base := entity.baseEntity
		// 添加实体本身
		base.foreignTables.Add(createInfo.instance)

		// 实例化实体字段
		for k, fieldClsMap := range createInfo.fieldConstructors {
			obj := k.(reflect.Value)
			for _, fieldCls := range fieldClsMap {
				index := fieldCls.FieldIndex()
				value := fieldCls.NewField(createInfo.instance)
				obj.FieldByIndex(index).Set(reflect.ValueOf(value))
				if _, ok := createInfo.fields[fieldCls.DefineName()]; !ok {
					createInfo.fields[fieldCls.DefineName()] = value.(xqi.EntField)
				}
			}
		}
		tmpMp := make(map[string]xqi.EntField)
		for _, field := range createInfo.fields {
			// 去掉重复字段
			if _, ok := tmpMp[field.DefineName()]; ok {
				continue
			}
			tmpMp[field.DefineName()] = field

			if field.IsForeign() {
				if foreign, ok := field.This().(xqi.EFForeign); ok {
					base.foreignFields = append(base.foreignFields, foreign)
					base.foreignTables.Add(foreign.JoinTable())
				}
			}
			if field.IsPrimary() {
				entity.keyField = field
			}
			// fields 为所有字段列表，饱含外联字段
			// 字段的consructor中已经调用table的AddField方法添加字段，此处不需要现添加
			//base.fields = append(base.fields, field)
			base.fieldMaps[field.DefineName()] = field
			if _, ok := base.fieldMaps[field.FieldName()]; !ok {
				base.fieldMaps[field.FieldName()] = field
			}
			if _, ok := base.fieldMaps[field.AliasName()]; !ok {
				base.fieldMaps[field.AliasName()] = field
			}
			if valid, ok := createInfo.insertValidItems[field.DefineName()]; ok {
				entity.insertFieldValidRule = append(entity.insertFieldValidRule, &entityFieldValidRule{field: field, valid: valid})
				entity.insertFieldRuleIndex[field.FieldName()] = len(entity.insertFieldValidRule) - 1
			}
			if valid, ok := createInfo.updateValidItems[field.DefineName()]; ok {
				entity.updateFieldValidRule = append(entity.updateFieldValidRule, &entityFieldValidRule{field: field, valid: valid})
				entity.updateFieldRuleIndex[field.FieldName()] = len(entity.updateFieldValidRule) - 1
			}
			// 处理字段格式化器
			if field.Formatter() != "" {
				if fmt := formatStore2FormatFunc(field.Formatter()); fmt != nil {
					entity.fieldFormatter[field.DefineName()] = fmt
				}
			}
		}
	}
	createInfo = nil
	// 调用构造方法
	return inst.Constructor(entityCls.Attributes(), tableAlias, inst)
}

func formatStore2FormatFunc(formatter string) func(old interface{}) interface{} {
	store := getEntityFieldFormatter(formatter)
	if store == nil {
		return nil
	}
	return func(old interface{}) interface{} {
		return store.formatter(old, store.options)
	}
}

func getEntityFieldFormatter(formatter string) *fieldFormatterStore {
	infos := xstring.Split(formatter, ":", 1)
	fmtName := infos[0]

	fmt := GetFieldFormatter(fmtName)
	if fmt == nil {
		return nil
	}
	opts := make([]interface{}, 0)
	if len(infos) > 1 {
		items := xstring.Split(infos[1], ",")
		for _, item := range items {
			s := xstring.Trim(item)
			if s != "" {
				opts = append(opts, s)
			}
		}
	}
	fmtOptions := fmt.OptionFromSlice(opts...)
	return newfieldFormatterStore(fmtOptions, fmt.Formatter())
}
