// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/array_job.proto

package plugins

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Describes a job that can process independent pieces of data concurrently. Multiple copies of the runnable component
// will be executed concurrently.
type ArrayJob struct {
	// Defines the minimum number of instances to bring up concurrently at any given point. Note that this is an
	// optimistic restriction and that, due to network partitioning or other failures, the actual number of currently
	// running instances might be more. This has to be a positive number if assigned. Default value is size.
	Parallelism int64 `protobuf:"varint,1,opt,name=parallelism,proto3" json:"parallelism,omitempty"`
	// Defines the number of instances to launch at most. This number should match the size of the input if the job
	// requires processing of all input data. This has to be a positive number.
	// In the case this is not defined, the back-end will determine the size at run-time by reading the inputs.
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	// Types that are valid to be assigned to SuccessCriteria:
	//	*ArrayJob_MinSuccesses
	//	*ArrayJob_MinSuccessRatio
	SuccessCriteria      isArrayJob_SuccessCriteria `protobuf_oneof:"success_criteria"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ArrayJob) Reset()         { *m = ArrayJob{} }
func (m *ArrayJob) String() string { return proto.CompactTextString(m) }
func (*ArrayJob) ProtoMessage()    {}
func (*ArrayJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_794211c91ec6cd2b, []int{0}
}

func (m *ArrayJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrayJob.Unmarshal(m, b)
}
func (m *ArrayJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrayJob.Marshal(b, m, deterministic)
}
func (m *ArrayJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrayJob.Merge(m, src)
}
func (m *ArrayJob) XXX_Size() int {
	return xxx_messageInfo_ArrayJob.Size(m)
}
func (m *ArrayJob) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrayJob.DiscardUnknown(m)
}

var xxx_messageInfo_ArrayJob proto.InternalMessageInfo

func (m *ArrayJob) GetParallelism() int64 {
	if m != nil {
		return m.Parallelism
	}
	return 0
}

func (m *ArrayJob) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type isArrayJob_SuccessCriteria interface {
	isArrayJob_SuccessCriteria()
}

type ArrayJob_MinSuccesses struct {
	MinSuccesses int64 `protobuf:"varint,3,opt,name=min_successes,json=minSuccesses,proto3,oneof"`
}

type ArrayJob_MinSuccessRatio struct {
	MinSuccessRatio float32 `protobuf:"fixed32,4,opt,name=min_success_ratio,json=minSuccessRatio,proto3,oneof"`
}

func (*ArrayJob_MinSuccesses) isArrayJob_SuccessCriteria() {}

func (*ArrayJob_MinSuccessRatio) isArrayJob_SuccessCriteria() {}

func (m *ArrayJob) GetSuccessCriteria() isArrayJob_SuccessCriteria {
	if m != nil {
		return m.SuccessCriteria
	}
	return nil
}

func (m *ArrayJob) GetMinSuccesses() int64 {
	if x, ok := m.GetSuccessCriteria().(*ArrayJob_MinSuccesses); ok {
		return x.MinSuccesses
	}
	return 0
}

func (m *ArrayJob) GetMinSuccessRatio() float32 {
	if x, ok := m.GetSuccessCriteria().(*ArrayJob_MinSuccessRatio); ok {
		return x.MinSuccessRatio
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ArrayJob) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ArrayJob_MinSuccesses)(nil),
		(*ArrayJob_MinSuccessRatio)(nil),
	}
}

func init() {
	proto.RegisterType((*ArrayJob)(nil), "flyteidl.plugins.ArrayJob")
}

func init() { proto.RegisterFile("flyteidl/plugins/array_job.proto", fileDescriptor_794211c91ec6cd2b) }

var fileDescriptor_794211c91ec6cd2b = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd0, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x05, 0xe0, 0xa6, 0xad, 0x10, 0x32, 0x20, 0x8a, 0xa7, 0x8c, 0x11, 0x12, 0x52, 0x07, 0x88,
	0x07, 0x06, 0xc4, 0x48, 0xa7, 0x8a, 0x31, 0x6c, 0x2c, 0x91, 0x6d, 0x8c, 0x39, 0x64, 0xfb, 0xac,
	0xb3, 0x33, 0x94, 0x7f, 0xc4, 0xbf, 0x44, 0xb1, 0x08, 0x44, 0xdd, 0x4e, 0xef, 0x7d, 0xcb, 0x3b,
	0xd6, 0xbc, 0xbb, 0x43, 0x36, 0xf0, 0xe6, 0x44, 0x74, 0x83, 0x85, 0x90, 0x84, 0x24, 0x92, 0x87,
	0xfe, 0x13, 0x55, 0x1b, 0x09, 0x33, 0xf2, 0xcd, 0x24, 0xda, 0x5f, 0x71, 0xfd, 0x5d, 0xb1, 0xd3,
	0xa7, 0x51, 0x3d, 0xa3, 0xe2, 0x0d, 0x3b, 0x8b, 0x92, 0xa4, 0x73, 0xc6, 0x41, 0xf2, 0x75, 0xd5,
	0x54, 0xdb, 0x55, 0x37, 0x8f, 0x38, 0x67, 0xeb, 0x04, 0x5f, 0xa6, 0x5e, 0x96, 0xaa, 0xdc, 0xfc,
	0x86, 0x5d, 0x78, 0x08, 0x7d, 0x1a, 0xb4, 0x36, 0x29, 0x99, 0x54, 0xaf, 0xc6, 0x72, 0xbf, 0xe8,
	0xce, 0x3d, 0x84, 0x97, 0x29, 0xe5, 0xb7, 0xec, 0x6a, 0xc6, 0x7a, 0x92, 0x19, 0xb0, 0x5e, 0x37,
	0xd5, 0x76, 0xb9, 0x5f, 0x74, 0x97, 0xff, 0xb4, 0x1b, 0x8b, 0x1d, 0x67, 0x9b, 0x49, 0x6a, 0x82,
	0x6c, 0x08, 0xe4, 0xee, 0xf1, 0xf5, 0xc1, 0x42, 0xfe, 0x18, 0x54, 0xab, 0xd1, 0x8b, 0x32, 0x05,
	0xc9, 0x8a, 0xbf, 0xd5, 0xd6, 0x04, 0x11, 0xd5, 0x9d, 0x45, 0x71, 0xfc, 0x08, 0x75, 0x52, 0xf6,
	0xdf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x26, 0xec, 0x61, 0x11, 0x23, 0x01, 0x00, 0x00,
}
