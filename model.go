package ZhihuZhuanlanCrawler

import "time"

type Zhuanlan struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type Author struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	Headline    string `json:"headline"`
	Description string `json:"description"`
	AvatarUrl   string `json:"avatar_url"`
}

type PinnedArticleAndAuthor struct {
	Type       string `json:"type"`
	ID         int    `json:"id"`
	Title      string `json:"title"`
	TitleImage string `json:"title_image"`
	URL        string `json:"url"`
	Excerpt    string `json:"excerpt"`
	Created    int    `json:"created"`
	Updated    int    `json:"updated"`
	Author     Author
}

type Topic struct {
	Url  string `json:"url"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Article struct {
	IsPinned    int       `json:"is_pinned"`
	Pid         int       `json:"pid"`
	CoverImg    string    `json:"cover_img"`
	Title       string    `json:"title"`
	PreviewText string    `json:"preview_text"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
	Topics      []Topic   `json:"topics"`
}

type ArticleList struct {
	Paging struct {
		IsEnd   bool `json:"is_end"`
		Totals  int  `json:"totals"`
		IsStart bool `json:"is_start"`
	} `json:"paging"`
	Data []struct {
		ID int `json:"id"`
	} `json:"data"`
}
