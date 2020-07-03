package xqi

// sql编译状态
type SCPState int

const (
	// 编译查询字段
	SCPUnknown SCPState = iota
	SCPQrSelectFieldsState
	SCPQrSelectFieldState
	// 查询表达式字段状态
	SCPQrSelectExprFieldState
	// 查询表
	SCPQrSelectFromState
	SCPQrSelectWhereCondState
	SCPQrSelectJoinTableState
	SCPQrSelectJoinCondState
	SCPQrSelectJoinItemState
	SCPQrSelectGroupFieldState
	SCPQrSelectHavingCondState
	SCPQrSelectOrderFieldState
	// 生成参数名称
	SCPQrSelectParamNameState
	SCPQrSelectFunParamState
	// ====== 添加、更新查卷状态 ==== //

	// 创建插入表名状态
	SCPBuildInsertTableState
	// 创建插入字段名状态
	SCPBuildInsertFieldNameState
	// 创建插入字段值状态
	SCPBuildInsertFieldValueState
	// 创建从其他表源插入的from部份脚本状态
	SCPBuildInsertFromState
	// 创建更新表的表名状态
	SCPBuildUpdateTableState
	// 创建联合更新的表名
	SCPBuildUpdateTableWithFromState

	// 创建更新字段状态
	SCPBuildUpdateFieldState
	// 创建联合被更新的字段名
	SCPBuildUpdateFieldWithFromState
	// 创建更新字段值表达式使用：表别名.字段名
	SCPBuildUpdateFieldValueState
	// 创建联合更新字段值表达式
	SCPBuildUpdateFieldValueFromState
	// 创建关联更新join部份表名状态
	SCPBuildUpdateJoinTableState
	// 创建联合更新关联关系表达式状态
	SCPBuildUpdateJoinOnState
	// 创建更新条件时的状态
	SCPBuildUpdateWhereState
	// 创建联合更新时的条件表达式状态
	SCPBuildUpdateWhereFromState

	// 创建被删除表名状态
	SCPBuildDeleteTableState
	// 删除表条件表达式状态
	SCPBuildDeleteWhereState
)

var scpStateNames = map[SCPState]string{
	SCPUnknown:                "unknownState",
	SCPQrSelectFieldsState:    "selectFieldsState",
	SCPQrSelectFieldState:     "selectFieldState",
	SCPQrSelectExprFieldState: "selectExprFieldState",
	// 查询表
	SCPQrSelectFromState:       "selectFromState",
	SCPQrSelectWhereCondState:  "selectWhereCondState",
	SCPQrSelectJoinTableState:  "selectJoinTableState",
	SCPQrSelectJoinCondState:   "selectJoinCondState",
	SCPQrSelectJoinItemState:   "selectJoinItemState",
	SCPQrSelectGroupFieldState: "selectGroupFieldState",
	SCPQrSelectHavingCondState: "selectHavingCondState",
	SCPQrSelectOrderFieldState: "selectOrderFieldState",
	// 生成参数名称
	SCPQrSelectParamNameState: "selectParamNameState",
	SCPQrSelectFunParamState:  "selectFunParamState",
	// ====== 添加、更新查卷状态 ==== //

	// 创建插入表名状态
	SCPBuildInsertTableState: "buildInsertTableState",
	// 创建插入字段名状态
	SCPBuildInsertFieldNameState: "buildInsertFieldNameState",
	// 创建插入字段值状态
	SCPBuildInsertFieldValueState: "buildInsertFieldValueState",
	// 创建从其他表源插入的from部份脚本状态
	SCPBuildInsertFromState: "buildInsertFromState",
	// 创建更新表的表名状态
	SCPBuildUpdateTableState: "buildUpdateTableState",
	// 创建联合更新的表名
	SCPBuildUpdateTableWithFromState: "buildUpdateTableWithFromState",

	// 创建更新字段状态
	SCPBuildUpdateFieldState: "buildUpdateFieldState",
	// 创建联合被更新的字段名
	SCPBuildUpdateFieldWithFromState: "buildUpdateFieldWithFromState",
	// 创建更新字段值表达式使用：表别名.字段名
	SCPBuildUpdateFieldValueState:     "buildUpdateFieldValueState",
	SCPBuildUpdateFieldValueFromState: "updateFieldValueFromState",
	// 创建关联更新join部份表名状态
	SCPBuildUpdateJoinTableState: "buildJoinTableState",
	// 创建联合更新关联关系表达式状态
	SCPBuildUpdateJoinOnState: "buildUpdateJoinOnState",
	// 创建更新条件时的状态
	SCPBuildUpdateWhereState: "buildUpdateWhereState",
	// 创建联合更新时的条件表达式状态
	SCPBuildUpdateWhereFromState: "buildUpdateWhereFromState",

	// 创建被删除表名状态
	SCPBuildDeleteTableState: "buildDeleteTableState",
	// 删除表条件表达式状态
	SCPBuildDeleteWhereState: "buildDeleteWhereState",
}

func (scs SCPState) String() string {
	if s, ok := scpStateNames[scs]; ok {
		return s
	}
	return "unknownState"
}
