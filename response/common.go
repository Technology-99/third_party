package response

type CommonResponse struct {
	Code int32       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data interface{} `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}
