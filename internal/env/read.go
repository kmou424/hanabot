package env

import (
	"github.com/gookit/goutil/strutil"
	"os"
)

func readString(env string) (envString string) {
	envString = os.Getenv(env)
	if envString == `""` {
		envString = ""
	}
	return
}

func readInt64(env string) (envInt64 int64) {
	src := readString(env)
	envInt64 = strutil.Int64(src)
	return
}

func readStringSliceWithTrim(env string, trimFunc func(string) string, split ...string) (envSlice []string) {
	src := readString(env)
	splitChar := ","
	if len(split) > 0 {
		splitChar = split[0]
	}
	envSlice = strutil.Split(src, splitChar)
	if trimFunc != nil {
		for i := len(envSlice) - 1; i >= 0; i-- {
			envSlice[i] = trimFunc(envSlice[i])
			if envSlice[i] == "" {
				envSlice = append(envSlice[:i], envSlice[i+1:]...)
			}
		}
	}
	return
}
