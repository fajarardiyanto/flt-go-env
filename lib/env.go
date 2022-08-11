package lib

import (
	"os"
	"strconv"
	"time"
)

// EnvString returns the env string value if the key exists.
// Otherwise, returns initial value.
func EnvString(key, initial string) string {
	v, exists := os.LookupEnv(key)
	if !exists {
		return initial
	}
	return v
}

// EnvInt returns the env integer value if the key exists.
// Otherwise, returns initial value.
func EnvInt(key string, initial int) int {
	v := EnvString(key, "")
	if v == "" {
		return initial
	}

	n, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	return n
}

// EnvDuration returns the env duration value if the key exists.
// Otherwise, returns initial value.
func EnvDuration(key string, initial time.Duration) time.Duration {
	v := EnvString(key, "")
	if v == "" {
		return initial
	}

	d, err := time.ParseDuration(v)
	if err != nil {
		panic(err)
	}

	return d
}

// EnvBool returns the env boolean value if the key exists.
// Otherwise, returns initial value.
func EnvBool(key string, initial bool) bool {
	v := EnvString(key, "")
	if v == "" {
		return initial
	}

	b, err := strconv.ParseBool(v)
	if err != nil {
		panic(err)
	}

	return b
}
