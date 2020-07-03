package comm

type (
	ArrForEachFunc func(index int, item interface{})
	ArrMapFunc     func(index int, item interface{}) interface{}
	ArrSearchFunc  func(index int, item interface{}) bool
	// a < b时 < 0, a > b 时 > 0, a == b时, = 0
	ArrSortCompareFunc func(aIndex int, bIndex int) int

	StrArrForEachFunc func(index int, item string)
	StrArrMapFunc     func(index int, item string) string
	StrArrSearchFunc  func(index int, item string) bool

	IntArrForEachFunc func(index int, item int)
	IntArrMapFunc     func(index int, item int) int
	IntArrSearchFunc  func(index int, item int) bool

	Int8ArrForEachFunc func(index int, item int8)
	Int8ArrMapFunc     func(index int, item int8) int8
	Int8ArrSearchFunc  func(index int, item int8) bool

	Int16ArrForEachFunc func(index int, item int16)
	Int16ArrMapFunc     func(index int, item int16) int16
	Int16ArrSearchFunc  func(index int, item int16) bool

	Int32ArrForEachFunc func(index int, item int32)
	Int32ArrMapFunc     func(index int, item int32) int32
	Int32ArrSearchFunc  func(index int, item int32) bool

	Int64ArrForEachFunc func(index int, item int64)
	Int64ArrMapFunc     func(index int, item int64) int64
	Int64ArrSearchFunc  func(index int, item int64) bool

	UintArrForEachFunc func(index int, item uint)
	UintArrMapFunc     func(index int, item uint) uint
	UintArrSearchFunc  func(index int, item uint) bool

	Uint8ArrForEachFunc func(index int, item uint8)
	Uint8ArrMapFunc     func(index int, item uint8) uint8
	Uint8ArrSearchFunc  func(index int, item uint8) bool

	Uint16ArrForEachFunc func(index int, item uint16)
	Uint16ArrMapFunc     func(index int, item uint16) uint16
	Uint16ArrSearchFunc  func(index int, item uint16) bool

	Uint32ArrForEachFunc func(index int, item uint32)
	Uint32ArrMapFunc     func(index int, item uint32) uint32
	Uint32ArrSearchFunc  func(index int, item uint32) bool

	Uint64ArrForEachFunc func(index int, item uint64)
	Uint64ArrMapFunc     func(index int, item uint64) uint64
	Uint64ArrSearchFunc  func(index int, item uint64) bool

	FloatArrForEachFunc func(index int, item float32)
	FloatArrMapFunc     func(index int, item float32) float32
	FloatArrSearchFunc  func(index int, item float32) bool

	Float64ArrForEachFunc func(index int, item float64)
	Float64ArrMapFunc     func(index int, item float64) float64
	Float64ArrSearchFunc  func(index int, item float64) bool

	BoolArrForEachFunc func(index int, item bool)
	BoolArrMapFunc     func(index int, item bool) bool
)
