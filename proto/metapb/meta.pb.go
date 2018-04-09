// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: meta.proto

/*
	Package metapb is a generated protocol buffer package.

	It is generated from these files:
		meta.proto

	It has these top-level messages:
		DB
		Space
		Partition
		ReplicaGroup
		PartitionServer
		Route
*/
package metapb

import (
	"fmt"
	"io"
	"math"

	proto "github.com/golang/protobuf/proto"

	_ "github.com/gogo/protobuf/gogoproto"
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

type SpaceStatus int32

const (
	SpaceStatus_SS_Invalid SpaceStatus = 0
	// 初始状态，space刚刚创建，分片还不能提供服务
	SpaceStatus_SS_Init SpaceStatus = 1
	// 准备状态，等待space的初始分片补足三个副本
	SpaceStatus_SS_Prepare SpaceStatus = 2
	// 正常状态，可以提供完全的服务
	SpaceStatus_SS_Running SpaceStatus = 3
	// 标记删除，元数据都保留，允许分片参与调度,但不能分裂
	SpaceStatus_SS_Delete SpaceStatus = 4
	// 正在删除
	SpaceStatus_SS_Deleting SpaceStatus = 5
)

var SpaceStatus_name = map[int32]string{
	0: "SS_Invalid",
	1: "SS_Init",
	2: "SS_Prepare",
	3: "SS_Running",
	4: "SS_Delete",
	5: "SS_Deleting",
}
var SpaceStatus_value = map[string]int32{
	"SS_Invalid":  0,
	"SS_Init":     1,
	"SS_Prepare":  2,
	"SS_Running":  3,
	"SS_Delete":   4,
	"SS_Deleting": 5,
}

func (x SpaceStatus) String() string {
	return proto.EnumName(SpaceStatus_name, int32(x))
}
func (SpaceStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptorMeta, []int{0} }

type PSStatus int32

const (
	PSStatus_PS_Invalid PSStatus = 0
	PSStatus_PS_Initial PSStatus = 1
	PSStatus_PS_Login   PSStatus = 2
	PSStatus_PS_Offline PSStatus = 3
	PSStatus_PS_Logout  PSStatus = 4
)

var PSStatus_name = map[int32]string{
	0: "PS_Invalid",
	1: "PS_Initial",
	2: "PS_Login",
	3: "PS_Offline",
	4: "PS_Logout",
}
var PSStatus_value = map[string]int32{
	"PS_Invalid": 0,
	"PS_Initial": 1,
	"PS_Login":   2,
	"PS_Offline": 3,
	"PS_Logout":  4,
}

func (x PSStatus) String() string {
	return proto.EnumName(PSStatus_name, int32(x))
}
func (PSStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptorMeta, []int{1} }

type DB struct {
	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *DB) Reset()                    { *m = DB{} }
func (m *DB) String() string            { return proto.CompactTextString(m) }
func (*DB) ProtoMessage()               {}
func (*DB) Descriptor() ([]byte, []int) { return fileDescriptorMeta, []int{0} }

func (m *DB) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DB) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Space struct {
	Type   string      `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id     uint32      `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name   string      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	DbId   uint32      `protobuf:"varint,4,opt,name=db_id,json=dbId,proto3" json:"db_id,omitempty"`
	DbName string      `protobuf:"bytes,5,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	Status SpaceStatus `protobuf:"varint,6,opt,name=status,proto3,enum=metapb.SpaceStatus" json:"status,omitempty"`
}

func (m *Space) Reset()                    { *m = Space{} }
func (m *Space) String() string            { return proto.CompactTextString(m) }
func (*Space) ProtoMessage()               {}
func (*Space) Descriptor() ([]byte, []int) { return fileDescriptorMeta, []int{1} }

func (m *Space) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Space) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Space) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Space) GetDbId() uint32 {
	if m != nil {
		return m.DbId
	}
	return 0
}

func (m *Space) GetDbName() string {
	if m != nil {
		return m.DbName
	}
	return ""
}

func (m *Space) GetStatus() SpaceStatus {
	if m != nil {
		return m.Status
	}
	return SpaceStatus_SS_Invalid
}

type Partition struct {
	Type      string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	DbId      uint32 `protobuf:"varint,3,opt,name=db_id,json=dbId,proto3" json:"db_id,omitempty"`
	SpaceId   uint32 `protobuf:"varint,4,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	StartSlot uint32 `protobuf:"varint,5,opt,name=start_slot,json=startSlot,proto3" json:"start_slot,omitempty"`
	EndSlot   uint32 `protobuf:"varint,6,opt,name=end_slot,json=endSlot,proto3" json:"end_slot,omitempty"`
}

func (m *Partition) Reset()                    { *m = Partition{} }
func (m *Partition) String() string            { return proto.CompactTextString(m) }
func (*Partition) ProtoMessage()               {}
func (*Partition) Descriptor() ([]byte, []int) { return fileDescriptorMeta, []int{2} }

func (m *Partition) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Partition) GetDbId() uint32 {
	if m != nil {
		return m.DbId
	}
	return 0
}

func (m *Partition) GetSpaceId() uint32 {
	if m != nil {
		return m.SpaceId
	}
	return 0
}

func (m *Partition) GetStartSlot() uint32 {
	if m != nil {
		return m.StartSlot
	}
	return 0
}

func (m *Partition) GetEndSlot() uint32 {
	if m != nil {
		return m.EndSlot
	}
	return 0
}

type ReplicaGroup struct {
	Id       uint32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Replicas []*PartitionServer `protobuf:"bytes,2,rep,name=replicas" json:"replicas,omitempty"`
}

func (m *ReplicaGroup) Reset()                    { *m = ReplicaGroup{} }
func (m *ReplicaGroup) String() string            { return proto.CompactTextString(m) }
func (*ReplicaGroup) ProtoMessage()               {}
func (*ReplicaGroup) Descriptor() ([]byte, []int) { return fileDescriptorMeta, []int{3} }

func (m *ReplicaGroup) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReplicaGroup) GetReplicas() []*PartitionServer {
	if m != nil {
		return m.Replicas
	}
	return nil
}

type PartitionServer struct {
	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Zone string `protobuf:"bytes,3,opt,name=zone,proto3" json:"zone,omitempty"`
	Ip   string `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	Port string `protobuf:"bytes,5,opt,name=port,proto3" json:"port,omitempty"`
}

func (m *PartitionServer) Reset()                    { *m = PartitionServer{} }
func (m *PartitionServer) String() string            { return proto.CompactTextString(m) }
func (*PartitionServer) ProtoMessage()               {}
func (*PartitionServer) Descriptor() ([]byte, []int) { return fileDescriptorMeta, []int{4} }

func (m *PartitionServer) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PartitionServer) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *PartitionServer) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *PartitionServer) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

type Route struct {
	Partition  *Partition    `protobuf:"bytes,1,opt,name=partition" json:"partition,omitempty"`
	Replicas   *ReplicaGroup `protobuf:"bytes,2,opt,name=replicas" json:"replicas,omitempty"`
	LeaderPSId uint32        `protobuf:"varint,3,opt,name=leaderPSId,proto3" json:"leaderPSId,omitempty"`
}

func (m *Route) Reset()                    { *m = Route{} }
func (m *Route) String() string            { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()               {}
func (*Route) Descriptor() ([]byte, []int) { return fileDescriptorMeta, []int{5} }

func (m *Route) GetPartition() *Partition {
	if m != nil {
		return m.Partition
	}
	return nil
}

func (m *Route) GetReplicas() *ReplicaGroup {
	if m != nil {
		return m.Replicas
	}
	return nil
}

func (m *Route) GetLeaderPSId() uint32 {
	if m != nil {
		return m.LeaderPSId
	}
	return 0
}

func init() {
	proto.RegisterType((*DB)(nil), "metapb.DB")
	proto.RegisterType((*Space)(nil), "metapb.Space")
	proto.RegisterType((*Partition)(nil), "metapb.Partition")
	proto.RegisterType((*ReplicaGroup)(nil), "metapb.ReplicaGroup")
	proto.RegisterType((*PartitionServer)(nil), "metapb.PartitionServer")
	proto.RegisterType((*Route)(nil), "metapb.Route")
	proto.RegisterEnum("metapb.SpaceStatus", SpaceStatus_name, SpaceStatus_value)
	proto.RegisterEnum("metapb.PSStatus", PSStatus_name, PSStatus_value)
}
func (m *DB) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DB) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.Id))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMeta(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *Space) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Space) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMeta(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if m.Id != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.Id))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMeta(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.DbId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.DbId))
	}
	if len(m.DbName) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMeta(dAtA, i, uint64(len(m.DbName)))
		i += copy(dAtA[i:], m.DbName)
	}
	if m.Status != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.Status))
	}
	return i, nil
}

func (m *Partition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Partition) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMeta(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if m.DbId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.DbId))
	}
	if m.SpaceId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.SpaceId))
	}
	if m.StartSlot != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.StartSlot))
	}
	if m.EndSlot != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.EndSlot))
	}
	return i, nil
}

func (m *ReplicaGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplicaGroup) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.Id))
	}
	if len(m.Replicas) > 0 {
		for _, msg := range m.Replicas {
			dAtA[i] = 0x12
			i++
			i = encodeVarintMeta(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *PartitionServer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PartitionServer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.Id))
	}
	if len(m.Zone) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMeta(dAtA, i, uint64(len(m.Zone)))
		i += copy(dAtA[i:], m.Zone)
	}
	if len(m.Ip) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMeta(dAtA, i, uint64(len(m.Ip)))
		i += copy(dAtA[i:], m.Ip)
	}
	if len(m.Port) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMeta(dAtA, i, uint64(len(m.Port)))
		i += copy(dAtA[i:], m.Port)
	}
	return i, nil
}

func (m *Route) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Route) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Partition != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.Partition.Size()))
		n1, err := m.Partition.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Replicas != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.Replicas.Size()))
		n2, err := m.Replicas.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.LeaderPSId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMeta(dAtA, i, uint64(m.LeaderPSId))
	}
	return i, nil
}

func encodeVarintMeta(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DB) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMeta(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMeta(uint64(l))
	}
	return n
}

func (m *Space) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovMeta(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovMeta(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMeta(uint64(l))
	}
	if m.DbId != 0 {
		n += 1 + sovMeta(uint64(m.DbId))
	}
	l = len(m.DbName)
	if l > 0 {
		n += 1 + l + sovMeta(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovMeta(uint64(m.Status))
	}
	return n
}

func (m *Partition) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovMeta(uint64(l))
	}
	if m.DbId != 0 {
		n += 1 + sovMeta(uint64(m.DbId))
	}
	if m.SpaceId != 0 {
		n += 1 + sovMeta(uint64(m.SpaceId))
	}
	if m.StartSlot != 0 {
		n += 1 + sovMeta(uint64(m.StartSlot))
	}
	if m.EndSlot != 0 {
		n += 1 + sovMeta(uint64(m.EndSlot))
	}
	return n
}

func (m *ReplicaGroup) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMeta(uint64(m.Id))
	}
	if len(m.Replicas) > 0 {
		for _, e := range m.Replicas {
			l = e.Size()
			n += 1 + l + sovMeta(uint64(l))
		}
	}
	return n
}

func (m *PartitionServer) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMeta(uint64(m.Id))
	}
	l = len(m.Zone)
	if l > 0 {
		n += 1 + l + sovMeta(uint64(l))
	}
	l = len(m.Ip)
	if l > 0 {
		n += 1 + l + sovMeta(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovMeta(uint64(l))
	}
	return n
}

func (m *Route) Size() (n int) {
	var l int
	_ = l
	if m.Partition != nil {
		l = m.Partition.Size()
		n += 1 + l + sovMeta(uint64(l))
	}
	if m.Replicas != nil {
		l = m.Replicas.Size()
		n += 1 + l + sovMeta(uint64(l))
	}
	if m.LeaderPSId != 0 {
		n += 1 + sovMeta(uint64(m.LeaderPSId))
	}
	return n
}

func sovMeta(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMeta(x uint64) (n int) {
	return sovMeta(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DB) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: DB: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DB: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Space) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: Space: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Space: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DbId", wireType)
			}
			m.DbId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DbId |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DbName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DbName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (SpaceStatus(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Partition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: Partition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Partition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DbId", wireType)
			}
			m.DbId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DbId |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceId", wireType)
			}
			m.SpaceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpaceId |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartSlot", wireType)
			}
			m.StartSlot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartSlot |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndSlot", wireType)
			}
			m.EndSlot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndSlot |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReplicaGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: ReplicaGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplicaGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replicas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Replicas = append(m.Replicas, &PartitionServer{})
			if err := m.Replicas[len(m.Replicas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PartitionServer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: PartitionServer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PartitionServer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Zone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Zone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Route) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: Route: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Route: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Partition", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Partition == nil {
				m.Partition = &Partition{}
			}
			if err := m.Partition.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replicas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Replicas == nil {
				m.Replicas = &ReplicaGroup{}
			}
			if err := m.Replicas.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaderPSId", wireType)
			}
			m.LeaderPSId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LeaderPSId |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMeta(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMeta
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
					return 0, ErrIntOverflowMeta
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
					return 0, ErrIntOverflowMeta
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
				return 0, ErrInvalidLengthMeta
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMeta
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
				next, err := skipMeta(dAtA[start:])
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
	ErrInvalidLengthMeta = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMeta   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("meta.proto", fileDescriptorMeta) }

var fileDescriptorMeta = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xdf, 0x6a, 0xdb, 0x3e,
	0x14, 0xc7, 0x2b, 0x27, 0x71, 0xe2, 0x93, 0x26, 0xf5, 0x4f, 0x2d, 0x34, 0xbf, 0xc1, 0x42, 0xf0,
	0x55, 0xe8, 0x20, 0x1d, 0xe9, 0x1b, 0x94, 0xc2, 0x08, 0x8c, 0xcd, 0x58, 0x17, 0x63, 0x57, 0x46,
	0xae, 0xd4, 0xa0, 0xe1, 0x4a, 0x42, 0x56, 0x0a, 0xdb, 0x1b, 0x6c, 0x4f, 0x30, 0xf6, 0x44, 0xbb,
	0xdc, 0x23, 0x8c, 0xec, 0x45, 0x86, 0xe4, 0x38, 0xf1, 0x96, 0xdd, 0x9d, 0xa3, 0xef, 0xc7, 0xe7,
	0x7b, 0xfe, 0x60, 0x80, 0x47, 0x6e, 0xe9, 0x42, 0x1b, 0x65, 0x15, 0x0e, 0x5d, 0xac, 0x8b, 0x67,
	0x17, 0x6b, 0xb5, 0x56, 0xfe, 0xe9, 0xda, 0x45, 0xb5, 0x9a, 0xcc, 0x21, 0xb8, 0xbb, 0xc5, 0x63,
	0x08, 0x04, 0x9b, 0xa0, 0x19, 0x9a, 0x8f, 0xb2, 0x40, 0x30, 0x8c, 0xa1, 0x2b, 0xe9, 0x23, 0x9f,
	0x04, 0x33, 0x34, 0x8f, 0x32, 0x1f, 0x27, 0xdf, 0x10, 0xf4, 0x88, 0xa6, 0xf7, 0xdc, 0xa9, 0xf6,
	0xa3, 0xe6, 0x9e, 0x8f, 0x32, 0x1f, 0xef, 0x2a, 0x04, 0x47, 0x15, 0x3a, 0x87, 0x0a, 0xf8, 0x1c,
	0x7a, 0xac, 0xc8, 0x05, 0x9b, 0x74, 0x3d, 0xd6, 0x65, 0xc5, 0x8a, 0xe1, 0x4b, 0xe8, 0xb3, 0x22,
	0xf7, 0x6c, 0xcf, 0xb3, 0x21, 0x2b, 0xde, 0x38, 0xfa, 0x05, 0x84, 0x95, 0xa5, 0x76, 0x53, 0x4d,
	0xc2, 0x19, 0x9a, 0x8f, 0x97, 0xe7, 0x8b, 0x7a, 0x90, 0x85, 0x6f, 0x82, 0x78, 0x29, 0xdb, 0x21,
	0xc9, 0x67, 0x04, 0x51, 0x4a, 0x8d, 0x15, 0x56, 0x28, 0xb9, 0x6f, 0x30, 0x68, 0x35, 0xb8, 0x37,
	0xef, 0xb4, 0xcc, 0xff, 0x87, 0x41, 0xe5, 0xaa, 0x1d, 0x9a, 0xea, 0xfb, 0x7c, 0xc5, 0xf0, 0x73,
	0x80, 0xca, 0x52, 0x63, 0xf3, 0xaa, 0x54, 0xd6, 0xb7, 0x36, 0xca, 0x22, 0xff, 0x42, 0x4a, 0x65,
	0xdd, 0x97, 0x5c, 0xb2, 0x5a, 0x0c, 0xeb, 0x2f, 0xb9, 0x64, 0x4e, 0x4a, 0x08, 0x9c, 0x66, 0x5c,
	0x97, 0xe2, 0x9e, 0xbe, 0x32, 0x6a, 0xa3, 0x8f, 0x96, 0x7b, 0x03, 0x03, 0x53, 0xeb, 0xd5, 0x24,
	0x98, 0x75, 0xe6, 0xc3, 0xe5, 0x65, 0x33, 0xda, 0x7e, 0x04, 0xc2, 0xcd, 0x13, 0x37, 0xd9, 0x1e,
	0x4c, 0xde, 0xc3, 0xd9, 0x5f, 0xe2, 0xbf, 0x8e, 0xf6, 0x49, 0xc9, 0xfd, 0xca, 0x5d, 0xec, 0x19,
	0xed, 0x47, 0x8b, 0xb2, 0x40, 0x68, 0xc7, 0x68, 0x65, 0xec, 0x6e, 0xd5, 0x3e, 0x4e, 0xbe, 0x20,
	0xe8, 0x65, 0x6a, 0x63, 0x39, 0xbe, 0x86, 0x48, 0x37, 0x26, 0xbe, 0xf0, 0x70, 0xf9, 0xdf, 0x51,
	0x6b, 0xd9, 0x81, 0xc1, 0x2f, 0xff, 0x18, 0xc5, 0xf1, 0x17, 0x0d, 0xdf, 0x5e, 0xc1, 0x61, 0x0e,
	0x3c, 0x05, 0x28, 0x39, 0x65, 0xdc, 0xa4, 0x64, 0xd5, 0xdc, 0xa2, 0xf5, 0x72, 0xf5, 0x01, 0x86,
	0xad, 0xfb, 0xe2, 0x31, 0x00, 0x21, 0xf9, 0x4a, 0x3e, 0xd1, 0x52, 0xb0, 0xf8, 0x04, 0x0f, 0xa1,
	0xef, 0x73, 0x61, 0x63, 0xb4, 0x13, 0x53, 0xc3, 0x35, 0x35, 0x3c, 0x0e, 0x76, 0x79, 0xb6, 0x91,
	0x52, 0xc8, 0x75, 0xdc, 0xc1, 0x23, 0x88, 0x08, 0xc9, 0xef, 0x78, 0xc9, 0x2d, 0x8f, 0xbb, 0xf8,
	0x0c, 0x86, 0x4d, 0xea, 0xf4, 0xde, 0xd5, 0x3b, 0x18, 0xa4, 0xe4, 0x60, 0x94, 0xb6, 0x8d, 0x9a,
	0x5c, 0x58, 0x41, 0xcb, 0x18, 0xe1, 0x53, 0xc7, 0xe6, 0xaf, 0xd5, 0x5a, 0xc8, 0xda, 0x29, 0x25,
	0xf9, 0xdb, 0x87, 0x87, 0x52, 0x48, 0x5e, 0x3b, 0xd5, 0xaa, 0xda, 0xd8, 0xb8, 0x7b, 0x1b, 0x7f,
	0xdf, 0x4e, 0xd1, 0x8f, 0xed, 0x14, 0xfd, 0xdc, 0x4e, 0xd1, 0xd7, 0x5f, 0xd3, 0x93, 0x22, 0xf4,
	0x7f, 0xdb, 0xcd, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xc2, 0xce, 0x51, 0x99, 0x03, 0x00,
	0x00,
}
