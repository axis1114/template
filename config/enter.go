package config

import (
	"strconv"
)

type Config struct {
	Mysql  Mysql  `mapstructure:"mysql"`
	Redis  Redis  `mapstructure:"redis"`
	Log    Log    `mapstructure:"log"`
	System System `mapstructure:"system"`
	Email  Email  `mapstructure:"email"`
	Es     Es     `mapstructure:"es"`
	Jwt    Jwt    `mapstructure:"jwt"`
}

func (m Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func (e Es) Dsn() string {
	return "http://" + e.Host + ":" + strconv.Itoa(e.Port)
}
