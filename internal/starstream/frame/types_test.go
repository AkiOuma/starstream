package frame

import "testing"

func TestConvertGoType(t *testing.T) {
	source, target := "time", "time.Time"
	if ans := ConvertGoType(source); ans != nil || ans.TypeName != target {
		t.Error("convert go type do not pass")
	}
}

func TestConvertProtoType(t *testing.T) {
	source, target := "time", "google.protobuf.Timestamp"
	if ans := ConvertProtoType(source); ans != nil || ans.TypeName != target {
		t.Error("convert go type do not pass")
	}
}
