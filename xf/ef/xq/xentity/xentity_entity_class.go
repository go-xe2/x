package xentity

import (
	"fmt"
	"github.com/go-xe2/x/type/slice/xstrSlice"
	"github.com/go-xe2/x/type/xstring"
	"github.com/go-xe2/x/xf/anno"
	"github.com/go-xe2/x/xf/ef/xqi"
	"reflect"
)

// 框架解析的tag名称
var XFrameworkTagNames = []string{"ef"}

// 实体字段类型
type entityFieldClass struct {
	fieldIndex []int
	// 定义的字段名
	defineName string
	// 字段关联的元注解
	annotations map[string]anno.AnnotationContainer
	// 字段数据类型
	fieldType xqi.FieldDataType
	// 是否外联字段
	isForeign bool
	// 字段构造方法
	constructor xqi.FieldConstructor
}

var _ xqi.EntityFieldClass = (*entityFieldClass)(nil)

func (efc *entityFieldClass) NewField(entity xqi.Entity, params ...interface{}) interface{} {
	if efc.constructor == nil {
		return nil
	}
	// 实例化元注解的实际类
	annotations := make(map[string]interface{})
	var fieldAttrs = make([]xqi.XqAttribute, 0)
	for k, c := range efc.annotations {
		if ann := c.Create(entity, params...); ann != nil {
			annotations[k] = ann
			if attr, ok := ann.(xqi.XqAttribute); ok {
				fieldAttrs = append(fieldAttrs, attr)
			}
		}
	}
	return efc.constructor(entity, efc.defineName, fieldAttrs, annotations, params...)
}

func (efc *entityFieldClass) DefineName() string {
	return efc.defineName
}

func (efc *entityFieldClass) Annotations() map[string]anno.AnnotationContainer {
	return efc.annotations
}

func (efc *entityFieldClass) FieldType() xqi.FieldDataType {
	return efc.fieldType
}

func (efc *entityFieldClass) IsForeign() bool {
	return efc.isForeign
}

func (efc *entityFieldClass) Constructor() xqi.FieldConstructor {
	return efc.constructor
}

func (efc *entityFieldClass) FieldIndex() []int {
	return efc.fieldIndex
}

func newEntityFieldClass(fieldIndex []int, defineName string, defineType reflect.Type, annotations map[string]anno.AnnotationContainer) *entityFieldClass {
	isForeign := false
	isExpr := false
	var fieldType = xqi.FDTUnknown
	switch defineType {
	case efStringType:
		fieldType = xqi.FDTString
		break
	case efIntType:
		fieldType = xqi.FDTInt
		break
	case efInt8Type:
		fieldType = xqi.FDTInt8
		break
	case efInt16Type:
		fieldType = xqi.FDTInt16
		break
	case efInt32Type:
		fieldType = xqi.FDTInt32
		break
	case efInt64Type:
		fieldType = xqi.FDTInt64
		break
	case efUintType:
		fieldType = xqi.FDTUint
		break
	case efUint8Type:
		fieldType = xqi.FDTUint8
		break
	case efUint16Type:
		fieldType = xqi.FDTUint16
		break
	case efUint32Type:
		fieldType = xqi.FDTUint32
		break
	case efUint64Type:
		fieldType = xqi.FDTUint64
		break
	case efFloatType:
		fieldType = xqi.FDTFloat
		break
	case efDoubleType:
		fieldType = xqi.FDTDouble
		break
	case efBoolType:
		fieldType = xqi.FDTBool
		break
	case efDateType:
		fieldType = xqi.FDTDatetime
		break
	case efByteType:
		fieldType = xqi.FDTByte
		break
	case efBinaryType:
		fieldType = xqi.FDTBinary
		break
	case efForeignType:
		isForeign = true
		break
	case efExprType:
		isExpr = true
		break
	default:
		fieldType = xqi.FieldDataType(efTypeConstructorEntry.GetUserFieldType(defineType))
	}
	if !isForeign && !isExpr && fieldType == xqi.FDTUnknown {
		// 未注册的类型，框架不会自动初始化该字段
		return nil
	}
	result := &entityFieldClass{fieldIndex: fieldIndex, fieldType: fieldType, isForeign: isForeign, defineName: defineName, annotations: annotations}
	if !isForeign && !isExpr && fieldType != xqi.FDTUnknown {
		result.constructor = func(entity xqi.Entity, fieldName string, attrs []xqi.XqAttribute, annotations map[string]interface{}, callParams ...interface{}) interface{} {
			return CreateField(fieldType, entity, fieldName, attrs, annotations, callParams...)
		}
	} else if isForeign {
		// 外联字段
		result.constructor = func(entity xqi.Entity, defineName string, attrs []xqi.XqAttribute, annotations map[string]interface{}, callParams ...interface{}) interface{} {
			return newEFForeignField(entity, defineName, attrs, annotations, callParams...)
		}
	} else if isExpr {
		result.constructor = func(entity xqi.Entity, defineName string, attrs []xqi.XqAttribute, annotations map[string]interface{}, callParams ...interface{}) interface{} {
			return newEFExprField(entity, defineName, attrs, annotations, callParams...)
		}
	}
	return result
}

type tEntityClass struct {
	entityAttribute xqi.EntityAttribute
	// 实体类型
	entType reflect.Type
	// 实体类型创建方法
	constructor func() xqi.Entity
	// 实体类型的tag
	attributes []xqi.XqAttribute
	// 字段的元注解
	fieldAnnotations map[string]map[string]anno.AnnotationContainer
	fields           map[string]xqi.EntityFieldClass
	// 字段定义名 => 规则
	// 插入字段对应规则
	insertFieldValid map[string]xqi.FieldValid
	// 更新字段对应规则
	// 字段定义名 => 规则
	updateFieldValid map[string]xqi.FieldValid
}

var _ xqi.EntityClass = (*tEntityClass)(nil)

func ClassOfEntity(constructor func() xqi.Entity, attr []xqi.XqAttribute, foreignFieldTypes ...*EntForeignFieldDefine) xqi.EntityClass {
	tmp := constructor()
	cls := &tEntityClass{attributes: attr}
	cls.constructor = constructor
	cls.fields = make(map[string]xqi.EntityFieldClass)
	clsV := reflect.ValueOf(tmp)
	cls.entType = clsV.Type()
	cls.fieldAnnotations = make(map[string]map[string]anno.AnnotationContainer)
	cls.insertFieldValid = make(map[string]xqi.FieldValid)
	cls.updateFieldValid = make(map[string]xqi.FieldValid)

	for cls.entType.Kind() == reflect.Ptr {
		cls.entType = cls.entType.Elem()
		clsV = clsV.Elem()
	}
	for _, foreignField := range foreignFieldTypes {
		registerFieldDefine(foreignField)
	}
	buildEntityFieldClass(cls, clsV)
	registerEntityClasses(cls.entType, cls)

	return cls
}

func buildEntityFieldClass(cls *tEntityClass, entV reflect.Value) {
	entT := entV.Type()
	count := entV.NumField()
	for i := 0; i < count; i++ {
		field := entT.Field(i)
		defineName := field.Name
		// 不处理私有字段
		if xstring.IsFirstLetterLower(defineName) && defineName[0] != '_' {
			continue
		}
		// 解析元注解annotation
		fieldTag := ""
		for _, tagName := range XFrameworkTagNames {
			if fieldTag = field.Tag.Get(tagName); fieldTag != "" && fieldTag != "-" {
				break
			}
		}
		annotations := make(map[string]anno.AnnotationContainer)
		// 此处理改成在需要的时间再创建元注解
		if fieldTag != "" {
			szAnnItems := xstring.Split(fieldTag, ";")
			for _, szAnn := range szAnnItems {
				szAnn = xstring.Trim(szAnn)
				if !anno.IsAnnotationString(szAnn) {
					continue
				}
				if annContainer := anno.Create(szAnn); annContainer != nil {

					// 字段检查规则，不需要添加到运行时
					if annContainer.AnnName() == FieldValidAnnName {
						fieldValid := annContainer.Create(entV.Interface()).(xqi.FieldValid)
						operations := fieldValid.Operation()
						if len(operations) == 0 || operations[0] == "" {
							cls.insertFieldValid[defineName] = fieldValid
							cls.updateFieldValid[defineName] = fieldValid
						} else {
							if xstrSlice.Contain(operations, "I") {
								cls.insertFieldValid[defineName] = fieldValid
							}
							if xstrSlice.Contain(operations, "U") {
								cls.updateFieldValid[defineName] = fieldValid
							}
						}
					} else {
						annotations[annContainer.AnnName()] = annContainer
					}
				}
			}
		}

		fieldCls := newEntityFieldClass(field.Index, defineName, field.Type, annotations)
		if fieldCls == nil {
			continue
		}
		cls.fields[defineName] = fieldCls
		cls.fieldAnnotations[defineName] = annotations
	}
}

func (ec *tEntityClass) EntType() reflect.Type {
	return ec.entType
}

func (ec *tEntityClass) Constructor() func() xqi.Entity {
	return ec.constructor
}

func (ec *tEntityClass) Attributes() []xqi.XqAttribute {
	return ec.attributes
}

func (ec *tEntityClass) Annotations() map[string]map[string]anno.AnnotationContainer {
	return ec.fieldAnnotations
}

func (ec *tEntityClass) Create() interface{} {
	return NewEntity(ec)
}

func (ec *tEntityClass) initEntityAttr() {
	if ec.entityAttribute == nil && len(ec.attributes) > 0 {
		for _, attr := range ec.attributes {
			if entAttr, ok := attr.This().(xqi.EntityAttribute); ok {
				ec.entityAttribute = entAttr
				break
			}
		}
	}
}

func (ec *tEntityClass) TableName() string {
	ec.initEntityAttr()
	if ec.entityAttribute != nil {
		return ec.entityAttribute.TableName()
	}
	return ""
}

func (ec *tEntityClass) InsertValidItems() map[string]xqi.FieldValid {
	return ec.insertFieldValid
}

func (ec *tEntityClass) UpdateValidItems() map[string]xqi.FieldValid {
	return ec.updateFieldValid
}

func (ec *tEntityClass) TableAlias() string {
	ec.initEntityAttr()
	if ec.entityAttribute != nil {
		return ec.entityAttribute.TableAlias()
	}
	return ""
}

func (ec *tEntityClass) Fields() map[string]xqi.EntityFieldClass {
	return ec.fields
}

func (ec *tEntityClass) String() string {
	if ec.entType == nil {
		return "nil"
	}
	return fmt.Sprintf("{ class of entity(%s %s) }", ec.TableName(), ec.TableAlias())
}
