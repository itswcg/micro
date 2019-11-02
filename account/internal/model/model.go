package model

type Info struct {
	Mid  int64  `protobuf:"varint,1,opt,name=mid,proto3" json:"mid"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Sex  string `protobuf:"bytes,3,opt,name=sex,proto3" json:"sex"`
	Face string `protobuf:"bytes,4,opt,name=face,proto3" json:"face"`
}

type TokenInfo struct {
	Info
	Token string `json:"token"`
}
