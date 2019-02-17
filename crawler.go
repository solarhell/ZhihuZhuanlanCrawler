package ZhihuZhuanlanCrawler

import (
	"encoding/json"
	"fmt"
	"log"
)

func (c *Client) GetPinnedArticlePidAndAuthor(columnName string) (*PinnedArticleAndAuthor, error) {
	if columnName == "" {
		return nil, ColumnNameCanNotBeEmpty
	}
	u := fmt.Sprintf("https://zhuanlan.zhihu.com/api/columns/%s/pinned-article", columnName)
	res, err := c.SendNewZhihuRequest(u)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	pinnedArticleAndAuthor := PinnedArticleAndAuthor{}
	err = json.Unmarshal(res, &pinnedArticleAndAuthor)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pinnedArticleAndAuthor, nil
}

func (c *Client) GetSingleArticle(pid int) (*Article, error) {
	if pid == 0 {
		return nil, PidCanNotBeEmpty
	}
	u := fmt.Sprintf("https://zhuanlan.zhihu.com/api/posts/%d", pid)
	res, err := c.SendNewZhihuRequest(u)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	article := Article{}
	err = json.Unmarshal(res, &article)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &article, nil
}

func (c *Client) GetArticlesListPids(columnName string) ([]int, error) {
	if columnName == "" {
		return nil, ColumnNameCanNotBeEmpty
	}

	var limit = 20
	var offset = 0

	u := fmt.Sprintf("https://zhuanlan.zhihu.com/api/columns/%s/articles?limit=%d&offset=%d", columnName, limit, offset)
	res, err := c.SendNewZhihuRequest(u)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	articleList := ArticleList{}
	err = json.Unmarshal(res, &articleList)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var articleIds = []int{}

	for _, entry := range articleList.Data {
		articleIds = append(articleIds, entry.ID)
	}

	for offset = offset + limit; offset < articleList.Paging.Totals; offset = offset + limit {
		u := fmt.Sprintf("https://zhuanlan.zhihu.com/api/columns/%s/articles?limit=%d&offset=%d", columnName, limit, offset)
		res, err := c.SendNewZhihuRequest(u)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		articleList := ArticleList{}
		err = json.Unmarshal(res, &articleList)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		for _, entry := range articleList.Data {
			articleIds = append(articleIds, entry.ID)
		}
	}

	return articleIds, nil
}
