package utils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns hash value of a string
func HashPassword(password string) (hashedPassword string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPassword = string(hash)
	return
}

// ComparePassword returns check password result
func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

// ShuLogin login with services.shu.edu.cn
func ShuLogin(username string, password string) (result bool, err error) {
	authURL := "http://services.shu.edu.cn/Login.aspx"
	client := &http.Client{}
	// use cookie jar to save cookies during the request session
	jar, err := cookiejar.New(nil)
	if err != nil {
		return
	}
	client.Jar = jar
	resp, err := client.Get(authURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// <input type="hidden" name="__VIEWSTATE" value="dDwtMTIwMjUxOTIxNDs7PsH7y+VEVR/6VfZ5PNSi21UwSxxy" />
	// Get the value of __VIEWSTATE from html
	reg := regexp.MustCompile(`<input type="hidden" name="__VIEWSTATE" value="(.*?)" />`)
	// if cannot find asp viewstate from html, means school server is down
	if len(reg.FindStringSubmatch(string(body))) != 2 {
		err = errors.New("OauthUnavailable")
		return
	}
	viewState := reg.FindStringSubmatch(string(body))[1]
	// compose form data
	data := url.Values{
		"txtUserName":     {username},
		"txtPassword":     {password},
		"btnOk":           {"提交(Submit)"},
		"__EVENTTARGET":   {""},
		"__EVENTARGUMENT": {""},
		"__VIEWSTATE":     {viewState},
	}
	req, err := http.NewRequest("POST", authURL, strings.NewReader(data.Encode()))
	if err != nil {
		return
	}
	// mock chrome browser headers
	headers := map[string]string{
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"Accept-Language":           "zh-CN,zh;q=0.9",
		"Cache-Control":             "max-age=0",
		"Content-Type":              "application/x-www-form-urlencoded",
		"Host":                      "services.shu.edu.cn",
		"Origin":                    "http://services.shu.edu.cn",
		"Proxy-Connection":          "keep-alive",
		"Referer":                   "http://services.shu.edu.cn/Login.aspx",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 Safari/537.36",
	}
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	// send auth request
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	body, _ = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	// if response contains string "修改密码", then password is correct
	result = (strings.Contains(string(body), "修改密码"))
	return
}
