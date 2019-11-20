package utils

import (
	"reflect"
	"regexp"
	"runtime"

	"github.com/kataras/iris"
)

// LogInfo logs information about what is happening on the server
func LogInfo(context iris.Context, info string) {
	_, fileName, lineNumber, ok := runtime.Caller(1)
	if !ok {
		fileName = "UNKNOWN"
		lineNumber = 0
	}
	context.Application().Logger().Infof("[%s:%d] %s - ", fileName, lineNumber, info)
}

// MakeUpdateData returns a map of update data model
func MakeUpdateData(dataStruct interface{}) map[string]interface{} {
	// Using this helper function to generate updated data map
	// We don't use model struct to update it because gorm library cannot update a structed model with default "zero" value
	// This means you cannot update a string to empty string nor update a number to 0
	// If we use map[string]interface{} to update it, we can set the fields that we don't need update to null pointer
	// Then we can update a value to "zero" value
	t := reflect.TypeOf(dataStruct)
	v := reflect.ValueOf(dataStruct)
	updateData := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			key := t.Field(i).Name
			value := v.Field(i)
			if !value.IsNil() {
				// Elem() returns the value that the interface v contains or that the pointer v points to.
				updateData[key] = reflect.ValueOf(value.Interface()).Elem()
			}
		}
	}
	return updateData
}

// IsShuUser returns if given username is SHU account
func IsShuUser(username string) bool {
	r, _ := regexp.Compile("^[0-9]{8}$")
	return r.MatchString(username)
}
