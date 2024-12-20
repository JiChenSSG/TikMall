// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: cart_f.proto

package cart

import (
	_ "github.com/jichenssg/tikmall/app/gateway/biz/model/api"
	common "github.com/jichenssg/tikmall/app/gateway/biz/model/frontend/common"
	product "github.com/jichenssg/tikmall/app/gateway/biz/model/frontend/product"
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

type RespCartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product  *product.Product `protobuf:"bytes,1,opt,name=product,proto3" form:"product" json:"product,omitempty" query:"product"`
	Quantity int32            `protobuf:"varint,2,opt,name=quantity,proto3" form:"quantity" json:"quantity,omitempty" query:"quantity"`
}

func (x *RespCartItem) Reset() {
	*x = RespCartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_f_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespCartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespCartItem) ProtoMessage() {}

func (x *RespCartItem) ProtoReflect() protoreflect.Message {
	mi := &file_cart_f_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespCartItem.ProtoReflect.Descriptor instead.
func (*RespCartItem) Descriptor() ([]byte, []int) {
	return file_cart_f_proto_rawDescGZIP(), []int{0}
}

func (x *RespCartItem) GetProduct() *product.Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *RespCartItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type AddItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64 `protobuf:"varint,2,opt,name=product_id,proto3" form:"product_id" json:"product_id,omitempty" query:"product_id"`
	Quantity  int32 `protobuf:"varint,3,opt,name=quantity,proto3" form:"quantity" json:"quantity,omitempty" query:"quantity"`
}

func (x *AddItemReq) Reset() {
	*x = AddItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_f_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemReq) ProtoMessage() {}

func (x *AddItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_f_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemReq.ProtoReflect.Descriptor instead.
func (*AddItemReq) Descriptor() ([]byte, []int) {
	return file_cart_f_proto_rawDescGZIP(), []int{1}
}

func (x *AddItemReq) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *AddItemReq) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type AddItemResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" form:"message" json:"message,omitempty" query:"message"`
	Id      int64  `protobuf:"varint,2,opt,name=id,proto3" form:"id" json:"id,omitempty" query:"id"`
}

func (x *AddItemResp) Reset() {
	*x = AddItemResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_f_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemResp) ProtoMessage() {}

func (x *AddItemResp) ProtoReflect() protoreflect.Message {
	mi := &file_cart_f_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemResp.ProtoReflect.Descriptor instead.
func (*AddItemResp) Descriptor() ([]byte, []int) {
	return file_cart_f_proto_rawDescGZIP(), []int{2}
}

func (x *AddItemResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AddItemResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCartResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" form:"message" json:"message,omitempty" query:"message"`
	Cart    *Cart  `protobuf:"bytes,2,opt,name=cart,proto3" form:"cart" json:"cart,omitempty" query:"cart"`
}

func (x *GetCartResp) Reset() {
	*x = GetCartResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_f_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartResp) ProtoMessage() {}

func (x *GetCartResp) ProtoReflect() protoreflect.Message {
	mi := &file_cart_f_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartResp.ProtoReflect.Descriptor instead.
func (*GetCartResp) Descriptor() ([]byte, []int) {
	return file_cart_f_proto_rawDescGZIP(), []int{3}
}

func (x *GetCartResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetCartResp) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64           `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" form:"user_id" json:"user_id,omitempty" query:"user_id"`
	Items  []*RespCartItem `protobuf:"bytes,2,rep,name=items,proto3" form:"items" json:"items,omitempty" query:"items"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_f_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_cart_f_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_cart_f_proto_rawDescGZIP(), []int{4}
}

func (x *Cart) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Cart) GetItems() []*RespCartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_cart_f_proto protoreflect.FileDescriptor

var file_cart_f_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x1a, 0x12, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x66, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0x48, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x37, 0x0a,
	0x0b, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x50, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x27, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x52, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xf5, 0x01, 0x0a,
	0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x19, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x22, 0x09,
	0xd2, 0xc1, 0x18, 0x05, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x12, 0x48, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x09, 0xca, 0xc1, 0x18, 0x05, 0x2f, 0x63,
	0x61, 0x72, 0x74, 0x12, 0x4f, 0x0a, 0x09, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x72, 0x74,
	0x12, 0x16, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x0f, 0xd2, 0xc1, 0x18, 0x0b, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6a, 0x69, 0x63, 0x68, 0x65, 0x6e, 0x73, 0x73, 0x67, 0x2f, 0x74, 0x69, 0x6b,
	0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cart_f_proto_rawDescOnce sync.Once
	file_cart_f_proto_rawDescData = file_cart_f_proto_rawDesc
)

func file_cart_f_proto_rawDescGZIP() []byte {
	file_cart_f_proto_rawDescOnce.Do(func() {
		file_cart_f_proto_rawDescData = protoimpl.X.CompressGZIP(file_cart_f_proto_rawDescData)
	})
	return file_cart_f_proto_rawDescData
}

var file_cart_f_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cart_f_proto_goTypes = []interface{}{
	(*RespCartItem)(nil),    // 0: frontend.cart.RespCartItem
	(*AddItemReq)(nil),      // 1: frontend.cart.AddItemReq
	(*AddItemResp)(nil),     // 2: frontend.cart.AddItemResp
	(*GetCartResp)(nil),     // 3: frontend.cart.GetCartResp
	(*Cart)(nil),            // 4: frontend.cart.Cart
	(*product.Product)(nil), // 5: frontend.product.Product
	(*common.Empty)(nil),    // 6: frontend.common.Empty
	(*common.Response)(nil), // 7: frontend.common.Response
}
var file_cart_f_proto_depIdxs = []int32{
	5, // 0: frontend.cart.RespCartItem.product:type_name -> frontend.product.Product
	4, // 1: frontend.cart.GetCartResp.cart:type_name -> frontend.cart.Cart
	0, // 2: frontend.cart.Cart.items:type_name -> frontend.cart.RespCartItem
	1, // 3: frontend.cart.CartService.AddItem:input_type -> frontend.cart.AddItemReq
	6, // 4: frontend.cart.CartService.GetCart:input_type -> frontend.common.Empty
	6, // 5: frontend.cart.CartService.EmptyCart:input_type -> frontend.common.Empty
	2, // 6: frontend.cart.CartService.AddItem:output_type -> frontend.cart.AddItemResp
	3, // 7: frontend.cart.CartService.GetCart:output_type -> frontend.cart.GetCartResp
	7, // 8: frontend.cart.CartService.EmptyCart:output_type -> frontend.common.Response
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cart_f_proto_init() }
func file_cart_f_proto_init() {
	if File_cart_f_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cart_f_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespCartItem); i {
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
		file_cart_f_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemReq); i {
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
		file_cart_f_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemResp); i {
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
		file_cart_f_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartResp); i {
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
		file_cart_f_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart); i {
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
			RawDescriptor: file_cart_f_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_f_proto_goTypes,
		DependencyIndexes: file_cart_f_proto_depIdxs,
		MessageInfos:      file_cart_f_proto_msgTypes,
	}.Build()
	File_cart_f_proto = out.File
	file_cart_f_proto_rawDesc = nil
	file_cart_f_proto_goTypes = nil
	file_cart_f_proto_depIdxs = nil
}
