package main

import (
	"Go-000/Week02/biz"
	"fmt"
)

func main() {

	_, err := biz.NewMockBiz().QueryCount()
	if err != nil {
		fmt.Printf("QueryCountFailed||err=%+v", err)
	}
}
