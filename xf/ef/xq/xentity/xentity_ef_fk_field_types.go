package xentity

import (
	"github.com/go-xe2/x/xf/ef/xqi"
)

type EntForeignFieldDefine struct {
	ForeignKey  string
	JoinType    xqi.SqlJoinType
	JoinEnt     xqi.EntityClass
	Look        func(ent interface{}) xqi.SqlField
	OnCondition func(on xqi.SqlCondition, joinEnt interface{}, leftEntities xqi.SqlTables) xqi.SqlCondition
}

// 定义
func TypeLeftJoinField(foreignKey string, joinEnt xqi.EntityClass, look func(ent interface{}) xqi.SqlField, on func(on xqi.SqlCondition, joinEnt interface{}, leftEntities xqi.SqlTables) xqi.SqlCondition) *EntForeignFieldDefine {
	return defineJoinField(xqi.SqlLeftJoinType, foreignKey, joinEnt, look, on)
}

func TypeRightJoinField(foreignKey string, joinEnt xqi.EntityClass, look func(ent interface{}) xqi.SqlField, on func(on xqi.SqlCondition, joinEnt interface{}, leftEntities xqi.SqlTables) xqi.SqlCondition) *EntForeignFieldDefine {
	return defineJoinField(xqi.SqlRightJoinType, foreignKey, joinEnt, look, on)
}

func TypeCrossJoinField(foreignKey string, joinEnt xqi.EntityClass, look func(ent interface{}) xqi.SqlField, on func(on xqi.SqlCondition, joinEnt interface{}, leftEntities xqi.SqlTables) xqi.SqlCondition) *EntForeignFieldDefine {
	return defineJoinField(xqi.SqlCrossJoinType, foreignKey, joinEnt, look, on)
}

func TypeJoinField(foreignKey string, joinEnt xqi.EntityClass, look func(ent interface{}) xqi.SqlField, on func(on xqi.SqlCondition, joinEnt interface{}, leftEntities xqi.SqlTables) xqi.SqlCondition) *EntForeignFieldDefine {
	return defineJoinField(xqi.SqlInnerJoinType, foreignKey, joinEnt, look, on)
}

func defineJoinField(joinType xqi.SqlJoinType, foreignKey string, joinEnt xqi.EntityClass, look func(ent interface{}) xqi.SqlField, on func(on xqi.SqlCondition, joinEnt interface{}, leftEntities xqi.SqlTables) xqi.SqlCondition) *EntForeignFieldDefine {
	return &EntForeignFieldDefine{
		JoinType:    joinType,
		ForeignKey:  foreignKey,
		JoinEnt:     joinEnt,
		Look:        look,
		OnCondition: on,
	}
}
