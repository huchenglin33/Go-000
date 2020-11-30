package main

import (
	"fmt"
	"jike/biz"
)

func main() {
	_, err := biz.NewMockBiz().QueryCount()
	if err != nil {
		fmt.Printf("QueryCountFailed||err=%+v", err)
	}
}
