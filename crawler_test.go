package ZhihuZhuanlanCrawler

import (
	"testing"
)

const columnName = "OTalk" // https://zhuanlan.zhihu.com/Otalk
const pid = 60968502       // https://zhuanlan.zhihu.com/p/60968502

func TestClient_GetPinnedArticlePidAndAuthor(t *testing.T) {
	info, err := GetPinnedArticlePidAndAuthor(columnName)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("%+v\n", *info)
}

func TestClient_GetArticlesListPids(t *testing.T) {
	pids, err := GetArticlesListPids(columnName)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("%v\n", pids)
}

func TestClient_GetSingleArticle(t *testing.T) {
	article, err := GetSingleArticle(pid)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("%+v\n", *article)
}

func TestCrawl(t *testing.T) {
	const columnName = "OTalk"

	pinnedArticlePidAndAuthor, err := GetPinnedArticlePidAndAuthor(columnName)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", *pinnedArticlePidAndAuthor)

	pinnedArticle, err := GetSingleArticle(pinnedArticlePidAndAuthor.ID)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", *pinnedArticle)

	pids, err := GetArticlesListPids(columnName)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("pids: %v\n", pids)

	for _, pid := range pids {
		if pid == pinnedArticle.ID {
			continue
		}
		article, err := GetSingleArticle(pid)
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("%+v\n", *article)
	}
}
