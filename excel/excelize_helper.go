package excel

import (
	"github.com/linxlib/conv"
	"github.com/xuri/excelize/v2"
)

func MustGetValue(f *excelize.File, row, index int) string {
	name := f.GetSheetName(0)
	axis := GetAxis(row, index)
	v, _ := f.GetCellValue(name, axis)
	return v
}

func GetAxis(row, index int) string {
	axis, _ := excelize.ColumnNumberToName(index + 1)
	axis = axis + conv.String(row)
	return axis
}
