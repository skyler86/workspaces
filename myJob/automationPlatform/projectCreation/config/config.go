package config

import (
	"github.com/spf13/viper"
	"log"
)

func YamlInfo() *viper.Viper {
	//创建viper实例
	config := viper.New()

	//指定配置文件的信息
	config.AddConfigPath("./config/") // 设置配置文件的路径
	config.SetConfigName("config")    // 读取配置文件文件名
	config.SetConfigType("yaml")      // 读取配置文件格式

	//判断读取配置文件是否有误
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("找不到配置文件..")
		} else {
			log.Println("配置文件出错..")
		}
		log.Fatal(err)
	}
	return config
}

func EnVarInfo() {
	//env_url := os.Getenv("HARBORURL")
	//fmt.Println(env_url)
	//env_username := os.Getenv("HARBORUSERNAME")
	//env_password := os.Getenv("HARBORPASSWORD")

	////自动从环境变量读取匹配的参数
	//config.AutomaticEnv()
	////读取环境变量前缀,有这个前缀的环境变量才会读取连接
	//viper.SetEnvPrefix("HARBORURL")
	//viper.SetEnvPrefix("HARBORADDRESS")
	//viper.SetEnvPrefix("HARBORPASSWORD")
}
