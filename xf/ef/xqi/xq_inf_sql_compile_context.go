package xqi

import "github.com/go-xe2/x/xf/ef/xdriveri"

// sql参数化查询、预处理参数收集器
type SqlCompileContext interface {
	// 添加参数到当前上下文
	Add(val ...xdriveri.SqlParam) SqlCompileContext
	// 添加参数到编译上下文件
	AddParam(name string, val interface{}) xdriveri.SqlParam
	// 所有已被收到的实参
	Params() []xdriveri.SqlParam
	// 获取指定序号的参数
	ParamByIndex(index int) xdriveri.SqlParam
	// 删除指定序号的参数
	DeleteParam(index int) xdriveri.SqlParam
	// 根据参数名删除参数
	DeleteParamByName(name string) xdriveri.SqlParam
	// 清空实参
	ClearParams() SqlCompileContext
	// 清空参数、表及状态
	Clear() SqlCompileContext
	// 拷贝实参
	Copy(src SqlCompileContext) SqlCompileContext
	// 参数个数
	ParamCount() int
	// 编译状态入栈
	PushState(state SCPState) SqlCompileContext
	// 编辑状态出栈
	PopState() SCPState
	// 获取最后状态
	State() SCPState
	// 条件表达式左则表达式入规模
	PushLExp(exp SqlCompiler) SqlCompileContext
	// 左则表达式出栈
	PopLExp() SqlCompiler
	// 最后出栈的表达式
	LExp() SqlCompiler
	// 生成参数名称
	MakeParamId() string
	// 生成标识Id
	MakeIndentId() string
	AssignTables(src SqlTables)
	UseTable(table SqlTable)
	Tables() SqlTables
	SetDatabase(db Database)
	Database() Database
	TablePrefix() string
}
