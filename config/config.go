package config

import (
	"gopkg.in/ini.v1"
	"log"
)

type ConfigList struct {
	SQLDriver  string
	DbName     string
	DbUser     string
	DbPassword string
	DbPort     string
	DBHost     string
}

var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		SQLDriver:  cfg.Section("db").Key("driver").String(),
		DbName:     cfg.Section("db").Key("name").String(),
		DbUser:     cfg.Section("db").Key("user").String(),
		DbPassword: cfg.Section("db").Key("password").String(),
		DbPort:     cfg.Section("db").Key("port").MustString("5432"),
		DBHost:     cfg.Section("db").Key("host").String(),
	}
}
