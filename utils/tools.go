package utils

import (
	"ccsl/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	"regexp"
	"runtime"

	"github.com/kataras/iris/v12"
)

// LogInfo logs information about what is happening on the server
func LogInfo(context iris.Context, info string) {
	_, fileName, lineNumber, ok := runtime.Caller(1)
	if !ok {
		fileName = "UNKNOWN"
		lineNumber = 0
	}
	context.Application().Logger().Infof("[%s:%d] %s ", fileName, lineNumber, info)
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

// IsPublicIP checks if given ip is public ip address
func IsPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

// GetIPInfo returns info of given ip address
func GetIPInfo(ipAddress string) (info models.LoginHistory, err error) {
	ip := net.ParseIP(ipAddress)
	if IsPublicIP(ip) {
		errInfo := models.LoginHistory{
			IP:          ipAddress,
			Country:     "Unkown Area",
			CountryCode: "UNKNOWN",
			Status:      "unknown",
		}
		url := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,message,country,countryCode,region,regionName,city,district,lat,lon,timezone,isp,org,query", ipAddress)
		var req *http.Request
		if req, err = http.NewRequest("GET", url, nil); err != nil {
			info = errInfo
			return
		}
		var res *http.Response
		if res, err = http.DefaultClient.Do(req); err != nil {
			info = errInfo
			return
		}
		defer res.Body.Close()
		var body []byte
		if body, err = ioutil.ReadAll(res.Body); err != nil {
			info = errInfo
			return
		}
		if err = json.Unmarshal([]byte(body), &info); err != nil {
			info = errInfo
			return
		}
		if info.Status != "success" {
			info = errInfo
			return
		}
		info.IP = ipAddress
		return
	}
	info = models.LoginHistory{
		IP:          ipAddress,
		Country:     "Internal IP Address",
		CountryCode: "INTERNAL",
		Status:      "internal",
	}
	return
}
