package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func findBook(writer http.ResponseWriter, request *http.Request) {

	//打印request信息
	fmt.Println(request.URL)
	fmt.Println(request.Method)

	queryParam := request.URL.Query()
	id := queryParam.Get("id")
	name := queryParam.Get("author")

	fmt.Printf("the request param id:%s,name:%s", id, name)

	book := Book{
		Id:     10,
		Name:   "鲁滨孙漂流记",
		Author: "david",
	}
	bookStr, err := json.Marshal(book)
	fmt.Println(string(bookStr))
	if err != nil {
		fmt.Println("序列化失败", err)
	}
	writer.Write(bookStr)
}

func HttpServer() {

	http.HandleFunc("/demo/find_book", findBook)
	http.ListenAndServe("127.0.0.1:9090", nil)
}

func HttpClientDirectly() {

	resp, err := http.Get("http://127.0.0.1:9090/demo/find_book?id=10&author=david")

	if err != nil {
		fmt.Println("get url failed, err:", err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("resp:%s", string(b))
	}
}

func HttpClient() {

	data := url.Values{}
	data.Set("id", "10")
	data.Set("author", "david")
	queryStr := data.Encode()
	fmt.Println(queryStr)
	urlObj, _ := url.Parse("http://127.0.0.1:9090/demo/find_book/")
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	if err != nil {
		fmt.Println("new request failed", err)
	}
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("client do failed")
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("resp:%s", string(b))
	}
}

func queryBook(writer http.ResponseWriter, request *http.Request) {

}

func main() {
	HttpServer()
}
