package pagination_v1

import (
	"errors"
	"github.com/azdbaaaaaa/util/log"
	"math"
	"reflect"
)

const (
	MaxPageSize     = 5000
	DefaultPageSize = 20
	DefaultPage     = 1
)

func New(page, pageSize int32) *PaginationReq {
	pag := &PaginationReq{
		Page:     page,
		PageSize: pageSize,
	}
	pag.checkValid()
	return pag
}

func NewPagination(pag *PaginationReq) *PaginationReq {
	if pag == nil {
		pag = &PaginationReq{
			Page:     DefaultPage,
			PageSize: DefaultPageSize,
		}
	}
	pag.checkValid()
	return pag
}

func (m *PaginationReq) checkValid() {
	if m.PageSize < 0 || m.PageSize > MaxPageSize {
		m.PageSize = MaxPageSize
	} else if m.PageSize == 0 {
		m.PageSize = DefaultPageSize
	}

	if m.Page <= 0 {
		m.Page = DefaultPage
	}
	return
}

func (m *PaginationReq) Offset() int32 {
	m.checkValid()
	return (m.Page - 1) * m.PageSize
}

func (m *PaginationReq) Limit() int32 {
	m.checkValid()
	return m.PageSize
}

func (m *PaginationReq) Start() int64 {
	m.checkValid()
	return int64((m.Page - 1) * m.PageSize)
}

func (m *PaginationReq) End() int64 {
	m.checkValid()
	return int64(m.Page*m.PageSize) - 1
}

func (m *PaginationReq) Index(total int32) (from, to int32, err error) {
	if total == 0 {
		err = errors.New("total is zero")
		return
	}

	var totalPage int32
	totalPage = int32(math.Ceil(float64(total) / float64(m.PageSize)))
	if m.Page > totalPage {
		err = errors.New("page is over max page")
		return
	}
	from = (m.Page - 1) * m.PageSize
	if m.Page != totalPage {
		to = from + m.PageSize
	} else {
		to = total
	}
	return
}

// GetIndex to be deprecated, Use Index instead
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

// 完整数据自动分页
func (m *PaginationReq) Split(arr interface{}) (page, pageSize int32, total int64, ret []interface{}, err error) {

	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		err = errors.New("src is not a slice")
		return
	}
	l := v.Len()
	ret = make([]interface{}, 0, m.PageSize)
	if l == 0 {
		return m.Page, m.PageSize, 0, ret, nil

	}

	from, to, err := m.Index(int32(l))
	if err != nil {
		log.Errorw("Index", "err", err)
		return
	}
	//dst = src[from:to]

	for i := 0; i < l; i++ {
		if i >= int(from) && i < int(to) {
			ret = append(ret, v.Index(i).Interface())
		}
	}
	//return ret

	return m.Page, m.PageSize, int64(l), ret, nil
}
