package rosreestrcn

import "github.com/xuri/excelize/v2"

// xlsxOpenFile opens xslx-file
func xlsxOpenFile(filename string) (*excelize.File, error) {
	return excelize.OpenFile(filename)
}

// xlsxAddRow adds row of data in our template
func xlsxAddRow(wb *excelize.File, obj *ObjectInfo, i int) error {

	return nil
}
