// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: proto/weather.proto

package pb

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

type RequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	City string `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *RequestData) Reset() {
	*x = RequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestData) ProtoMessage() {}

func (x *RequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestData.ProtoReflect.Descriptor instead.
func (*RequestData) Descriptor() ([]byte, []int) {
	return file_proto_weather_proto_rawDescGZIP(), []int{0}
}

func (x *RequestData) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type ResponseBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *ResponseBody_LocationBody `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Current  *ResponseBody_Current      `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
}

func (x *ResponseBody) Reset() {
	*x = ResponseBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weather_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseBody) ProtoMessage() {}

func (x *ResponseBody) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weather_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseBody.ProtoReflect.Descriptor instead.
func (*ResponseBody) Descriptor() ([]byte, []int) {
	return file_proto_weather_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseBody) GetLocation() *ResponseBody_LocationBody {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *ResponseBody) GetCurrent() *ResponseBody_Current {
	if x != nil {
		return x.Current
	}
	return nil
}

type ResponseBody_LocationBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region    string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Country   string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Localtime string `protobuf:"bytes,4,opt,name=localtime,proto3" json:"localtime,omitempty"`
}

func (x *ResponseBody_LocationBody) Reset() {
	*x = ResponseBody_LocationBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weather_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseBody_LocationBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseBody_LocationBody) ProtoMessage() {}

func (x *ResponseBody_LocationBody) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weather_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseBody_LocationBody.ProtoReflect.Descriptor instead.
func (*ResponseBody_LocationBody) Descriptor() ([]byte, []int) {
	return file_proto_weather_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ResponseBody_LocationBody) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *ResponseBody_LocationBody) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResponseBody_LocationBody) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *ResponseBody_LocationBody) GetLocaltime() string {
	if x != nil {
		return x.Localtime
	}
	return ""
}

type ResponseBody_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TempC float32 `protobuf:"fixed32,1,opt,name=temp_c,json=tempC,proto3" json:"temp_c,omitempty"`
	TempF float32 `protobuf:"fixed32,2,opt,name=temp_f,json=tempF,proto3" json:"temp_f,omitempty"`
}

func (x *ResponseBody_Current) Reset() {
	*x = ResponseBody_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weather_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseBody_Current) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseBody_Current) ProtoMessage() {}

func (x *ResponseBody_Current) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weather_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseBody_Current.ProtoReflect.Descriptor instead.
func (*ResponseBody_Current) Descriptor() ([]byte, []int) {
	return file_proto_weather_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ResponseBody_Current) GetTempC() float32 {
	if x != nil {
		return x.TempC
	}
	return 0
}

func (x *ResponseBody_Current) GetTempF() float32 {
	if x != nil {
		return x.TempF
	}
	return 0
}

var File_proto_weather_proto protoreflect.FileDescriptor

var file_proto_weather_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x21, 0x0a, 0x0b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x22, 0xae,
	0x02, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12,
	0x3b, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x6f, 0x64, 0x79, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f,
	0x64, 0x79, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64,
	0x79, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x1a, 0x72, 0x0a, 0x0c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f,
	0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x37, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x74, 0x65, 0x6d, 0x70, 0x43, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x65, 0x6d, 0x70,
	0x5f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x65, 0x6d, 0x70, 0x46, 0x32,
	0x4b, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x39, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_weather_proto_rawDescOnce sync.Once
	file_proto_weather_proto_rawDescData = file_proto_weather_proto_rawDesc
)

func file_proto_weather_proto_rawDescGZIP() []byte {
	file_proto_weather_proto_rawDescOnce.Do(func() {
		file_proto_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_weather_proto_rawDescData)
	})
	return file_proto_weather_proto_rawDescData
}

var file_proto_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_weather_proto_goTypes = []interface{}{
	(*RequestData)(nil),               // 0: main.RequestData
	(*ResponseBody)(nil),              // 1: main.ResponseBody
	(*ResponseBody_LocationBody)(nil), // 2: main.ResponseBody.LocationBody
	(*ResponseBody_Current)(nil),      // 3: main.ResponseBody.Current
}
var file_proto_weather_proto_depIdxs = []int32{
	2, // 0: main.ResponseBody.location:type_name -> main.ResponseBody.LocationBody
	3, // 1: main.ResponseBody.current:type_name -> main.ResponseBody.Current
	0, // 2: main.WeatherService.WeatherRequest:input_type -> main.RequestData
	1, // 3: main.WeatherService.WeatherRequest:output_type -> main.ResponseBody
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_weather_proto_init() }
func file_proto_weather_proto_init() {
	if File_proto_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_weather_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestData); i {
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
		file_proto_weather_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseBody); i {
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
		file_proto_weather_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseBody_LocationBody); i {
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
		file_proto_weather_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseBody_Current); i {
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
			RawDescriptor: file_proto_weather_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_weather_proto_goTypes,
		DependencyIndexes: file_proto_weather_proto_depIdxs,
		MessageInfos:      file_proto_weather_proto_msgTypes,
	}.Build()
	File_proto_weather_proto = out.File
	file_proto_weather_proto_rawDesc = nil
	file_proto_weather_proto_goTypes = nil
	file_proto_weather_proto_depIdxs = nil
}
