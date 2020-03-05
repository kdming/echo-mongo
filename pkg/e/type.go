package e

// 常用的类型

// map
type Map map[string]interface{}

// page
type Page struct {
	Page  int
	Limit int
}

func (p *Page) GetSkip() int {
	if p.Page == 0 {
		return 0
	} else {
		return (p.Page - 1) * p.Limit
	}
}
