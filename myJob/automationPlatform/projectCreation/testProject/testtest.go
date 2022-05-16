package main

import (
"encoding/json"
"fmt"
)

// Actress 女演员
type Actress struct {
	Name       string
	Birthday   string
	BirthPlace string
	Opus       []string
}

func main() {
	// 普通JSON
	// 因为json.UnMarshal() 函数接收的参数是字节切片
	// 所以需要把JSON字符串转换成字节切片。
	jsonData := []byte(`{
      "name":"木村拓展",
      "birthday":"1972-11-13",
      "birthPlace":"日本东京都",
      "opus":[
         "《恋爱世纪》",
         "《HERO》",
         "《冰上恋人》"
      ]
   }`)

	var actress Actress
	err := json.Unmarshal(jsonData, &actress)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("姓名：%s\n", actress.Name)
	fmt.Printf("生日：%s\n", actress.Birthday)
	fmt.Printf("出生地：%s\n", actress.BirthPlace)
	fmt.Println("作品：")
	for _, val := range actress.Opus {
		fmt.Println("\t", val)
	}
}
