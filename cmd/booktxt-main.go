package main

import (
	"fmt"
	"github.com/wtifs/go-crawler/spiders"
)

func main() {
	fmt.Println("start")
	s, _ := spiders.NewSpider("booktxt")

	err := s.SpiderUrl("http://www.booktxt.net/2_2219/")
	if (err != nil) {
		fmt.Println("new Document error: ", err.Error())
	}
}