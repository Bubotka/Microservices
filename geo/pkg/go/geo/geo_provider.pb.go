// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: api/protoc/user.proto

package geo

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type ListLevenshteinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Column string `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ListLevenshteinRequest) Reset() {
	*x = ListLevenshteinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protoc_geo_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLevenshteinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLevenshteinRequest) ProtoMessage() {}

func (x *ListLevenshteinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protoc_geo_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLevenshteinRequest.ProtoReflect.Descriptor instead.
func (*ListLevenshteinRequest) Descriptor() ([]byte, []int) {
	return file_api_protoc_geo_provider_proto_rawDescGZIP(), []int{0}
}

func (x *ListLevenshteinRequest) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *ListLevenshteinRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ListLevenshteinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sha *SearchHistoryAddress `protobuf:"bytes,1,opt,name=sha,proto3" json:"sha,omitempty"`
}

func (x *ListLevenshteinResponse) Reset() {
	*x = ListLevenshteinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protoc_geo_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLevenshteinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLevenshteinResponse) ProtoMessage() {}

func (x *ListLevenshteinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protoc_geo_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLevenshteinResponse.ProtoReflect.Descriptor instead.
func (*ListLevenshteinResponse) Descriptor() ([]byte, []int) {
	return file_api_protoc_geo_provider_proto_rawDescGZIP(), []int{1}
}

func (x *ListLevenshteinResponse) GetSha() *SearchHistoryAddress {
	if x != nil {
		return x.Sha
	}
	return nil
}

type SearchHistoryAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SearchRequest   string `protobuf:"bytes,2,opt,name=search_request,json=searchRequest,proto3" json:"search_request,omitempty"`
	AddressResponse string `protobuf:"bytes,3,opt,name=address_response,json=addressResponse,proto3" json:"address_response,omitempty"`
}

func (x *SearchHistoryAddress) Reset() {
	*x = SearchHistoryAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protoc_geo_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchHistoryAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchHistoryAddress) ProtoMessage() {}

func (x *SearchHistoryAddress) ProtoReflect() protoreflect.Message {
	mi := &file_api_protoc_geo_provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchHistoryAddress.ProtoReflect.Descriptor instead.
func (*SearchHistoryAddress) Descriptor() ([]byte, []int) {
	return file_api_protoc_geo_provider_proto_rawDescGZIP(), []int{2}
}

func (x *SearchHistoryAddress) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SearchHistoryAddress) GetSearchRequest() string {
	if x != nil {
		return x.SearchRequest
	}
	return ""
}

func (x *SearchHistoryAddress) GetAddressResponse() string {
	if x != nil {
		return x.AddressResponse
	}
	return ""
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Place string `protobuf:"bytes,1,opt,name=place,proto3" json:"place,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protoc_geo_provider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protoc_geo_provider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_api_protoc_geo_provider_proto_rawDescGZIP(), []int{3}
}

func (x *SearchRequest) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sha *SearchHistoryAddress `protobuf:"bytes,1,opt,name=sha,proto3" json:"sha,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protoc_geo_provider_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protoc_geo_provider_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_api_protoc_geo_provider_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRequest) GetSha() *SearchHistoryAddress {
	if x != nil {
		return x.Sha
	}
	return nil
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elements []*AddressElementSearch `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protoc_geo_provider_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protoc_geo_provider_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_api_protoc_geo_provider_proto_rawDescGZIP(), []int{5}
}

func (x *SearchResponse) GetElements() []*AddressElementSearch {
	if x != nil {
		return x.Elements
	}
	return nil
}

type GeoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coordinates string `protobuf:"bytes,1,opt,name=coordinates,proto3" json:"coordinates,omitempty"`
}

func (x *GeoRequest) Reset() {
	*x = GeoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protoc_geo_provider_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoRequest) ProtoMessage() {}

func (x *GeoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protoc_geo_provider_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoRequest.ProtoReflect.Descriptor instead.
func (*GeoRequest) Descriptor() ([]byte, []int) {
	return file_api_protoc_geo_provider_proto_rawDescGZIP(), []int{6}
}

func (x *GeoRequest) GetCoordinates() string {
	if x != nil {
		return x.Coordinates
	}
	return ""
}

type GeoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elements []*AddressElementSearch `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements,omitempty"`
}

func (x *GeoResponse) Reset() {
	*x = GeoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protoc_geo_provider_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoResponse) ProtoMessage() {}

func (x *GeoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protoc_geo_provider_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoResponse.ProtoReflect.Descriptor instead.
func (*GeoResponse) Descriptor() ([]byte, []int) {
	return file_api_protoc_geo_provider_proto_rawDescGZIP(), []int{7}
}

func (x *GeoResponse) GetElements() []*AddressElementSearch {
	if x != nil {
		return x.Elements
	}
	return nil
}

type AddressElementSearch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	GeoLat string `protobuf:"bytes,2,opt,name=geo_lat,json=geoLat,proto3" json:"geo_lat,omitempty"`
	GeoLon string `protobuf:"bytes,3,opt,name=geo_lon,json=geoLon,proto3" json:"geo_lon,omitempty"`
}

func (x *AddressElementSearch) Reset() {
	*x = AddressElementSearch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protoc_geo_provider_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressElementSearch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressElementSearch) ProtoMessage() {}

func (x *AddressElementSearch) ProtoReflect() protoreflect.Message {
	mi := &file_api_protoc_geo_provider_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressElementSearch.ProtoReflect.Descriptor instead.
func (*AddressElementSearch) Descriptor() ([]byte, []int) {
	return file_api_protoc_geo_provider_proto_rawDescGZIP(), []int{8}
}

func (x *AddressElementSearch) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *AddressElementSearch) GetGeoLat() string {
	if x != nil {
		return x.GeoLat
	}
	return ""
}

func (x *AddressElementSearch) GetGeoLon() string {
	if x != nil {
		return x.GeoLon
	}
	return ""
}

var File_api_protoc_geo_provider_proto protoreflect.FileDescriptor

var file_api_protoc_geo_provider_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x67, 0x65, 0x6f,
	0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x67, 0x65, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x16, 0x4c, 0x69,
	0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6e, 0x73, 0x68, 0x74, 0x65, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x22, 0x4f, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6e, 0x73, 0x68, 0x74,
	0x65, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x03, 0x73,
	0x68, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x65, 0x6f, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x03, 0x73, 0x68,
	0x61, 0x22, 0x78, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x29, 0x0a, 0x10, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x22, 0x45, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x03, 0x73, 0x68, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x67, 0x65, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x03, 0x73, 0x68, 0x61, 0x22, 0x50, 0x0a, 0x0e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x67, 0x65, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x2e, 0x0a, 0x0a, 0x47,
	0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x22, 0x4d, 0x0a, 0x0b, 0x47,
	0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67,
	0x65, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x60, 0x0a, 0x14, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x65,
	0x6f, 0x5f, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6f,
	0x4c, 0x61, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x65, 0x6f, 0x5f, 0x6c, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6f, 0x4c, 0x6f, 0x6e, 0x32, 0xb9, 0x02, 0x0a,
	0x0b, 0x47, 0x65, 0x6f, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x60, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6e, 0x73, 0x68, 0x74, 0x65, 0x69, 0x6e, 0x12,
	0x24, 0x2e, 0x67, 0x65, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6e, 0x73, 0x68, 0x74, 0x65, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67, 0x65, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6e, 0x73, 0x68,
	0x74, 0x65, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x67, 0x65, 0x6f, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x67, 0x65, 0x6f, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x65, 0x6f, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x2e, 0x67, 0x65, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x65,
	0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2e, 0x2f, 0x2e,
	0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protoc_geo_provider_proto_rawDescOnce sync.Once
	file_api_protoc_geo_provider_proto_rawDescData = file_api_protoc_geo_provider_proto_rawDesc
)

func file_api_protoc_geo_provider_proto_rawDescGZIP() []byte {
	file_api_protoc_geo_provider_proto_rawDescOnce.Do(func() {
		file_api_protoc_geo_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protoc_geo_provider_proto_rawDescData)
	})
	return file_api_protoc_geo_provider_proto_rawDescData
}

var file_api_protoc_geo_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_protoc_geo_provider_proto_goTypes = []interface{}{
	(*ListLevenshteinRequest)(nil),  // 0: geo_provider.ListLevenshteinRequest
	(*ListLevenshteinResponse)(nil), // 1: geo_provider.ListLevenshteinResponse
	(*SearchHistoryAddress)(nil),    // 2: geo_provider.SearchHistoryAddress
	(*SearchRequest)(nil),           // 3: geo_provider.SearchRequest
	(*CreateRequest)(nil),           // 4: geo_provider.CreateRequest
	(*SearchResponse)(nil),          // 5: geo_provider.SearchResponse
	(*GeoRequest)(nil),              // 6: geo_provider.GeoRequest
	(*GeoResponse)(nil),             // 7: geo_provider.GeoResponse
	(*AddressElementSearch)(nil),    // 8: geo_provider.AddressElementSearch
	(*empty.Empty)(nil),             // 9: google.protobuf.Empty
}
var file_api_protoc_geo_provider_proto_depIdxs = []int32{
	2, // 0: geo_provider.ListLevenshteinResponse.sha:type_name -> geo_provider.SearchHistoryAddress
	2, // 1: geo_provider.CreateRequest.sha:type_name -> geo_provider.SearchHistoryAddress
	8, // 2: geo_provider.SearchResponse.elements:type_name -> geo_provider.AddressElementSearch
	8, // 3: geo_provider.GeoResponse.elements:type_name -> geo_provider.AddressElementSearch
	0, // 4: geo_provider.GeoProvider.ListLevenshtein:input_type -> geo_provider.ListLevenshteinRequest
	4, // 5: geo_provider.GeoProvider.Create:input_type -> geo_provider.CreateRequest
	3, // 6: geo_provider.GeoProvider.Search:input_type -> geo_provider.SearchRequest
	6, // 7: geo_provider.GeoProvider.GeoCode:input_type -> geo_provider.GeoRequest
	1, // 8: geo_provider.GeoProvider.ListLevenshtein:output_type -> geo_provider.ListLevenshteinResponse
	9, // 9: geo_provider.GeoProvider.Create:output_type -> google.protobuf.Empty
	5, // 10: geo_provider.GeoProvider.Search:output_type -> geo_provider.SearchResponse
	7, // 11: geo_provider.GeoProvider.GeoCode:output_type -> geo_provider.GeoResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_protoc_geo_provider_proto_init() }
func file_api_protoc_geo_provider_proto_init() {
	if File_api_protoc_geo_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_protoc_geo_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLevenshteinRequest); i {
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
		file_api_protoc_geo_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLevenshteinResponse); i {
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
		file_api_protoc_geo_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchHistoryAddress); i {
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
		file_api_protoc_geo_provider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_api_protoc_geo_provider_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_api_protoc_geo_provider_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_api_protoc_geo_provider_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoRequest); i {
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
		file_api_protoc_geo_provider_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoResponse); i {
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
		file_api_protoc_geo_provider_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressElementSearch); i {
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
			RawDescriptor: file_api_protoc_geo_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_protoc_geo_provider_proto_goTypes,
		DependencyIndexes: file_api_protoc_geo_provider_proto_depIdxs,
		MessageInfos:      file_api_protoc_geo_provider_proto_msgTypes,
	}.Build()
	File_api_protoc_geo_provider_proto = out.File
	file_api_protoc_geo_provider_proto_rawDesc = nil
	file_api_protoc_geo_provider_proto_goTypes = nil
	file_api_protoc_geo_provider_proto_depIdxs = nil
}
