package douban

import (
	"../pkg/loader"
	"../pkg/writer"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"

)


type info struct {
	text string
	img []byte
}

var total uint64
var url string =  "https://movie.douban.com/chart"


//解析并且返回"https://movie.douban.com/chart"的10部电影信息
func parse(docs *goquery.Document) []info {
	pages := make([]string, 0)
	var page string
	var exist bool
	docs.Find("a[class=nbg]").Each(func(i int, selection *goquery.Selection) {
		page, exist = selection.Attr("href")
		if exist {
			pages  = append(pages, page)
		}
	})
	var wg sync.WaitGroup
	wg.Add(len(pages))
	pageInfos:= make([]info, len(pages))
	for index, pageUrl := range pages {
		doc, err := loader.Load(pageUrl, "GET")
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			defer wg.Done()
			//doc.Find("div#dale_movie_subject_top_icon").Each(func(i int, selection *goquery.Selection) {
			//	pageInfos[index].text = selection.Text()
			//	println(selection.Text())
			//})
			doc.Find("div#info").Each(func(i int, selection *goquery.Selection) {
				pageInfos[index].text += selection.Text()
			})
			doc.Find("div#mainpic").Each(func(i int, selection *goquery.Selection) {
				imgUrl, exist := selection.Find("img").Attr("src")
				if exist {
					rsp, err :=http.Get(imgUrl)
					if err != nil {
						log.Fatal("download img error: ", err)
					}
					imgByte, err := ioutil.ReadAll(rsp.Body)
					if err != nil {
						log.Fatal("read imgInfo error: ", err)
					}
					//imgInfo := bytes.NewReader([]byte(imgBody))
					pageInfos[index].img = imgByte
				}

			})
		}()
		total++
	}
	wg.Wait()
	return pageInfos
}

func Start() error {
	doc, err := loader.Load(url, "GET")
	if err != nil {
		log.Fatal("load html error: ", err)
		return err
	}
	result := parse(doc)
	err = writer.MakeDir("./download"); if err != nil {
		log.Fatal("MakeDir error: ", err)
		return err
	}
	for index, page := range result {
		_ = writer.Save([]byte(page.text), "./download/"+strconv.Itoa(index)+".txt")
		_ = writer.Save([]byte(page.img), "./download/"+strconv.Itoa(index)+".jpg")
	}
	println("----------------------------")
	println("download total: ", total)
	println("----------------------------")

	return nil
}

