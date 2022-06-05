package template

var TransportHelperTemplate = `package helper

import (
	v1 "%s"
	"%s"
	"%s"
	"google.golang.org/protobuf/types/known/timestamppb"
)

%s

%s

%s

func bulkInt32ToInt(source []int32) []int {
	ans := make([]int, 0, len(source))
	for _, v := range source {
		ans = append(ans, int(v))
	}
	return ans
}

`

var Proto2EntityTemplate = `func %s(source ...*%s) []*%s {
	ans := make([]*%s, 0, len(source))
	for _, v := range source {
		ans = append(ans, entity.New%s(%s, &valueobject.%s{
			%s
		}))
	}
	return ans
}
`

var Entity2ProtoTemplate = `func %s(source ...*%s) []*%s {
	ans := make([]*%s, 0, len(source))
	for _, v := range source {
		ans = append(ans, &v1.%s{
			%s
		})
	}
	return ans
}
`

var ConvertQuerierTemplate = `func %s(source *%s) *%s {
	return &valueobject.%s{
		%s
	}
}
`
