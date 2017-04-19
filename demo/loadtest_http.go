package main

import (
	"fmt"
	"github.com/linsongze/loadtest"
	"io/ioutil"
	"net/http"
)
func httpGet() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("request erorr")
	}

	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("request erorr")
	}
	if body !=nil {
	}



}
func main() {
	runner := loadtest.New()
	runner.SetThreadNum(5)
	runner.SetRunFunction(func() {
		//time.Sleep(time.Second)
		httpGet()
	})
	runner.Start()

}
