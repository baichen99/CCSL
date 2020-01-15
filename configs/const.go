package configs

// AppConf is iris application config file
const AppConf string = "./configs/iris.yml"

// EnvName is the environment variable name
const (
	EnvName                     = "CCSL_ENV"
	EnvDevelopment              = "dev"
	EnvProduction               = "prod"
	RoleSuperUser               = "super"
	RoleAdminUser               = "admin"
	RoleDatabaseUser            = "user"
	RoleLearningPlatformStudent = "learner"
	RoleLearningPlatformTeacher = "teacher"
)

var c Config

// Conf reads configs from yaml file according to the environment variable
var Conf = c.ReadConfig()
