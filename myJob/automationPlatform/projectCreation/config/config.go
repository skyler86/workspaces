package config

import (
	"fmt"
	//"github.com/spf13/viper"
	//"log"
	//"strings"
	"os"
)

//func YamlInfo() *viper.Viper {
//	//创建viper实例
//	config := viper.New()
//
//	//指定配置文件的信息
//	config.AddConfigPath("./config/") // 设置配置文件的路径
//	config.SetConfigName("config")    // 读取配置文件文件名
//	config.SetConfigType("yaml")      // 读取配置文件格式
//
//	//判断读取配置文件是否有误
//	if err := config.ReadInConfig(); err != nil {
//		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
//			log.Println("找不到配置文件..")
//		} else {
//			log.Println("配置文件出错..")
//		}
//		log.Fatal(err)
//	}
//	return config
//}

func EnVarInfo(HARBORURL, HARBORUSERNAME, HARBORPASSWORD string) (string, string, string) {
	//for _, v := range os.Environ() {
	//	//输出系统所有环境变量的值
	//	fmt.Println(v)
	//}
	envIP := os.Getenv(HARBORURL)
	envUserName := os.Getenv(HARBORUSERNAME)
	envPassWord := os.Getenv(HARBORPASSWORD)
	if envIP == "" {
		fmt.Println("missing address")
	} else {
		fmt.Println("IP是:", envIP)
	}
	if envUserName == "" {
		fmt.Println("missing UserName")
	} else {
		fmt.Println(envUserName)
	}
	if envPassWord == "" {
		fmt.Println("missing PassWord")
	} else {
		fmt.Println(envPassWord)
	}
	return envIP, envUserName, envPassWord
}
