package rosreestrcn

/*
RosreestrData gets cadastral numbers
and returns a file.
*/
func RosreestrData(cns []string) (file []byte, err error) {

	file, err = xlsx("template.xlsx", cns)
	return
}
