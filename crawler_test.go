package ZhihuZhuanlanCrawler

import "testing"

const columnName = "OTalk" // https://zhuanlan.zhihu.com/Otalk
const pid = 60968502       // https://zhuanlan.zhihu.com/p/60968502

func TestClient_GetPinnedArticlePidAndAuthor(t *testing.T) {
	_, err := GetPinnedArticlePidAndAuthor(columnName)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestClient_GetArticlesListPids(t *testing.T) {
	_, err := GetArticlesListPids(columnName)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestClient_GetSingleArticle(t *testing.T) {
	_, err := GetSingleArticle(pid)
	if err != nil {
		t.Error(err.Error())
	}
}
