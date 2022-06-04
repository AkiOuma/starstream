package template

var ProtoTemplate = `syntax = "proto3";

package %s;

option go_package = "%s";

%s
message %s {
%s}

message %s {
%s}
`
