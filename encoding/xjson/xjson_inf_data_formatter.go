package xjson

// json序列化接口
type JsonDataFormatter interface {
	Format(v interface{}) interface{}
}
