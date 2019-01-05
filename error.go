package ZhihuZhuanlanCrawler

import "errors"

var (
	ColumnNameCanNotBeEmpty = errors.New("专栏名不能为空")
	PidCanNotBeEmpty        = errors.New("pid 不能为空")
)
