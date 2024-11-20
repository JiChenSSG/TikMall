// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: order_f.proto

package order

import (
	_ "github.com/jichenssg/tikmall/app/gateway/biz/model/api"
	common "github.com/jichenssg/tikmall/app/gateway/biz/model/frontend/common"
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

type CancelOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" form:"order_id,required" json:"order_id,required,omitempty"`
}

func (x *CancelOrderReq) Reset() {
	*x = CancelOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_f_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderReq) ProtoMessage() {}

func (x *CancelOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_f_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderReq.ProtoReflect.Descriptor instead.
func (*CancelOrderReq) Descriptor() ([]byte, []int) {
	return file_order_f_proto_rawDescGZIP(), []int{0}
}

func (x *CancelOrderReq) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type PlaceOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreetAddress string `protobuf:"bytes,6,opt,name=street_address,json=streetAddress,proto3" form:"street_address,required" json:"street_address,required,omitempty"`
	City          string `protobuf:"bytes,2,opt,name=city,proto3" form:"city,required" json:"city,required,omitempty"`
	State         string `protobuf:"bytes,3,opt,name=state,proto3" form:"state,required" json:"state,required,omitempty"`
	Country       string `protobuf:"bytes,4,opt,name=country,proto3" form:"country,required" json:"country,required,omitempty"`
	ZipCode       int32  `protobuf:"varint,5,opt,name=zip_code,json=zipCode,proto3" form:"zip_code,required" json:"zip_code,required,omitempty"`
	Email         string `protobuf:"bytes,8,opt,name=email,proto3" form:"email,required" json:"email,required,omitempty"`
}

func (x *PlaceOrderReq) Reset() {
	*x = PlaceOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_f_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderReq) ProtoMessage() {}

func (x *PlaceOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_f_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderReq.ProtoReflect.Descriptor instead.
func (*PlaceOrderReq) Descriptor() ([]byte, []int) {
	return file_order_f_proto_rawDescGZIP(), []int{1}
}

func (x *PlaceOrderReq) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *PlaceOrderReq) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *PlaceOrderReq) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *PlaceOrderReq) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *PlaceOrderReq) GetZipCode() int32 {
	if x != nil {
		return x.ZipCode
	}
	return 0
}

func (x *PlaceOrderReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UpdateOrderInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId       string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" form:"order_id,required" json:"order_id,required,omitempty"`
	Email         string `protobuf:"bytes,7,opt,name=email,proto3" form:"email" json:"email,omitempty"`
	StreetAddress string `protobuf:"bytes,6,opt,name=street_address,json=streetAddress,proto3" form:"street_address" json:"street_address,omitempty"`
	City          string `protobuf:"bytes,2,opt,name=city,proto3" form:"city" json:"city,omitempty"`
	State         string `protobuf:"bytes,3,opt,name=state,proto3" form:"state" json:"state,omitempty"`
	Country       string `protobuf:"bytes,4,opt,name=country,proto3" form:"country" json:"country,omitempty"`
	ZipCode       int32  `protobuf:"varint,5,opt,name=zip_code,json=zipCode,proto3" form:"zip_code" json:"zip_code,omitempty"`
}

func (x *UpdateOrderInfoReq) Reset() {
	*x = UpdateOrderInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_f_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderInfoReq) ProtoMessage() {}

func (x *UpdateOrderInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_f_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderInfoReq.ProtoReflect.Descriptor instead.
func (*UpdateOrderInfoReq) Descriptor() ([]byte, []int) {
	return file_order_f_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateOrderInfoReq) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *UpdateOrderInfoReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateOrderInfoReq) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *UpdateOrderInfoReq) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UpdateOrderInfoReq) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *UpdateOrderInfoReq) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *UpdateOrderInfoReq) GetZipCode() int32 {
	if x != nil {
		return x.ZipCode
	}
	return 0
}

type GetOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty" query:"order_id,required"`
}

func (x *GetOrderReq) Reset() {
	*x = GetOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_f_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderReq) ProtoMessage() {}

func (x *GetOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_f_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderReq.ProtoReflect.Descriptor instead.
func (*GetOrderReq) Descriptor() ([]byte, []int) {
	return file_order_f_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderReq) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

var File_order_f_proto protoreflect.FileDescriptor

var file_order_f_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x1a,
	0x12, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0e, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x30, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15,
	0xca, 0xbb, 0x18, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0xb0,
	0x02, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x42, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xca, 0xbb, 0x18, 0x17, 0x73, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2c, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x11, 0xca, 0xbb, 0x18, 0x0d, 0x63, 0x69, 0x74, 0x79, 0x2c, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0xca, 0xbb, 0x18, 0x0e,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xca, 0xbb, 0x18, 0x10, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x2c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x30, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x15, 0xca, 0xbb, 0x18, 0x11, 0x7a, 0x69, 0x70,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x2c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x07,
	0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0xca, 0xbb, 0x18, 0x0e, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x2c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x22, 0xb1, 0x02, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x30, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xca, 0xbb, 0x18, 0x11,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xca, 0xbb, 0x18, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x39, 0x0a, 0x0e, 0x73,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x12, 0xca, 0xbb, 0x18, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xca, 0xbb, 0x18, 0x04, 0x63, 0x69, 0x74, 0x79, 0x52, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xca, 0xbb, 0x18, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xca, 0xbb, 0x18, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x08,
	0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0c,
	0xca, 0xbb, 0x18, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x7a, 0x69,
	0x70, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x30, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xb2, 0xbb, 0x18, 0x11, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x2c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x32, 0xfc, 0x02, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0xca, 0xc1, 0x18, 0x0b, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x59, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x11, 0xd2, 0xc1, 0x18, 0x0d, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x5b, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x1e, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x19, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0xd2,
	0xc1, 0x18, 0x0d, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x12, 0x63, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x22, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x11, 0xd2, 0xc1, 0x18, 0x0d, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x69, 0x63, 0x68, 0x65, 0x6e, 0x73, 0x73, 0x67, 0x2f, 0x74, 0x69,
	0x6b, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_order_f_proto_rawDescOnce sync.Once
	file_order_f_proto_rawDescData = file_order_f_proto_rawDesc
)

func file_order_f_proto_rawDescGZIP() []byte {
	file_order_f_proto_rawDescOnce.Do(func() {
		file_order_f_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_f_proto_rawDescData)
	})
	return file_order_f_proto_rawDescData
}

var file_order_f_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_order_f_proto_goTypes = []interface{}{
	(*CancelOrderReq)(nil),     // 0: frontend.order.CancelOrderReq
	(*PlaceOrderReq)(nil),      // 1: frontend.order.PlaceOrderReq
	(*UpdateOrderInfoReq)(nil), // 2: frontend.order.UpdateOrderInfoReq
	(*GetOrderReq)(nil),        // 3: frontend.order.GetOrderReq
	(*common.Empty)(nil),       // 4: frontend.common.Empty
	(*common.Response)(nil),    // 5: frontend.common.Response
}
var file_order_f_proto_depIdxs = []int32{
	4, // 0: frontend.order.OrderService.OrderList:input_type -> frontend.common.Empty
	1, // 1: frontend.order.OrderService.PlaceOrder:input_type -> frontend.order.PlaceOrderReq
	0, // 2: frontend.order.OrderService.CancelOrder:input_type -> frontend.order.CancelOrderReq
	2, // 3: frontend.order.OrderService.UpdateOrderInfo:input_type -> frontend.order.UpdateOrderInfoReq
	5, // 4: frontend.order.OrderService.OrderList:output_type -> frontend.common.Response
	5, // 5: frontend.order.OrderService.PlaceOrder:output_type -> frontend.common.Response
	5, // 6: frontend.order.OrderService.CancelOrder:output_type -> frontend.common.Response
	5, // 7: frontend.order.OrderService.UpdateOrderInfo:output_type -> frontend.common.Response
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_order_f_proto_init() }
func file_order_f_proto_init() {
	if File_order_f_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_f_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOrderReq); i {
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
		file_order_f_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceOrderReq); i {
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
		file_order_f_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderInfoReq); i {
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
		file_order_f_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderReq); i {
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
			RawDescriptor: file_order_f_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_f_proto_goTypes,
		DependencyIndexes: file_order_f_proto_depIdxs,
		MessageInfos:      file_order_f_proto_msgTypes,
	}.Build()
	File_order_f_proto = out.File
	file_order_f_proto_rawDesc = nil
	file_order_f_proto_goTypes = nil
	file_order_f_proto_depIdxs = nil
}
