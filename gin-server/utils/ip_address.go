package utils

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/errors/gerror"
	"net/http"
)

func GetIpAddress(ip string) (string, string, error) {
	var address = ""
	var ipStr = ""
	res, err := http.Get("https://q.ip5.me/?s=" + ip)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", "", gerror.New("res.StatusCode != 200")
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", "", err
	}
	doc.Find(".main div").Each(func(i int, s *goquery.Selection) {
		//s.Find("div").Each(func(i int, s1 *goquery.Selection) {
		content := s.Text()
		if i == 0 {
			ipStr = content
		}
		if i == 1 {
			address = content
		}
		//})
	})
	return ipStr, address, nil
}
