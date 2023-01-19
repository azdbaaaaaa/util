package pagination_v1

import "testing"

type Person struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func TestPaginationReq_Spilt(t *testing.T) {
	arr := []Person{{Name: "j", Age: 1}, {Name: "a", Age: 2}, {Name: "b", Age: 3}, {Name: "c", Age: 4}}

	pag := New(2, 3)
	page, pageSize, total, ret, _ := pag.Split(arr)
	t.Log(page, pageSize, total, ret)
}
