// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commands/board.proto

package commands

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

type BoardDetailsReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Fqbn                 string    `protobuf:"bytes,2,opt,name=fqbn,proto3" json:"fqbn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BoardDetailsReq) Reset()         { *m = BoardDetailsReq{} }
func (m *BoardDetailsReq) String() string { return proto.CompactTextString(m) }
func (*BoardDetailsReq) ProtoMessage()    {}
func (*BoardDetailsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{0}
}

func (m *BoardDetailsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardDetailsReq.Unmarshal(m, b)
}
func (m *BoardDetailsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardDetailsReq.Marshal(b, m, deterministic)
}
func (m *BoardDetailsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardDetailsReq.Merge(m, src)
}
func (m *BoardDetailsReq) XXX_Size() int {
	return xxx_messageInfo_BoardDetailsReq.Size(m)
}
func (m *BoardDetailsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardDetailsReq.DiscardUnknown(m)
}

var xxx_messageInfo_BoardDetailsReq proto.InternalMessageInfo

func (m *BoardDetailsReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *BoardDetailsReq) GetFqbn() string {
	if m != nil {
		return m.Fqbn
	}
	return ""
}

type BoardDetailsResp struct {
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ConfigOptions        []*ConfigOption `protobuf:"bytes,3,rep,name=config_options,json=configOptions,proto3" json:"config_options,omitempty"`
	RequiredTools        []*RequiredTool `protobuf:"bytes,4,rep,name=required_tools,json=requiredTools,proto3" json:"required_tools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BoardDetailsResp) Reset()         { *m = BoardDetailsResp{} }
func (m *BoardDetailsResp) String() string { return proto.CompactTextString(m) }
func (*BoardDetailsResp) ProtoMessage()    {}
func (*BoardDetailsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{1}
}

func (m *BoardDetailsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardDetailsResp.Unmarshal(m, b)
}
func (m *BoardDetailsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardDetailsResp.Marshal(b, m, deterministic)
}
func (m *BoardDetailsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardDetailsResp.Merge(m, src)
}
func (m *BoardDetailsResp) XXX_Size() int {
	return xxx_messageInfo_BoardDetailsResp.Size(m)
}
func (m *BoardDetailsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardDetailsResp.DiscardUnknown(m)
}

var xxx_messageInfo_BoardDetailsResp proto.InternalMessageInfo

func (m *BoardDetailsResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BoardDetailsResp) GetConfigOptions() []*ConfigOption {
	if m != nil {
		return m.ConfigOptions
	}
	return nil
}

func (m *BoardDetailsResp) GetRequiredTools() []*RequiredTool {
	if m != nil {
		return m.RequiredTools
	}
	return nil
}

type ConfigOption struct {
	Option               string         `protobuf:"bytes,1,opt,name=option,proto3" json:"option,omitempty"`
	OptionLabel          string         `protobuf:"bytes,2,opt,name=option_label,json=optionLabel,proto3" json:"option_label,omitempty"`
	Values               []*ConfigValue `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ConfigOption) Reset()         { *m = ConfigOption{} }
func (m *ConfigOption) String() string { return proto.CompactTextString(m) }
func (*ConfigOption) ProtoMessage()    {}
func (*ConfigOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{2}
}

func (m *ConfigOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigOption.Unmarshal(m, b)
}
func (m *ConfigOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigOption.Marshal(b, m, deterministic)
}
func (m *ConfigOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigOption.Merge(m, src)
}
func (m *ConfigOption) XXX_Size() int {
	return xxx_messageInfo_ConfigOption.Size(m)
}
func (m *ConfigOption) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigOption.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigOption proto.InternalMessageInfo

func (m *ConfigOption) GetOption() string {
	if m != nil {
		return m.Option
	}
	return ""
}

func (m *ConfigOption) GetOptionLabel() string {
	if m != nil {
		return m.OptionLabel
	}
	return ""
}

func (m *ConfigOption) GetValues() []*ConfigValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type ConfigValue struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	ValueLabel           string   `protobuf:"bytes,2,opt,name=value_label,json=valueLabel,proto3" json:"value_label,omitempty"`
	Selected             bool     `protobuf:"varint,3,opt,name=selected,proto3" json:"selected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigValue) Reset()         { *m = ConfigValue{} }
func (m *ConfigValue) String() string { return proto.CompactTextString(m) }
func (*ConfigValue) ProtoMessage()    {}
func (*ConfigValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{3}
}

func (m *ConfigValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigValue.Unmarshal(m, b)
}
func (m *ConfigValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigValue.Marshal(b, m, deterministic)
}
func (m *ConfigValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigValue.Merge(m, src)
}
func (m *ConfigValue) XXX_Size() int {
	return xxx_messageInfo_ConfigValue.Size(m)
}
func (m *ConfigValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigValue.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigValue proto.InternalMessageInfo

func (m *ConfigValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ConfigValue) GetValueLabel() string {
	if m != nil {
		return m.ValueLabel
	}
	return ""
}

func (m *ConfigValue) GetSelected() bool {
	if m != nil {
		return m.Selected
	}
	return false
}

type RequiredTool struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Packager             string   `protobuf:"bytes,3,opt,name=packager,proto3" json:"packager,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequiredTool) Reset()         { *m = RequiredTool{} }
func (m *RequiredTool) String() string { return proto.CompactTextString(m) }
func (*RequiredTool) ProtoMessage()    {}
func (*RequiredTool) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{4}
}

func (m *RequiredTool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequiredTool.Unmarshal(m, b)
}
func (m *RequiredTool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequiredTool.Marshal(b, m, deterministic)
}
func (m *RequiredTool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequiredTool.Merge(m, src)
}
func (m *RequiredTool) XXX_Size() int {
	return xxx_messageInfo_RequiredTool.Size(m)
}
func (m *RequiredTool) XXX_DiscardUnknown() {
	xxx_messageInfo_RequiredTool.DiscardUnknown(m)
}

var xxx_messageInfo_RequiredTool proto.InternalMessageInfo

func (m *RequiredTool) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RequiredTool) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RequiredTool) GetPackager() string {
	if m != nil {
		return m.Packager
	}
	return ""
}

type BoardAttachReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	BoardUri             string    `protobuf:"bytes,2,opt,name=board_uri,json=boardUri,proto3" json:"board_uri,omitempty"`
	SketchPath           string    `protobuf:"bytes,3,opt,name=sketch_path,json=sketchPath,proto3" json:"sketch_path,omitempty"`
	SearchTimeout        string    `protobuf:"bytes,4,opt,name=search_timeout,json=searchTimeout,proto3" json:"search_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BoardAttachReq) Reset()         { *m = BoardAttachReq{} }
func (m *BoardAttachReq) String() string { return proto.CompactTextString(m) }
func (*BoardAttachReq) ProtoMessage()    {}
func (*BoardAttachReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{5}
}

func (m *BoardAttachReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardAttachReq.Unmarshal(m, b)
}
func (m *BoardAttachReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardAttachReq.Marshal(b, m, deterministic)
}
func (m *BoardAttachReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardAttachReq.Merge(m, src)
}
func (m *BoardAttachReq) XXX_Size() int {
	return xxx_messageInfo_BoardAttachReq.Size(m)
}
func (m *BoardAttachReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardAttachReq.DiscardUnknown(m)
}

var xxx_messageInfo_BoardAttachReq proto.InternalMessageInfo

func (m *BoardAttachReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *BoardAttachReq) GetBoardUri() string {
	if m != nil {
		return m.BoardUri
	}
	return ""
}

func (m *BoardAttachReq) GetSketchPath() string {
	if m != nil {
		return m.SketchPath
	}
	return ""
}

func (m *BoardAttachReq) GetSearchTimeout() string {
	if m != nil {
		return m.SearchTimeout
	}
	return ""
}

type BoardAttachResp struct {
	TaskProgress         *TaskProgress `protobuf:"bytes,1,opt,name=task_progress,json=taskProgress,proto3" json:"task_progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BoardAttachResp) Reset()         { *m = BoardAttachResp{} }
func (m *BoardAttachResp) String() string { return proto.CompactTextString(m) }
func (*BoardAttachResp) ProtoMessage()    {}
func (*BoardAttachResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{6}
}

func (m *BoardAttachResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardAttachResp.Unmarshal(m, b)
}
func (m *BoardAttachResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardAttachResp.Marshal(b, m, deterministic)
}
func (m *BoardAttachResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardAttachResp.Merge(m, src)
}
func (m *BoardAttachResp) XXX_Size() int {
	return xxx_messageInfo_BoardAttachResp.Size(m)
}
func (m *BoardAttachResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardAttachResp.DiscardUnknown(m)
}

var xxx_messageInfo_BoardAttachResp proto.InternalMessageInfo

func (m *BoardAttachResp) GetTaskProgress() *TaskProgress {
	if m != nil {
		return m.TaskProgress
	}
	return nil
}

type BoardListReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BoardListReq) Reset()         { *m = BoardListReq{} }
func (m *BoardListReq) String() string { return proto.CompactTextString(m) }
func (*BoardListReq) ProtoMessage()    {}
func (*BoardListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{7}
}

func (m *BoardListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardListReq.Unmarshal(m, b)
}
func (m *BoardListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardListReq.Marshal(b, m, deterministic)
}
func (m *BoardListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardListReq.Merge(m, src)
}
func (m *BoardListReq) XXX_Size() int {
	return xxx_messageInfo_BoardListReq.Size(m)
}
func (m *BoardListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardListReq.DiscardUnknown(m)
}

var xxx_messageInfo_BoardListReq proto.InternalMessageInfo

func (m *BoardListReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

type BoardListResp struct {
	Ports                []*DetectedPort `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BoardListResp) Reset()         { *m = BoardListResp{} }
func (m *BoardListResp) String() string { return proto.CompactTextString(m) }
func (*BoardListResp) ProtoMessage()    {}
func (*BoardListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{8}
}

func (m *BoardListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardListResp.Unmarshal(m, b)
}
func (m *BoardListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardListResp.Marshal(b, m, deterministic)
}
func (m *BoardListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardListResp.Merge(m, src)
}
func (m *BoardListResp) XXX_Size() int {
	return xxx_messageInfo_BoardListResp.Size(m)
}
func (m *BoardListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardListResp.DiscardUnknown(m)
}

var xxx_messageInfo_BoardListResp proto.InternalMessageInfo

func (m *BoardListResp) GetPorts() []*DetectedPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

type BoardInstallReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Board                string    `protobuf:"bytes,2,opt,name=board,proto3" json:"board,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BoardInstallReq) Reset()         { *m = BoardInstallReq{} }
func (m *BoardInstallReq) String() string { return proto.CompactTextString(m) }
func (*BoardInstallReq) ProtoMessage()    {}
func (*BoardInstallReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{9}
}

func (m *BoardInstallReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardInstallReq.Unmarshal(m, b)
}
func (m *BoardInstallReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardInstallReq.Marshal(b, m, deterministic)
}
func (m *BoardInstallReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardInstallReq.Merge(m, src)
}
func (m *BoardInstallReq) XXX_Size() int {
	return xxx_messageInfo_BoardInstallReq.Size(m)
}
func (m *BoardInstallReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardInstallReq.DiscardUnknown(m)
}

var xxx_messageInfo_BoardInstallReq proto.InternalMessageInfo

func (m *BoardInstallReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *BoardInstallReq) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

type BoardInstallResp struct {
	DownloadProgress     *DownloadProgress `protobuf:"bytes,1,opt,name=download_progress,json=downloadProgress,proto3" json:"download_progress,omitempty"`
	TaskProgress         *TaskProgress     `protobuf:"bytes,2,opt,name=task_progress,json=taskProgress,proto3" json:"task_progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BoardInstallResp) Reset()         { *m = BoardInstallResp{} }
func (m *BoardInstallResp) String() string { return proto.CompactTextString(m) }
func (*BoardInstallResp) ProtoMessage()    {}
func (*BoardInstallResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{10}
}

func (m *BoardInstallResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardInstallResp.Unmarshal(m, b)
}
func (m *BoardInstallResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardInstallResp.Marshal(b, m, deterministic)
}
func (m *BoardInstallResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardInstallResp.Merge(m, src)
}
func (m *BoardInstallResp) XXX_Size() int {
	return xxx_messageInfo_BoardInstallResp.Size(m)
}
func (m *BoardInstallResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardInstallResp.DiscardUnknown(m)
}

var xxx_messageInfo_BoardInstallResp proto.InternalMessageInfo

func (m *BoardInstallResp) GetDownloadProgress() *DownloadProgress {
	if m != nil {
		return m.DownloadProgress
	}
	return nil
}

func (m *BoardInstallResp) GetTaskProgress() *TaskProgress {
	if m != nil {
		return m.TaskProgress
	}
	return nil
}

type BoardUninstallReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Board                string    `protobuf:"bytes,2,opt,name=board,proto3" json:"board,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BoardUninstallReq) Reset()         { *m = BoardUninstallReq{} }
func (m *BoardUninstallReq) String() string { return proto.CompactTextString(m) }
func (*BoardUninstallReq) ProtoMessage()    {}
func (*BoardUninstallReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{11}
}

func (m *BoardUninstallReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardUninstallReq.Unmarshal(m, b)
}
func (m *BoardUninstallReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardUninstallReq.Marshal(b, m, deterministic)
}
func (m *BoardUninstallReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardUninstallReq.Merge(m, src)
}
func (m *BoardUninstallReq) XXX_Size() int {
	return xxx_messageInfo_BoardUninstallReq.Size(m)
}
func (m *BoardUninstallReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardUninstallReq.DiscardUnknown(m)
}

var xxx_messageInfo_BoardUninstallReq proto.InternalMessageInfo

func (m *BoardUninstallReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *BoardUninstallReq) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

type BoardUninstallResp struct {
	TaskProgress         *TaskProgress `protobuf:"bytes,1,opt,name=task_progress,json=taskProgress,proto3" json:"task_progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BoardUninstallResp) Reset()         { *m = BoardUninstallResp{} }
func (m *BoardUninstallResp) String() string { return proto.CompactTextString(m) }
func (*BoardUninstallResp) ProtoMessage()    {}
func (*BoardUninstallResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{12}
}

func (m *BoardUninstallResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardUninstallResp.Unmarshal(m, b)
}
func (m *BoardUninstallResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardUninstallResp.Marshal(b, m, deterministic)
}
func (m *BoardUninstallResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardUninstallResp.Merge(m, src)
}
func (m *BoardUninstallResp) XXX_Size() int {
	return xxx_messageInfo_BoardUninstallResp.Size(m)
}
func (m *BoardUninstallResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardUninstallResp.DiscardUnknown(m)
}

var xxx_messageInfo_BoardUninstallResp proto.InternalMessageInfo

func (m *BoardUninstallResp) GetTaskProgress() *TaskProgress {
	if m != nil {
		return m.TaskProgress
	}
	return nil
}

type BoardSearchReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Query                string    `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BoardSearchReq) Reset()         { *m = BoardSearchReq{} }
func (m *BoardSearchReq) String() string { return proto.CompactTextString(m) }
func (*BoardSearchReq) ProtoMessage()    {}
func (*BoardSearchReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{13}
}

func (m *BoardSearchReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardSearchReq.Unmarshal(m, b)
}
func (m *BoardSearchReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardSearchReq.Marshal(b, m, deterministic)
}
func (m *BoardSearchReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardSearchReq.Merge(m, src)
}
func (m *BoardSearchReq) XXX_Size() int {
	return xxx_messageInfo_BoardSearchReq.Size(m)
}
func (m *BoardSearchReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardSearchReq.DiscardUnknown(m)
}

var xxx_messageInfo_BoardSearchReq proto.InternalMessageInfo

func (m *BoardSearchReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *BoardSearchReq) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type BoardSearchResp struct {
	Fqbn                 []string `protobuf:"bytes,1,rep,name=fqbn,proto3" json:"fqbn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoardSearchResp) Reset()         { *m = BoardSearchResp{} }
func (m *BoardSearchResp) String() string { return proto.CompactTextString(m) }
func (*BoardSearchResp) ProtoMessage()    {}
func (*BoardSearchResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{14}
}

func (m *BoardSearchResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardSearchResp.Unmarshal(m, b)
}
func (m *BoardSearchResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardSearchResp.Marshal(b, m, deterministic)
}
func (m *BoardSearchResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardSearchResp.Merge(m, src)
}
func (m *BoardSearchResp) XXX_Size() int {
	return xxx_messageInfo_BoardSearchResp.Size(m)
}
func (m *BoardSearchResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardSearchResp.DiscardUnknown(m)
}

var xxx_messageInfo_BoardSearchResp proto.InternalMessageInfo

func (m *BoardSearchResp) GetFqbn() []string {
	if m != nil {
		return m.Fqbn
	}
	return nil
}

type DetectedPort struct {
	Address              string           `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Protocol             string           `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	ProtocolLabel        string           `protobuf:"bytes,3,opt,name=protocol_label,json=protocolLabel,proto3" json:"protocol_label,omitempty"`
	Boards               []*BoardListItem `protobuf:"bytes,4,rep,name=boards,proto3" json:"boards,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DetectedPort) Reset()         { *m = DetectedPort{} }
func (m *DetectedPort) String() string { return proto.CompactTextString(m) }
func (*DetectedPort) ProtoMessage()    {}
func (*DetectedPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{15}
}

func (m *DetectedPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectedPort.Unmarshal(m, b)
}
func (m *DetectedPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectedPort.Marshal(b, m, deterministic)
}
func (m *DetectedPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectedPort.Merge(m, src)
}
func (m *DetectedPort) XXX_Size() int {
	return xxx_messageInfo_DetectedPort.Size(m)
}
func (m *DetectedPort) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectedPort.DiscardUnknown(m)
}

var xxx_messageInfo_DetectedPort proto.InternalMessageInfo

func (m *DetectedPort) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DetectedPort) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *DetectedPort) GetProtocolLabel() string {
	if m != nil {
		return m.ProtocolLabel
	}
	return ""
}

func (m *DetectedPort) GetBoards() []*BoardListItem {
	if m != nil {
		return m.Boards
	}
	return nil
}

type BoardListAllReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	SearchArgs           []string  `protobuf:"bytes,2,rep,name=search_args,json=searchArgs,proto3" json:"search_args,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BoardListAllReq) Reset()         { *m = BoardListAllReq{} }
func (m *BoardListAllReq) String() string { return proto.CompactTextString(m) }
func (*BoardListAllReq) ProtoMessage()    {}
func (*BoardListAllReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{16}
}

func (m *BoardListAllReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardListAllReq.Unmarshal(m, b)
}
func (m *BoardListAllReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardListAllReq.Marshal(b, m, deterministic)
}
func (m *BoardListAllReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardListAllReq.Merge(m, src)
}
func (m *BoardListAllReq) XXX_Size() int {
	return xxx_messageInfo_BoardListAllReq.Size(m)
}
func (m *BoardListAllReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardListAllReq.DiscardUnknown(m)
}

var xxx_messageInfo_BoardListAllReq proto.InternalMessageInfo

func (m *BoardListAllReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *BoardListAllReq) GetSearchArgs() []string {
	if m != nil {
		return m.SearchArgs
	}
	return nil
}

type BoardListAllResp struct {
	Boards               []*BoardListItem `protobuf:"bytes,1,rep,name=boards,proto3" json:"boards,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BoardListAllResp) Reset()         { *m = BoardListAllResp{} }
func (m *BoardListAllResp) String() string { return proto.CompactTextString(m) }
func (*BoardListAllResp) ProtoMessage()    {}
func (*BoardListAllResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{17}
}

func (m *BoardListAllResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardListAllResp.Unmarshal(m, b)
}
func (m *BoardListAllResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardListAllResp.Marshal(b, m, deterministic)
}
func (m *BoardListAllResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardListAllResp.Merge(m, src)
}
func (m *BoardListAllResp) XXX_Size() int {
	return xxx_messageInfo_BoardListAllResp.Size(m)
}
func (m *BoardListAllResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardListAllResp.DiscardUnknown(m)
}

var xxx_messageInfo_BoardListAllResp proto.InternalMessageInfo

func (m *BoardListAllResp) GetBoards() []*BoardListItem {
	if m != nil {
		return m.Boards
	}
	return nil
}

type BoardListItem struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	FQBN                 string   `protobuf:"bytes,2,opt,name=FQBN,proto3" json:"FQBN,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoardListItem) Reset()         { *m = BoardListItem{} }
func (m *BoardListItem) String() string { return proto.CompactTextString(m) }
func (*BoardListItem) ProtoMessage()    {}
func (*BoardListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_0882eeddaa6507ab, []int{18}
}

func (m *BoardListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardListItem.Unmarshal(m, b)
}
func (m *BoardListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardListItem.Marshal(b, m, deterministic)
}
func (m *BoardListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardListItem.Merge(m, src)
}
func (m *BoardListItem) XXX_Size() int {
	return xxx_messageInfo_BoardListItem.Size(m)
}
func (m *BoardListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardListItem.DiscardUnknown(m)
}

var xxx_messageInfo_BoardListItem proto.InternalMessageInfo

func (m *BoardListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BoardListItem) GetFQBN() string {
	if m != nil {
		return m.FQBN
	}
	return ""
}

func init() {
	proto.RegisterType((*BoardDetailsReq)(nil), "cc.arduino.cli.commands.BoardDetailsReq")
	proto.RegisterType((*BoardDetailsResp)(nil), "cc.arduino.cli.commands.BoardDetailsResp")
	proto.RegisterType((*ConfigOption)(nil), "cc.arduino.cli.commands.ConfigOption")
	proto.RegisterType((*ConfigValue)(nil), "cc.arduino.cli.commands.ConfigValue")
	proto.RegisterType((*RequiredTool)(nil), "cc.arduino.cli.commands.RequiredTool")
	proto.RegisterType((*BoardAttachReq)(nil), "cc.arduino.cli.commands.BoardAttachReq")
	proto.RegisterType((*BoardAttachResp)(nil), "cc.arduino.cli.commands.BoardAttachResp")
	proto.RegisterType((*BoardListReq)(nil), "cc.arduino.cli.commands.BoardListReq")
	proto.RegisterType((*BoardListResp)(nil), "cc.arduino.cli.commands.BoardListResp")
	proto.RegisterType((*BoardInstallReq)(nil), "cc.arduino.cli.commands.BoardInstallReq")
	proto.RegisterType((*BoardInstallResp)(nil), "cc.arduino.cli.commands.BoardInstallResp")
	proto.RegisterType((*BoardUninstallReq)(nil), "cc.arduino.cli.commands.BoardUninstallReq")
	proto.RegisterType((*BoardUninstallResp)(nil), "cc.arduino.cli.commands.BoardUninstallResp")
	proto.RegisterType((*BoardSearchReq)(nil), "cc.arduino.cli.commands.BoardSearchReq")
	proto.RegisterType((*BoardSearchResp)(nil), "cc.arduino.cli.commands.BoardSearchResp")
	proto.RegisterType((*DetectedPort)(nil), "cc.arduino.cli.commands.DetectedPort")
	proto.RegisterType((*BoardListAllReq)(nil), "cc.arduino.cli.commands.BoardListAllReq")
	proto.RegisterType((*BoardListAllResp)(nil), "cc.arduino.cli.commands.BoardListAllResp")
	proto.RegisterType((*BoardListItem)(nil), "cc.arduino.cli.commands.BoardListItem")
}

func init() { proto.RegisterFile("commands/board.proto", fileDescriptor_0882eeddaa6507ab) }

var fileDescriptor_0882eeddaa6507ab = []byte{
	// 763 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x6e, 0x13, 0x49,
	0x10, 0xd6, 0xc4, 0x8e, 0xd7, 0x2e, 0xdb, 0xd9, 0xa4, 0x95, 0xdd, 0xb5, 0xb2, 0x87, 0x75, 0x46,
	0x9b, 0x95, 0x57, 0x28, 0xb6, 0x14, 0x0e, 0x1c, 0xf8, 0x91, 0x12, 0x22, 0xa4, 0x20, 0x03, 0xa1,
	0x49, 0x22, 0x84, 0x84, 0x9c, 0x76, 0x4f, 0xc7, 0x1e, 0x79, 0x3c, 0x3d, 0xee, 0x6e, 0x07, 0xf1,
	0x04, 0x3c, 0x0c, 0x57, 0xc4, 0x03, 0xf0, 0x64, 0xa8, 0xff, 0x46, 0x93, 0x80, 0x89, 0x14, 0xcc,
	0x69, 0xaa, 0x6a, 0xaa, 0xab, 0xea, 0xab, 0x5f, 0xd8, 0xa4, 0x7c, 0x3a, 0x25, 0x69, 0x24, 0x7b,
	0x43, 0x4e, 0x44, 0xd4, 0xcd, 0x04, 0x57, 0x1c, 0xfd, 0x45, 0x69, 0x97, 0x88, 0x68, 0x1e, 0xa7,
	0xbc, 0x4b, 0x93, 0xb8, 0xeb, 0x95, 0xb6, 0xfe, 0xc8, 0xd5, 0x35, 0xc1, 0x53, 0xab, 0x1f, 0x46,
	0xf0, 0xfb, 0x81, 0x7e, 0x7e, 0xc8, 0x14, 0x89, 0x13, 0x89, 0xd9, 0x0c, 0x3d, 0x84, 0x6a, 0x9c,
	0x4a, 0x45, 0x52, 0xca, 0x5a, 0x41, 0x3b, 0xe8, 0xd4, 0xf7, 0xb6, 0xbb, 0x0b, 0xac, 0x76, 0x8f,
	0x9c, 0x22, 0xce, 0x9f, 0x20, 0x04, 0xe5, 0x8b, 0xd9, 0x30, 0x6d, 0xad, 0xb4, 0x83, 0x4e, 0x0d,
	0x1b, 0x3a, 0xfc, 0x12, 0xc0, 0xfa, 0x55, 0x37, 0x32, 0xd3, 0x8a, 0x29, 0x99, 0x32, 0xaf, 0xa8,
	0x69, 0xd4, 0x87, 0x35, 0xca, 0xd3, 0x8b, 0x78, 0x34, 0xe0, 0x99, 0x8a, 0x79, 0x2a, 0x5b, 0xa5,
	0x76, 0xa9, 0x53, 0xdf, 0xdb, 0x59, 0x18, 0xc1, 0x63, 0xa3, 0xfe, 0xc2, 0x68, 0xe3, 0x26, 0x2d,
	0x70, 0x52, 0x5b, 0x13, 0x6c, 0x36, 0x8f, 0x05, 0x8b, 0x06, 0x8a, 0xf3, 0x44, 0xb6, 0xca, 0x37,
	0x58, 0xc3, 0x4e, 0xfd, 0x84, 0xf3, 0x04, 0x37, 0x45, 0x81, 0x93, 0xe1, 0x87, 0x00, 0x1a, 0x45,
	0x6f, 0xe8, 0x4f, 0xa8, 0xd8, 0x28, 0x4d, 0x9a, 0x6a, 0xd8, 0x71, 0x68, 0x1b, 0x1a, 0x96, 0x1a,
	0x24, 0x64, 0xc8, 0x12, 0x07, 0xb0, 0x6e, 0x65, 0x7d, 0x2d, 0x42, 0x0f, 0xa0, 0x72, 0x49, 0x92,
	0x39, 0xf3, 0xf8, 0xfe, 0xbd, 0x01, 0xdf, 0x99, 0x56, 0xc6, 0xee, 0x4d, 0x78, 0x0e, 0xf5, 0x82,
	0x18, 0x6d, 0xc2, 0xaa, 0xf9, 0xe1, 0xc2, 0xb0, 0x0c, 0xfa, 0x07, 0xea, 0x86, 0xb8, 0x12, 0x04,
	0x18, 0x91, 0x8d, 0x61, 0x0b, 0xaa, 0x92, 0x25, 0x8c, 0x2a, 0x16, 0xb5, 0x4a, 0xed, 0xa0, 0x53,
	0xc5, 0x39, 0x1f, 0xbe, 0x86, 0x46, 0x31, 0x15, 0x79, 0xad, 0x82, 0x42, 0xad, 0x5a, 0xf0, 0xdb,
	0x25, 0x13, 0x52, 0xe3, 0xb7, 0xc6, 0x3d, 0xab, 0x2d, 0x67, 0x84, 0x4e, 0xc8, 0x88, 0x09, 0x63,
	0xb9, 0x86, 0x73, 0x3e, 0xfc, 0x14, 0xc0, 0x9a, 0x69, 0x85, 0x7d, 0xa5, 0x08, 0x1d, 0x2f, 0xa1,
	0xe1, 0xfe, 0x86, 0x9a, 0x99, 0x80, 0xc1, 0x5c, 0xc4, 0x2e, 0x92, 0xaa, 0x11, 0x9c, 0x8a, 0x58,
	0x67, 0x41, 0x4e, 0x98, 0xa2, 0xe3, 0x41, 0x46, 0xd4, 0xd8, 0x45, 0x03, 0x56, 0x74, 0x4c, 0xd4,
	0x18, 0xed, 0xc0, 0x9a, 0x64, 0x44, 0xd0, 0xf1, 0x40, 0xc5, 0x53, 0xc6, 0xe7, 0xaa, 0x55, 0x36,
	0x3a, 0x4d, 0x2b, 0x3d, 0xb1, 0xc2, 0xf0, 0xad, 0x9b, 0x13, 0x1f, 0xb5, 0xcc, 0xd0, 0x53, 0x68,
	0x2a, 0x22, 0x27, 0x83, 0x4c, 0xf0, 0x91, 0x60, 0x52, 0xba, 0xd8, 0x17, 0x37, 0xd7, 0x09, 0x91,
	0x93, 0x63, 0xa7, 0x8c, 0x1b, 0xaa, 0xc0, 0x85, 0xcf, 0xa0, 0x61, 0xcc, 0xf7, 0x63, 0xa9, 0x7e,
	0x3e, 0x25, 0x61, 0x1f, 0x9a, 0x05, 0x73, 0x32, 0x43, 0xf7, 0x61, 0x35, 0xe3, 0x42, 0xe9, 0x18,
	0x7f, 0x3c, 0x00, 0x87, 0x4c, 0x99, 0x0e, 0x38, 0xe6, 0x42, 0x61, 0xfb, 0x26, 0xbc, 0x70, 0xd8,
	0x8d, 0xa3, 0x24, 0x59, 0x42, 0xc9, 0x36, 0x61, 0xd5, 0x54, 0xc8, 0x95, 0xcb, 0x32, 0xe1, 0x67,
	0xbf, 0x25, 0x72, 0x47, 0x32, 0x43, 0x67, 0xb0, 0x11, 0xf1, 0x77, 0x69, 0xc2, 0x49, 0x74, 0x3d,
	0xd3, 0xff, 0x2f, 0x46, 0xe1, 0x5e, 0xe4, 0xd9, 0x5e, 0x8f, 0xae, 0x49, 0xbe, 0xad, 0xde, 0xca,
	0xed, 0xab, 0x37, 0x86, 0x0d, 0x13, 0xf7, 0x69, 0x1a, 0xff, 0xe2, 0x14, 0x9d, 0x03, 0xba, 0xee,
	0x69, 0xc9, 0x9d, 0xc8, 0xdc, 0x78, 0xbe, 0x32, 0xed, 0xbf, 0x1c, 0x20, 0xb3, 0x39, 0x13, 0xef,
	0x3d, 0x10, 0xc3, 0x84, 0x3b, 0xae, 0xa7, 0xbc, 0x1b, 0x7b, 0x0f, 0xcc, 0xe1, 0xd0, 0x2d, 0xea,
	0x0f, 0xc7, 0xc7, 0x00, 0x1a, 0xc5, 0x96, 0xd4, 0x4b, 0x87, 0x44, 0x51, 0x0e, 0xb2, 0x86, 0x3d,
	0x6b, 0x96, 0x8e, 0x3e, 0x69, 0x94, 0xfb, 0x65, 0x97, 0xf3, 0x7a, 0xc8, 0x3d, 0xed, 0xd6, 0xa1,
	0x5d, 0x04, 0x4d, 0x2f, 0xb5, 0x1b, 0xf1, 0x11, 0x54, 0x4c, 0x9a, 0xfd, 0x9d, 0xf8, 0x6f, 0x21,
	0xce, 0x7c, 0xba, 0x8e, 0x14, 0x9b, 0x62, 0xf7, 0x2a, 0x9c, 0x39, 0x50, 0xfa, 0xc7, 0xfe, 0x52,
	0xba, 0x40, 0xaf, 0x2f, 0xbb, 0x9d, 0x88, 0x18, 0xe9, 0x1e, 0x2d, 0x99, 0xf5, 0x65, 0x44, 0xfb,
	0x62, 0x24, 0x43, 0xec, 0x46, 0x26, 0x77, 0x29, 0xb3, 0x02, 0x8c, 0xe0, 0x56, 0x30, 0xee, 0x15,
	0xb6, 0x87, 0xfe, 0xf1, 0xdd, 0xed, 0x8f, 0xa0, 0xfc, 0xe4, 0xe5, 0xc1, 0x73, 0x7f, 0xbd, 0x35,
	0x7d, 0xb0, 0xfb, 0xe6, 0xce, 0x28, 0x56, 0xe3, 0xf9, 0x50, 0x7b, 0xe8, 0x39, 0x8f, 0xfe, 0xbb,
	0x4b, 0x93, 0xb8, 0x27, 0x32, 0xda, 0xf3, 0xde, 0x87, 0x15, 0x93, 0xfd, 0xbb, 0x5f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x92, 0x20, 0x27, 0x45, 0xca, 0x08, 0x00, 0x00,
}
