package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var result *Results

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
	Error   bool        `json:"error"`
	Results []*GankMain `json:"results"`
}

func AppList(w http.ResponseWriter, r *http.Request) {
	//生成要访问的url
	response, err := http.Get("http://gank.io/api/data/Android/10/1")
	//判断是否出错
	if err != nil {
		return
	}

	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	//stdout := os.Stdout
	//_,  err = io.Copy(stdout, response.Body)
	//func NewDecoder(r io.Reader) *Decoder
	eror := json.NewDecoder(response.Body)
	if err := eror.Decode(&result); err != nil {
		return
	}

	//判断是否请求成功
	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		fmt.Println("请求出错：%s", response.Status)
		return
	}

	//mm := map[string]interface{}{}
	//t_r := []GankMain{}
	//mm["error"] = result.Error
	func NewEncoder(w io.Writer) *Encoder
	enc := json.NewEncoder(w)
	for _, v := range result.Results {
		if v.Id == "59eff4ad421aa90fef2034cb" {
			v.Who = "12313131231"
		}
	}
	if err := enc.Encode(&result); err != nil {
		//log.Println(err)
		return
	}
	//res, _ := json.Marshal(result)
	//io.WriteString(w, string(res))
}
func main() {
	http.HandleFunc("/test", AppList)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
