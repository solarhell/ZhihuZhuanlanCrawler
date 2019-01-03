# ZhihuZhuanlanCrawler
知乎专栏爬虫

![screenshot1.jpg](misc/screenshot1.jpg)

## usage

```go
package main

import (
	"fmt"
	"github.com/solarhell/ZhihuZhuanlanCrawler"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	const debug = true
	c := ZhihuZhuanlanCrawler.NewClient(&http.Client{
		Timeout: 30 * time.Second,
		Transport: &ZhihuZhuanlanCrawler.DebugRequestTransport{
			RequestHeader:  debug,
			RequestBody:    debug,
			ResponseHeader: debug,
			ResponseBody:   debug,
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
	fmt.Printf("%+v\n", *pinnedArticlePidAndAuthor)
}
```
