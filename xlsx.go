package rosreestrcn

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

var (
	wb *excelize.File
)

func xlsx(filename string, cns []string) (file []byte, err error) {

	if len(cns) == 0 {
		err = fmt.Errorf(`"There are no cadastral numbers"`)
		return
	}

	var obj *ObjectInfo

	wb, err = excelize.OpenFile("template.xlsx")
	if err != nil {
		return
	}
	defer wb.Close()

	for i, cn := range cns {
		obj, err = rosreestrRequest(cn)
		if err != nil {
			return
		}

		err = xlsxAddRow(obj, i)
		if err != nil {
			return
		}
	}

	// save in file

	// make zip

	return
}

// xlsxAddRow adds row of data in our template
func xlsxAddRow(obj *ObjectInfo, i int) error {

	return nil
}
