package department

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func BenchmarkGetOrgEnrolling(b *testing.B) {
	// 测试并发
	for i := 0; i < 10; i++ {
		go GetRegData()
	}
	m := make(chan int)
	m <- 10
	b.StopTimer()
}

func GetRegData() {

	url := "http://xdqy.top/dep/reg-data"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXBpZCI6NDMsImRlcG5hbWUiOiJiYWdhIiwiZXhwIjoxNjQxNDcyNjc1LCJpYXQiOjE2NDA4Njc4NzUsIm9yZ2lkIjoxLCJvcmduYW1lIjoi6Z2S5Y2PIn0.jfP7yTRwl5FfTyVeiAZ5EudqVvnS3IfDygOA2cq2K-o")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
