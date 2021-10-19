package utils

//新浪图床go语言版本

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type sinaUpload struct {
	Username, Password, Sub string
	Time                    time.Time
	temp_Pid string
}

func NewSinaUp() *sinaUpload {
	sina := new(sinaUpload)

	sina.Time = time.Now().Add(-time.Hour*24)
	return sina
}

func (this *sinaUpload) SetUsername(username string) *sinaUpload {
	this.Username = username
	return this
}

func (this *sinaUpload) SetPassword(password string) *sinaUpload {
	this.Password = password
	return this
}

func (this *sinaUpload) GetCookieDumpTime() time.Time {
	return this.Time
}

func (this *sinaUpload) LoginGetSub() error {
	username := strings.TrimSpace(this.Username)
	username = base64.StdEncoding.EncodeToString([]byte(this.Username))
	password := strings.TrimSpace(this.Password)
	client := new(http.Client)
	postData := url.Values{
		"entry":      {"sso"},
		"gateway":    {"1"},
		"from":       {"null"},
		"savestate":  {"30"},
		"useticket":  {"0"},
		"pagerefer":  {""},
		"vsnf":       {"1"},
		"su":         {username},
		"service":    {"sso"},
		"sp":         {password},
		"sr":         {"1920*1080"},
		"encoding":   {"UTF-8"},
		"cdult":      {"3"},
		"domain":     {"sina.com.cn"},
		"prelt":      {"0"},
		"returntype": {"TEXT"},
	}
	urlLink := "https://login.sina.com.cn/sso/login.php?client=ssologin.js(v1.4.15)&_=1403138799543"
	req, err := http.NewRequest(http.MethodPost, urlLink, strings.NewReader(postData.Encode()))
	if err != nil {
		return fmt.Errorf("请求建立失败：%v\n", err)
	}
	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;param=value")

	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("访问网页失败：%v\n", err)
	}
	defer resp.Body.Close()
	for _, v := range resp.Cookies() {
		if v.Name == "SUB" {
			this.Sub = v.Value
			this.Time = time.Now()
			return nil
		}
	}
	return errors.New("不知道发生了什么！")
}

func (this *sinaUpload) UploadImg(fileLink string) (string, error) {
	for {
		if time.Now().Sub(this.Time) > 12*time.Hour {
			this.LoginGetSub()
		} else {
			break
		}
	}

	urlLink := "http://picupload.service.weibo.com/interface/pic_upload.php?mime=image%2Fjpeg&data=base64&url=0&markpos=1&logo=&nick=0&marks=1&app=miniblog"
	isMultipart := strings.HasPrefix(fileLink, "http")
	postData := url.Values{}

	if !isMultipart {
		f, err := os.Open(fileLink)
		if err != nil {
			return "", fmt.Errorf("打开本地文件出错：%v\n", err)
		}
		defer f.Close()
		buf, err := ioutil.ReadAll(f)
		if err != nil {
			return "", fmt.Errorf("读取本地文件出错：%v\n", err)
		}
		postData.Set("b64_data", base64.StdEncoding.EncodeToString([]byte(buf)))
	} else {
		f, err := http.Get(fileLink)
		if err != nil {
			return "", fmt.Errorf("打开远程文件出错：%v\n", err)
		}
		defer f.Body.Close()
		buf, err := ioutil.ReadAll(f.Body)
		if err != nil {
			return "", fmt.Errorf("读取远程文件出错：%v\n", err)
		}
		postData.Set("b64_data", base64.StdEncoding.EncodeToString([]byte(buf)))
	}

	client := new(http.Client)

	req, err := http.NewRequest(http.MethodPost, urlLink, strings.NewReader(postData.Encode()))
	if err != nil {
		return "", fmt.Errorf("上传文件建立请求失败：%v\n", err)
	}

	req.AddCookie(&http.Cookie{Name: "SUB", Value: this.Sub, Path: "/"})

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;param=value")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("上传文件post访问失败：%v\n", err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	sb := string(b)
	pid := sb[strings.Index(sb, "pid")+6 : strings.LastIndex(sb, "}}}}")-1]
	if len(pid) > 32 {
		return "", fmt.Errorf("请求数据非法！")
	}
	this.temp_Pid = pid
	return pid, nil
}

func (this *sinaUpload) GetTypeImg(i uint8) string{
	t := make([]string,10)
	t = []string{
		"square/",
		"thumb150/",
		"orj360/",
		"orj480/",
		"mw690/",
		"mw1024/",
		"mw2048/",
		"small/",
		"bmiddle/",
		"large/",
	}
	if i> 9{
		i = 9
	}
	return "https://ww1.sinaimg.cn/" + t[i] + this.temp_Pid
}

