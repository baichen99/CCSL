package configs

// AppConf is iris application config file
const AppConf string = "./configs/iris.yml"

// EnvName is the environment variable name
const (
	EnvName           = "CCSL_ENV"
	EnvDevelopment    = "dev"
	EnvProduction     = "prod"
	RoleSuperUser     = "super"
	RoleAdminUser     = "admin"
	RoleDatabaseUser  = "dbuser"
	RoleStudent       = "student"
	RoleTeacher       = "teacher"
	UserStateActive   = "active"
	UserStateInactive = "inactive"
)

var c Config

// Conf reads configs from yaml file according to the environment variable
var Conf = c.ReadConfig()
