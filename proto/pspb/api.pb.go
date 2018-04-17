// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

/*
	Package pspb is a generated protocol buffer package.

	It is generated from these files:
		api.proto

	It has these top-level messages:
		ActionRequestHeader
		IndexRequest
		IndexResponse
*/
package pspb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import meta "github.com/tiglabs/baud/proto/metapb"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RequestContentType int32

const (
	RequestContentType_JSON RequestContentType = 0
)

var RequestContentType_name = map[int32]string{
	0: "JSON",
}
var RequestContentType_value = map[string]int32{
	"JSON": 0,
}

func (x RequestContentType) String() string {
	return proto.EnumName(RequestContentType_name, int32(x))
}
func (RequestContentType) EnumDescriptor() ([]byte, []int) { return fileDescriptorApi, []int{0} }

type WriteResult int32

const (
	WriteResult_CREATED WriteResult = 0
	WriteResult_UPDATED WriteResult = 1
	WriteResult_DELETED WriteResult = 2
)

var WriteResult_name = map[int32]string{
	0: "CREATED",
	1: "UPDATED",
	2: "DELETED",
}
var WriteResult_value = map[string]int32{
	"CREATED": 0,
	"UPDATED": 1,
	"DELETED": 2,
}

func (x WriteResult) String() string {
	return proto.EnumName(WriteResult_name, int32(x))
}
func (WriteResult) EnumDescriptor() ([]byte, []int) { return fileDescriptorApi, []int{1} }

type ActionRequestHeader struct {
	SpaceID     SpaceID     `protobuf:"varint,1,opt,name=space_id,json=spaceId,proto3,casttype=SpaceID" json:"space_id,omitempty"`
	PartitionID PartitionID `protobuf:"varint,2,opt,name=partition_id,json=partitionId,proto3,casttype=PartitionID" json:"partition_id,omitempty"`
}

func (m *ActionRequestHeader) Reset()                    { *m = ActionRequestHeader{} }
func (*ActionRequestHeader) ProtoMessage()               {}
func (*ActionRequestHeader) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{0} }

type IndexRequest struct {
	ActionRequestHeader `protobuf:"bytes,1,opt,name=header,embedded=header" json:"header"`
	Source              []byte             `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	ContentYpe          RequestContentType `protobuf:"varint,3,opt,name=content_ype,json=contentYpe,proto3,enum=RequestContentType" json:"content_ype,omitempty"`
}

func (m *IndexRequest) Reset()                    { *m = IndexRequest{} }
func (*IndexRequest) ProtoMessage()               {}
func (*IndexRequest) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{1} }

type IndexResponse struct {
	meta.ResponseHeader `protobuf:"bytes,1,opt,name=header,embedded=header" json:"header"`
	Id                  string      `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Result              WriteResult `protobuf:"varint,3,opt,name=result,proto3,enum=WriteResult" json:"result,omitempty"`
}

func (m *IndexResponse) Reset()                    { *m = IndexResponse{} }
func (*IndexResponse) ProtoMessage()               {}
func (*IndexResponse) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{2} }

func init() {
	proto.RegisterType((*ActionRequestHeader)(nil), "ActionRequestHeader")
	proto.RegisterType((*IndexRequest)(nil), "IndexRequest")
	proto.RegisterType((*IndexResponse)(nil), "IndexResponse")
	proto.RegisterEnum("RequestContentType", RequestContentType_name, RequestContentType_value)
	proto.RegisterEnum("WriteResult", WriteResult_name, WriteResult_value)
}
func (this *ActionRequestHeader) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ActionRequestHeader)
	if !ok {
		that2, ok := that.(ActionRequestHeader)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.SpaceID != that1.SpaceID {
		return false
	}
	if this.PartitionID != that1.PartitionID {
		return false
	}
	return true
}
func (this *IndexRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IndexRequest)
	if !ok {
		that2, ok := that.(IndexRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ActionRequestHeader.Equal(&that1.ActionRequestHeader) {
		return false
	}
	if !bytes.Equal(this.Source, that1.Source) {
		return false
	}
	if this.ContentYpe != that1.ContentYpe {
		return false
	}
	return true
}
func (this *IndexResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IndexResponse)
	if !ok {
		that2, ok := that.(IndexResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ResponseHeader.Equal(&that1.ResponseHeader) {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.Result != that1.Result {
		return false
	}
	return true
}
func (this *ActionRequestHeader) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pspb.ActionRequestHeader{")
	s = append(s, "SpaceID: "+fmt.Sprintf("%#v", this.SpaceID)+",\n")
	s = append(s, "PartitionID: "+fmt.Sprintf("%#v", this.PartitionID)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *IndexRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&pspb.IndexRequest{")
	s = append(s, "ActionRequestHeader: "+strings.Replace(this.ActionRequestHeader.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "Source: "+fmt.Sprintf("%#v", this.Source)+",\n")
	s = append(s, "ContentYpe: "+fmt.Sprintf("%#v", this.ContentYpe)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *IndexResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&pspb.IndexResponse{")
	s = append(s, "ResponseHeader: "+strings.Replace(this.ResponseHeader.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "Result: "+fmt.Sprintf("%#v", this.Result)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringApi(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ApiGrpc service

type ApiGrpcClient interface {
	Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error)
}

type apiGrpcClient struct {
	cc *grpc.ClientConn
}

func NewApiGrpcClient(cc *grpc.ClientConn) ApiGrpcClient {
	return &apiGrpcClient{cc}
}

func (c *apiGrpcClient) Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error) {
	out := new(IndexResponse)
	err := grpc.Invoke(ctx, "/ApiGrpc/Index", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApiGrpc service

type ApiGrpcServer interface {
	Index(context.Context, *IndexRequest) (*IndexResponse, error)
}

func RegisterApiGrpcServer(s *grpc.Server, srv ApiGrpcServer) {
	s.RegisterService(&_ApiGrpc_serviceDesc, srv)
}

func _ApiGrpc_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiGrpcServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApiGrpc/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiGrpcServer).Index(ctx, req.(*IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApiGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ApiGrpc",
	HandlerType: (*ApiGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _ApiGrpc_Index_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func (m *ActionRequestHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActionRequestHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SpaceID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.SpaceID))
	}
	if m.PartitionID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.PartitionID))
	}
	return i, nil
}

func (m *IndexRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintApi(dAtA, i, uint64(m.ActionRequestHeader.Size()))
	n1, err := m.ActionRequestHeader.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Source) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Source)))
		i += copy(dAtA[i:], m.Source)
	}
	if m.ContentYpe != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.ContentYpe))
	}
	return i, nil
}

func (m *IndexResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintApi(dAtA, i, uint64(m.ResponseHeader.Size()))
	n2, err := m.ResponseHeader.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.Id) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.Result != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Result))
	}
	return i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedActionRequestHeader(r randyApi, easy bool) *ActionRequestHeader {
	this := &ActionRequestHeader{}
	this.SpaceID = SpaceID(r.Uint32())
	this.PartitionID = PartitionID(r.Uint32())
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedIndexRequest(r randyApi, easy bool) *IndexRequest {
	this := &IndexRequest{}
	v1 := NewPopulatedActionRequestHeader(r, easy)
	this.ActionRequestHeader = *v1
	v2 := r.Intn(100)
	this.Source = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.Source[i] = byte(r.Intn(256))
	}
	this.ContentYpe = RequestContentType([]int32{0}[r.Intn(1)])
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedIndexResponse(r randyApi, easy bool) *IndexResponse {
	this := &IndexResponse{}
	v3 := meta.NewPopulatedResponseHeader(r, easy)
	this.ResponseHeader = *v3
	this.Id = string(randStringApi(r))
	this.Result = WriteResult([]int32{0, 1, 2}[r.Intn(3)])
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyApi interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneApi(r randyApi) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringApi(r randyApi) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneApi(r)
	}
	return string(tmps)
}
func randUnrecognizedApi(r randyApi, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldApi(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldApi(dAtA []byte, r randyApi, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateApi(dAtA, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateApi(dAtA, uint64(v5))
	case 1:
		dAtA = encodeVarintPopulateApi(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateApi(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateApi(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateApi(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateApi(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ActionRequestHeader) Size() (n int) {
	var l int
	_ = l
	if m.SpaceID != 0 {
		n += 1 + sovApi(uint64(m.SpaceID))
	}
	if m.PartitionID != 0 {
		n += 1 + sovApi(uint64(m.PartitionID))
	}
	return n
}

func (m *IndexRequest) Size() (n int) {
	var l int
	_ = l
	l = m.ActionRequestHeader.Size()
	n += 1 + l + sovApi(uint64(l))
	l = len(m.Source)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.ContentYpe != 0 {
		n += 1 + sovApi(uint64(m.ContentYpe))
	}
	return n
}

func (m *IndexResponse) Size() (n int) {
	var l int
	_ = l
	l = m.ResponseHeader.Size()
	n += 1 + l + sovApi(uint64(l))
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.Result != 0 {
		n += 1 + sovApi(uint64(m.Result))
	}
	return n
}

func sovApi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ActionRequestHeader) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ActionRequestHeader{`,
		`SpaceID:` + fmt.Sprintf("%v", this.SpaceID) + `,`,
		`PartitionID:` + fmt.Sprintf("%v", this.PartitionID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IndexRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IndexRequest{`,
		`ActionRequestHeader:` + strings.Replace(strings.Replace(this.ActionRequestHeader.String(), "ActionRequestHeader", "ActionRequestHeader", 1), `&`, ``, 1) + `,`,
		`Source:` + fmt.Sprintf("%v", this.Source) + `,`,
		`ContentYpe:` + fmt.Sprintf("%v", this.ContentYpe) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IndexResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IndexResponse{`,
		`ResponseHeader:` + strings.Replace(strings.Replace(this.ResponseHeader.String(), "ResponseHeader", "meta.ResponseHeader", 1), `&`, ``, 1) + `,`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`Result:` + fmt.Sprintf("%v", this.Result) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringApi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ActionRequestHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: ActionRequestHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActionRequestHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceID", wireType)
			}
			m.SpaceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpaceID |= (SpaceID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartitionID", wireType)
			}
			m.PartitionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PartitionID |= (PartitionID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
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
func (m *IndexRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: IndexRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionRequestHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ActionRequestHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = append(m.Source[:0], dAtA[iNdEx:postIndex]...)
			if m.Source == nil {
				m.Source = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentYpe", wireType)
			}
			m.ContentYpe = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ContentYpe |= (RequestContentType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
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
func (m *IndexResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: IndexResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ResponseHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= (WriteResult(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
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
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
				return 0, ErrInvalidLengthApi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApi
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
				next, err := skipApi(dAtA[start:])
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
	ErrInvalidLengthApi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api.proto", fileDescriptorApi) }

var fileDescriptorApi = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x3f, 0x6f, 0x13, 0x31,
	0x18, 0xc6, 0xed, 0x50, 0x92, 0xf4, 0xbd, 0x24, 0x54, 0x2e, 0xaa, 0xaa, 0x0e, 0x4e, 0x15, 0x31,
	0x44, 0x95, 0xb8, 0x53, 0xd3, 0x8a, 0x89, 0x25, 0x69, 0x22, 0x1a, 0x84, 0xa0, 0x72, 0x8b, 0x10,
	0x2c, 0xd5, 0xfd, 0x31, 0xe9, 0x49, 0xed, 0xd9, 0xdc, 0xf9, 0xa4, 0x76, 0x63, 0x67, 0x61, 0xe0,
	0x43, 0xf0, 0x11, 0x18, 0x19, 0x33, 0x76, 0x64, 0x8a, 0x7a, 0xe6, 0x0b, 0x30, 0x22, 0x26, 0x74,
	0x8e, 0x0b, 0xe1, 0x4f, 0xa7, 0xf3, 0xf3, 0xbc, 0xcf, 0x63, 0xfd, 0xfc, 0xea, 0x60, 0xd9, 0x97,
	0xb1, 0x2b, 0x53, 0xa1, 0xc4, 0xc6, 0xfd, 0x49, 0xac, 0x4e, 0xf2, 0xc0, 0x0d, 0xc5, 0x99, 0x37,
	0x11, 0x13, 0xe1, 0x19, 0x3b, 0xc8, 0x5f, 0x1b, 0x65, 0x84, 0x39, 0xd9, 0xb8, 0xb7, 0x10, 0x57,
	0xf1, 0xe4, 0xd4, 0x0f, 0x32, 0x2f, 0xf0, 0xf3, 0x68, 0x5e, 0xf3, 0xce, 0xb8, 0xf2, 0x65, 0x60,
	0x3e, 0xf3, 0x42, 0xe7, 0x1d, 0x86, 0xd5, 0x7e, 0xa8, 0x62, 0x91, 0x30, 0xfe, 0x26, 0xe7, 0x99,
	0xda, 0xe7, 0x7e, 0xc4, 0x53, 0xb2, 0x0d, 0xf5, 0x4c, 0xfa, 0x21, 0x3f, 0x8e, 0xa3, 0x75, 0xbc,
	0x89, 0xbb, 0xcd, 0xc1, 0x9a, 0x9e, 0xb5, 0x6b, 0x87, 0xa5, 0x37, 0x1e, 0xfe, 0xf8, 0x7d, 0x64,
	0x35, 0x93, 0x1b, 0x47, 0xa4, 0x0f, 0x0d, 0xe9, 0xa7, 0x2a, 0x2e, 0x2f, 0x2b, 0x6b, 0x15, 0x53,
	0xa3, 0x7a, 0xd6, 0x76, 0x0e, 0xae, 0x7d, 0x53, 0x5d, 0x94, 0xcc, 0xf9, 0xd5, 0x19, 0x47, 0x9d,
	0x0f, 0x18, 0x1a, 0xe3, 0x24, 0xe2, 0xe7, 0x16, 0x86, 0x3c, 0x80, 0xea, 0x89, 0x01, 0x32, 0x10,
	0x4e, 0xef, 0xae, 0xfb, 0x1f, 0xd8, 0x41, 0x7d, 0x3a, 0x6b, 0xa3, 0xcb, 0x59, 0x1b, 0x33, 0x9b,
	0x26, 0x6b, 0x50, 0xcd, 0x44, 0x9e, 0x86, 0xdc, 0x50, 0x34, 0x98, 0x55, 0x64, 0x17, 0x9c, 0x50,
	0x24, 0x8a, 0x27, 0xea, 0xf8, 0x42, 0xf2, 0xf5, 0x5b, 0x9b, 0xb8, 0xdb, 0xea, 0xad, 0xba, 0xf6,
	0xba, 0xbd, 0xf9, 0xe8, 0xe8, 0x42, 0x72, 0x06, 0x36, 0xf7, 0x52, 0xf2, 0xce, 0x39, 0x34, 0x2d,
	0x55, 0x26, 0x45, 0x92, 0x71, 0xb2, 0xfd, 0x17, 0xd6, 0x1d, 0xf7, 0x7a, 0x74, 0x23, 0x51, 0x0b,
	0x2a, 0x76, 0x27, 0xcb, 0xac, 0x12, 0x47, 0xe4, 0x1e, 0x54, 0x53, 0x9e, 0xe5, 0xa7, 0xca, 0x42,
	0x34, 0xdc, 0x17, 0x69, 0xac, 0x38, 0x33, 0x1e, 0xb3, 0xb3, 0x2d, 0x0a, 0xe4, 0x5f, 0x36, 0x52,
	0x87, 0xa5, 0xc7, 0x87, 0xcf, 0x9e, 0xae, 0xa0, 0xad, 0x5d, 0x70, 0x16, 0x6a, 0xc4, 0x81, 0xda,
	0x1e, 0x1b, 0xf5, 0x8f, 0x46, 0xc3, 0x15, 0x54, 0x8a, 0xe7, 0x07, 0x43, 0x23, 0x70, 0x29, 0x86,
	0xa3, 0x27, 0xa3, 0x52, 0x54, 0x7a, 0x3b, 0x50, 0xeb, 0xcb, 0xf8, 0x51, 0x2a, 0x43, 0xd2, 0x85,
	0xdb, 0xe6, 0x69, 0xa4, 0xe9, 0x2e, 0x2e, 0x7e, 0xa3, 0xe5, 0xfe, 0xf1, 0xe2, 0x0e, 0x1a, 0x3c,
	0x9c, 0x16, 0x14, 0x7d, 0x29, 0x28, 0xba, 0x2a, 0x28, 0xfa, 0x56, 0x50, 0xfc, 0xbd, 0xa0, 0xf8,
	0xad, 0xa6, 0xf8, 0xa3, 0xa6, 0xf8, 0x93, 0xa6, 0xe8, 0xb3, 0xa6, 0x68, 0xaa, 0x29, 0xbe, 0xd4,
	0x14, 0x5f, 0x69, 0x8a, 0xdf, 0x7f, 0xa5, 0x68, 0x1f, 0xbf, 0x5a, 0x92, 0x99, 0x0c, 0x82, 0xaa,
	0xf9, 0xdd, 0x76, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x64, 0xd0, 0xb4, 0xf9, 0xdb, 0x02, 0x00,
	0x00,
}