package config

import (
	"strconv"

	"github.com/spf13/viper"
)

type EnvConfig struct {
	v *viper.Viper
}

const (
	TRUE  = "true"
	FALSE = "false"
)

func newEnvConfig() *EnvConfig {
	var e EnvConfig
	e.v = viper.New()
	e.v.BindEnv(DB_USER)
	e.v.BindEnv(DB_DRIVER)
	e.v.BindEnv(DB_PASSWORD)
	e.v.BindEnv(DB_HOST)
	e.v.BindEnv(DB_PORT)
	e.v.BindEnv(DB_DATABASE)
	e.v.BindEnv(MIGRATION)
	return &e
}

func (e *EnvConfig) getString(key string) string {
	str := e.v.Get(key)
	if str != nil {
		return str.(string)
	}
	return ""
}

func (e *EnvConfig) getBool(key string) bool {
	return e.v.Get(key) == TRUE
}

func (e *EnvConfig) getInt(key string) int {
	value := e.getString(key)
	if value == "" {
		return 0
	}
	i, err := strconv.ParseInt(value, 0, 32)
	if err != nil {
		return int(i)
	}
	return 0
}
