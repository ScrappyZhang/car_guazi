package spiders

import (
	"regexp"
	"github.com/PuerkitoBio/goquery" //页面解析
	"strconv"
	"log"
	"strings"
	"os"
	"io"
	"fmt"
)

var (
	compileNumber = regexp.MustCompile("\\d+\\.\\d+")

)




//车结构体字段
type QcCar struct {
	CityName string //城市名称
	Title string  //车标题
	Price float64  //价格
	OldPrice float64 //原价
	Kilometer float64  //行程公里数
	Year int //年份
	TransferCity string //过户城市
}


func ReadFile(path string) (string) {
	// 1.打开文件
	f, err:=os.Open(path)
	if err != nil && err != io.EOF{
		fmt.Println("OpenReadFile. err = ",err)
		return ""
	}
	//2. 关闭文件
	defer f.Close()
	result := ""
	for {

		buf := make([]byte, 1024*2)
		// n 代表读取的内容长度，err代表错误
		n, err := f.Read(buf)
		if err != nil{
			if err == io.EOF{
				result += string(buf[:n])
				return result
			} else {
				fmt.Println("ReadFile.err = ",err)
				return ""
			}
		}
		if n == 0{
			break
		}
		result += string(buf[:n])
	}
	return result
}

//获取城市列表
func GetCitys() map[string]string {
//func main(){
	citys := map[string]string{}
	citySeo :=  ReadFile("./citys.html")
	if citySeo == "" {
		log.Println("无seo城市信息")
	} else {
		//1.解析内容
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(citySeo))
		if err != nil {
			log.Printf("citys.Get err: %v", err)
		}
		//2.获取城市信息
		//<a href="/zz/" target="_blank">郑州二手车</a>
		doc.Find("a").Each(func(i int, selection *goquery.Selection) {
			//名称+二手车
			cityName := selection.Text()
			// 链接
			href, exist := selection.Attr("href")
			if exist{
				//若存在， 将城市添加到citys
				//1.数据处理
				hreflist := strings.Split(href, "/")
				href = hreflist[1]
				//2. 添加地区信息
				citys[cityName] = href
			}


		})

	}

	//fmt.Println(citys)
	return citys
	//for _, v := range citys {
	//	fmt.Println(v)
	//}
}




//<div class="citycont">
//<div class="fn-left" style="font-weight:bold; color:#ff7700"></div>
//<div class="citycont-city" id="sltArea">
//<a href="javascript:void(0);" id="city01">切换城市<i class="icon10 iog10-sjb"></i></a>
//</div>
//
//</div>

//获取页面城市信息
func GetCityName(doc *goquery.Document) string {
	return doc.Find(".city-curr").Text()
}

//<a href="/www/buy/o3/#bread" data-gzlog="tracking_type=click&amp;eventid=0060000000000059" class="next"><span>下一页</span><em>&gt;</em></a>
//获取下一页链接
func GetNextPageUrl(doc *goquery.Document) (val string, exists bool) {
	return doc.Find(".pageBox .next").Attr("href")
}

//<li class="link-on"><a><span>2</span></a></li>
//当前页页码
func GetCurrentPage(doc *goquery.Document) (page int) {
	pageS := doc.Find(".link-on a span").Text()
	fmt.Println("11111111111111>",pageS)
	if pageS != "" {
		var err error
		page, err = strconv.Atoi(pageS)
		if err != nil {
			log.Printf("spiders.GetCurrentPage err: %v", err)
		}
	}

	return page
}


func GetCars(doc *goquery.Document) (cars []QcCar) {
	cityName := GetCityName(doc) //当前页面城市名称
	doc.Find(".carlist li").Each(func(i int, selection *goquery.Selection) {
		//标题
		title := selection.Find("a .t").Text()
		// 价格
		price := selection.Find("a .t-price p").Text()
		//  年份、行程、城市
		tinfo := selection.Find("a .t-i").Text()
		tinfoList := strings.Split(tinfo,"|")
		//年份
		year := tinfoList[1]
		//行程
		kilometer := tinfoList[2]
		// 城市
		//city := tinfoList[3]
		//city = strings.Trim(city, `"`)
		//原价
		oldPrice := selection.Find("a .t-price .line-through").Text()

		// 数据处理
		kilometer = strings.Join(compileNumber.FindAllString(kilometer, -1), "")
		year = strings.Join(compileNumber.FindAllString(strings.TrimSpace(year), -1), "")
		price = strings.Join(compileNumber.FindAllString(strings.TrimSpace(price), -1), "")
		oldPrice = strings.Join(compileNumber.FindAllString(strings.TrimSpace(oldPrice), -1), "")
		priceS, _ := strconv.ParseFloat(price, 64)
		oldPriceS, _ := strconv.ParseFloat(oldPrice, 64)
		kilometerS, _ := strconv.ParseFloat(kilometer, 64)
		yearS, _ := strconv.Atoi(year)

		//将车信息添加进cars切片
		cars = append(cars, QcCar{
			CityName: cityName,
			Title: title,
			Price: priceS,
			OldPrice:oldPriceS,
			Kilometer: kilometerS,
			Year: yearS,
		})
	})

	return cars
}

