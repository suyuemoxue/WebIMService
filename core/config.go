package core

import (
	"WebIM/global"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// InitConfig 读取yaml文件配置
func InitConfig() {
	// 老方法
	//const ConfigFile = "config.yaml"         //声明一个常量保存文件名
	//c := &config.Config{}                    //需要把配置文件放进结构体里面，所以声明一个配置文件的结构体
	//yamlConf, err := os.ReadFile(ConfigFile) //读取yaml配置文件
	//if err != nil {
	//	panic(fmt.Errorf("get config error: %s", err))
	//}
	//err = yaml.Unmarshal(yamlConf, c) //将读取到的配置文件解析到上面声明的结构体里面
	//if err != nil {
	//	log.Fatalf("config Init Unmarshal: %v", err)
	//}
	//log.Println("config yamlFile load Init success.")
	//global.Config = c //将保存配置文件的结构体变量赋值给全局变量

	// 新方法
	// 初始化 Viper
	viper.SetConfigFile("config.yaml") // 设置配置文件名及路径
	viper.SetConfigType("yaml")        // 设置配置文件类型
	err := viper.ReadInConfig()        // 读取配置文件
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}
	// 将配置文件中的值绑定到 Config 结构体变量
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %w", err))
	}
	log.Println("config yamlFile load Init success.")
}
