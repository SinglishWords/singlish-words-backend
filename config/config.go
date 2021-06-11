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

func readMySqlConfig(viper *viper.Viper) error {
	conn := viper.GetStringMapString("mysql")

	user := conn["user"]
	password := conn["password"]
	host := conn["host"]
	dbname := conn["dbname"]

	fmt.Printf("%s:%s@tcp(%s)/%s", user, password, host, dbname)

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

func init() {
	viper := viper.New()

	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := readMySqlConfig(viper)
	if err != nil {
		panic("Config of mysql is wrong." + err.Error())
	}
	err = readRedisConfig(viper)
	if err != nil {
		panic("Config of redis is wrong." + err.Error())
	}
}
