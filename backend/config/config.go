package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Mysql MysqlConfig
}
type MysqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Charset  string
}

func Mysql() MysqlConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(viper.GetString("mysql.host"))
	return MysqlConfig{
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetInt("mysql.port"),
		User:     viper.GetString("mysql.user"),
		Password: viper.GetString("mysql.password"),
		Database: viper.GetString("mysql.database"),
		Charset:  viper.GetString("mysql.charset"),
	}
}
func Redis() {

}
