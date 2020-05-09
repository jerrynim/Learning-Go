package main

import (
	"fmt"
	"net/http"
)

func main() {
	var results = make(map[string]string)
	urls := []string{"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.cos/"}

	apiCheckChannel := make(chan responseResult)

	//* url마다 비동기 checkUrl 함수실행
	for _, url := range urls {
		go checkURL(url, apiCheckChannel)
	}

	//* url마다 비동기 checkUrl 함수실행
	for index := 0; index < len(urls); index++ {
		result := <-apiCheckChannel
		results[result.url] = result.status
	}
	//* 채널의 결과를 출력
	for url, status := range results {
		fmt.Println(url, status)
	}

}

type responseResult struct {
	url    string
	status string
}

//* Url 체크하기
func checkURL(url string, channel chan<- responseResult) {
	fmt.Println("start checking", url)
	_, err := http.Get(url)
	if err != nil {
		channel <- responseResult{url, "Failed"}
		return
	}
	channel <- responseResult{url, "OK"}
	return

}
