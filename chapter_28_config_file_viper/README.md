# GO 配置文件读取

## viper 库读取配置文件
github 地址: https://github.com/spf13/viper

```go
package main

import (
    "fmt"
    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
    "path"
)

type MysqlConfig struct {
    Endpoint string `mapstructure:"endpoint"`
}

type ServerConfig struct {
    Name        string      `mapstructure:"name"`
    Port        int         `mapstructure:"port"`
    MysqlConfig MysqlConfig `mapstructure:"mysql"`
}

// GetBoolEnv 通过viper获取布尔类型的环境变量
func GetBoolEnv(env string) bool {
    viper.AutomaticEnv()
    return viper.GetBool(env)
}

// GetConfigFile 更加 WEB_DEBUG 环境变量获取配置文件
func GetConfigFile() string {
    configPath := "02_viper_env_and_dynamic_loading/configs"
    configFileName := "config_pro.yaml"
    debug := GetBoolEnv("WEB_DEBUG")
    if debug {
        configFileName = "config_debug.yaml"
    }
    return path.Join(configPath, configFileName)
}

func viperReadYaml() {
    // 获取配置文件路径
    configName := GetConfigFile()
    fmt.Printf("config name: %s\n", configName)

    // 加载配置文件
    v := viper.New()
    v.SetConfigFile(configName)
    if err := v.ReadInConfig(); err != nil {
        panic(err)
    }

    // 配置绑定到结构体上
    conf := ServerConfig{}
    if err := v.Unmarshal(&conf); err != nil {
        panic(err)
    }
    fmt.Printf("CONF: %+v\n", conf)

    // 动态加载配置文件
    v.WatchConfig()
    v.OnConfigChange(func(e fsnotify.Event) {
        fmt.Printf("the config file [%s] has been modified.\n", e.Name)
        _ = v.ReadInConfig()
        _ = v.Unmarshal(&conf)
        fmt.Printf("New CONF: %+v\n", conf)
    })
    ch := make(chan bool)
    <-ch
}

func main() {
    viperReadYaml()
}

/* 修改环境变量，代码使用不同的配置文件
(Golang) ➜  chapter_28_config_file_viper git:(master) ✗ export WEB_DEBUG=true
(Golang) ➜  chapter_28_config_file_viper git:(master) ✗ echo $WEB_DEBUG
true
(Golang) ➜  chapter_28_config_file_viper git:(master) ✗ go run 02_viper_env_and_dynamic_loading/main.go
config name: 02_viper_env_and_dynamic_loading/configs/config_debug.yaml
conf: {Name:web Port:8888 MysqlConfig:{Endpoint:127.168.0.1:3306}}
(Golang) ➜  chapter_28_config_file_viper git:(master) ✗ WEB_DEBUG=false
(Golang) ➜  chapter_28_config_file_viper git:(master) ✗ echo $WEB_DEBUG
false
(Golang) ➜  chapter_28_config_file_viper git:(master) ✗ go run 02_viper_env_and_dynamic_loading/main.go
config name: 02_viper_env_and_dynamic_loading/configs/config_pro.yaml
conf: {Name:web Port:9000 MysqlConfig:{Endpoint:127.168.0.1:13306}}
*/


/* 动态加载
(Golang) ➜  chapter_28_config_file_viper git:(master) ✗ go run 02_viper_env_and_dynamic_loading/main.go
config name: 02_viper_env_and_dynamic_loading/configs/config_debug.yaml
CONF: {Name:web Port:8888 MysqlConfig:{Endpoint:127.168.0.1:3306}}
the config file [02_viper_env_and_dynamic_loading/configs/config_debug.yaml] has been modified.
New CONF: {Name:web Port:8888 MysqlConfig:{Endpoint:127.168.0.1:3306}}
the config file [02_viper_env_and_dynamic_loading/configs/config_debug.yaml] has been modified.
New CONF: {Name:web Port:8888 MysqlConfig:{Endpoint:127.168.0.1:13306}}
the config file [02_viper_env_and_dynamic_loading/configs/config_debug.yaml] has been modified.
New CONF: {Name:web Port:7888 MysqlConfig:{Endpoint:127.168.0.1:13306}}
*/
```


