package config

import (
	"grpc-gateway/common"
)

// type AppInfo struct {
// 	Addr string
// }
//
// type LogInfo struct {
// 	Dir       string
// 	Name      string
// 	KeepDay   int
// 	RateHours int
// }
//
// type SessionInfo struct {
// 	CookieName      string
// 	EnableSetCookie bool
// 	MaxLifeTime     int64
// 	CookieLifeTime  int
// 	Secure          bool
// 	ProviderType    string
// 	ProviderConfig  string
//
// 	NoNeedLogin   []string
// 	NoNeedRelease []string
// }

type Config struct {
	AppCnf struct {
		Addr string
	}
	LogCnf struct {
		Dir       string
		Name      string
		KeepDay   int
		RateHours int
	}
	SessionCnf struct {
		CookieName      string
		EnableSetCookie bool
		MaxLifeTime     int64
		CookieLifeTime  int
		Secure          bool
		ProviderType    string
		ProviderConfig  string

		NoNeedLogin   []string
		NoNeedRelease []string
	}
}

var cnf Config

func InitConfig() error {
	return common.UnmarshalTomlFromFile("config/config.toml", &cnf)
}

func GetConfig() *Config {
	return &cnf
}
