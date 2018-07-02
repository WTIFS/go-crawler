package spiders

import (
"errors"
)

type SBook struct{
	Name string
	Image string
	Url string
	Chapters []*SChapter
}

type SChapter struct{
	Title string
	Url string
	Order int
	Pre int
	Next int
	Content string
}

type Spider interface{
	SpiderUrl(url string) error
}

func NewSpider(from string) (Spider, error){
	switch from {
	case "booktxt":
		return new(BookTextSpider), nil
	default:
		return nil, errors.New("系统暂未处理该类型的配置文件")
	}
}