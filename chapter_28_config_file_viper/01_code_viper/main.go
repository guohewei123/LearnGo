package main

import (
    "fmt"
    "github.com/spf13/viper"
)

type MysqlConfig struct {
    Endpoint string `mapstructure:"endpoint"`
}

type ServerConfig struct {
    Name        string      `mapstructure:"name"`
    Port        int         `mapstructure:"port"`
    MysqlConfig MysqlConfig `mapstructure:"mysql"`
}

func viperReadYaml() {
    v := viper.New()
    v.SetConfigFile("01_code_viper/configs/config.yaml")
    if err := v.ReadInConfig(); err != nil {
        panic(err)
    }
    conf := ServerConfig{}
    if err := v.Unmarshal(&conf); err != nil {
        panic(err)
    }
    fmt.Printf("conf: %+v\n", conf)
}

func main() {
    viperReadYaml()
}
