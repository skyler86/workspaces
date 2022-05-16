package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"projectcreation/createHarbor"
)

type HarborProject struct {
	PN string
}

func myfunc(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hi")
	//fmt.Printf("hi")
	defer fmt.Fprint(w, "OK")
	fmt.Println("method:", r.Method)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err:%v", err)
	} else {
		println("json内容是:\n", string(body))
	}
	jsonData := []byte(string(body))

	var name HarborProject
	err = json.Unmarshal(jsonData, &name)
	if err != nil {
		fmt.Printf("error:%v", err)

	} else {
		fmt.Printf("web端输入的项目名是:%v \n", name.PN)
	}
	createHarbor.Logon(name.PN)
}

func main() {
	server := http.Server{
		Addr:         "10.109.80.14:8080",
		ReadTimeout:  0,
		WriteTimeout: 0,
	}
	http.HandleFunc("/", myfunc)
	server.ListenAndServe()

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { })
	//http.ListenAndServe("localhost:8000", nil)
	//log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
