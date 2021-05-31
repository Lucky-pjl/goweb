package model

type Page struct {
	Books       []*Book
	PageNo      int64
	PageSize    int64
	TotalPageNo int64 // 总页数
	TotalRecord int64 // 总记录数
	MinPrice    string
	MaxPrice    string
	IsLogin     bool
	Username    string
}

func (p *Page) IsHasPre() bool {
	return p.PageNo > 1
}

func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

func (p *Page) GetPrePageNo() int64 {
	if p.IsHasPre() {
		return p.PageNo - 1
	} else {
		return 1
	}
}

func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.TotalPageNo
	}
}
