package utils

import (
	"crypto/tls"
	"ccsl/configs"
	"fmt"
	"io/ioutil"
	"net"
	"net/smtp"
	"os"
	"strings"
)

// Dial creates a new SMTP client
func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		return nil, err
	}
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

// SendMailUsingTLS sends a email using TLS
func SendMailUsingTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) (err error) {
	c, err := Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

// SendEmail replace the {{param}} in template with given parameters
func SendEmail(address string, param string, template string, subject string) (err error) {
	host := configs.Conf.App.Domain
	port := configs.Conf.Email.Port
	email := configs.Conf.Email.Account
	password := configs.Conf.Email.Password
	header := make(map[string]string)
	header["From"] = configs.Conf.App.Name + "<" + email + ">"
	header["To"] = address
	header["Subject"] = subject
	header["Content-Type"] = "text/html; charset=UTF-8"
	workPath, err := os.Getwd()
	if err != nil {
		return err
	}
	htmlFile := workPath + "/assets/email/" + template + ".html"
	content, err := ioutil.ReadFile(htmlFile)
	if err != nil {
		return err
	}
	body := string(content)
	body = strings.Replace(body, "{{param}}", param, -1)
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	auth := smtp.PlainAuth(
		"",
		email,
		password,
		host,
	)
	err = SendMailUsingTLS(
		fmt.Sprintf("%s:%d", host, port),
		auth,
		email,
		[]string{address},
		[]byte(message),
	)
	if err != nil {
		return err
	}
	return nil
}
