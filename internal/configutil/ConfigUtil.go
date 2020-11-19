package configutil

import (
	"github.com/spf13/viper"
)

var _config = make(map[string]*viper.Viper)

// LoadConfigs 加载多个配置文件
// configPath 配置文件目录
// names 文件名列表
func LoadConfigs(configPath string, names []string) error {
	for _, name := range names {
		err := LoadConfig(configPath, name)
		if err != nil {
			return err
		}
	}
	return nil
}

// LoadConfig 加载单个配置文件
// configPath 配置文件目录
// name 配置名称
func LoadConfig(configPath string, name string) error {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(name)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	_config[name] = v
	return nil
}

// GetConfigString 获取配置值
// name 配置的名称
// key 配置项的key
func GetConfigString(name string, key string) string {
	if config, ok := _config[name]; ok {
		return config.GetString(key)
	}
	return ""
}

func GetConfigBool(name string, key string) bool {
	if config, ok := _config[name]; ok {
		return config.GetBool(key)
	}
	return false
}

func GetConfigStringMapString(name string, key string) map[string]string {
	if config, ok := _config[name]; ok {
		return config.GetStringMapString(key)
	}
	return nil
}
