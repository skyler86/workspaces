package createHarbor

import (
	"context"
	"fmt"
	"github.com/mittwald/goharbor-client/v4/apiv1"
	"log"
)

type HC struct {
	uri        string
	userName   string
	userPasswd string
}

func Logon(projectName string) {
	var hc HC
	hc.uri = "https://harborsy.lenovo.com.cn/api"
	hc.userName = "jenkins-harbor"
	hc.userPasswd = "jenkins-Harbor123"
	//hc.uri = "https://192.168.31.150/api"
	//hc.userName = "admin"
	//hc.userPasswd = "rainbow123"

	harborClient, err := apiv1.NewRESTClientForHost(hc.uri, hc.userName, hc.userPasswd)
	if err != nil {
		log.Println(err)
	}
	//value, err := harborClient.ListProjects(context.TODO(), "")
	//for _, v := range value {
	//	println(v.Name)
	//}

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

	result, err := harborClient.NewProject(context.TODO(), projectName, countLimit, storageLimit)

	if err != nil {
		log.Println("创建项目错误", err.Error())
	}
	fmt.Println(result)
}
