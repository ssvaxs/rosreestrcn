package rosreestrcn

import (
	"fmt"
)

type (
	cadastralNumbers interface {
		GetCN() ([]string, error) // slice of cadastral numbers
	}
)

/*
RosreestrData gets cadastral numbers
and returns a link to the file.
Type inData is interface and must support method GetCN() that returns slice []string
*/
func RosreestrData(data cadastralNumbers) (link string, err error) {

	cns, err := data.GetCN()
	if err != nil {
		return
	}

	if len(cns) == 0 {
		err = fmt.Errorf(`"There are no cadastral numbers"`)
		return
	}

	var obj *ObjectInfo

	wb, err := xlsxOpenFile("")
	if err != nil {
		return
	}
	defer wb.Close()

	for i, cn := range cns {
		obj, err = rosreestrRequest(cn)
		if err != nil {
			return
		}

		err = xlsxAddRow(wb, obj, i)
		if err != nil {
			return
		}
	}

	return
}
