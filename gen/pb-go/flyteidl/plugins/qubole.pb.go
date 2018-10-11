// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/qubole.proto

package plugins // import "github.com/lyft/flyteidl/gen/pb-go/plugins"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Defines a query to execute on a hive cluster.
type HiveQuery struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	TimeoutSec           uint32   `protobuf:"varint,2,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`
	RetryCount           uint32   `protobuf:"varint,3,opt,name=retryCount,proto3" json:"retryCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HiveQuery) Reset()         { *m = HiveQuery{} }
func (m *HiveQuery) String() string { return proto.CompactTextString(m) }
func (*HiveQuery) ProtoMessage()    {}
func (*HiveQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_qubole_744e52afeb73081d, []int{0}
}
func (m *HiveQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiveQuery.Unmarshal(m, b)
}
func (m *HiveQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiveQuery.Marshal(b, m, deterministic)
}
func (dst *HiveQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiveQuery.Merge(dst, src)
}
func (m *HiveQuery) XXX_Size() int {
	return xxx_messageInfo_HiveQuery.Size(m)
}
func (m *HiveQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_HiveQuery.DiscardUnknown(m)
}

var xxx_messageInfo_HiveQuery proto.InternalMessageInfo

func (m *HiveQuery) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *HiveQuery) GetTimeoutSec() uint32 {
	if m != nil {
		return m.TimeoutSec
	}
	return 0
}

func (m *HiveQuery) GetRetryCount() uint32 {
	if m != nil {
		return m.RetryCount
	}
	return 0
}

// Defines a collection of hive queries.
type HiveQueryCollection struct {
	Queries              []*HiveQuery `protobuf:"bytes,2,rep,name=queries,proto3" json:"queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HiveQueryCollection) Reset()         { *m = HiveQueryCollection{} }
func (m *HiveQueryCollection) String() string { return proto.CompactTextString(m) }
func (*HiveQueryCollection) ProtoMessage()    {}
func (*HiveQueryCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_qubole_744e52afeb73081d, []int{1}
}
func (m *HiveQueryCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiveQueryCollection.Unmarshal(m, b)
}
func (m *HiveQueryCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiveQueryCollection.Marshal(b, m, deterministic)
}
func (dst *HiveQueryCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiveQueryCollection.Merge(dst, src)
}
func (m *HiveQueryCollection) XXX_Size() int {
	return xxx_messageInfo_HiveQueryCollection.Size(m)
}
func (m *HiveQueryCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_HiveQueryCollection.DiscardUnknown(m)
}

var xxx_messageInfo_HiveQueryCollection proto.InternalMessageInfo

func (m *HiveQueryCollection) GetQueries() []*HiveQuery {
	if m != nil {
		return m.Queries
	}
	return nil
}

type QuboleHiveJob struct {
	ClusterLabel         string               `protobuf:"bytes,1,opt,name=cluster_label,json=clusterLabel,proto3" json:"cluster_label,omitempty"`
	QueryCollection      *HiveQueryCollection `protobuf:"bytes,2,opt,name=query_collection,json=queryCollection,proto3" json:"query_collection,omitempty"`
	Tags                 []string             `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *QuboleHiveJob) Reset()         { *m = QuboleHiveJob{} }
func (m *QuboleHiveJob) String() string { return proto.CompactTextString(m) }
func (*QuboleHiveJob) ProtoMessage()    {}
func (*QuboleHiveJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_qubole_744e52afeb73081d, []int{2}
}
func (m *QuboleHiveJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuboleHiveJob.Unmarshal(m, b)
}
func (m *QuboleHiveJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuboleHiveJob.Marshal(b, m, deterministic)
}
func (dst *QuboleHiveJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuboleHiveJob.Merge(dst, src)
}
func (m *QuboleHiveJob) XXX_Size() int {
	return xxx_messageInfo_QuboleHiveJob.Size(m)
}
func (m *QuboleHiveJob) XXX_DiscardUnknown() {
	xxx_messageInfo_QuboleHiveJob.DiscardUnknown(m)
}

var xxx_messageInfo_QuboleHiveJob proto.InternalMessageInfo

func (m *QuboleHiveJob) GetClusterLabel() string {
	if m != nil {
		return m.ClusterLabel
	}
	return ""
}

func (m *QuboleHiveJob) GetQueryCollection() *HiveQueryCollection {
	if m != nil {
		return m.QueryCollection
	}
	return nil
}

func (m *QuboleHiveJob) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*HiveQuery)(nil), "flyteidl.plugins.HiveQuery")
	proto.RegisterType((*HiveQueryCollection)(nil), "flyteidl.plugins.HiveQueryCollection")
	proto.RegisterType((*QuboleHiveJob)(nil), "flyteidl.plugins.QuboleHiveJob")
}

func init() {
	proto.RegisterFile("flyteidl/plugins/qubole.proto", fileDescriptor_qubole_744e52afeb73081d)
}

var fileDescriptor_qubole_744e52afeb73081d = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0xe3, 0x1f, 0x3a, 0xb5, 0x58, 0x56, 0x0f, 0x51, 0x51, 0x43, 0x45, 0x08, 0xa2,
	0x09, 0x54, 0x7c, 0x01, 0x7b, 0x11, 0xe9, 0xc1, 0xc6, 0x9b, 0x97, 0xd2, 0x5d, 0xa7, 0x71, 0x71,
	0x9b, 0x6d, 0x77, 0x67, 0x85, 0xbc, 0x8a, 0x4f, 0x2b, 0x49, 0x93, 0x14, 0x7a, 0xf0, 0x36, 0xfb,
	0x7d, 0x33, 0xcc, 0xef, 0xdb, 0x81, 0xcb, 0x85, 0x2a, 0x08, 0xe5, 0xa7, 0x4a, 0x56, 0xca, 0x65,
	0x32, 0xb7, 0xc9, 0xda, 0x71, 0xad, 0x30, 0x5e, 0x19, 0x4d, 0x9a, 0x0d, 0x1a, 0x3b, 0xae, 0xed,
	0xf3, 0xb3, 0x76, 0x40, 0x68, 0x83, 0x09, 0xcd, 0xed, 0xb7, 0xdd, 0x34, 0x0f, 0x39, 0x74, 0x5f,
	0xe4, 0x0f, 0x4e, 0x1d, 0x9a, 0x82, 0x9d, 0xc2, 0xfe, 0xba, 0x2c, 0x02, 0x2f, 0xf4, 0xa2, 0x6e,
	0xba, 0x79, 0xb0, 0x6b, 0xe8, 0x91, 0x5c, 0xa2, 0x76, 0x34, 0xb3, 0x28, 0x82, 0x4e, 0xe8, 0x45,
	0xfd, 0x14, 0x6a, 0xe9, 0x1d, 0x05, 0xbb, 0x02, 0x30, 0x48, 0xa6, 0x18, 0x6b, 0x97, 0x53, 0xe0,
	0x6f, 0xfc, 0xad, 0x32, 0x9c, 0xc0, 0x49, 0xbb, 0x63, 0xac, 0x95, 0x42, 0x41, 0x52, 0xe7, 0xec,
	0x09, 0x0e, 0xcb, 0x05, 0x12, 0x6d, 0xd0, 0x09, 0xfd, 0xa8, 0x37, 0xba, 0x88, 0x77, 0xc9, 0xe3,
	0x76, 0x2e, 0x6d, 0x7a, 0x87, 0xbf, 0x1e, 0xf4, 0xa7, 0x55, 0xde, 0xd2, 0x7c, 0xd5, 0x9c, 0xdd,
	0x40, 0x5f, 0x28, 0x67, 0x09, 0xcd, 0x4c, 0xcd, 0x39, 0xaa, 0x1a, 0xff, 0xa8, 0x16, 0x27, 0xa5,
	0xc6, 0xde, 0x60, 0x50, 0xc5, 0x99, 0x89, 0x96, 0xa0, 0x8a, 0xd2, 0x1b, 0xdd, 0xfe, 0xb3, 0x76,
	0x8b, 0x9b, 0x1e, 0xaf, 0x77, 0xf8, 0x19, 0xec, 0xd1, 0x3c, 0xb3, 0x81, 0x1f, 0xfa, 0x51, 0x37,
	0xad, 0xea, 0xe7, 0xfb, 0x8f, 0xbb, 0x4c, 0xd2, 0x97, 0xe3, 0xb1, 0xd0, 0xcb, 0x44, 0x15, 0x0b,
	0x4a, 0xda, 0xbf, 0xcf, 0x30, 0x4f, 0x56, 0xfc, 0x21, 0xd3, 0xcd, 0xd9, 0xf8, 0x41, 0x75, 0x83,
	0xc7, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0x51, 0x84, 0x46, 0xd1, 0x01, 0x00, 0x00,
}
