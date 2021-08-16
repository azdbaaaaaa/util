package pagination

import "errors"

func (m *PaginationReq) GetIndex(totalNum int32) (from, to int32, err error) {
	var totalPage int32
	if totalNum%m.PageSize == 0 {
		totalPage = totalNum / m.PageSize
	} else {
		totalPage = totalNum/m.PageSize + 1
	}
	if totalNum == 0 || m.Page > totalPage {
		return from, to, errors.New("page error")
	}
	from = (m.Page - 1) * m.PageSize
	if m.Page != totalPage {
		to = from + m.PageSize
	} else {
		to = totalNum
	}
	return from, to, nil
}
