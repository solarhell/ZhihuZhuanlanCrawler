# ZhihuZhuanlanCrawler
知乎专栏爬虫

![screenshot1.png](misc/screenshot1.png)

## usage

```go
package main

import (
	"fmt"
	"log"
	"os"
	Zhihu "github.com/solarhell/ZhihuZhuanlanCrawler"
)

func main() {
	const columnName = "OTalk"

	pinnedArticlePidAndAuthor, err := Zhihu.GetPinnedArticlePidAndAuthor(columnName)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", *pinnedArticlePidAndAuthor)

	pinnedArticle, err := Zhihu.GetSingleArticle(pinnedArticlePidAndAuthor.ID)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", *pinnedArticle)

	pids, err := Zhihu.GetArticlesListPids(columnName)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, pid := range pids {
		if pid == pinnedArticle.ID {
			continue
		}
		article, err := Zhihu.GetSingleArticle(pid)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%+v\n", *article)
	}
}
```
