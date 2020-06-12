// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/event/event.proto

package event

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
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

// Indicates the status of CatalogCaching. The reason why this is not embeded in TaskNodeMetadata is, that we may use for other types of nodes as well in the future
type CatalogCacheStatus int32

const (
	// Used to indicate that caching was disabled
	CatalogCacheStatus_CACHE_DISABED CatalogCacheStatus = 0
	// Used to indicate that the cache lookup resulted in no matches
	CatalogCacheStatus_CACHE_MISS CatalogCacheStatus = 1
	// used to indicate that the associated artifact was a result of a previous execution
	CatalogCacheStatus_CACHE_HIT CatalogCacheStatus = 2
	// used to indicate that the resultant artifact was added to the cache
	CatalogCacheStatus_CACHE_POPULATED CatalogCacheStatus = 3
	// Used to indicate that cache lookup failed because of an error
	CatalogCacheStatus_CACHE_LOOKUP_FAILURE CatalogCacheStatus = 4
	// Used to indicate that cache lookup failed because of an error
	CatalogCacheStatus_CACHE_PUT_FAILURE CatalogCacheStatus = 5
)

var CatalogCacheStatus_name = map[int32]string{
	0: "CACHE_DISABED",
	1: "CACHE_MISS",
	2: "CACHE_HIT",
	3: "CACHE_POPULATED",
	4: "CACHE_LOOKUP_FAILURE",
	5: "CACHE_PUT_FAILURE",
}

var CatalogCacheStatus_value = map[string]int32{
	"CACHE_DISABED":        0,
	"CACHE_MISS":           1,
	"CACHE_HIT":            2,
	"CACHE_POPULATED":      3,
	"CACHE_LOOKUP_FAILURE": 4,
	"CACHE_PUT_FAILURE":    5,
}

func (x CatalogCacheStatus) String() string {
	return proto.EnumName(CatalogCacheStatus_name, int32(x))
}

func (CatalogCacheStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{0}
}

type WorkflowExecutionEvent struct {
	// Workflow execution id
	ExecutionId *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	// the id of the originator (Propeller) of the event
	ProducerId string                       `protobuf:"bytes,2,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	Phase      core.WorkflowExecution_Phase `protobuf:"varint,3,opt,name=phase,proto3,enum=flyteidl.core.WorkflowExecution_Phase" json:"phase,omitempty"`
	// This timestamp represents when the original event occurred, it is generated
	// by the executor of the workflow.
	OccurredAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	// Types that are valid to be assigned to OutputResult:
	//	*WorkflowExecutionEvent_OutputUri
	//	*WorkflowExecutionEvent_Error
	OutputResult         isWorkflowExecutionEvent_OutputResult `protobuf_oneof:"output_result"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *WorkflowExecutionEvent) Reset()         { *m = WorkflowExecutionEvent{} }
func (m *WorkflowExecutionEvent) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecutionEvent) ProtoMessage()    {}
func (*WorkflowExecutionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{0}
}

func (m *WorkflowExecutionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecutionEvent.Unmarshal(m, b)
}
func (m *WorkflowExecutionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecutionEvent.Marshal(b, m, deterministic)
}
func (m *WorkflowExecutionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionEvent.Merge(m, src)
}
func (m *WorkflowExecutionEvent) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecutionEvent.Size(m)
}
func (m *WorkflowExecutionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionEvent proto.InternalMessageInfo

func (m *WorkflowExecutionEvent) GetExecutionId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.ExecutionId
	}
	return nil
}

func (m *WorkflowExecutionEvent) GetProducerId() string {
	if m != nil {
		return m.ProducerId
	}
	return ""
}

func (m *WorkflowExecutionEvent) GetPhase() core.WorkflowExecution_Phase {
	if m != nil {
		return m.Phase
	}
	return core.WorkflowExecution_UNDEFINED
}

func (m *WorkflowExecutionEvent) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

type isWorkflowExecutionEvent_OutputResult interface {
	isWorkflowExecutionEvent_OutputResult()
}

type WorkflowExecutionEvent_OutputUri struct {
	OutputUri string `protobuf:"bytes,5,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type WorkflowExecutionEvent_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,6,opt,name=error,proto3,oneof"`
}

func (*WorkflowExecutionEvent_OutputUri) isWorkflowExecutionEvent_OutputResult() {}

func (*WorkflowExecutionEvent_Error) isWorkflowExecutionEvent_OutputResult() {}

func (m *WorkflowExecutionEvent) GetOutputResult() isWorkflowExecutionEvent_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *WorkflowExecutionEvent) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*WorkflowExecutionEvent_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *WorkflowExecutionEvent) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*WorkflowExecutionEvent_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*WorkflowExecutionEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*WorkflowExecutionEvent_OutputUri)(nil),
		(*WorkflowExecutionEvent_Error)(nil),
	}
}

type NodeExecutionEvent struct {
	// Unique identifier for this node execution
	Id *core.NodeExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// the id of the originator (Propeller) of the event
	ProducerId string                   `protobuf:"bytes,2,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	Phase      core.NodeExecution_Phase `protobuf:"varint,3,opt,name=phase,proto3,enum=flyteidl.core.NodeExecution_Phase" json:"phase,omitempty"`
	// This timestamp represents when the original event occurred, it is generated
	// by the executor of the node.
	OccurredAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	InputUri   string               `protobuf:"bytes,5,opt,name=input_uri,json=inputUri,proto3" json:"input_uri,omitempty"`
	// Types that are valid to be assigned to OutputResult:
	//	*NodeExecutionEvent_OutputUri
	//	*NodeExecutionEvent_Error
	OutputResult isNodeExecutionEvent_OutputResult `protobuf_oneof:"output_result"`
	// Additional metadata to do with this event's node target based
	// on the node type
	//
	// Types that are valid to be assigned to TargetMetadata:
	//	*NodeExecutionEvent_WorkflowNodeMetadata
	//	*NodeExecutionEvent_TaskNodeMetadata
	TargetMetadata isNodeExecutionEvent_TargetMetadata `protobuf_oneof:"target_metadata"`
	// Specifies which task (if any) launched this node.
	ParentTaskMetadata   *ParentTaskExecutionMetadata `protobuf:"bytes,9,opt,name=parent_task_metadata,json=parentTaskMetadata,proto3" json:"parent_task_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *NodeExecutionEvent) Reset()         { *m = NodeExecutionEvent{} }
func (m *NodeExecutionEvent) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionEvent) ProtoMessage()    {}
func (*NodeExecutionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{1}
}

func (m *NodeExecutionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionEvent.Unmarshal(m, b)
}
func (m *NodeExecutionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionEvent.Marshal(b, m, deterministic)
}
func (m *NodeExecutionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionEvent.Merge(m, src)
}
func (m *NodeExecutionEvent) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionEvent.Size(m)
}
func (m *NodeExecutionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionEvent proto.InternalMessageInfo

func (m *NodeExecutionEvent) GetId() *core.NodeExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *NodeExecutionEvent) GetProducerId() string {
	if m != nil {
		return m.ProducerId
	}
	return ""
}

func (m *NodeExecutionEvent) GetPhase() core.NodeExecution_Phase {
	if m != nil {
		return m.Phase
	}
	return core.NodeExecution_UNDEFINED
}

func (m *NodeExecutionEvent) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

func (m *NodeExecutionEvent) GetInputUri() string {
	if m != nil {
		return m.InputUri
	}
	return ""
}

type isNodeExecutionEvent_OutputResult interface {
	isNodeExecutionEvent_OutputResult()
}

type NodeExecutionEvent_OutputUri struct {
	OutputUri string `protobuf:"bytes,6,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type NodeExecutionEvent_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,7,opt,name=error,proto3,oneof"`
}

func (*NodeExecutionEvent_OutputUri) isNodeExecutionEvent_OutputResult() {}

func (*NodeExecutionEvent_Error) isNodeExecutionEvent_OutputResult() {}

func (m *NodeExecutionEvent) GetOutputResult() isNodeExecutionEvent_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *NodeExecutionEvent) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*NodeExecutionEvent_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *NodeExecutionEvent) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*NodeExecutionEvent_Error); ok {
		return x.Error
	}
	return nil
}

type isNodeExecutionEvent_TargetMetadata interface {
	isNodeExecutionEvent_TargetMetadata()
}

type NodeExecutionEvent_WorkflowNodeMetadata struct {
	WorkflowNodeMetadata *WorkflowNodeMetadata `protobuf:"bytes,8,opt,name=workflow_node_metadata,json=workflowNodeMetadata,proto3,oneof"`
}

type NodeExecutionEvent_TaskNodeMetadata struct {
	TaskNodeMetadata *TaskNodeMetadata `protobuf:"bytes,10,opt,name=task_node_metadata,json=taskNodeMetadata,proto3,oneof"`
}

func (*NodeExecutionEvent_WorkflowNodeMetadata) isNodeExecutionEvent_TargetMetadata() {}

func (*NodeExecutionEvent_TaskNodeMetadata) isNodeExecutionEvent_TargetMetadata() {}

func (m *NodeExecutionEvent) GetTargetMetadata() isNodeExecutionEvent_TargetMetadata {
	if m != nil {
		return m.TargetMetadata
	}
	return nil
}

func (m *NodeExecutionEvent) GetWorkflowNodeMetadata() *WorkflowNodeMetadata {
	if x, ok := m.GetTargetMetadata().(*NodeExecutionEvent_WorkflowNodeMetadata); ok {
		return x.WorkflowNodeMetadata
	}
	return nil
}

func (m *NodeExecutionEvent) GetTaskNodeMetadata() *TaskNodeMetadata {
	if x, ok := m.GetTargetMetadata().(*NodeExecutionEvent_TaskNodeMetadata); ok {
		return x.TaskNodeMetadata
	}
	return nil
}

func (m *NodeExecutionEvent) GetParentTaskMetadata() *ParentTaskExecutionMetadata {
	if m != nil {
		return m.ParentTaskMetadata
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NodeExecutionEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NodeExecutionEvent_OutputUri)(nil),
		(*NodeExecutionEvent_Error)(nil),
		(*NodeExecutionEvent_WorkflowNodeMetadata)(nil),
		(*NodeExecutionEvent_TaskNodeMetadata)(nil),
	}
}

// For Workflow Nodes we need to send information about the workflow that's launched
type WorkflowNodeMetadata struct {
	ExecutionId          *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *WorkflowNodeMetadata) Reset()         { *m = WorkflowNodeMetadata{} }
func (m *WorkflowNodeMetadata) String() string { return proto.CompactTextString(m) }
func (*WorkflowNodeMetadata) ProtoMessage()    {}
func (*WorkflowNodeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{2}
}

func (m *WorkflowNodeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowNodeMetadata.Unmarshal(m, b)
}
func (m *WorkflowNodeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowNodeMetadata.Marshal(b, m, deterministic)
}
func (m *WorkflowNodeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowNodeMetadata.Merge(m, src)
}
func (m *WorkflowNodeMetadata) XXX_Size() int {
	return xxx_messageInfo_WorkflowNodeMetadata.Size(m)
}
func (m *WorkflowNodeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowNodeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowNodeMetadata proto.InternalMessageInfo

func (m *WorkflowNodeMetadata) GetExecutionId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.ExecutionId
	}
	return nil
}

// Catalog artifact information with specific metadata
type CatalogMetadata struct {
	// Dataset ID in the catalog
	DatasetId string `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// Artifact tag in the catalog
	ArtifactTag string `protobuf:"bytes,2,opt,name=artifact_tag,json=artifactTag,proto3" json:"artifact_tag,omitempty"`
	// Optional: Source Workflow Execution identifier, if this dataset was generated by another execution in Flyte
	SourceExecutionId    *core.WorkflowExecutionIdentifier `protobuf:"bytes,3,opt,name=source_execution_id,json=sourceExecutionId,proto3" json:"source_execution_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *CatalogMetadata) Reset()         { *m = CatalogMetadata{} }
func (m *CatalogMetadata) String() string { return proto.CompactTextString(m) }
func (*CatalogMetadata) ProtoMessage()    {}
func (*CatalogMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{3}
}

func (m *CatalogMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatalogMetadata.Unmarshal(m, b)
}
func (m *CatalogMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatalogMetadata.Marshal(b, m, deterministic)
}
func (m *CatalogMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatalogMetadata.Merge(m, src)
}
func (m *CatalogMetadata) XXX_Size() int {
	return xxx_messageInfo_CatalogMetadata.Size(m)
}
func (m *CatalogMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CatalogMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CatalogMetadata proto.InternalMessageInfo

func (m *CatalogMetadata) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *CatalogMetadata) GetArtifactTag() string {
	if m != nil {
		return m.ArtifactTag
	}
	return ""
}

func (m *CatalogMetadata) GetSourceExecutionId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.SourceExecutionId
	}
	return nil
}

type TaskNodeMetadata struct {
	// Captures the status of caching for this execution.
	CacheStatus CatalogCacheStatus `protobuf:"varint,1,opt,name=cache_status,json=cacheStatus,proto3,enum=flyteidl.event.CatalogCacheStatus" json:"cache_status,omitempty"`
	// This structure carries the catalog artifact information
	CatalogKey           *CatalogMetadata `protobuf:"bytes,2,opt,name=catalog_key,json=catalogKey,proto3" json:"catalog_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TaskNodeMetadata) Reset()         { *m = TaskNodeMetadata{} }
func (m *TaskNodeMetadata) String() string { return proto.CompactTextString(m) }
func (*TaskNodeMetadata) ProtoMessage()    {}
func (*TaskNodeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{4}
}

func (m *TaskNodeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskNodeMetadata.Unmarshal(m, b)
}
func (m *TaskNodeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskNodeMetadata.Marshal(b, m, deterministic)
}
func (m *TaskNodeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskNodeMetadata.Merge(m, src)
}
func (m *TaskNodeMetadata) XXX_Size() int {
	return xxx_messageInfo_TaskNodeMetadata.Size(m)
}
func (m *TaskNodeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskNodeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TaskNodeMetadata proto.InternalMessageInfo

func (m *TaskNodeMetadata) GetCacheStatus() CatalogCacheStatus {
	if m != nil {
		return m.CacheStatus
	}
	return CatalogCacheStatus_CACHE_DISABED
}

func (m *TaskNodeMetadata) GetCatalogKey() *CatalogMetadata {
	if m != nil {
		return m.CatalogKey
	}
	return nil
}

type ParentTaskExecutionMetadata struct {
	Id                   *core.TaskExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ParentTaskExecutionMetadata) Reset()         { *m = ParentTaskExecutionMetadata{} }
func (m *ParentTaskExecutionMetadata) String() string { return proto.CompactTextString(m) }
func (*ParentTaskExecutionMetadata) ProtoMessage()    {}
func (*ParentTaskExecutionMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{5}
}

func (m *ParentTaskExecutionMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParentTaskExecutionMetadata.Unmarshal(m, b)
}
func (m *ParentTaskExecutionMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParentTaskExecutionMetadata.Marshal(b, m, deterministic)
}
func (m *ParentTaskExecutionMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParentTaskExecutionMetadata.Merge(m, src)
}
func (m *ParentTaskExecutionMetadata) XXX_Size() int {
	return xxx_messageInfo_ParentTaskExecutionMetadata.Size(m)
}
func (m *ParentTaskExecutionMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ParentTaskExecutionMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ParentTaskExecutionMetadata proto.InternalMessageInfo

func (m *ParentTaskExecutionMetadata) GetId() *core.TaskExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

// Plugin specific execution event information. For tasks like Python, Hive, Spark, DynamicJob.
type TaskExecutionEvent struct {
	// ID of the task. In combination with the retryAttempt this will indicate
	// the task execution uniquely for a given parent node execution.
	TaskId *core.Identifier `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// A task execution is always kicked off by a node execution, the event consumer
	// will use the parent_id to relate the task to it's parent node execution
	ParentNodeExecutionId *core.NodeExecutionIdentifier `protobuf:"bytes,2,opt,name=parent_node_execution_id,json=parentNodeExecutionId,proto3" json:"parent_node_execution_id,omitempty"`
	// retry attempt number for this task, ie., 2 for the second attempt
	RetryAttempt uint32 `protobuf:"varint,3,opt,name=retry_attempt,json=retryAttempt,proto3" json:"retry_attempt,omitempty"`
	// Phase associated with the event
	Phase core.TaskExecution_Phase `protobuf:"varint,4,opt,name=phase,proto3,enum=flyteidl.core.TaskExecution_Phase" json:"phase,omitempty"`
	// id of the process that sent this event, mainly for trace debugging
	ProducerId string `protobuf:"bytes,5,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	// log information for the task execution
	Logs []*core.TaskLog `protobuf:"bytes,6,rep,name=logs,proto3" json:"logs,omitempty"`
	// This timestamp represents when the original event occurred, it is generated
	// by the executor of the task.
	OccurredAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	// URI of the input file, it encodes all the information
	// including Cloud source provider. ie., s3://...
	InputUri string `protobuf:"bytes,8,opt,name=input_uri,json=inputUri,proto3" json:"input_uri,omitempty"`
	// Types that are valid to be assigned to OutputResult:
	//	*TaskExecutionEvent_OutputUri
	//	*TaskExecutionEvent_Error
	OutputResult isTaskExecutionEvent_OutputResult `protobuf_oneof:"output_result"`
	// Custom data that the task plugin sends back. This is extensible to allow various plugins in the system.
	CustomInfo *_struct.Struct `protobuf:"bytes,11,opt,name=custom_info,json=customInfo,proto3" json:"custom_info,omitempty"`
	// Some phases, like RUNNING, can send multiple events with changed metadata (new logs, additional custom_info, etc)
	// that should be recorded regardless of the lack of phase change.
	// The version field should be incremented when metadata changes across the duration of an individual phase.
	PhaseVersion         uint32   `protobuf:"varint,12,opt,name=phase_version,json=phaseVersion,proto3" json:"phase_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskExecutionEvent) Reset()         { *m = TaskExecutionEvent{} }
func (m *TaskExecutionEvent) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionEvent) ProtoMessage()    {}
func (*TaskExecutionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{6}
}

func (m *TaskExecutionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionEvent.Unmarshal(m, b)
}
func (m *TaskExecutionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionEvent.Marshal(b, m, deterministic)
}
func (m *TaskExecutionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionEvent.Merge(m, src)
}
func (m *TaskExecutionEvent) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionEvent.Size(m)
}
func (m *TaskExecutionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionEvent proto.InternalMessageInfo

func (m *TaskExecutionEvent) GetTaskId() *core.Identifier {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *TaskExecutionEvent) GetParentNodeExecutionId() *core.NodeExecutionIdentifier {
	if m != nil {
		return m.ParentNodeExecutionId
	}
	return nil
}

func (m *TaskExecutionEvent) GetRetryAttempt() uint32 {
	if m != nil {
		return m.RetryAttempt
	}
	return 0
}

func (m *TaskExecutionEvent) GetPhase() core.TaskExecution_Phase {
	if m != nil {
		return m.Phase
	}
	return core.TaskExecution_UNDEFINED
}

func (m *TaskExecutionEvent) GetProducerId() string {
	if m != nil {
		return m.ProducerId
	}
	return ""
}

func (m *TaskExecutionEvent) GetLogs() []*core.TaskLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *TaskExecutionEvent) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

func (m *TaskExecutionEvent) GetInputUri() string {
	if m != nil {
		return m.InputUri
	}
	return ""
}

type isTaskExecutionEvent_OutputResult interface {
	isTaskExecutionEvent_OutputResult()
}

type TaskExecutionEvent_OutputUri struct {
	OutputUri string `protobuf:"bytes,9,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type TaskExecutionEvent_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,10,opt,name=error,proto3,oneof"`
}

func (*TaskExecutionEvent_OutputUri) isTaskExecutionEvent_OutputResult() {}

func (*TaskExecutionEvent_Error) isTaskExecutionEvent_OutputResult() {}

func (m *TaskExecutionEvent) GetOutputResult() isTaskExecutionEvent_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *TaskExecutionEvent) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*TaskExecutionEvent_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *TaskExecutionEvent) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*TaskExecutionEvent_Error); ok {
		return x.Error
	}
	return nil
}

func (m *TaskExecutionEvent) GetCustomInfo() *_struct.Struct {
	if m != nil {
		return m.CustomInfo
	}
	return nil
}

func (m *TaskExecutionEvent) GetPhaseVersion() uint32 {
	if m != nil {
		return m.PhaseVersion
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TaskExecutionEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TaskExecutionEvent_OutputUri)(nil),
		(*TaskExecutionEvent_Error)(nil),
	}
}

func init() {
	proto.RegisterEnum("flyteidl.event.CatalogCacheStatus", CatalogCacheStatus_name, CatalogCacheStatus_value)
	proto.RegisterType((*WorkflowExecutionEvent)(nil), "flyteidl.event.WorkflowExecutionEvent")
	proto.RegisterType((*NodeExecutionEvent)(nil), "flyteidl.event.NodeExecutionEvent")
	proto.RegisterType((*WorkflowNodeMetadata)(nil), "flyteidl.event.WorkflowNodeMetadata")
	proto.RegisterType((*CatalogMetadata)(nil), "flyteidl.event.CatalogMetadata")
	proto.RegisterType((*TaskNodeMetadata)(nil), "flyteidl.event.TaskNodeMetadata")
	proto.RegisterType((*ParentTaskExecutionMetadata)(nil), "flyteidl.event.ParentTaskExecutionMetadata")
	proto.RegisterType((*TaskExecutionEvent)(nil), "flyteidl.event.TaskExecutionEvent")
}

func init() { proto.RegisterFile("flyteidl/event/event.proto", fileDescriptor_4b035d24120b1b12) }

var fileDescriptor_4b035d24120b1b12 = []byte{
	// 928 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xd1, 0x6e, 0xe3, 0x44,
	0x14, 0xad, 0xd3, 0x24, 0x6d, 0xae, 0x93, 0x36, 0x9d, 0xed, 0x16, 0xd3, 0xa5, 0x34, 0x04, 0x84,
	0xaa, 0x22, 0x1c, 0x91, 0x15, 0x68, 0x25, 0x78, 0x20, 0x6d, 0x83, 0x62, 0x6d, 0x4b, 0x23, 0x27,
	0x01, 0x69, 0x05, 0xb2, 0xa6, 0xf6, 0xc4, 0xb5, 0x9a, 0x78, 0xac, 0xf1, 0x78, 0x4b, 0xfe, 0x61,
	0xff, 0x80, 0x8f, 0xe0, 0xc7, 0x78, 0xe7, 0x15, 0x79, 0xc6, 0x76, 0x6b, 0xa7, 0x04, 0x82, 0xf6,
	0x25, 0xd1, 0x9c, 0x39, 0x73, 0xe6, 0xce, 0x99, 0x3b, 0x47, 0x86, 0xc3, 0xe9, 0x6c, 0xc1, 0x89,
	0xe7, 0xcc, 0x3a, 0xe4, 0x2d, 0xf1, 0xb9, 0xfc, 0xd5, 0x03, 0x46, 0x39, 0x45, 0x3b, 0xe9, 0x9c,
	0x2e, 0xd0, 0xc3, 0xa3, 0x8c, 0x6b, 0x53, 0x46, 0x3a, 0xe4, 0x37, 0x62, 0x47, 0xdc, 0xa3, 0xbe,
	0xa4, 0x1f, 0x7e, 0x9c, 0x9f, 0xf6, 0x1c, 0xe2, 0x73, 0x6f, 0xea, 0x11, 0x96, 0xcc, 0x1f, 0xbb,
	0x94, 0xba, 0x33, 0xd2, 0x11, 0xa3, 0x9b, 0x68, 0xda, 0xe1, 0xde, 0x9c, 0x84, 0x1c, 0xcf, 0x83,
	0x84, 0xf0, 0x51, 0x91, 0x10, 0x72, 0x16, 0xd9, 0x49, 0x35, 0xed, 0x3f, 0x4b, 0x70, 0xf0, 0x33,
	0x65, 0x77, 0xd3, 0x19, 0xbd, 0xef, 0xa7, 0x5b, 0xf7, 0xe3, 0xc2, 0xd0, 0x15, 0xd4, 0xb3, 0x62,
	0x2c, 0xcf, 0xd1, 0x94, 0x96, 0x72, 0xa2, 0x76, 0x4f, 0xf5, 0xac, 0xfe, 0xb8, 0x20, 0x7d, 0x69,
	0xb1, 0x91, 0x55, 0x68, 0xaa, 0xe4, 0x01, 0x44, 0xc7, 0xa0, 0x06, 0x8c, 0x3a, 0x91, 0x4d, 0x58,
	0xac, 0x56, 0x6a, 0x29, 0x27, 0x35, 0x13, 0x52, 0xc8, 0x70, 0xd0, 0x77, 0x50, 0x09, 0x6e, 0x71,
	0x48, 0xb4, 0xcd, 0x96, 0x72, 0xb2, 0xd3, 0xfd, 0xfc, 0xdf, 0x36, 0xd2, 0x87, 0x31, 0xdb, 0x94,
	0x8b, 0xd0, 0xb7, 0xa0, 0x52, 0xdb, 0x8e, 0x18, 0x23, 0x8e, 0x85, 0xb9, 0x56, 0x16, 0xc5, 0x1e,
	0xea, 0xf2, 0xf0, 0x7a, 0x7a, 0x78, 0x7d, 0x9c, 0xba, 0x63, 0x42, 0x4a, 0xef, 0x71, 0x74, 0x0c,
	0x40, 0x23, 0x1e, 0x44, 0xdc, 0x8a, 0x98, 0xa7, 0x55, 0xe2, 0xd2, 0x06, 0x1b, 0x66, 0x4d, 0x62,
	0x13, 0xe6, 0xa1, 0xaf, 0xa1, 0x42, 0x18, 0xa3, 0x4c, 0xab, 0x0a, 0xdd, 0xa3, 0x42, 0x6d, 0x0f,
	0xce, 0xc5, 0xa4, 0xc1, 0x86, 0x29, 0xd9, 0x67, 0xbb, 0xd0, 0x48, 0x74, 0x19, 0x09, 0xa3, 0x19,
	0x6f, 0xbf, 0xab, 0x00, 0xfa, 0x91, 0x3a, 0xa4, 0x60, 0xf5, 0x37, 0x50, 0xca, 0x0c, 0x2e, 0x9e,
	0x3b, 0x47, 0x7f, 0x64, 0x6e, 0xc9, 0xfb, 0x0f, 0x9e, 0xbe, 0xca, 0x7b, 0xda, 0x5e, 0xa5, 0xfd,
	0x1e, 0xfd, 0x7c, 0x01, 0x35, 0xcf, 0xcf, 0xd9, 0x69, 0x6e, 0x0b, 0x20, 0xf6, 0x32, 0x6f, 0x76,
	0x75, 0x85, 0xd9, 0x5b, 0xeb, 0x98, 0x8d, 0x7e, 0x81, 0x83, 0xfb, 0xa4, 0x47, 0x2c, 0x9f, 0x3a,
	0xc4, 0x9a, 0x13, 0x8e, 0x1d, 0xcc, 0xb1, 0xb6, 0x2d, 0x74, 0x3e, 0xd3, 0xf3, 0x2f, 0x2f, 0xeb,
	0xa8, 0xd8, 0x85, 0xab, 0x84, 0x3b, 0x50, 0xcc, 0xfd, 0xfb, 0x27, 0x70, 0x34, 0x04, 0xc4, 0x71,
	0x78, 0x57, 0x50, 0x06, 0xa1, 0xdc, 0x2a, 0x2a, 0x8f, 0x71, 0x78, 0x57, 0x50, 0x6d, 0xf2, 0x02,
	0x86, 0x7e, 0x85, 0xfd, 0x00, 0x33, 0xe2, 0x73, 0x4b, 0x08, 0x67, 0x9a, 0x35, 0xa1, 0xf9, 0x45,
	0x51, 0x73, 0x28, 0xb8, 0xb1, 0x72, 0x66, 0x40, 0x2a, 0x65, 0xa2, 0x20, 0x9b, 0x4c, 0xb1, 0xa5,
	0xde, 0x3b, 0xdb, 0x83, 0x5d, 0x8e, 0x99, 0x4b, 0x78, 0xb6, 0x55, 0x9b, 0xc0, 0xfe, 0x53, 0x26,
	0xbc, 0xe7, 0xa7, 0xdf, 0xfe, 0x43, 0x81, 0xdd, 0x73, 0xcc, 0xf1, 0x8c, 0xba, 0xd9, 0x16, 0x47,
	0x00, 0xf1, 0x7f, 0x48, 0x78, 0xba, 0x41, 0xcd, 0xac, 0x25, 0x88, 0xe1, 0xa0, 0x4f, 0xa0, 0x8e,
	0x19, 0xf7, 0xa6, 0xd8, 0x8e, 0xed, 0x71, 0x93, 0xd6, 0x56, 0x53, 0x6c, 0x8c, 0x5d, 0xf4, 0x06,
	0x9e, 0x85, 0x34, 0x62, 0x36, 0xb1, 0x72, 0xb5, 0x6e, 0xae, 0x5d, 0xeb, 0x9e, 0x94, 0x79, 0x34,
	0xd5, 0xfe, 0x5d, 0x81, 0x66, 0xf1, 0x12, 0x51, 0x1f, 0xea, 0x36, 0xb6, 0x6f, 0x89, 0x15, 0x72,
	0xcc, 0xa3, 0x50, 0x14, 0x9d, 0x7b, 0x53, 0xf2, 0xa2, 0x92, 0x93, 0x9e, 0xc7, 0xd4, 0x91, 0x60,
	0x9a, 0xaa, 0xfd, 0x30, 0x40, 0xdf, 0x83, 0x6a, 0x4b, 0x8a, 0x75, 0x47, 0x16, 0xe2, 0x64, 0x6a,
	0xf7, 0xf8, 0x1f, 0x54, 0xb2, 0x2b, 0x86, 0x64, 0xcd, 0x6b, 0xb2, 0x68, 0x4f, 0xe0, 0xc5, 0x8a,
	0x6e, 0x58, 0x99, 0x26, 0xb9, 0x15, 0xf9, 0x34, 0x69, 0xff, 0x55, 0x06, 0x94, 0x9b, 0x97, 0xe1,
	0xd4, 0x85, 0x2d, 0xd1, 0xa0, 0x99, 0xe6, 0x87, 0x05, 0xcd, 0x47, 0x32, 0xd5, 0x98, 0x69, 0x38,
	0xc8, 0x02, 0x2d, 0xe9, 0x6d, 0xf1, 0x5e, 0x72, 0x17, 0x54, 0x5a, 0x2b, 0xe6, 0x9e, 0x4b, 0x9d,
	0xc2, 0x34, 0xfa, 0x14, 0x1a, 0x8c, 0x70, 0xb6, 0xb0, 0x30, 0xe7, 0x64, 0x1e, 0x70, 0x71, 0xed,
	0x0d, 0xb3, 0x2e, 0xc0, 0x9e, 0xc4, 0x1e, 0xd2, 0xaf, 0xfc, 0x64, 0xfa, 0xe5, 0xce, 0x9a, 0x4f,
	0xbf, 0x42, 0xb0, 0x56, 0x96, 0x82, 0xf5, 0x14, 0xca, 0x33, 0xea, 0x86, 0x5a, 0xb5, 0xb5, 0x79,
	0xa2, 0x76, 0x0f, 0x9e, 0x50, 0xbe, 0xa4, 0xae, 0x29, 0x38, 0xc5, 0x28, 0xdd, 0xfa, 0xff, 0x51,
	0xba, 0xbd, 0x32, 0x4a, 0x6b, 0x2b, 0xa2, 0x14, 0xd6, 0x8a, 0xd2, 0x57, 0xa0, 0xda, 0x51, 0xc8,
	0xe9, 0xdc, 0xf2, 0xfc, 0x29, 0xd5, 0x54, 0xb1, 0xf8, 0x83, 0xa5, 0x8a, 0x47, 0xe2, 0x4b, 0xc2,
	0x04, 0xc9, 0x35, 0xfc, 0x29, 0x8d, 0xef, 0x45, 0x38, 0x68, 0xbd, 0x25, 0x2c, 0xf4, 0xa8, 0xaf,
	0xd5, 0xe5, 0xbd, 0x08, 0xf0, 0x27, 0x89, 0x2d, 0x45, 0xd3, 0xe9, 0x3b, 0x05, 0xd0, 0xf2, 0xb3,
	0x41, 0x7b, 0xd0, 0x38, 0xef, 0x9d, 0x0f, 0xfa, 0xd6, 0x85, 0x31, 0xea, 0x9d, 0xf5, 0x2f, 0x9a,
	0x1b, 0x68, 0x07, 0x40, 0x42, 0x57, 0xc6, 0x68, 0xd4, 0x54, 0x50, 0x03, 0x6a, 0x72, 0x3c, 0x30,
	0xc6, 0xcd, 0x12, 0x7a, 0x06, 0xbb, 0x72, 0x38, 0xbc, 0x1e, 0x4e, 0x2e, 0x7b, 0xe3, 0xfe, 0x45,
	0x73, 0x13, 0x69, 0xb0, 0x2f, 0xc1, 0xcb, 0xeb, 0xeb, 0xd7, 0x93, 0xa1, 0xf5, 0x43, 0xcf, 0xb8,
	0x9c, 0x98, 0xfd, 0x66, 0x19, 0x3d, 0x87, 0xbd, 0x84, 0x3e, 0x19, 0x67, 0x70, 0xe5, 0xec, 0xe5,
	0x9b, 0xaf, 0x5c, 0x8f, 0xdf, 0x46, 0x37, 0xba, 0x4d, 0xe7, 0x9d, 0xd9, 0x62, 0xca, 0x3b, 0xd9,
	0x57, 0x98, 0x4b, 0xfc, 0x4e, 0x70, 0xf3, 0xa5, 0x4b, 0x3b, 0xf9, 0x6f, 0xbc, 0x9b, 0xaa, 0xb0,
	0xe5, 0xe5, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x24, 0x95, 0xd2, 0x52, 0xfc, 0x09, 0x00, 0x00,
}
