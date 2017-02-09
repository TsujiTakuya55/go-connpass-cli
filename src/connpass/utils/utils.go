package utils

import (
	"net/http"
	"fmt"
)

const BaseUrl = "https://connpass.com/api/v1/event"

func NewResponse(url string) (*http.Response, error) {
	req, err := NewRequest(url)

	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}

func NewRequest(url string) (*http.Request, error) {

	//Requestを生成
	req, err := http.NewRequest("GET", url, nil)

	fmt.Println(req)
	//TODO 何かリクエストに処理を加えたい場合はここに記載
	return req, err
}