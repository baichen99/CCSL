package utils

import (
	"math/rand"
	"reflect"
	"regexp"
	"runtime"
	"time"

	"github.com/kataras/iris"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

// LogInfo logs information about what is happening on the server
func LogInfo(context iris.Context, info string) {
	// TODO: If this gets too slow, might want to have a setting in the config file to enable/disable
	//       caller information.
	_, fileName, lineNumber, ok := runtime.Caller(1)
	if !ok {
		fileName = "UNKNOWN"
		lineNumber = 0
	}
	context.Application().Logger().Infof("[%s:%d] %s - ", fileName, lineNumber, info)
}

// MakeUpdateData returns a map of update data model
func MakeUpdateData(dataStruct interface{}) map[string]interface{} {
	t := reflect.TypeOf(dataStruct)
	v := reflect.ValueOf(dataStruct)
	updateData := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			key := t.Field(i).Name
			value := v.Field(i)
			if !value.IsNil() {
				updateData[key] = reflect.ValueOf(value.Interface()).Elem()
			}
		}
	}
	return updateData
}

// RandomString returns given bytes of radom string
func RandomString(n int) string {
	seed := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i, cache, remain := n-1, seed.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = seed.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// IsShuUser returns if given username is SHU account
func IsShuUser(username string) bool {
	r, _ := regexp.Compile("^[0-9]{8}$")
	return r.MatchString(username)
}
