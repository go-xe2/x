package xcache

var cache = New()

func Set(key interface{}, value interface{}, duration interface{}) {
	cache.Set(key, value, duration)
}

func SetIfNotExist(key interface{}, value interface{}, duration interface{}) bool {
	return cache.SetIfNotExist(key, value, duration)
}

func Sets(data map[interface{}]interface{}, duration interface{}) {
	cache.Sets(data, duration)
}

func Get(key interface{}) interface{} {
	return cache.Get(key)
}

func GetOrSet(key interface{}, value interface{}, duration interface{}) interface{} {
	return cache.GetOrSet(key, value, duration)
}

func GetOrSetFunc(key interface{}, f func() interface{}, duration interface{}) interface{} {
	return cache.GetOrSetFunc(key, f, duration)
}

func GetOrSetFuncLock(key interface{}, f func() interface{}, duration interface{}) interface{} {
	return cache.GetOrSetFuncLock(key, f, duration)
}

func Contains(key interface{}) bool {
	return cache.Contains(key)
}

func Remove(key interface{}) interface{} {
	return cache.Remove(key)
}

func Removes(keys []interface{}) {
	cache.Removes(keys)
}

func Data() map[interface{}]interface{} {
	return cache.Data()
}

func Keys() []interface{} {
	return cache.Keys()
}

func KeyStrings() []string {
	return cache.KeyStrings()
}

func Values() []interface{} {
	return cache.Values()
}

func Size() int {
	return cache.Size()
}
