package xfile

import "github.com/tealeg/xlsx/v3"

type XlsxData struct {
	Headers []string
	Data    [][]string
}

func CreateXlsxFile(path, filename string, xData XlsxData) error {
	wb := xlsx.NewFile()
	sh, err := wb.AddSheet("Sheet1")
	if err != nil {
		return err
	}
	defer sh.Close()
	headerRow, err := sh.Row(0)
	if err != nil {
		return err
	}
	headerRow.SetHeight(20)
	for _, header := range xData.Headers {
		cell := headerRow.AddCell()
		cell.SetString(header)
	}

	for i, data := range xData.Data {
		dataRow, err := sh.Row(i + 1)
		if err != nil {
			return err
		}
		dataRow.SetHeight(20)
		for _, c := range data {
			cell := dataRow.AddCell()
			cell.SetString(c)
		}
	}
	err = wb.Save(path + filename + ".xlsx")
	if err != nil {
		return err
	}
	return nil
}
