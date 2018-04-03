package scheduler

var URLs = []string{}

//抛出URL
func PopUrl() string {
	length := len(URLs)
	if length < 1 {
		return ""
	}

	url := URLs[length - 1]
	URLs = URLs[:length - 1]
	return url
}

//添加URL
func AppendUrl(url string) {
	URLs = append(URLs, url)
}
