package ZhihuZhuanlanCrawler

type Zhuanlan struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type Author struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	Headline    string `json:"headline"`
	Description string `json:"description"`
	AvatarUrl   string `json:"avatar_url"`
	Type        string `json:"type"`
	UID         string `json:"uid"`
	URL         string `json:"url"`
	URLToken    string `json:"url_token"`
	UserType    string `json:"user_type"`
	AvatarURL   string `json:"avatar_url"`
}

type PinnedArticleAndAuthor struct {
	Type     string `json:"type"`
	ID       int    `json:"id"`
	Updated  int64  `json:"updated"`
	Created  int64  `json:"created"`
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
	URL      string `json:"url"`
	Excerpt  string `json:"excerpt"`
	Author   Author
}

type Topic struct {
	Url  string `json:"url"`
	Type string `json:"type"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Article struct {
	ID       int     `json:"id"`
	Type     string  `json:"type"`
	Title    string  `json:"title"`
	URL      string  `json:"url"`
	Updated  int64   `json:"updated"`
	Created  int64   `json:"created"`
	Excerpt  string  `json:"excerpt"`
	Content  string  `json:"content"`
	ImageURL string  `json:"image_url"`
	Topics   []Topic `json:"topics"`
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
