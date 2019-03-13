package ZhihuZhuanlanCrawler

func checkHttpStatusCode(code int) bool {
	if (code >= 200 && code <= 300) || code == 304 {
		return true
	}
	return false
}
