# ZhihuZhuanlanCrawler
知乎专栏爬虫

## usage

```go
package main

import (
	"github.com/solarhell/ZhihuZhuanlanCrawler"
	"net/http"
	"time"
	"log"
	"os"
)

func main() {
	c := ZhihuZhuanlanCrawler.NewClient(&http.Client{
		Timeout: 30 * time.Second,
		Transport: &ZhihuZhuanlanCrawler.DebugRequestTransport{
			RequestHeader:  true,
			RequestBody:    true,
			ResponseHeader: true,
			ResponseBody:   true,
			Transport: &http.Transport{
				IdleConnTimeout: 30 * time.Second,
	        },
		},
	})

    pinnedArticlePidAndAuthor, err := c.GetPinnedArticlePidAndAuthor("OTalk")
    if err != nil {
    	log.Println(err)
    	os.Exit(1)
    }
    log.Println(pinnedArticlePidAndAuthor)
}
```
