package common

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const DefaultPagingCount = 10

// ex: "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
type Database struct {
	Dialect   string `yaml:"dialect"`
	Host      string `yaml:"host"`
	Port      uint32 `yaml:"port"`
	Name      string `yaml:"name"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Option    string `yaml:"option"`
	MaxIdle   int    `yaml:"maxIdle"`
	MaxActive int    `yaml:"maxActive"`
}

type Server struct {
	Host         string `yaml:"host"`
	Port         uint32 `yaml:"port"`
	WriteTimeout int64  `yaml:"writeTimeout"`
	ReadTimeout  int64  `yaml:"readTimeout"`
}

type RabbitMQ struct {
	Host     string `yaml:"host"`
	Port     uint32 `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Queue    string `yaml:"queue"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     uint32 `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type Config struct {
	Database              Database `yaml:"database"`
	Server                Server   `yaml:"server"`
	RabbitMQ              RabbitMQ `yaml:"rabbitMQ"`
	Redis                 Redis    `yaml:"redis"`
	FirebaseAdminJsonPath string   `yaml:"firebaseAdminJsonPath"`
	SQLMigrateSourcePath  string   `yaml:"sqlMigrateSourcePath"`
	AccessTokenTtl        int64    `yaml:"accessTokenTtl"`
}

var config *Config = nil

func GetConfig() *Config {
	logger := GetLogger()
	if config == nil {
		logger.Fatal("config not init.")
	}
	return config
}

func InitConfig(yamlFile string) *Config {
	fileContent, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		panic(fmt.Sprintf("error: read fail - %v\n", yamlFile))
	}

	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		panic(fmt.Sprintf("error: parse yaml fail. - %v", err))
	}
	return config
}
