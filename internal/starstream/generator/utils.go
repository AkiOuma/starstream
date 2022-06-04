package generator

import (
	"bytes"
	"fmt"

	"github.com/AkiOuma/starstream/internal/starstream/frame"
)

func buildParam(params ...*frame.Param) string {
	param := bytes.NewBufferString("")
	for i, v := range params {
		param.WriteString(fmt.Sprintf("%s %s", v.Name, v.DataType))
		if i < len(params)-1 {
			param.WriteString(", ")
		}
	}
	return param.String()
}
