package downloader

import (
	"io"
	"net/http"
	"log"
	"routinego/car_guazi/fake"
	"fmt"
	"net/http/cookiejar"
)

//获取页面信息
func Get(url string) (io.Reader, error ){
	//1.创建一个http请求客户端
	//1.1初始化cookiejar
	var cookieJar *cookiejar.Jar
	cookieJar, _ = cookiejar.New(nil)
	//1.2初始化客户端
	client := http.Client{
		Jar: cookieJar,
	}
	//2. GET请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("http.NewRequest err: %v", err)
	}
	// 3. 增加请求头进行伪装
	req.Header.Add("User-Agent", fake.GetUserAgent())
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("Accept-Encoding","gzip, deflate, br")
	req.Header.Add("Accept-Language","zh-CN,zh;q=0.9")
	req.Header.Add("Cache-Control","no-cache")
	req.Header.Add("Upgrade-Insecure-Requests","1")
	req.Header.Add("Referer", "https://www.guazi.com/yiwugz/")
	req.Header.Add("Cookie","antipas=7a3a076S74649789x1715906; uuid=9ea6f6bd-d0b5-4460-d005-60bdfbc3a194; cityDomain=sz; cainfo=%7B%22ca_s%22%3A%22self%22%2C%22ca_n%22%3A%22self%22%2C%22ca_i%22%3A%22-%22%2C%22ca_medium%22%3A%22-%22%2C%22ca_term%22%3A%22-%22%2C%22ca_kw%22%3A%22-%22%2C%22keyword%22%3A%22-%22%2C%22ca_keywordid%22%3A%22-%22%2C%22scode%22%3A%22-%22%2C%22version%22%3A1%2C%22platform%22%3A%221%22%2C%22client_ab%22%3A%22-%22%2C%22guid%22%3A%229ea6f6bd-d0b5-4460-d005-60bdfbc3a194%22%2C%22sessionid%22%3A%22-%22%7D; preTime=%7B%22last%22%3A1522740193%2C%22this%22%3A1522740177%2C%22pre%22%3A1522740177%7D; clueSourceCode=%2A%2300")
	// 4.发送请求，获得响应
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client.Do err: %v", err)
	}
	// 5.响应BODY解码并返回
	//defer resp.Body.Close()
	//var result io.Reader
	fmt.Println(resp.Cookies())
	content := resp.Body
	statusCode := resp.StatusCode
	if statusCode == 200 {
		buf := make([]byte, 2048)
		var result string

		for{
			n, err := content.Read(buf)
			if n == 0{
				if err == io.EOF {
					break
				}
				fmt.Println("read err = ", err)
				break
			}
			result += string(buf[:n])
		}
		fmt.Println("result : ", result)

		return resp.Body, nil
	} else {
		fmt.Println("页面访问出错:",statusCode)
		err = fmt.Errorf("页面%s请求失败...", url )
		return resp.Body, err
	}

}
