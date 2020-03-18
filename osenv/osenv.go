package osenv

import (
	"os"
	"strconv"
)

// String ...
func String(key string) string {
	return getEnv(key)
}

// Int64 ...
func Int64(key string) (int64, error) {
	i, err := strconv.ParseInt(getEnv(key), 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// Int32 ...
func Int32(key string) (int32, error) {
	i, err := strconv.ParseInt(getEnv(key), 10, 64)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

// Int ...
func Int(key string) (int, error) {
	i, err := strconv.Atoi(getEnv(key))
	if err != nil {
		return 0, err
	}
	return i, nil
}

// Bool ...
func Bool(key string) (bool, error) {
	b, err := strconv.ParseBool(getEnv(key))
	if err != nil {
		return false, err
	}
	return b, nil
}

func getEnv(key string) string {
	return os.Getenv(key)
}
