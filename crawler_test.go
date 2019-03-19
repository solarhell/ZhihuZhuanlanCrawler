package ZhihuZhuanlanCrawler

import "testing"

const columnName = "OTalk" // https://zhuanlan.zhihu.com/Otalk
const pid = 41604227 // https://zhuanlan.zhihu.com/p/41604227


func TestClient_GetPinnedArticlePidAndAuthor(t *testing.T) {
	c := NewClient(nil)
	_, err := c.GetPinnedArticlePidAndAuthor(columnName)
	if err != nil {
		t.Error(err.Error())
	}

	t.Log("GetPinnedArticlePidAndAuthor ok")

}

func TestClient_GetArticlesListPids(t *testing.T) {
	c := NewClient(nil)
	_, err := c.GetArticlesListPids(columnName)
	if err != nil {
		t.Error(err.Error())
	}

	t.Log("GetArticlesListPids ok")
}

func TestClient_GetSingleArticle(t *testing.T) {
	c := NewClient(nil)
	_, err := c.GetSingleArticle(pid)
	if err != nil {
		t.Error(err.Error())
	}

	t.Log("GetSingleArticle ok")
}
