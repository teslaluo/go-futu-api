// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: Trd_GetMarginRatio.proto

package trdgetmarginratio

import (
	qotcommon "teslaluo/go-futu-api/pb/qotcommon"
	trdcommon "teslaluo/go-futu-api/pb/trdcommon"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MarginRatioInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Security        *qotcommon.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`                 //股票
	IsLongPermit    *bool               `protobuf:"varint,2,opt,name=isLongPermit" json:"isLongPermit,omitempty"`        //是否允许融资
	IsShortPermit   *bool               `protobuf:"varint,3,opt,name=isShortPermit" json:"isShortPermit,omitempty"`      //是否允许融券
	ShortPoolRemain *float64            `protobuf:"fixed64,4,opt,name=shortPoolRemain" json:"shortPoolRemain,omitempty"` //卖空池剩余量
	ShortFeeRate    *float64            `protobuf:"fixed64,5,opt,name=shortFeeRate" json:"shortFeeRate,omitempty"`       //融券参考利率
	AlertLongRatio  *float64            `protobuf:"fixed64,6,opt,name=alertLongRatio" json:"alertLongRatio,omitempty"`   //融资预警比率
	AlertShortRatio *float64            `protobuf:"fixed64,7,opt,name=alertShortRatio" json:"alertShortRatio,omitempty"` //融券预警比率
	ImLongRatio     *float64            `protobuf:"fixed64,8,opt,name=imLongRatio" json:"imLongRatio,omitempty"`         //融资初始保证金率
	ImShortRatio    *float64            `protobuf:"fixed64,9,opt,name=imShortRatio" json:"imShortRatio,omitempty"`       //融券初始保证金率
	McmLongRatio    *float64            `protobuf:"fixed64,10,opt,name=mcmLongRatio" json:"mcmLongRatio,omitempty"`      //融资 margin call 保证金率
	McmShortRatio   *float64            `protobuf:"fixed64,11,opt,name=mcmShortRatio" json:"mcmShortRatio,omitempty"`    //融券 margin call 保证金率
	MmLongRatio     *float64            `protobuf:"fixed64,12,opt,name=mmLongRatio" json:"mmLongRatio,omitempty"`        //融资维持保证金率
	MmShortRatio    *float64            `protobuf:"fixed64,13,opt,name=mmShortRatio" json:"mmShortRatio,omitempty"`      //融券维持保证金率
}

func (x *MarginRatioInfo) Reset() {
	*x = MarginRatioInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Trd_GetMarginRatio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarginRatioInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarginRatioInfo) ProtoMessage() {}

func (x *MarginRatioInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Trd_GetMarginRatio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarginRatioInfo.ProtoReflect.Descriptor instead.
func (*MarginRatioInfo) Descriptor() ([]byte, []int) {
	return file_Trd_GetMarginRatio_proto_rawDescGZIP(), []int{0}
}

func (x *MarginRatioInfo) GetSecurity() *qotcommon.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *MarginRatioInfo) GetIsLongPermit() bool {
	if x != nil && x.IsLongPermit != nil {
		return *x.IsLongPermit
	}
	return false
}

func (x *MarginRatioInfo) GetIsShortPermit() bool {
	if x != nil && x.IsShortPermit != nil {
		return *x.IsShortPermit
	}
	return false
}

func (x *MarginRatioInfo) GetShortPoolRemain() float64 {
	if x != nil && x.ShortPoolRemain != nil {
		return *x.ShortPoolRemain
	}
	return 0
}

func (x *MarginRatioInfo) GetShortFeeRate() float64 {
	if x != nil && x.ShortFeeRate != nil {
		return *x.ShortFeeRate
	}
	return 0
}

func (x *MarginRatioInfo) GetAlertLongRatio() float64 {
	if x != nil && x.AlertLongRatio != nil {
		return *x.AlertLongRatio
	}
	return 0
}

func (x *MarginRatioInfo) GetAlertShortRatio() float64 {
	if x != nil && x.AlertShortRatio != nil {
		return *x.AlertShortRatio
	}
	return 0
}

func (x *MarginRatioInfo) GetImLongRatio() float64 {
	if x != nil && x.ImLongRatio != nil {
		return *x.ImLongRatio
	}
	return 0
}

func (x *MarginRatioInfo) GetImShortRatio() float64 {
	if x != nil && x.ImShortRatio != nil {
		return *x.ImShortRatio
	}
	return 0
}

func (x *MarginRatioInfo) GetMcmLongRatio() float64 {
	if x != nil && x.McmLongRatio != nil {
		return *x.McmLongRatio
	}
	return 0
}

func (x *MarginRatioInfo) GetMcmShortRatio() float64 {
	if x != nil && x.McmShortRatio != nil {
		return *x.McmShortRatio
	}
	return 0
}

func (x *MarginRatioInfo) GetMmLongRatio() float64 {
	if x != nil && x.MmLongRatio != nil {
		return *x.MmLongRatio
	}
	return 0
}

func (x *MarginRatioInfo) GetMmShortRatio() float64 {
	if x != nil && x.MmShortRatio != nil {
		return *x.MmShortRatio
	}
	return 0
}

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header       *trdcommon.TrdHeader  `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`             //交易公共参数头
	SecurityList []*qotcommon.Security `protobuf:"bytes,2,rep,name=securityList" json:"securityList,omitempty"` //股票
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Trd_GetMarginRatio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_Trd_GetMarginRatio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S.ProtoReflect.Descriptor instead.
func (*C2S) Descriptor() ([]byte, []int) {
	return file_Trd_GetMarginRatio_proto_rawDescGZIP(), []int{1}
}

func (x *C2S) GetHeader() *trdcommon.TrdHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *C2S) GetSecurityList() []*qotcommon.Security {
	if x != nil {
		return x.SecurityList
	}
	return nil
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header              *trdcommon.TrdHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`                           //交易公共参数头
	MarginRatioInfoList []*MarginRatioInfo   `protobuf:"bytes,2,rep,name=marginRatioInfoList" json:"marginRatioInfoList,omitempty"` //账户资金
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Trd_GetMarginRatio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Trd_GetMarginRatio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C.ProtoReflect.Descriptor instead.
func (*S2C) Descriptor() ([]byte, []int) {
	return file_Trd_GetMarginRatio_proto_rawDescGZIP(), []int{2}
}

func (x *S2C) GetHeader() *trdcommon.TrdHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *S2C) GetMarginRatioInfoList() []*MarginRatioInfo {
	if x != nil {
		return x.MarginRatioInfoList
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Trd_GetMarginRatio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_Trd_GetMarginRatio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_Trd_GetMarginRatio_proto_rawDescGZIP(), []int{3}
}

func (x *Request) GetC2S() *C2S {
	if x != nil {
		return x.C2S
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//以下3个字段每条协议都有，注释说明在InitConnect.proto中
	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg  *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C     *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

// Default values for Response fields.
const (
	Default_Response_RetType = int32(-400)
)

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Trd_GetMarginRatio_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Trd_GetMarginRatio_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_Trd_GetMarginRatio_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetRetType() int32 {
	if x != nil && x.RetType != nil {
		return *x.RetType
	}
	return Default_Response_RetType
}

func (x *Response) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *Response) GetErrCode() int32 {
	if x != nil && x.ErrCode != nil {
		return *x.ErrCode
	}
	return 0
}

func (x *Response) GetS2C() *S2C {
	if x != nil {
		return x.S2C
	}
	return nil
}

var File_Trd_GetMarginRatio_proto protoreflect.FileDescriptor

var file_Trd_GetMarginRatio_proto_rawDesc = []byte{
	0x0a, 0x18, 0x54, 0x72, 0x64, 0x5f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52,
	0x61, 0x74, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x54, 0x72, 0x64, 0x5f,
	0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x1a, 0x10,
	0x54, 0x72, 0x64, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x83, 0x04, 0x0a, 0x0f, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x61, 0x74,
	0x69, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x4c, 0x6f,
	0x6e, 0x67, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x69, 0x73, 0x4c, 0x6f, 0x6e, 0x67, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x12, 0x24, 0x0a, 0x0d,
	0x69, 0x73, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x52,
	0x65, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x46, 0x65, 0x65, 0x52, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x46, 0x65, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x61, 0x74,
	0x69, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x4c,
	0x6f, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x74,
	0x69, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6d, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x69,
	0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x69, 0x6d, 0x4c, 0x6f, 0x6e, 0x67, 0x52,
	0x61, 0x74, 0x69, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6d, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52,
	0x61, 0x74, 0x69, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x69, 0x6d, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x63, 0x6d, 0x4c,
	0x6f, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c,
	0x6d, 0x63, 0x6d, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x24, 0x0a, 0x0d,
	0x6d, 0x63, 0x6d, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0d, 0x6d, 0x63, 0x6d, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x74,
	0x69, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x6d, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x69,
	0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x6d, 0x6d, 0x4c, 0x6f, 0x6e, 0x67, 0x52,
	0x61, 0x74, 0x69, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x6d, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52,
	0x61, 0x74, 0x69, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x6d, 0x6d, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x22, 0x6e, 0x0a, 0x03, 0x43, 0x32, 0x53, 0x12,
	0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x54, 0x72, 0x64, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x64,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x38,
	0x0a, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x8b, 0x01, 0x0a, 0x03, 0x53, 0x32, 0x43,
	0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x54, 0x72, 0x64, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72,
	0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x55, 0x0a, 0x13, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x49, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x54,
	0x72, 0x64, 0x5f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x61, 0x74, 0x69,
	0x6f, 0x2e, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x13, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x49, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x03, 0x63, 0x32, 0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x54, 0x72, 0x64, 0x5f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x61,
	0x74, 0x69, 0x6f, 0x2e, 0x43, 0x32, 0x53, 0x52, 0x03, 0x63, 0x32, 0x73, 0x22, 0x87, 0x01, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x3a, 0x04, 0x2d, 0x34, 0x30, 0x30,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74,
	0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x03, 0x73,
	0x32, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x54, 0x72, 0x64, 0x5f, 0x47,
	0x65, 0x74, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x2e, 0x53, 0x32,
	0x43, 0x52, 0x03, 0x73, 0x32, 0x63, 0x42, 0x4c, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75,
	0x74, 0x75, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x72, 0x69, 0x73, 0x68,
	0x65, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x66, 0x75, 0x74, 0x75, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x62, 0x2f, 0x74, 0x72, 0x64, 0x67, 0x65, 0x74, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x72,
	0x61, 0x74, 0x69, 0x6f,
}

var (
	file_Trd_GetMarginRatio_proto_rawDescOnce sync.Once
	file_Trd_GetMarginRatio_proto_rawDescData = file_Trd_GetMarginRatio_proto_rawDesc
)

func file_Trd_GetMarginRatio_proto_rawDescGZIP() []byte {
	file_Trd_GetMarginRatio_proto_rawDescOnce.Do(func() {
		file_Trd_GetMarginRatio_proto_rawDescData = protoimpl.X.CompressGZIP(file_Trd_GetMarginRatio_proto_rawDescData)
	})
	return file_Trd_GetMarginRatio_proto_rawDescData
}

var file_Trd_GetMarginRatio_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Trd_GetMarginRatio_proto_goTypes = []interface{}{
	(*MarginRatioInfo)(nil),     // 0: Trd_GetMarginRatio.MarginRatioInfo
	(*C2S)(nil),                 // 1: Trd_GetMarginRatio.C2S
	(*S2C)(nil),                 // 2: Trd_GetMarginRatio.S2C
	(*Request)(nil),             // 3: Trd_GetMarginRatio.Request
	(*Response)(nil),            // 4: Trd_GetMarginRatio.Response
	(*qotcommon.Security)(nil),  // 5: Qot_Common.Security
	(*trdcommon.TrdHeader)(nil), // 6: Trd_Common.TrdHeader
}
var file_Trd_GetMarginRatio_proto_depIdxs = []int32{
	5, // 0: Trd_GetMarginRatio.MarginRatioInfo.security:type_name -> Qot_Common.Security
	6, // 1: Trd_GetMarginRatio.C2S.header:type_name -> Trd_Common.TrdHeader
	5, // 2: Trd_GetMarginRatio.C2S.securityList:type_name -> Qot_Common.Security
	6, // 3: Trd_GetMarginRatio.S2C.header:type_name -> Trd_Common.TrdHeader
	0, // 4: Trd_GetMarginRatio.S2C.marginRatioInfoList:type_name -> Trd_GetMarginRatio.MarginRatioInfo
	1, // 5: Trd_GetMarginRatio.Request.c2s:type_name -> Trd_GetMarginRatio.C2S
	2, // 6: Trd_GetMarginRatio.Response.s2c:type_name -> Trd_GetMarginRatio.S2C
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_Trd_GetMarginRatio_proto_init() }
func file_Trd_GetMarginRatio_proto_init() {
	if File_Trd_GetMarginRatio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Trd_GetMarginRatio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarginRatioInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Trd_GetMarginRatio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Trd_GetMarginRatio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Trd_GetMarginRatio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Trd_GetMarginRatio_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Trd_GetMarginRatio_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Trd_GetMarginRatio_proto_goTypes,
		DependencyIndexes: file_Trd_GetMarginRatio_proto_depIdxs,
		MessageInfos:      file_Trd_GetMarginRatio_proto_msgTypes,
	}.Build()
	File_Trd_GetMarginRatio_proto = out.File
	file_Trd_GetMarginRatio_proto_rawDesc = nil
	file_Trd_GetMarginRatio_proto_goTypes = nil
	file_Trd_GetMarginRatio_proto_depIdxs = nil
}
