// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/imports/test_import_all.proto

package imports

import (
	fmt "github.com/bitquery/protobuf-sql/cmd/protoc-gen-go/testdata/imports/fmt"
	test_a_1 "github.com/bitquery/protobuf-sql/cmd/protoc-gen-go/testdata/imports/test_a_1"
	_ "github.com/bitquery/protobuf-sql/cmd/protoc-gen-go/testdata/imports/test_a_2"
	test_b_1 "github.com/bitquery/protobuf-sql/cmd/protoc-gen-go/testdata/imports/test_b_1"
	protoreflect "github.com/bitquery/protobuf-sql/reflect/protoreflect"
	protoimpl "github.com/bitquery/protobuf-sql/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type All struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Am1 *test_a_1.M1 `protobuf:"bytes,1,opt,name=am1,proto3" json:"am1,omitempty"`
	Am2 *test_a_1.M2 `protobuf:"bytes,2,opt,name=am2,proto3" json:"am2,omitempty"`
	Bm1 *test_b_1.M1 `protobuf:"bytes,5,opt,name=bm1,proto3" json:"bm1,omitempty"`
	Bm2 *test_b_1.M2 `protobuf:"bytes,6,opt,name=bm2,proto3" json:"bm2,omitempty"`
	Fmt *fmt.M       `protobuf:"bytes,7,opt,name=fmt,proto3" json:"fmt,omitempty"`
}

func (x *All) Reset() {
	*x = All{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *All) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*All) ProtoMessage() {}

func (x *All) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use All.ProtoReflect.Descriptor instead.
func (*All) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_rawDescGZIP(), []int{0}
}

func (x *All) GetAm1() *test_a_1.M1 {
	if x != nil {
		return x.Am1
	}
	return nil
}

func (x *All) GetAm2() *test_a_1.M2 {
	if x != nil {
		return x.Am2
	}
	return nil
}

func (x *All) GetBm1() *test_b_1.M1 {
	if x != nil {
		return x.Bm1
	}
	return nil
}

func (x *All) GetBm2() *test_b_1.M2 {
	if x != nil {
		return x.Bm2
	}
	return nil
}

func (x *All) GetFmt() *fmt.M {
	if x != nil {
		return x.Fmt
	}
	return nil
}

var File_cmd_protoc_gen_go_testdata_imports_test_import_all_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_rawDesc = []byte{
	0x0a, 0x38, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73, 0x74,
	0x1a, 0x34, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x5f, 0x31, 0x2f, 0x6d, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x61, 0x5f, 0x31, 0x2f, 0x6d, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6d,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x5f, 0x32, 0x2f, 0x6d, 0x33, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x5f, 0x32, 0x2f,
	0x6d, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x62, 0x5f, 0x31, 0x2f, 0x6d, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34,
	0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x5f, 0x31, 0x2f, 0x6d, 0x32, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x66, 0x6d, 0x74, 0x2f, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x12, 0x1c, 0x0a, 0x03,
	0x61, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x61, 0x2e, 0x4d, 0x31, 0x52, 0x03, 0x61, 0x6d, 0x31, 0x12, 0x1c, 0x0a, 0x03, 0x61, 0x6d,
	0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x61,
	0x2e, 0x4d, 0x32, 0x52, 0x03, 0x61, 0x6d, 0x32, 0x12, 0x22, 0x0a, 0x03, 0x62, 0x6d, 0x31, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x62, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x31, 0x2e, 0x4d, 0x31, 0x52, 0x03, 0x62, 0x6d, 0x31, 0x12, 0x22, 0x0a, 0x03,
	0x62, 0x6d, 0x32, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x62, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x32, 0x2e, 0x4d, 0x32, 0x52, 0x03, 0x62, 0x6d, 0x32,
	0x12, 0x18, 0x0a, 0x03, 0x66, 0x6d, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e,
	0x66, 0x6d, 0x74, 0x2e, 0x4d, 0x52, 0x03, 0x66, 0x6d, 0x74, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_rawDescData = file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_rawDesc
)

func file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_goTypes = []interface{}{
	(*All)(nil),         // 0: test.All
	(*test_a_1.M1)(nil), // 1: test.a.M1
	(*test_a_1.M2)(nil), // 2: test.a.M2
	(*test_b_1.M1)(nil), // 3: test.b.part1.M1
	(*test_b_1.M2)(nil), // 4: test.b.part2.M2
	(*fmt.M)(nil),       // 5: fmt.M
}
var file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_depIdxs = []int32{
	1, // 0: test.All.am1:type_name -> test.a.M1
	2, // 1: test.All.am2:type_name -> test.a.M2
	3, // 2: test.All.bm1:type_name -> test.b.part1.M1
	4, // 3: test.All.bm2:type_name -> test.b.part2.M2
	5, // 4: test.All.fmt:type_name -> fmt.M
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_init() }
func file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_init() {
	if File_cmd_protoc_gen_go_testdata_imports_test_import_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*All); i {
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
			RawDescriptor: file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_depIdxs,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_imports_test_import_all_proto = out.File
	file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_rawDesc = nil
	file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_imports_test_import_all_proto_depIdxs = nil
}
