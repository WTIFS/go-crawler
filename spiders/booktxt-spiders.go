package spiders

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/wtifs/go-crawler/common"
	"fmt"
)

type BookTextSpider struct {

}

func (self *BookTextSpider) SpiderUrl(url string) (error) {
	book := SBook{}
	book.Url = url
	doc, err := goquery.NewDocument(url)

	if (err != nil) {
		return err
	}

	bookname := common.GbkToUtf8(doc.Find("#info h1").Text())
	fmt.Println(bookname)

	// 先收集章节名&url
	doc.Find("#list dd").Each(func(i int, contentSelection *goquery.Selection) {
		// 前9章不要
		if (i<9) {
			return
		}
		pre := i-9
		next := i-7
		title := common.GbkToUtf8(contentSelection.Find("a").Text())
		fmt.Println(title)
		href, _ := contentSelection.Find("a").Attr("href")
		chapter := SChapter{Title: title, Url: "http://www.booktxt.net" + href, Order: i-8, Pre: pre, Next: next}
		book.Chapters = append(book.Chapters, &chapter)
	});

	// 100个协程爬取各个章节
	channel := make(chan struct{}, 100)
	for _, chapter := range book.Chapters {
		channel <- struct{}{}
		go SpiderChapter(chapter, channel)
	}

	for i:=0; i<100; i++ {
		channel <- struct{}{}
	}

	close(channel)
	return nil
}

type ChanTag struct{}

func SpiderChapter(chapter *SChapter, c chan struct{}) {
	defer func(){<-c}()
	doc, err := goquery.NewDocument(chapter.Url)
	if err != nil {
		fmt.Println("get chapter details error: ", err.Error())
		return
	}

	content := doc.Find("#content").Text()
	content = common.GbkToUtf8(content)
	fmt.Println("拉取 " + chapter.Title + " 完成")
}

