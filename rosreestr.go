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
and returns a file.
Type inData is interface and must support method GetCN() that returns slice []string
*/
func RosreestrData(data cadastralNumbers) (file *[]byte, err error) {

	cns, err := data.GetCN()
	if err != nil {
		return
	}

	if len(cns) == 0 {
		err = fmt.Errorf(`"There are no cadastral numbers"`)
		return
	}

	file, err = xlsx("template.xlsx", &cns)
	return
}
