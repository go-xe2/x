package xenv

import "os"

func All() []string {
	return os.Environ()
}

func Get(key string, def ...string) string {
	v, ok := os.LookupEnv(key)
	if !ok && len(def) > 0 {
		return def[0]
	}
	return v
}

func Set(key, value string) error {
	return os.Setenv(key, value)
}

func Remove(key string) error {
	return os.Unsetenv(key)
}
