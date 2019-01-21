// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/node/kind.proto

package node

import (
	network "berty.tech/core/network"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	io "io"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Kind int32

const (
	// Unknown events should never happen
	Kind_Unknown         Kind = 0
	Kind_NodeStarted     Kind = 1
	Kind_NodeStopped     Kind = 2
	Kind_NodeIsAlive     Kind = 3
	Kind_BackgroundError Kind = 4
	Kind_BackgroundWarn  Kind = 5
	Kind_Debug           Kind = 6
	Kind_Statistics      Kind = 7
)

var Kind_name = map[int32]string{
	0: "Unknown",
	1: "NodeStarted",
	2: "NodeStopped",
	3: "NodeIsAlive",
	4: "BackgroundError",
	5: "BackgroundWarn",
	6: "Debug",
	7: "Statistics",
}

var Kind_value = map[string]int32{
	"Unknown":         0,
	"NodeStarted":     1,
	"NodeStopped":     2,
	"NodeIsAlive":     3,
	"BackgroundError": 4,
	"BackgroundWarn":  5,
	"Debug":           6,
	"Statistics":      7,
}

func (x Kind) String() string {
	return proto.EnumName(Kind_name, int32(x))
}

func (Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_209eb881e555280a, []int{0}
}

type NodeStartedAttrs struct {
	T                    bool     `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeStartedAttrs) Reset()         { *m = NodeStartedAttrs{} }
func (m *NodeStartedAttrs) String() string { return proto.CompactTextString(m) }
func (*NodeStartedAttrs) ProtoMessage()    {}
func (*NodeStartedAttrs) Descriptor() ([]byte, []int) {
	return fileDescriptor_209eb881e555280a, []int{0}
}
func (m *NodeStartedAttrs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeStartedAttrs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeStartedAttrs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodeStartedAttrs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStartedAttrs.Merge(m, src)
}
func (m *NodeStartedAttrs) XXX_Size() int {
	return m.Size()
}
func (m *NodeStartedAttrs) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStartedAttrs.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStartedAttrs proto.InternalMessageInfo

func (m *NodeStartedAttrs) GetT() bool {
	if m != nil {
		return m.T
	}
	return false
}

type NodeStoppedAttrs struct {
	ErrMsg               string   `protobuf:"bytes,1,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeStoppedAttrs) Reset()         { *m = NodeStoppedAttrs{} }
func (m *NodeStoppedAttrs) String() string { return proto.CompactTextString(m) }
func (*NodeStoppedAttrs) ProtoMessage()    {}
func (*NodeStoppedAttrs) Descriptor() ([]byte, []int) {
	return fileDescriptor_209eb881e555280a, []int{1}
}
func (m *NodeStoppedAttrs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeStoppedAttrs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeStoppedAttrs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodeStoppedAttrs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStoppedAttrs.Merge(m, src)
}
func (m *NodeStoppedAttrs) XXX_Size() int {
	return m.Size()
}
func (m *NodeStoppedAttrs) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStoppedAttrs.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStoppedAttrs proto.InternalMessageInfo

func (m *NodeStoppedAttrs) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type NodeIsAliveAttrs struct {
	T                    bool     `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeIsAliveAttrs) Reset()         { *m = NodeIsAliveAttrs{} }
func (m *NodeIsAliveAttrs) String() string { return proto.CompactTextString(m) }
func (*NodeIsAliveAttrs) ProtoMessage()    {}
func (*NodeIsAliveAttrs) Descriptor() ([]byte, []int) {
	return fileDescriptor_209eb881e555280a, []int{2}
}
func (m *NodeIsAliveAttrs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeIsAliveAttrs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeIsAliveAttrs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodeIsAliveAttrs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeIsAliveAttrs.Merge(m, src)
}
func (m *NodeIsAliveAttrs) XXX_Size() int {
	return m.Size()
}
func (m *NodeIsAliveAttrs) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeIsAliveAttrs.DiscardUnknown(m)
}

var xxx_messageInfo_NodeIsAliveAttrs proto.InternalMessageInfo

func (m *NodeIsAliveAttrs) GetT() bool {
	if m != nil {
		return m.T
	}
	return false
}

type BackgroundErrorAttrs struct {
	ErrMsg               string   `protobuf:"bytes,1,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackgroundErrorAttrs) Reset()         { *m = BackgroundErrorAttrs{} }
func (m *BackgroundErrorAttrs) String() string { return proto.CompactTextString(m) }
func (*BackgroundErrorAttrs) ProtoMessage()    {}
func (*BackgroundErrorAttrs) Descriptor() ([]byte, []int) {
	return fileDescriptor_209eb881e555280a, []int{3}
}
func (m *BackgroundErrorAttrs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BackgroundErrorAttrs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BackgroundErrorAttrs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BackgroundErrorAttrs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackgroundErrorAttrs.Merge(m, src)
}
func (m *BackgroundErrorAttrs) XXX_Size() int {
	return m.Size()
}
func (m *BackgroundErrorAttrs) XXX_DiscardUnknown() {
	xxx_messageInfo_BackgroundErrorAttrs.DiscardUnknown(m)
}

var xxx_messageInfo_BackgroundErrorAttrs proto.InternalMessageInfo

func (m *BackgroundErrorAttrs) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type BackgroundWarnAttrs struct {
	ErrMsg               string   `protobuf:"bytes,1,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackgroundWarnAttrs) Reset()         { *m = BackgroundWarnAttrs{} }
func (m *BackgroundWarnAttrs) String() string { return proto.CompactTextString(m) }
func (*BackgroundWarnAttrs) ProtoMessage()    {}
func (*BackgroundWarnAttrs) Descriptor() ([]byte, []int) {
	return fileDescriptor_209eb881e555280a, []int{4}
}
func (m *BackgroundWarnAttrs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BackgroundWarnAttrs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BackgroundWarnAttrs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BackgroundWarnAttrs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackgroundWarnAttrs.Merge(m, src)
}
func (m *BackgroundWarnAttrs) XXX_Size() int {
	return m.Size()
}
func (m *BackgroundWarnAttrs) XXX_DiscardUnknown() {
	xxx_messageInfo_BackgroundWarnAttrs.DiscardUnknown(m)
}

var xxx_messageInfo_BackgroundWarnAttrs proto.InternalMessageInfo

func (m *BackgroundWarnAttrs) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type DebugAttrs struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugAttrs) Reset()         { *m = DebugAttrs{} }
func (m *DebugAttrs) String() string { return proto.CompactTextString(m) }
func (*DebugAttrs) ProtoMessage()    {}
func (*DebugAttrs) Descriptor() ([]byte, []int) {
	return fileDescriptor_209eb881e555280a, []int{5}
}
func (m *DebugAttrs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DebugAttrs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DebugAttrs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DebugAttrs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugAttrs.Merge(m, src)
}
func (m *DebugAttrs) XXX_Size() int {
	return m.Size()
}
func (m *DebugAttrs) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugAttrs.DiscardUnknown(m)
}

var xxx_messageInfo_DebugAttrs proto.InternalMessageInfo

func (m *DebugAttrs) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StatisticsAttrs struct {
	ErrMsg                string                  `protobuf:"bytes,1,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	TotalNetworkBandwidth *network.BandwidthStats `protobuf:"bytes,2,opt,name=total_network_bandwidth,json=totalNetworkBandwidth,proto3" json:"total_network_bandwidth,omitempty"`
	PeersCount            int32                   `protobuf:"varint,3,opt,name=peers_count,json=peersCount,proto3" json:"peers_count,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                `json:"-"`
	XXX_unrecognized      []byte                  `json:"-"`
	XXX_sizecache         int32                   `json:"-"`
}

func (m *StatisticsAttrs) Reset()         { *m = StatisticsAttrs{} }
func (m *StatisticsAttrs) String() string { return proto.CompactTextString(m) }
func (*StatisticsAttrs) ProtoMessage()    {}
func (*StatisticsAttrs) Descriptor() ([]byte, []int) {
	return fileDescriptor_209eb881e555280a, []int{6}
}
func (m *StatisticsAttrs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StatisticsAttrs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StatisticsAttrs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StatisticsAttrs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatisticsAttrs.Merge(m, src)
}
func (m *StatisticsAttrs) XXX_Size() int {
	return m.Size()
}
func (m *StatisticsAttrs) XXX_DiscardUnknown() {
	xxx_messageInfo_StatisticsAttrs.DiscardUnknown(m)
}

var xxx_messageInfo_StatisticsAttrs proto.InternalMessageInfo

func (m *StatisticsAttrs) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *StatisticsAttrs) GetTotalNetworkBandwidth() *network.BandwidthStats {
	if m != nil {
		return m.TotalNetworkBandwidth
	}
	return nil
}

func (m *StatisticsAttrs) GetPeersCount() int32 {
	if m != nil {
		return m.PeersCount
	}
	return 0
}

func init() {
	proto.RegisterEnum("berty.node.Kind", Kind_name, Kind_value)
	proto.RegisterType((*NodeStartedAttrs)(nil), "berty.node.NodeStartedAttrs")
	proto.RegisterType((*NodeStoppedAttrs)(nil), "berty.node.NodeStoppedAttrs")
	proto.RegisterType((*NodeIsAliveAttrs)(nil), "berty.node.NodeIsAliveAttrs")
	proto.RegisterType((*BackgroundErrorAttrs)(nil), "berty.node.BackgroundErrorAttrs")
	proto.RegisterType((*BackgroundWarnAttrs)(nil), "berty.node.BackgroundWarnAttrs")
	proto.RegisterType((*DebugAttrs)(nil), "berty.node.DebugAttrs")
	proto.RegisterType((*StatisticsAttrs)(nil), "berty.node.StatisticsAttrs")
}

func init() { proto.RegisterFile("api/node/kind.proto", fileDescriptor_209eb881e555280a) }

var fileDescriptor_209eb881e555280a = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xd1, 0x6a, 0xd4, 0x40,
	0x14, 0x86, 0x3b, 0xdd, 0xee, 0xae, 0x3d, 0x2b, 0xdd, 0x61, 0xb6, 0xa5, 0xa1, 0x60, 0x0c, 0x7b,
	0xb5, 0x54, 0x48, 0x40, 0x9f, 0xa0, 0xab, 0x5e, 0x88, 0xd8, 0x8b, 0xb4, 0x45, 0xf0, 0x26, 0x24,
	0x99, 0x21, 0x1d, 0xd2, 0xce, 0x84, 0x33, 0x67, 0x5d, 0x7c, 0x05, 0x9f, 0xc0, 0x6b, 0x9f, 0xc6,
	0x4b, 0x1f, 0x41, 0xd6, 0x17, 0x91, 0x24, 0x1b, 0xa3, 0x22, 0xdd, 0xbb, 0xe1, 0xff, 0xff, 0xf3,
	0x87, 0x7c, 0xe7, 0xc0, 0x2c, 0xad, 0x74, 0x64, 0xac, 0x54, 0x51, 0xa9, 0x8d, 0x0c, 0x2b, 0xb4,
	0x64, 0x05, 0x64, 0x0a, 0xe9, 0x53, 0x58, 0xcb, 0x67, 0xc7, 0x85, 0x2d, 0x6c, 0x23, 0x47, 0xf5,
	0xab, 0x4d, 0x9c, 0x9d, 0x18, 0x45, 0x6b, 0x8b, 0x65, 0x74, 0xaf, 0x08, 0x75, 0xee, 0x5a, 0x79,
	0x1e, 0x00, 0xbf, 0xb4, 0x52, 0x5d, 0x51, 0x8a, 0xa4, 0xe4, 0x05, 0x11, 0x3a, 0xf1, 0x18, 0xd8,
	0xb5, 0xc7, 0x02, 0xb6, 0x78, 0x14, 0xb3, 0xeb, 0xf9, 0xb3, 0x2e, 0x61, 0xab, 0xaa, 0x4b, 0x9c,
	0xc2, 0x58, 0x21, 0x26, 0xf7, 0xae, 0x68, 0x72, 0x87, 0xf1, 0x48, 0x21, 0xbe, 0x73, 0x45, 0x57,
	0xf7, 0xc6, 0x5d, 0xdc, 0xe9, 0x8f, 0xea, 0x7f, 0x75, 0x11, 0x1c, 0x2f, 0xd3, 0xbc, 0x2c, 0xd0,
	0xae, 0x8c, 0x7c, 0x8d, 0x68, 0x71, 0x47, 0x65, 0x08, 0xb3, 0x7e, 0xe0, 0x7d, 0x8a, 0x66, 0x47,
	0xde, 0x07, 0x78, 0xa5, 0xb2, 0x55, 0xd1, 0xc6, 0x38, 0x0c, 0xfa, 0x48, 0xfd, 0x9c, 0x7f, 0x65,
	0x30, 0xbd, 0xa2, 0x94, 0xb4, 0x23, 0x9d, 0xbb, 0x87, 0xcb, 0xc4, 0x0d, 0x9c, 0x92, 0xa5, 0xf4,
	0x2e, 0xd9, 0xd2, 0x4b, 0xb2, 0xd4, 0xc8, 0xb5, 0x96, 0x74, 0xeb, 0xed, 0x07, 0x6c, 0x31, 0x79,
	0xfe, 0x24, 0xdc, 0x92, 0x6f, 0xfd, 0x70, 0xd9, 0xf9, 0xf5, 0x27, 0x5c, 0x7c, 0xd2, 0x4c, 0x5f,
	0xb6, 0xe6, 0x6f, 0x4f, 0x3c, 0x85, 0x49, 0xa5, 0x14, 0xba, 0x24, 0xb7, 0x2b, 0x43, 0xde, 0x20,
	0x60, 0x8b, 0x61, 0x0c, 0x8d, 0xf4, 0xb2, 0x56, 0xce, 0x3f, 0x33, 0x38, 0x78, 0xab, 0x8d, 0x14,
	0x13, 0x18, 0xdf, 0x98, 0xd2, 0xd8, 0xb5, 0xe1, 0x7b, 0x62, 0x0a, 0x93, 0x3f, 0x96, 0xc5, 0x59,
	0x2f, 0x34, 0xbb, 0xe1, 0xfb, 0x9d, 0xb0, 0xe5, 0xcf, 0x07, 0x62, 0x06, 0xd3, 0x7f, 0x70, 0xf3,
	0x03, 0x21, 0xe0, 0xe8, 0x6f, 0xa4, 0x7c, 0x28, 0x0e, 0x61, 0xd8, 0x60, 0xe3, 0x23, 0x71, 0x04,
	0xd0, 0x03, 0xe2, 0xe3, 0xe5, 0xf9, 0xb7, 0x8d, 0xcf, 0xbe, 0x6f, 0x7c, 0xf6, 0x63, 0xe3, 0xb3,
	0x2f, 0x3f, 0xfd, 0xbd, 0x0f, 0x5e, 0xfb, 0xd3, 0xa4, 0xf2, 0xdb, 0x28, 0xb7, 0xa8, 0xa2, 0xee,
	0x26, 0xb3, 0x51, 0x73, 0x56, 0x2f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x53, 0xc8, 0x53, 0xa1,
	0xa6, 0x02, 0x00, 0x00,
}

func (m *NodeStartedAttrs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeStartedAttrs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.T {
		dAtA[i] = 0x8
		i++
		if m.T {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *NodeStoppedAttrs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeStoppedAttrs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ErrMsg) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKind(dAtA, i, uint64(len(m.ErrMsg)))
		i += copy(dAtA[i:], m.ErrMsg)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *NodeIsAliveAttrs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeIsAliveAttrs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.T {
		dAtA[i] = 0x8
		i++
		if m.T {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BackgroundErrorAttrs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackgroundErrorAttrs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ErrMsg) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKind(dAtA, i, uint64(len(m.ErrMsg)))
		i += copy(dAtA[i:], m.ErrMsg)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BackgroundWarnAttrs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BackgroundWarnAttrs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ErrMsg) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKind(dAtA, i, uint64(len(m.ErrMsg)))
		i += copy(dAtA[i:], m.ErrMsg)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *DebugAttrs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DebugAttrs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKind(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *StatisticsAttrs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatisticsAttrs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ErrMsg) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKind(dAtA, i, uint64(len(m.ErrMsg)))
		i += copy(dAtA[i:], m.ErrMsg)
	}
	if m.TotalNetworkBandwidth != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintKind(dAtA, i, uint64(m.TotalNetworkBandwidth.Size()))
		n1, err := m.TotalNetworkBandwidth.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.PeersCount != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintKind(dAtA, i, uint64(m.PeersCount))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintKind(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *NodeStartedAttrs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.T {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NodeStoppedAttrs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ErrMsg)
	if l > 0 {
		n += 1 + l + sovKind(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NodeIsAliveAttrs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.T {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BackgroundErrorAttrs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ErrMsg)
	if l > 0 {
		n += 1 + l + sovKind(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BackgroundWarnAttrs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ErrMsg)
	if l > 0 {
		n += 1 + l + sovKind(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DebugAttrs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovKind(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StatisticsAttrs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ErrMsg)
	if l > 0 {
		n += 1 + l + sovKind(uint64(l))
	}
	if m.TotalNetworkBandwidth != nil {
		l = m.TotalNetworkBandwidth.Size()
		n += 1 + l + sovKind(uint64(l))
	}
	if m.PeersCount != 0 {
		n += 1 + sovKind(uint64(m.PeersCount))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovKind(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozKind(x uint64) (n int) {
	return sovKind(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NodeStartedAttrs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKind
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeStartedAttrs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeStartedAttrs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field T", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKind
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.T = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipKind(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKind
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NodeStoppedAttrs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKind
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeStoppedAttrs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeStoppedAttrs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKind
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKind
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKind(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKind
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NodeIsAliveAttrs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKind
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeIsAliveAttrs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeIsAliveAttrs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field T", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKind
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.T = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipKind(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKind
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BackgroundErrorAttrs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKind
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BackgroundErrorAttrs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackgroundErrorAttrs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKind
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKind
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKind(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKind
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BackgroundWarnAttrs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKind
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BackgroundWarnAttrs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BackgroundWarnAttrs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKind
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKind
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKind(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKind
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DebugAttrs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKind
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DebugAttrs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DebugAttrs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKind
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKind
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKind(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKind
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatisticsAttrs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKind
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatisticsAttrs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatisticsAttrs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKind
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthKind
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalNetworkBandwidth", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKind
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthKind
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TotalNetworkBandwidth == nil {
				m.TotalNetworkBandwidth = &network.BandwidthStats{}
			}
			if err := m.TotalNetworkBandwidth.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeersCount", wireType)
			}
			m.PeersCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKind
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PeersCount |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipKind(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKind
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipKind(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKind
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowKind
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowKind
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthKind
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowKind
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipKind(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthKind = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKind   = fmt.Errorf("proto: integer overflow")
)
