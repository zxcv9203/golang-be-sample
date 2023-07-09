package request

type Page struct {
	size int64
	page int64
}

func NewPageRequest(size int64, page int64) Page {
	if size <= 0 {
		size = 20
	}

	return Page{
		size: size,
		page: page,
	}
}

func (p *Page) Size() int64 {
	return p.size
}

func (p *Page) Page() int64 {
	return p.page
}
