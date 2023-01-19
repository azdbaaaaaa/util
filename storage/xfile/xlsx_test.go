package xfile

import "testing"

func TestCreateXlsxFile(t *testing.T) {
	xData := XlsxData{
		Headers: []string{"姓名", "年龄", "分数"},
		Data:    [][]string{{"jimmy", "20", "100"}, {"jack", "23", "99"}},
	}
	CreateXlsxFile("/tmp/test01", xData)
}
