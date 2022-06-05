package template

var ProtoTransportTemplate = `syntax = "proto3";

package %s;

option go_package = "%s";

%s
service %s {
%s}

%s
`
var ProtoTransportRpcTemplate = `	rpc %s(%s) returns (%s) {};
`

var ProtoMessageTemplate = `message %s {
%s}
`
