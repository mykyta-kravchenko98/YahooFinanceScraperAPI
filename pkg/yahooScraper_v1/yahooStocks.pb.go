// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: api/yahooScraper_v1/yahooStocks.proto

package yahooscraper_v1

import (
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

type GetAllValidStocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllValidStocksRequest) Reset() {
	*x = GetAllValidStocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_yahooScraper_v1_yahooStocks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllValidStocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllValidStocksRequest) ProtoMessage() {}

func (x *GetAllValidStocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_yahooScraper_v1_yahooStocks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllValidStocksRequest.ProtoReflect.Descriptor instead.
func (*GetAllValidStocksRequest) Descriptor() ([]byte, []int) {
	return file_api_yahooScraper_v1_yahooStocks_proto_rawDescGZIP(), []int{0}
}

type GetAllValidStocksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stocks []*Stocks `protobuf:"bytes,1,rep,name=stocks,proto3" json:"stocks,omitempty"`
}

func (x *GetAllValidStocksResponse) Reset() {
	*x = GetAllValidStocksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_yahooScraper_v1_yahooStocks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllValidStocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllValidStocksResponse) ProtoMessage() {}

func (x *GetAllValidStocksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_yahooScraper_v1_yahooStocks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllValidStocksResponse.ProtoReflect.Descriptor instead.
func (*GetAllValidStocksResponse) Descriptor() ([]byte, []int) {
	return file_api_yahooScraper_v1_yahooStocks_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllValidStocksResponse) GetStocks() []*Stocks {
	if x != nil {
		return x.Stocks
	}
	return nil
}

type Stocks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol              string  `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name                string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price               float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Change              float64 `protobuf:"fixed64,4,opt,name=change,proto3" json:"change,omitempty"`
	PercentChange       float64 `protobuf:"fixed64,5,opt,name=percent_change,json=percentChange,proto3" json:"percent_change,omitempty"`
	Volume              string  `protobuf:"bytes,6,opt,name=volume,proto3" json:"volume,omitempty"`
	MarketCap           string  `protobuf:"bytes,7,opt,name=market_cap,json=marketCap,proto3" json:"market_cap,omitempty"`
	PeRatio             string  `protobuf:"bytes,8,opt,name=pe_ratio,json=peRatio,proto3" json:"pe_ratio,omitempty"`
	CreatedAt           string  `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedUnix         int64   `protobuf:"varint,10,opt,name=created_unix,json=createdUnix,proto3" json:"created_unix,omitempty"`
	AvgVolumeFor_3Month string  `protobuf:"bytes,11,opt,name=avg_volume_for_3_month,json=avgVolumeFor3Month,proto3" json:"avg_volume_for_3_month,omitempty"`
}

func (x *Stocks) Reset() {
	*x = Stocks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_yahooScraper_v1_yahooStocks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stocks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stocks) ProtoMessage() {}

func (x *Stocks) ProtoReflect() protoreflect.Message {
	mi := &file_api_yahooScraper_v1_yahooStocks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stocks.ProtoReflect.Descriptor instead.
func (*Stocks) Descriptor() ([]byte, []int) {
	return file_api_yahooScraper_v1_yahooStocks_proto_rawDescGZIP(), []int{2}
}

func (x *Stocks) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Stocks) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Stocks) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Stocks) GetChange() float64 {
	if x != nil {
		return x.Change
	}
	return 0
}

func (x *Stocks) GetPercentChange() float64 {
	if x != nil {
		return x.PercentChange
	}
	return 0
}

func (x *Stocks) GetVolume() string {
	if x != nil {
		return x.Volume
	}
	return ""
}

func (x *Stocks) GetMarketCap() string {
	if x != nil {
		return x.MarketCap
	}
	return ""
}

func (x *Stocks) GetPeRatio() string {
	if x != nil {
		return x.PeRatio
	}
	return ""
}

func (x *Stocks) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Stocks) GetCreatedUnix() int64 {
	if x != nil {
		return x.CreatedUnix
	}
	return 0
}

func (x *Stocks) GetAvgVolumeFor_3Month() string {
	if x != nil {
		return x.AvgVolumeFor_3Month
	}
	return ""
}

var File_api_yahooScraper_v1_yahooStocks_proto protoreflect.FileDescriptor

var file_api_yahooScraper_v1_yahooStocks_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x79, 0x61, 0x68, 0x6f, 0x6f, 0x53, 0x63, 0x72, 0x61, 0x70,
	0x65, 0x72, 0x5f, 0x76, 0x31, 0x2f, 0x79, 0x61, 0x68, 0x6f, 0x6f, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x79, 0x61, 0x68, 0x6f, 0x6f, 0x73, 0x63,
	0x72, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x79, 0x61, 0x68, 0x6f, 0x6f, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72,
	0x5f, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x73, 0x22, 0xd1, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0d, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x5f, 0x63, 0x61, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x43, 0x61, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6f,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x6e, 0x69, 0x78, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x55, 0x6e,
	0x69, 0x78, 0x12, 0x32, 0x0a, 0x16, 0x61, 0x76, 0x67, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x33, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x61, 0x76, 0x67, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x46, 0x6f, 0x72,
	0x33, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x32, 0x80, 0x01, 0x0a, 0x12, 0x59, 0x61, 0x68, 0x6f, 0x6f,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x73, 0x12, 0x29, 0x2e, 0x79, 0x61, 0x68, 0x6f, 0x6f, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65,
	0x72, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x79, 0x61, 0x68, 0x6f, 0x6f, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x79, 0x61, 0x68,
	0x6f, 0x6f, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x79, 0x61,
	0x68, 0x6f, 0x6f, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x3b, 0x79, 0x61,
	0x68, 0x6f, 0x6f, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_yahooScraper_v1_yahooStocks_proto_rawDescOnce sync.Once
	file_api_yahooScraper_v1_yahooStocks_proto_rawDescData = file_api_yahooScraper_v1_yahooStocks_proto_rawDesc
)

func file_api_yahooScraper_v1_yahooStocks_proto_rawDescGZIP() []byte {
	file_api_yahooScraper_v1_yahooStocks_proto_rawDescOnce.Do(func() {
		file_api_yahooScraper_v1_yahooStocks_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_yahooScraper_v1_yahooStocks_proto_rawDescData)
	})
	return file_api_yahooScraper_v1_yahooStocks_proto_rawDescData
}

var file_api_yahooScraper_v1_yahooStocks_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_yahooScraper_v1_yahooStocks_proto_goTypes = []interface{}{
	(*GetAllValidStocksRequest)(nil),  // 0: yahooscraper_v1.GetAllValidStocksRequest
	(*GetAllValidStocksResponse)(nil), // 1: yahooscraper_v1.GetAllValidStocksResponse
	(*Stocks)(nil),                    // 2: yahooscraper_v1.Stocks
}
var file_api_yahooScraper_v1_yahooStocks_proto_depIdxs = []int32{
	2, // 0: yahooscraper_v1.GetAllValidStocksResponse.stocks:type_name -> yahooscraper_v1.Stocks
	0, // 1: yahooscraper_v1.YahooStocksService.GetAllValidStocks:input_type -> yahooscraper_v1.GetAllValidStocksRequest
	1, // 2: yahooscraper_v1.YahooStocksService.GetAllValidStocks:output_type -> yahooscraper_v1.GetAllValidStocksResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_yahooScraper_v1_yahooStocks_proto_init() }
func file_api_yahooScraper_v1_yahooStocks_proto_init() {
	if File_api_yahooScraper_v1_yahooStocks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_yahooScraper_v1_yahooStocks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllValidStocksRequest); i {
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
		file_api_yahooScraper_v1_yahooStocks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllValidStocksResponse); i {
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
		file_api_yahooScraper_v1_yahooStocks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stocks); i {
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
			RawDescriptor: file_api_yahooScraper_v1_yahooStocks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_yahooScraper_v1_yahooStocks_proto_goTypes,
		DependencyIndexes: file_api_yahooScraper_v1_yahooStocks_proto_depIdxs,
		MessageInfos:      file_api_yahooScraper_v1_yahooStocks_proto_msgTypes,
	}.Build()
	File_api_yahooScraper_v1_yahooStocks_proto = out.File
	file_api_yahooScraper_v1_yahooStocks_proto_rawDesc = nil
	file_api_yahooScraper_v1_yahooStocks_proto_goTypes = nil
	file_api_yahooScraper_v1_yahooStocks_proto_depIdxs = nil
}
