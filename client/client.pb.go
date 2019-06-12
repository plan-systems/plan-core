// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client.proto

/*
	Package client is a generated protocol buffer package.

	It is generated from these files:
		client.proto

	It has these top-level messages:
		LoginPB
		SessionPB
		Region
*/
package client

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import plan "github.com/plan-systems/go-plan/plan"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RgnPts int32

const (
	RgnPts_ORIGIN_X RgnPts = 0
	RgnPts_ORIGIN_Y RgnPts = 1
	RgnPts_ORIGIN_Z RgnPts = 2
)

var RgnPts_name = map[int32]string{
	0: "ORIGIN_X",
	1: "ORIGIN_Y",
	2: "ORIGIN_Z",
}
var RgnPts_value = map[string]int32{
	"ORIGIN_X": 0,
	"ORIGIN_Y": 1,
	"ORIGIN_Z": 2,
}

func (x RgnPts) String() string {
	return proto.EnumName(RgnPts_name, int32(x))
}
func (RgnPts) EnumDescriptor() ([]byte, []int) { return fileDescriptorClient, []int{0} }

type RegionType int32

const (
	// The first three points are the XYZ anchor point, and the 4th is the point gaussian radius (which is optional)
	RegionType_POINT RegionType = 0
	// The first three points are the XYZ anchor point and the the remaining are a sequence of UV cords that make a polygon.
	// †he polygon assumed to be closed, so the last point shouldn't be a repeat of the first point.
	RegionType_POLYGON_2D RegionType = 1
	// The first three points are the XYZ anchor point and the the remaining express a R3 spine.
	RegionType_SPLINE_2D RegionType = 2
)

var RegionType_name = map[int32]string{
	0: "POINT",
	1: "POLYGON_2D",
	2: "SPLINE_2D",
}
var RegionType_value = map[string]int32{
	"POINT":      0,
	"POLYGON_2D": 1,
	"SPLINE_2D":  2,
}

func (x RegionType) String() string {
	return proto.EnumName(RegionType_name, int32(x))
}
func (RegionType) EnumDescriptor() ([]byte, []int) { return fileDescriptorClient, []int{1} }

type LoginPB struct {
}

func (m *LoginPB) Reset()                    { *m = LoginPB{} }
func (m *LoginPB) String() string            { return proto.CompactTextString(m) }
func (*LoginPB) ProtoMessage()               {}
func (*LoginPB) Descriptor() ([]byte, []int) { return fileDescriptorClient, []int{0} }

type SessionPB struct {
}

func (m *SessionPB) Reset()                    { *m = SessionPB{} }
func (m *SessionPB) String() string            { return proto.CompactTextString(m) }
func (*SessionPB) ProtoMessage()               {}
func (*SessionPB) Descriptor() ([]byte, []int) { return fileDescriptorClient, []int{1} }

type Region struct {
	Type RegionType `protobuf:"varint,1,opt,name=type,proto3,enum=client.RegionType" json:"type,omitempty"`
	Pts  []float64  `protobuf:"fixed64,2,rep,packed,name=pts" json:"pts,omitempty"`
	// Specifies the stroke style and color, etc
	StrokeStyleId uint32     `protobuf:"varint,3,opt,name=stroke_style_id,json=strokeStyleId,proto3" json:"stroke_style_id,omitempty"`
	FillStyleId   uint32     `protobuf:"varint,4,opt,name=fill_style_id,json=fillStyleId,proto3" json:"fill_style_id,omitempty"`
	Link          *plan.Link `protobuf:"bytes,5,opt,name=link" json:"link,omitempty"`
	Subs          []*Region  `protobuf:"bytes,6,rep,name=subs" json:"subs,omitempty"`
}

func (m *Region) Reset()                    { *m = Region{} }
func (m *Region) String() string            { return proto.CompactTextString(m) }
func (*Region) ProtoMessage()               {}
func (*Region) Descriptor() ([]byte, []int) { return fileDescriptorClient, []int{2} }

func (m *Region) GetType() RegionType {
	if m != nil {
		return m.Type
	}
	return RegionType_POINT
}

func (m *Region) GetPts() []float64 {
	if m != nil {
		return m.Pts
	}
	return nil
}

func (m *Region) GetStrokeStyleId() uint32 {
	if m != nil {
		return m.StrokeStyleId
	}
	return 0
}

func (m *Region) GetFillStyleId() uint32 {
	if m != nil {
		return m.FillStyleId
	}
	return 0
}

func (m *Region) GetLink() *plan.Link {
	if m != nil {
		return m.Link
	}
	return nil
}

func (m *Region) GetSubs() []*Region {
	if m != nil {
		return m.Subs
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginPB)(nil), "client.LoginPB")
	proto.RegisterType((*SessionPB)(nil), "client.SessionPB")
	proto.RegisterType((*Region)(nil), "client.Region")
	proto.RegisterEnum("client.RgnPts", RgnPts_name, RgnPts_value)
	proto.RegisterEnum("client.RegionType", RegionType_name, RegionType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Workstation service

type WorkstationClient interface {
	// Bootstraps a repo for a given member.
	Login(ctx context.Context, in *LoginPB, opts ...grpc.CallOption) (*SessionPB, error)
}

type workstationClient struct {
	cc *grpc.ClientConn
}

func NewWorkstationClient(cc *grpc.ClientConn) WorkstationClient {
	return &workstationClient{cc}
}

func (c *workstationClient) Login(ctx context.Context, in *LoginPB, opts ...grpc.CallOption) (*SessionPB, error) {
	out := new(SessionPB)
	err := grpc.Invoke(ctx, "/client.Workstation/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Workstation service

type WorkstationServer interface {
	// Bootstraps a repo for a given member.
	Login(context.Context, *LoginPB) (*SessionPB, error)
}

func RegisterWorkstationServer(s *grpc.Server, srv WorkstationServer) {
	s.RegisterService(&_Workstation_serviceDesc, srv)
}

func _Workstation_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkstationServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.Workstation/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkstationServer).Login(ctx, req.(*LoginPB))
	}
	return interceptor(ctx, in, info, handler)
}

var _Workstation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "client.Workstation",
	HandlerType: (*WorkstationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Workstation_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client.proto",
}

func (m *LoginPB) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginPB) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *SessionPB) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionPB) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *Region) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Region) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintClient(dAtA, i, uint64(m.Type))
	}
	if len(m.Pts) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintClient(dAtA, i, uint64(len(m.Pts)*8))
		for _, num := range m.Pts {
			f1 := math.Float64bits(float64(num))
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(f1))
			i += 8
		}
	}
	if m.StrokeStyleId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintClient(dAtA, i, uint64(m.StrokeStyleId))
	}
	if m.FillStyleId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintClient(dAtA, i, uint64(m.FillStyleId))
	}
	if m.Link != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintClient(dAtA, i, uint64(m.Link.Size()))
		n2, err := m.Link.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Subs) > 0 {
		for _, msg := range m.Subs {
			dAtA[i] = 0x32
			i++
			i = encodeVarintClient(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintClient(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LoginPB) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *SessionPB) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *Region) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovClient(uint64(m.Type))
	}
	if len(m.Pts) > 0 {
		n += 1 + sovClient(uint64(len(m.Pts)*8)) + len(m.Pts)*8
	}
	if m.StrokeStyleId != 0 {
		n += 1 + sovClient(uint64(m.StrokeStyleId))
	}
	if m.FillStyleId != 0 {
		n += 1 + sovClient(uint64(m.FillStyleId))
	}
	if m.Link != nil {
		l = m.Link.Size()
		n += 1 + l + sovClient(uint64(l))
	}
	if len(m.Subs) > 0 {
		for _, e := range m.Subs {
			l = e.Size()
			n += 1 + l + sovClient(uint64(l))
		}
	}
	return n
}

func sovClient(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozClient(x uint64) (n int) {
	return sovClient(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LoginPB) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClient
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
			return fmt.Errorf("proto: LoginPB: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginPB: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipClient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClient
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
func (m *SessionPB) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClient
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
			return fmt.Errorf("proto: SessionPB: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionPB: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipClient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClient
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
func (m *Region) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClient
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
			return fmt.Errorf("proto: Region: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Region: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (RegionType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				v2 := float64(math.Float64frombits(v))
				m.Pts = append(m.Pts, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClient
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClient
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					v2 := float64(math.Float64frombits(v))
					m.Pts = append(m.Pts, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Pts", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StrokeStyleId", wireType)
			}
			m.StrokeStyleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StrokeStyleId |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FillStyleId", wireType)
			}
			m.FillStyleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FillStyleId |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Link", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Link == nil {
				m.Link = &plan.Link{}
			}
			if err := m.Link.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subs = append(m.Subs, &Region{})
			if err := m.Subs[len(m.Subs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClient
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
func skipClient(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClient
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
					return 0, ErrIntOverflowClient
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
					return 0, ErrIntOverflowClient
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
				return 0, ErrInvalidLengthClient
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowClient
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
				next, err := skipClient(dAtA[start:])
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
	ErrInvalidLengthClient = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClient   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("client.proto", fileDescriptorClient) }

var fileDescriptorClient = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x4f, 0x8b, 0x9b, 0x40,
	0x18, 0xc6, 0x33, 0xd1, 0xd8, 0xfa, 0xba, 0x66, 0xed, 0x9c, 0x64, 0x0f, 0x22, 0x1e, 0x16, 0xd9,
	0xb2, 0x0a, 0x16, 0x7a, 0xe8, 0x71, 0x69, 0x59, 0x04, 0x51, 0x99, 0x2c, 0xb4, 0xdb, 0x8b, 0x6c,
	0x92, 0xa9, 0x1d, 0x34, 0x8e, 0x64, 0x26, 0x07, 0xbf, 0x49, 0x3f, 0x52, 0x8f, 0xa5, 0x9f, 0xa0,
	0xa4, 0x5f, 0xa4, 0xa8, 0xf9, 0xd7, 0x8b, 0xf8, 0x7b, 0x9e, 0x67, 0x98, 0xf7, 0x7d, 0x06, 0xae,
	0x56, 0x35, 0xa3, 0x8d, 0x0c, 0xda, 0x2d, 0x97, 0x1c, 0x6b, 0x23, 0xdd, 0x84, 0x25, 0x93, 0xdf,
	0x77, 0xcb, 0x60, 0xc5, 0x37, 0x61, 0x5b, 0xbf, 0x34, 0xf7, 0xa2, 0x13, 0x92, 0x6e, 0x44, 0x58,
	0xf2, 0xfb, 0x9e, 0xc3, 0xd3, 0x67, 0x3c, 0xe8, 0xe9, 0xf0, 0x2a, 0xe1, 0x25, 0x6b, 0xf2, 0x07,
	0xcf, 0x00, 0x7d, 0x41, 0x85, 0x60, 0xbc, 0x87, 0xdf, 0x08, 0x34, 0x42, 0x4b, 0xc6, 0x1b, 0x7c,
	0x0b, 0xaa, 0xec, 0x5a, 0x6a, 0x23, 0x17, 0xf9, 0xf3, 0x08, 0x07, 0x87, 0x8b, 0x47, 0xf7, 0xa9,
	0x6b, 0x29, 0x19, 0x7c, 0x6c, 0x81, 0xd2, 0x4a, 0x61, 0x4f, 0x5d, 0xc5, 0x47, 0xa4, 0xff, 0xc5,
	0xb7, 0x70, 0x2d, 0xe4, 0x96, 0x57, 0xb4, 0x10, 0xb2, 0xab, 0x69, 0xc1, 0xd6, 0xb6, 0xe2, 0x22,
	0xdf, 0x24, 0xe6, 0x28, 0x2f, 0x7a, 0x35, 0x5e, 0x63, 0x0f, 0xcc, 0x6f, 0xac, 0xae, 0xcf, 0x29,
	0x75, 0x48, 0x19, 0xbd, 0x78, 0xcc, 0x38, 0xa0, 0xd6, 0xac, 0xa9, 0xec, 0x99, 0x8b, 0x7c, 0x23,
	0x82, 0x60, 0xd8, 0x21, 0x61, 0x4d, 0x45, 0x06, 0x1d, 0x7b, 0xa0, 0x8a, 0xdd, 0x52, 0xd8, 0x9a,
	0xab, 0xf8, 0x46, 0x34, 0xff, 0x7f, 0x4a, 0x32, 0x78, 0x77, 0x11, 0x68, 0xa4, 0x6c, 0x72, 0x29,
	0xf0, 0x15, 0xbc, 0xce, 0x48, 0xfc, 0x18, 0xa7, 0xc5, 0x17, 0x6b, 0x72, 0x41, 0xcf, 0x16, 0xba,
	0xa0, 0xaf, 0xd6, 0xf4, 0xee, 0x3d, 0xc0, 0x79, 0x53, 0xac, 0xc3, 0x2c, 0xcf, 0xe2, 0xf4, 0xc9,
	0x9a, 0xe0, 0x39, 0x40, 0x9e, 0x25, 0xcf, 0x8f, 0x59, 0x5a, 0x44, 0x1f, 0x2d, 0x84, 0x4d, 0xd0,
	0x17, 0x79, 0x12, 0xa7, 0x9f, 0x7a, 0x9c, 0x46, 0x1f, 0xc0, 0xf8, 0xcc, 0xb7, 0x95, 0x90, 0x2f,
	0xb2, 0x2f, 0xf1, 0x2d, 0xcc, 0x86, 0x9e, 0xf1, 0xf5, 0x71, 0xb2, 0x43, 0xed, 0x37, 0x6f, 0x8e,
	0xc2, 0xa9, 0xfc, 0x07, 0xeb, 0xe7, 0xde, 0x41, 0xbf, 0xf6, 0x0e, 0xfa, 0xb3, 0x77, 0xd0, 0x8f,
	0xbf, 0xce, 0x64, 0xa9, 0x0d, 0xaf, 0xf5, 0xee, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x06, 0x3d,
	0xda, 0xa0, 0xf6, 0x01, 0x00, 0x00,
}
