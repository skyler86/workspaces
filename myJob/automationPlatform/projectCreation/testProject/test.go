package main

//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func main() {
//	r := gin.Default()
//	r.POST("/", func(c *gin.Context) {
//		username := c.PostForm("xiao")
//		password := c.DefaultPostForm("123456", "000000")
//		//c.String(200, "Hello, Geektutu")
//		c.JSON(http.StatusOK, gin.H{
//			"username": username,
//			"password": password,
//		})
//	})
//	r.Run() // listen and serve on 0.0.0.0:8080
//}

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HarborName struct {
	HN string
}

func myfunc(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("hi")
	defer fmt.Fprint(w, "OK")
	fmt.Println("method:", r.Method)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Printf("read body err, %v", err)
	} else {
		println("json:", string(body))
	}
	jsonData := []byte(string(body))
	var name HarborName
	err = json.Unmarshal(jsonData, &name)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf(name.HN)
	//else {
	//	fmt.Printf("%+v", name.hn)
	//}
}


func main() {
	// 更多http.Server的字段可以根据情况初始化
	server := http.Server{
		Addr: ":8080",
		ReadTimeout:  0,
		WriteTimeout: 0,
	}
	http.HandleFunc("/", myfunc)
	server.ListenAndServe()
}