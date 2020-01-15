package configs

import (
	"database/sql/driver"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"time"
	"unicode"

	yaml "gopkg.in/yaml.v2"
)

// Config struct for prasing app configs
type Config struct {
	App struct {
		Name        string   `yaml:"Name"`
		Domain      string   `yaml:"Domain"`
		CORS        []string `yaml:"CORS"`
		AllowCookie bool     `yaml:"AllowCookie"`
	} `yaml:"App"`
	Log struct {
		Level string `yaml:"Level"`
	} `yaml:"Log"`
	Listener struct {
		Server string `yaml:"Server"`
		Port   int    `yaml:"Port"`
	} `yaml:"Listener"`
	JWT struct {
		ExpireHours int    `yaml:"ExpireHours"`
		Debug       bool   `yaml:"Debug"`
		PrivateKey  string `yaml:"PrivateKey"`
		PublicKey   string `yaml:"PublicKey"`
	} `yaml:"JWT"`
	Postgresql struct {
		Server   string `yaml:"Server"`
		Port     int    `yaml:"Port"`
		User     string `yaml:"User"`
		Password string `yaml:"Password"`
		Database string `yaml:"Database"`
	} `yaml:"Postgresql"`
	File struct {
		Dir string `yaml:"Dir"`
	} `yaml:"File"`
	Email struct {
		Port     int    `yaml:"Port"`
		Account  string `yaml:"Account"`
		Password string `yaml:"Password"`
	} `yaml:"Email"`
}

// ReadConfig read configs from a configuration file
func (conf *Config) ReadConfig() *Config {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	env := os.Getenv(EnvName)
	if env == "" {
		env = EnvDevelopment
	}
	confFileName := fmt.Sprintf("app.%s.yml", env)
	confFilePath := path.Join(workPath, "configs", confFileName)
	yamlFile, err := ioutil.ReadFile(confFilePath)
	if err != nil {
		panic(err)
	}
	err = yaml.UnmarshalStrict(yamlFile, conf)
	if err != nil {
		panic(err)
	}
	return conf
}

// GetLogger : stdout for dev environment and file for prod environment
func GetLogger(prefix string) io.Writer {
	if env := os.Getenv(EnvName); env == EnvDevelopment {
		return os.Stdout
	}
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filePath := path.Join(workPath, "logs", todayFileName(prefix))
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

// GetPostgresLogger returns logger for Postgresql Database
func GetPostgresLogger() SQLLogger {
	return SQLLogger{log.New(GetLogger("sql"), "\033[0;32m[SQLS]\033[0m ", 0)}
}

// LogWriter provides interface for sql logger
type LogWriter interface {
	Println(v ...interface{})
}

// SQLLogger is the logger for ORM library
type SQLLogger struct {
	LogWriter
}

// Print custom the logger for sql output
func (logger SQLLogger) Print(values ...interface{}) {
	if len(values) > 1 {
		var (
			sql             string
			formattedValues []string
		)
		level := values[0]
		currentTime := time.Now().Format("2006/01/02 15:04")
		messages := []interface{}{currentTime}
		if level == "sql" {
			messages = append(messages, fmt.Sprintf("%.2fms ", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0))
			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				if indirectValue.IsValid() {
					value = indirectValue.Interface()
					if t, ok := value.(time.Time); ok {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006/01/02 15:04:05")))
					} else if b, ok := value.([]byte); ok {
						if str := string(b); isPrintable(str) {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
						} else {
							formattedValues = append(formattedValues, "'<binary>'")
						}
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							formattedValues = append(formattedValues, "NULL")
						}
					} else {
						switch value.(type) {
						case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
							formattedValues = append(formattedValues, fmt.Sprintf("%v", value))
						default:
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						}
					}
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			}
			if regexp.MustCompile(`\$\d+`).MatchString(values[3].(string)) {
				sql = values[3].(string)
				for index, value := range formattedValues {
					placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
					sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
				}
			} else {
				formattedValuesLength := len(formattedValues)
				for index, value := range regexp.MustCompile(`\?`).Split(values[3].(string), -1) {
					sql += value
					if index < formattedValuesLength {
						sql += formattedValues[index]
					}
				}
			}
			messages = append(messages, sql)
			messages = append(messages, fmt.Sprintf("[%v] ", strconv.FormatInt(values[5].(int64), 10)+" rows"))
		} else {
			messages = append(messages, "\033[31;1m")
			messages = append(messages, values[2:]...)
			messages = append(messages, "\033[0m")
		}
		logger.Println(messages...)
	}
}

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

func todayFileName(logName string) string {
	return logName + "." + time.Now().Format("2006-01-02") + ".log"
}
