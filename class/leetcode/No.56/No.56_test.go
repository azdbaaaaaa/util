package No_56

import (
	"testing"
)

func Test49(t *testing.T) {
	data := [][]int{{1, 4}, {2, 3}}
	//data := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	ans := merge(data)
	t.Log(ans)
}
