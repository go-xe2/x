package xqi

// 实体编辑器
type EntityEditor interface {
	FieldEdit(field EntField, value FieldValue) FieldValue
}

// 实体插入数据监听器
type EntityInsertObserver interface {
	// 插入数据之前调用，返回false时，取消插入
	BeforeInsert(fields []FieldValue) bool
}

// 实体更新监听器
type EntityUpdateObserver interface {
	// 调更前调用，返回false时取消该次更新
	BeforeUpdate(fields []FieldValue) bool
}

type EntityDeleteObserver interface {
	// 删除之前调用，返回false取消该次删除
	BeforeDelete(where SqlCondition) bool
}

// 实体更新规则,如果实体实现该接口，ValidUpdate方法将先用该接口检查数据有效性，检查通过后才更新数据
type EntityUpdateRule interface {
	// 返回字段验证规则，规格格式为: [字段名@]规格名[#提示消息]
	UpdateRule(cate ...string) []string
}

// 实体插入数据规则检查,如果实体现实了该接口，ValidInsert方法将先调用该接口检查数据有效性，通过后才插入数据
type EntityInsertRule interface {
	// 返回字段验证规则，规格格式为: [字段名@]规格名[#提示消息]
	InsertRule(cate ...string) []string
}

// 实体更新规则,如果实体实现该接口，ValidUpdate方法将先用该接口检查数据有效性，检查通过后才更新数据
type EntityUpdateChecker interface {
	// 检查数据是否有效，返回error != nil 将取消该次更新
	CheckUpdateValid(values []FieldValue) ([]FieldValue, error)
}

// 实体插入数据规则检查,如果实体现实了该接口，ValidInsert方法将先调用该接口检查数据有效性，通过后才插入数据
type EntityInsertChecker interface {
	// 检查数据是否有效，返回error != nil 取消该次插入
	CheckInsertValid(values ...[]FieldValue) ([]FieldValue, error)
}

type BasicEntity interface {
	SqlTable
	// 关联的实体
	ForeignTables() SqlTables
	FieldByNameE(fieldName string) EntField
}

type EntityQuery interface {
	Select(fields ...SqlField) EntitySelect

	// 按指定规则字段返回查询结果
	SelectR(rule string) EntitySelect

	Open(fields ...SqlField) EntityQueryOpen

	// 只查询返回指定规则中的字段
	OpenR(rule string) EntityQueryOpen
	// 以实体自身方式返回数据查询，返回满足条件的记录行数，从实体自定字段读取数据
	OpenTop(count int, where SqlCondition, orders []SqlOrderField, fields ...SqlField) int
	// 获取一条数据，如果无指定条件的数据返回false,如果出错会panic
	OpenFirst(where SqlCondition, orders []SqlOrderField, fields ...SqlField) bool
	// 根据id获取一条数据，如果实体未设置主键或查询出错会panic,如果指定主键值不存在，返回false,其返回值从实体实例字段中读取
	OpenById(id interface{}, fields ...SqlField) bool

	// 检查数据是否存在
	Exists(where func(where SqlCondition, this interface{}) SqlCondition) (bool, error)
	// Exists的简化版本，出错时panic
	CheckExists(where ...SqlCondition) bool

	// 检查条件获取一条记录,简单查询方式
	First(where SqlCondition, orders []SqlOrderField, fields ...SqlField) (map[string]interface{}, error)
	// 查找数据，简单查询方式
	Find(where SqlCondition, orders []SqlOrderField, fields ...SqlField) ([]map[string]interface{}, error)
	// 根据主键ID获取一条map数据数据, 出错将会panic
	GetById(id interface{}, fields ...SqlField) map[string]interface{}
	// 根据条件获取指定条数的数据，出错将会panic
	GetTop(count int, where SqlCondition, orders []SqlOrderField, fields ...SqlField) []map[string]interface{}
}

type EntityInsertExecute interface {
	Execute() (int, error)
}

type EntityInsert interface {
	Execute() (int, error)
	From(table SqlTable) EntityInsertExecute
}

type EntityUpdateExecute interface {
	Execute() (int, error)
}

type EntityUpdate interface {
	Join(table SqlTable, on func(joinEnt SqlTable, leftTables SqlTables) SqlCondition) EntityUpdate
	LeftJoin(table SqlTable, on func(joinEnt SqlTable, leftTables SqlTables) SqlCondition) EntityUpdate
	RightJoin(table SqlTable, on func(joinEnt SqlTable, leftTables SqlTables) SqlCondition) EntityUpdate
	CrossJoin(table SqlTable, on func(joinEnt SqlTable, leftTables SqlTables) SqlCondition) EntityUpdate
	Where(where ...SqlCondition) EntityUpdateExecute
}

type EntityCUD interface {
	// 根据主键值删除
	DeleteByKey(keyVal interface{}) (int, error)
	// 根据条件删除
	Delete(where ...SqlCondition) (int, error)

	// 插入数据
	Insert(values ...FieldValue) EntityInsert
	// 使用map[string]interface数据源插入
	InsertByMap(maps map[string]interface{}) (int, error)
	// 分散式插入数据开始, 之后使用字字段的set设置值
	BeginInsert() bool
	// 提交插入数据
	CommitInsert() (int, error)
	// 更新
	Update(values ...FieldValue) EntityUpdate
	// 分散式更新开始, 使用字段的set相关方法设置值后更新
	BeginUpdate() bool
	// 提交更新
	CommitUpdate(where ...SqlCondition) (int, error)
	// 插入或更新,如果设置了主键则更新等于主键值的行，否则插入
	SaveMap(data map[string]interface{}) (int, error)
	// 插入或更新,如果设置了主键则更新等于主键值的行，否则插入
	Save(fields ...FieldValue) (int, error)

	// 在插入数据之前，先检查数据有效后才插入，实体需要实现EntityInsertChecker接口
	ValidInsert(values []FieldValue, validRule ...string) (int, error)
	// 在更新数据之前，先检查数据有效后才更新，实体需要实现EntityUpdateChecker接口
	ValidUpdate(values []FieldValue, where SqlCondition, validRule ...string) (int, error)

	// 使用参数数据源插入数据库，插入之前使用实体定义的插入规则过滤和检查数据有效性
	InsertFromParams(params map[string]interface{}) (int, error)

	// 使用参数数据源更新数据库，更新之前使用实体定义的更新规则过滤和检查数据有效性
	UpdateFromParams(params map[string]interface{}, where ...SqlCondition) (int, error)
}

type Entity interface {
	BasicEntity
	// 继承的父类
	Supper() Entity
	// 设置实体继承类实例, 供构架内部调用
	Implement(supper interface{})
	// 实体构造方法
	Constructor(attrs []XqAttribute, inherited ...interface{}) interface{}
	// 主键字段
	KeyField() EntField
	LastSql() string
	EntityQuery
	// 在特定数据库中进行操作
	Database(dbName string) Entity

	// 获取实体插入数据检查规则
	GetInsertRules(cate ...string) []string

	// 获取实体更新数据检查规则
	GetUpdateRules(cate ...string) []string

	// 使用更新规则过滤map参数
	// @param validParams 默认为false,为true时过滤并检查数据有效性，否则不检查
	FilterUpdateParams(params map[string]interface{}, validParams ...bool) (map[string]interface{}, error)

	// 使用插入规则过滤map参数
	// @param validParams 默认为false,为true时过滤并检查数据有效性，否则不检查
	FilterInsertParams(params map[string]interface{}, validParams ...bool) (map[string]interface{}, error)

	// 使用插入规则过滤map参数并转换成插入表达式中的值参数列表
	// @param validParams 默认为false,为true时过滤并检查数据有效性，否则不检查
	GetInsertValuesFromParams(params map[string]interface{}, validParams ...bool) ([]FieldValue, error)

	// 使用更新规则过滤map参数并转换成更新表达式中的值参数列表
	GetUpdateValuesFromParams(params map[string]interface{}, validParams ...bool) ([]FieldValue, error)
	FieldFormat(defineName string) func(old interface{}) interface{}
}
