package loader

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func Load(url string, method string) (*goquery.Document,error) {
	client := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal("http.NewRequest() error: ", err)
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36")
	rsp, err := client.Do(req)
	if err != nil {
		log.Fatal("client.Do() error: ", err)
		return nil , err
	}
	defer rsp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(rsp.Body)
	if err != nil {
		log.Fatal("html request error: ", err)
		return nil, err
	}
	return doc, nil
}
//func NewLoader() *loader {
//	return new(loader)
//}

