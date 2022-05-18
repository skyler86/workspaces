package createHarbor

import (
	"context"
	"fmt"
	"github.com/mittwald/goharbor-client/v4/apiv1"
	"github.com/spf13/viper"
	"log"
	"strings"
)

//type Configurations struct {
//	harbor harborConfig
//}
type harborConfig struct {
	url        string
	userName   string
	userPasswd string
}

func Logon(projectName string) {
	config := viper.New()

	////自动从环境变量读取匹配的参数
	//config.AutomaticEnv()
	////读取环境变量前缀,有这个前缀的环境变量才会读取连接
	//viper.SetEnvPrefix("HARBORURL")
	//viper.SetEnvPrefix("HARBORADDRESS")
	//viper.SetEnvPrefix("HARBORPASSWORD")

	//指定配置文件的信息
	config.AddConfigPath("./config/") // 设置配置文件的路径
	config.SetConfigName("app")       // 读取配置文件文件名
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

	var hc harborConfig
	hc.url = config.GetString("harbor.URL")
	hc.userName = config.GetString("harbor.USERNAME")
	hc.userPasswd = config.GetString("harbor.PASSWORD")

	//获取配置文件的值
	fmt.Println("viper load conf: ", hc.url)
	fmt.Println("viper load conf: ", hc.userName)
	fmt.Println("viper load conf: ", hc.userPasswd)

	harborClient, err := apiv1.NewRESTClientForHost(hc.url, hc.userName, hc.userPasswd)
	if err != nil {
		log.Println(err)
	}

	//将harbo仓库里项目的名字全部打印出来:
	//value, err := harborClient.ListProjects(context.TODO(), "")
	//for _, v := range value {
	//	println(v.Name)
	//}

	//在命令行添加参数:
	//var projectName string
	//flag.StringVar(
	//	&projectName,
	//	"u",
	//	"",
	//	"The project name is empty by default",
	//)
	//flag.Parse()

	CreateProject(projectName, harborClient)
}

func CreateProject(projectName string, harborClient *apiv1.RESTClient) {
	fmt.Printf("projectName=%v \n", projectName)
	var countLimit int
	var storageLimit int

	result1, err := harborClient.NewProject(context.TODO(), projectName, countLimit, storageLimit)
	result2, err := harborClient.NewProject(context.TODO(), CreateAgain(projectName), countLimit, storageLimit)

	if err != nil {
		log.Println("创建项目失败:", err.Error())
	}
	fmt.Println(result1)
	fmt.Println(result2)
}

func CreateAgain(pn string) string {
	npn := []string{"autosync", pn}
	newProjectName := strings.Join(npn, "-")
	return newProjectName
}
