package hytrix

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/spf13/viper"
)

//NewHytrixCommand new hytrix command
func NewHytrixCommand() hystrix.CommandConfig {
	return hystrix.CommandConfig{
		Timeout:                viper.GetInt("hytrix.hytrix_timeout"),
		MaxConcurrentRequests:  viper.GetInt("hytrix.hytrix_max_concurrent"),
		RequestVolumeThreshold: viper.GetInt("hytrix.hytrix_volumn_threshold"),
		SleepWindow:            viper.GetInt("hytrix.hytrix_sleep_window"),
		ErrorPercentThreshold:  viper.GetInt("hytrix.hytrix_error_percent_threshold"),
	}
}
