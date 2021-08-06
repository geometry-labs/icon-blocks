// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: block_raw.proto

package models

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

type BlockRaw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature        string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	ItemId           string `protobuf:"bytes,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	NextLeader       string `protobuf:"bytes,3,opt,name=next_leader,json=nextLeader,proto3" json:"next_leader,omitempty"`
	TransactionCount uint32 `protobuf:"varint,4,opt,name=transaction_count,json=transactionCount,proto3" json:"transaction_count,omitempty"`
	Type             string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Version          string `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	PeerId           string `protobuf:"bytes,7,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Number           uint32 `protobuf:"varint,8,opt,name=number,proto3" json:"number,omitempty"`
	MerkleRootHash   string `protobuf:"bytes,9,opt,name=merkle_root_hash,json=merkleRootHash,proto3" json:"merkle_root_hash,omitempty"`
	ItemTimestamp    string `protobuf:"bytes,10,opt,name=item_timestamp,json=itemTimestamp,proto3" json:"item_timestamp,omitempty"`
	Hash             string `protobuf:"bytes,11,opt,name=hash,proto3" json:"hash,omitempty"`
	ParentHash       string `protobuf:"bytes,12,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	Timestamp        uint64 `protobuf:"varint,13,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *BlockRaw) Reset() {
	*x = BlockRaw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_raw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRaw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRaw) ProtoMessage() {}

func (x *BlockRaw) ProtoReflect() protoreflect.Message {
	mi := &file_block_raw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRaw.ProtoReflect.Descriptor instead.
func (*BlockRaw) Descriptor() ([]byte, []int) {
	return file_block_raw_proto_rawDescGZIP(), []int{0}
}

func (x *BlockRaw) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *BlockRaw) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *BlockRaw) GetNextLeader() string {
	if x != nil {
		return x.NextLeader
	}
	return ""
}

func (x *BlockRaw) GetTransactionCount() uint32 {
	if x != nil {
		return x.TransactionCount
	}
	return 0
}

func (x *BlockRaw) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *BlockRaw) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *BlockRaw) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *BlockRaw) GetNumber() uint32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *BlockRaw) GetMerkleRootHash() string {
	if x != nil {
		return x.MerkleRootHash
	}
	return ""
}

func (x *BlockRaw) GetItemTimestamp() string {
	if x != nil {
		return x.ItemTimestamp
	}
	return ""
}

func (x *BlockRaw) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *BlockRaw) GetParentHash() string {
	if x != nil {
		return x.ParentHash
	}
	return ""
}

func (x *BlockRaw) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_block_raw_proto protoreflect.FileDescriptor

var file_block_raw_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x92, 0x03, 0x0a, 0x08, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x61, 0x77, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2b,
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x65,
	0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x74,
	0x65, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_block_raw_proto_rawDescOnce sync.Once
	file_block_raw_proto_rawDescData = file_block_raw_proto_rawDesc
)

func file_block_raw_proto_rawDescGZIP() []byte {
	file_block_raw_proto_rawDescOnce.Do(func() {
		file_block_raw_proto_rawDescData = protoimpl.X.CompressGZIP(file_block_raw_proto_rawDescData)
	})
	return file_block_raw_proto_rawDescData
}

var file_block_raw_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_block_raw_proto_goTypes = []interface{}{
	(*BlockRaw)(nil), // 0: models.BlockRaw
}
var file_block_raw_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_block_raw_proto_init() }
func file_block_raw_proto_init() {
	if File_block_raw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_block_raw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRaw); i {
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
			RawDescriptor: file_block_raw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_block_raw_proto_goTypes,
		DependencyIndexes: file_block_raw_proto_depIdxs,
		MessageInfos:      file_block_raw_proto_msgTypes,
	}.Build()
	File_block_raw_proto = out.File
	file_block_raw_proto_rawDesc = nil
	file_block_raw_proto_goTypes = nil
	file_block_raw_proto_depIdxs = nil
}
