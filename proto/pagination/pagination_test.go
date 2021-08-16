package pagination

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestPaginationReq_GetIndex(t *testing.T) {
	m := &PaginationReq{
		Page:     1,
		PageSize: 10,
	}
	from, to, err := m.GetIndex(5)
	assert.Equal(t, err, nil)
	assert.Equal(t, from, 0)
	assert.Equal(t, to, 5)
}
