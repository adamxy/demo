package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var result Results

type GankMain struct {
	Id          string    `json:"_id"`
	CreatedAt   time.Time `json:"createdAt"`
	Desc        string    `json:"desc"`
	PublishedAt time.Time `json:"publishAt"`
	Source      string    `json:"source"`
	Type        string    `json:"type"`
	Url         string    `json:"url"`
	Used        bool      `json:"used"`
	Who         string    `json:"who"`
}
type Results struct {
	Error   bool       `json:"error"`
	Results []GankMain `json:"results"`
}

func main() {
	//生成要访问的url
	response, err := http.Get("http://gank.io/api/data/Android/10/1")
	//判断是否出错
	if err != nil {
		return
	}

	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	//stdout := os.Stdout
	//_,  err = io.Copy(stdout, response.Body)
	eror := json.NewDecoder(response.Body)
	if err := eror.Decode(&result); err != nil {
		log.Println(err)
		return
	}

	//判断是否请求成功
	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		fmt.Println("请求出错：%s", response.Status)
		return
	}

	//fmt.Println(result)
	//res2B, err := json.Marshal(result)
	//if err != nil {
	//		fmt.Println(err)
	//}
	//fmt.Println(string(res2B))
	//fmt.Println(reflect.TypeOf(result))
	mm := map[string]interface{}{}
	t_r := []GankMain{}
	mm["error"] = result.Error
	enc := json.NewEncoder(os.Stdout)
	for _, v := range result.Results {
		//fmt.Println(v.Who)
		if v.Id == "59eff4ad421aa90fef2034cb" {
			fmt.Println("_______________________")
			v.Who = "12313131231"
		}
		t_r = append(t_r, v)
		//fmt.Println(v)
		//err := result.Error
		//data, _ := json.Marshal(v)
		//fmt.Println(string(data))
		//fmt.Println(err)
	}
	mm["results"] = t_r
	if err := enc.Encode(&mm); err != nil {
		//log.Println(err)
		return
	}

	//fmt.Println(reflect.TypeOf(res2B))
	//ress := json.Unmarshal(*result []byte, v interface{})
}
