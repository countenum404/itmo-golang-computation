// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/calculation.proto

package app

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Operation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Op            string                 `protobuf:"bytes,2,opt,name=op,proto3" json:"op,omitempty"`
	Var           string                 `protobuf:"bytes,3,opt,name=var,proto3" json:"var,omitempty"`
	Left          string                 `protobuf:"bytes,4,opt,name=left,proto3" json:"left,omitempty"`
	Right         string                 `protobuf:"bytes,5,opt,name=right,proto3" json:"right,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Operation) Reset() {
	*x = Operation{}
	mi := &file_proto_calculation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_proto_calculation_proto_rawDescGZIP(), []int{0}
}

func (x *Operation) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Operation) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *Operation) GetVar() string {
	if x != nil {
		return x.Var
	}
	return ""
}

func (x *Operation) GetLeft() string {
	if x != nil {
		return x.Left
	}
	return ""
}

func (x *Operation) GetRight() string {
	if x != nil {
		return x.Right
	}
	return ""
}

type PrintResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Var           string                 `protobuf:"bytes,1,opt,name=var,proto3" json:"var,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrintResult) Reset() {
	*x = PrintResult{}
	mi := &file_proto_calculation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrintResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintResult) ProtoMessage() {}

func (x *PrintResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintResult.ProtoReflect.Descriptor instead.
func (*PrintResult) Descriptor() ([]byte, []int) {
	return file_proto_calculation_proto_rawDescGZIP(), []int{1}
}

func (x *PrintResult) GetVar() string {
	if x != nil {
		return x.Var
	}
	return ""
}

func (x *PrintResult) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type CalcRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Operations    []*Operation           `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CalcRequest) Reset() {
	*x = CalcRequest{}
	mi := &file_proto_calculation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcRequest) ProtoMessage() {}

func (x *CalcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcRequest.ProtoReflect.Descriptor instead.
func (*CalcRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculation_proto_rawDescGZIP(), []int{2}
}

func (x *CalcRequest) GetOperations() []*Operation {
	if x != nil {
		return x.Operations
	}
	return nil
}

type SolutionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*PrintResult         `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SolutionResponse) Reset() {
	*x = SolutionResponse{}
	mi := &file_proto_calculation_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SolutionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolutionResponse) ProtoMessage() {}

func (x *SolutionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculation_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolutionResponse.ProtoReflect.Descriptor instead.
func (*SolutionResponse) Descriptor() ([]byte, []int) {
	return file_proto_calculation_proto_rawDescGZIP(), []int{3}
}

func (x *SolutionResponse) GetItems() []*PrintResult {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_proto_calculation_proto protoreflect.FileDescriptor

const file_proto_calculation_proto_rawDesc = "" +
	"\n" +
	"\x17proto/calculation.proto\x12\aservice\"k\n" +
	"\tOperation\x12\x12\n" +
	"\x04type\x18\x01 \x01(\tR\x04type\x12\x0e\n" +
	"\x02op\x18\x02 \x01(\tR\x02op\x12\x10\n" +
	"\x03var\x18\x03 \x01(\tR\x03var\x12\x12\n" +
	"\x04left\x18\x04 \x01(\tR\x04left\x12\x14\n" +
	"\x05right\x18\x05 \x01(\tR\x05right\"5\n" +
	"\vPrintResult\x12\x10\n" +
	"\x03var\x18\x01 \x01(\tR\x03var\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value\"A\n" +
	"\vCalcRequest\x122\n" +
	"\n" +
	"operations\x18\x01 \x03(\v2\x12.service.OperationR\n" +
	"operations\">\n" +
	"\x10SolutionResponse\x12*\n" +
	"\x05items\x18\x01 \x03(\v2\x14.service.PrintResultR\x05items2E\n" +
	"\tSolverSvc\x128\n" +
	"\x05Solve\x12\x14.service.CalcRequest\x1a\x19.service.SolutionResponseB\tZ\apkg/appb\x06proto3"

var (
	file_proto_calculation_proto_rawDescOnce sync.Once
	file_proto_calculation_proto_rawDescData []byte
)

func file_proto_calculation_proto_rawDescGZIP() []byte {
	file_proto_calculation_proto_rawDescOnce.Do(func() {
		file_proto_calculation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_calculation_proto_rawDesc), len(file_proto_calculation_proto_rawDesc)))
	})
	return file_proto_calculation_proto_rawDescData
}

var file_proto_calculation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_calculation_proto_goTypes = []any{
	(*Operation)(nil),        // 0: service.Operation
	(*PrintResult)(nil),      // 1: service.PrintResult
	(*CalcRequest)(nil),      // 2: service.CalcRequest
	(*SolutionResponse)(nil), // 3: service.SolutionResponse
}
var file_proto_calculation_proto_depIdxs = []int32{
	0, // 0: service.CalcRequest.operations:type_name -> service.Operation
	1, // 1: service.SolutionResponse.items:type_name -> service.PrintResult
	2, // 2: service.SolverSvc.Solve:input_type -> service.CalcRequest
	3, // 3: service.SolverSvc.Solve:output_type -> service.SolutionResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_calculation_proto_init() }
func file_proto_calculation_proto_init() {
	if File_proto_calculation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_calculation_proto_rawDesc), len(file_proto_calculation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_calculation_proto_goTypes,
		DependencyIndexes: file_proto_calculation_proto_depIdxs,
		MessageInfos:      file_proto_calculation_proto_msgTypes,
	}.Build()
	File_proto_calculation_proto = out.File
	file_proto_calculation_proto_goTypes = nil
	file_proto_calculation_proto_depIdxs = nil
}
