package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var MySqlDSN string
var Redis struct {
	Host     string
	Password string
	DB       int
}
var Swagger struct {
	Enable bool
	Path   string
}
var App struct {
	Addr    string
	Mode    string
	BaseURL string
}
var Log struct {
	Info  string `json:"info"`
	Route string `json:"route"`
	Error string `json:"error"`
}

func readMySqlConfig(viper *viper.Viper) error {
	conn := viper.GetStringMapString("mysql")

	user := conn["user"]
	password := conn["password"]
	host := conn["host"]
	dbname := conn["dbname"]

	if user == "" || password == "" || host == "" || dbname == "" {
		return fmt.Errorf("short of user, password, host, dbname parameters")
	}

	options := viper.GetStringSlice("mysql.options")
	MySqlDSN = fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", user, password, host, dbname, strings.Join(options, "&"))
	return nil
}

func readRedisConfig(viper *viper.Viper) error {
	return viper.UnmarshalKey("redis", &Redis)
}

func readSwaggerConfig(viper *viper.Viper) error {
	Swagger.Enable = true
	Swagger.Path = "/_/_swagger/*"
	return viper.UnmarshalKey("swagger", &Swagger)
}

func readAppConfig(viper *viper.Viper) error {
	App.Addr = "8080"
	App.Mode = "release"
	App.BaseURL = "/api/v1"
	return viper.UnmarshalKey("app", &App)
}

func readLogConfig(viper *viper.Viper) error {
	Log.Info = "stdout"
	Log.Route = "stdout"
	return viper.UnmarshalKey("log", &Log)
}

func init() {
	viper := viper.New()

	//viper.AddConfigPath("/etc/singlishwords/web/")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("/usr/local/etc/singlishwords/web/")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := readMySqlConfig(viper)
	if err != nil {
		panic("Config of mysql is wrong: " + err.Error())
	}
	err = readRedisConfig(viper)
	if err != nil {
		panic("Config of redis is wrong: " + err.Error())
	}
	err = readSwaggerConfig(viper)
	if err != nil {
		panic("Config of swagger is wrong: " + err.Error())
	}
	err = readAppConfig(viper)
	if err != nil {
		panic("Config of app is wrong: " + err.Error())
	}
	err = readLogConfig(viper)
	if err != nil {
		panic("Config of log is wrong: " + err.Error())
	}
}
