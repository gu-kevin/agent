package api

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"

	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PingRequest struct {
	// current time
	Time *types.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (m *PingRequest) Reset()      { *m = PingRequest{} }
func (*PingRequest) ProtoMessage() {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7672be1ea7f0e766, []int{0}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return m.Size()
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetTime() *types.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

type PingResponse struct {
	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (m *PingResponse) Reset()      { *m = PingResponse{} }
func (*PingResponse) ProtoMessage() {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7672be1ea7f0e766, []int{1}
}
func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return m.Size()
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "uber.proto.atg.buildkite.v1.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "uber.proto.atg.buildkite.v1.PingResponse")
}

func init() {
	proto.RegisterFile("uber/proto/atg/buildkite/v1/get_buildkite_api.proto", fileDescriptor_7672be1ea7f0e766)
}

var fileDescriptor_7672be1ea7f0e766 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x73, 0xa5, 0x88, 0x5e, 0x05, 0x21, 0x38, 0x48, 0x0a, 0xaf, 0xd2, 0xa9, 0x3a, 0xbc,
	0x90, 0x76, 0x76, 0x48, 0x16, 0x71, 0x0b, 0xc5, 0x76, 0xd0, 0x40, 0x49, 0xf4, 0x3c, 0x82, 0x6d,
	0x13, 0x9b, 0x4b, 0x67, 0x3f, 0x82, 0x9f, 0xc1, 0x49, 0xfc, 0x0e, 0xee, 0xe2, 0xd4, 0xb1, 0xa3,
	0xbd, 0x2e, 0x8e, 0xfd, 0x08, 0x92, 0x9c, 0x39, 0x74, 0x29, 0x6e, 0xef, 0xee, 0xfd, 0x7e, 0xef,
	0xff, 0x1e, 0xed, 0xe5, 0x11, 0x9b, 0xd9, 0xe9, 0x2c, 0x11, 0x89, 0x1d, 0x0a, 0x6e, 0x47, 0x79,
	0x3c, 0xbe, 0xbd, 0x8f, 0x05, 0xb3, 0xe7, 0x8e, 0xcd, 0x99, 0x18, 0xe9, 0x8f, 0x51, 0x98, 0xc6,
	0x58, 0x82, 0x66, 0xb3, 0x90, 0x54, 0x8d, 0xa1, 0xe0, 0xa8, 0x19, 0x9c, 0x3b, 0x56, 0x8b, 0x27,
	0x09, 0x1f, 0x33, 0x35, 0x33, 0xca, 0xef, 0x6c, 0x11, 0x4f, 0x58, 0x26, 0xc2, 0x49, 0xaa, 0x8c,
	0xf6, 0x19, 0x6d, 0xf8, 0xf1, 0x94, 0xf7, 0xd9, 0x43, 0xce, 0x32, 0x61, 0x22, 0xad, 0x17, 0xc4,
	0x11, 0x39, 0x26, 0x9d, 0x46, 0xd7, 0x42, 0xa5, 0x63, 0xa5, 0xe3, 0x65, 0xa5, 0xf7, 0x4b, 0xae,
	0x7d, 0x4a, 0xf7, 0x95, 0x9e, 0xa5, 0xc9, 0x34, 0x63, 0xa6, 0x45, 0x77, 0x67, 0x3f, 0x75, 0x39,
	0x63, 0xaf, 0xaf, 0xdf, 0xdd, 0x29, 0x3d, 0x38, 0x67, 0xc2, 0xab, 0xd6, 0x73, 0xfd, 0x0b, 0xf3,
	0x9a, 0xd6, 0x0b, 0xdd, 0xec, 0xe0, 0x96, 0x23, 0xf0, 0xd7, 0x82, 0xd6, 0xc9, 0x3f, 0x48, 0x95,
	0xe7, 0xbd, 0x91, 0xc5, 0x0a, 0x8c, 0xe5, 0x0a, 0x8c, 0xcd, 0x0a, 0xc8, 0xa3, 0x04, 0xf2, 0x22,
	0x81, 0xbc, 0x4b, 0x20, 0x0b, 0x09, 0xe4, 0x53, 0x02, 0xf9, 0x92, 0x60, 0x6c, 0x24, 0x90, 0xa7,
	0x35, 0x18, 0x8b, 0x35, 0x18, 0xcb, 0x35, 0x18, 0xb4, 0x75, 0x93, 0x4c, 0xb6, 0x85, 0x78, 0x87,
	0x7f, 0xae, 0x48, 0x63, 0xbf, 0xc0, 0x7c, 0x72, 0xd5, 0xd0, 0xd4, 0xdc, 0x79, 0xae, 0xd5, 0x07,
	0xbe, 0xeb, 0xbd, 0xd6, 0x9a, 0x83, 0x62, 0x52, 0x89, 0xa0, 0x2b, 0x38, 0x6a, 0x11, 0x87, 0xce,
	0x87, 0xea, 0x06, 0x65, 0x37, 0x70, 0x05, 0x0f, 0x74, 0x37, 0x18, 0x3a, 0xd1, 0x4e, 0x99, 0xdf,
	0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x99, 0xe5, 0x10, 0x61, 0x16, 0x02, 0x00, 0x00,
}

func (this *PingRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PingRequest)
	if !ok {
		that2, ok := that.(PingRequest)
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
	if !this.Time.Equal(that1.Time) {
		return false
	}
	return true
}
func (this *PingResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PingResponse)
	if !ok {
		that2, ok := that.(PingResponse)
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
	if this.Response != that1.Response {
		return false
	}
	return true
}
func (this *PingRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&buildkitev1.PingRequest{")
	if this.Time != nil {
		s = append(s, "Time: "+fmt.Sprintf("%#v", this.Time)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PingResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&buildkitev1.PingResponse{")
	s = append(s, "Response: "+fmt.Sprintf("%#v", this.Response)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGetBuildkiteApi(v interface{}, typ string) string {
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

// GetBuildkiteAPIClient is the client API for GetBuildkiteAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GetBuildkiteAPIClient interface {
	// Ping pings buildkite for the health status.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type getBuildkiteAPIClient struct {
	cc *grpc.ClientConn
}

func NewGetBuildkiteAPIClient(cc *grpc.ClientConn) GetBuildkiteAPIClient {
	return &getBuildkiteAPIClient{cc}
}

func (c *getBuildkiteAPIClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/uber.proto.atg.buildkite.v1.GetBuildkiteAPI/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetBuildkiteAPIServer is the server API for GetBuildkiteAPI service.
type GetBuildkiteAPIServer interface {
	// Ping pings buildkite for the health status.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterGetBuildkiteAPIServer(s *grpc.Server, srv GetBuildkiteAPIServer) {
	s.RegisterService(&_GetBuildkiteAPI_serviceDesc, srv)
}

func _GetBuildkiteAPI_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetBuildkiteAPIServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.proto.atg.buildkite.v1.GetBuildkiteAPI/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetBuildkiteAPIServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetBuildkiteAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "uber.proto.atg.buildkite.v1.GetBuildkiteAPI",
	HandlerType: (*GetBuildkiteAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _GetBuildkiteAPI_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "uber/proto/atg/buildkite/v1/get_buildkite_api.proto",
}

func (m *PingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PingRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Time != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGetBuildkiteApi(dAtA, i, uint64(m.Time.Size()))
		n1, err := m.Time.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *PingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PingResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Response) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGetBuildkiteApi(dAtA, i, uint64(len(m.Response)))
		i += copy(dAtA[i:], m.Response)
	}
	return i, nil
}

func encodeVarintGetBuildkiteApi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Time != nil {
		l = m.Time.Size()
		n += 1 + l + sovGetBuildkiteApi(uint64(l))
	}
	return n
}

func (m *PingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Response)
	if l > 0 {
		n += 1 + l + sovGetBuildkiteApi(uint64(l))
	}
	return n
}

func sovGetBuildkiteApi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGetBuildkiteApi(x uint64) (n int) {
	return sovGetBuildkiteApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PingRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PingRequest{`,
		`Time:` + strings.Replace(fmt.Sprintf("%v", this.Time), "Timestamp", "types.Timestamp", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PingResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PingResponse{`,
		`Response:` + fmt.Sprintf("%v", this.Response) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGetBuildkiteApi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGetBuildkiteApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGetBuildkiteApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGetBuildkiteApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGetBuildkiteApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Time == nil {
				m.Time = &types.Timestamp{}
			}
			if err := m.Time.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGetBuildkiteApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGetBuildkiteApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGetBuildkiteApi
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
func (m *PingResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGetBuildkiteApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGetBuildkiteApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGetBuildkiteApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGetBuildkiteApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Response = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGetBuildkiteApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGetBuildkiteApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGetBuildkiteApi
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
func skipGetBuildkiteApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGetBuildkiteApi
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
					return 0, ErrIntOverflowGetBuildkiteApi
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
					return 0, ErrIntOverflowGetBuildkiteApi
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
			if length < 0 {
				return 0, ErrInvalidLengthGetBuildkiteApi
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthGetBuildkiteApi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGetBuildkiteApi
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
				next, err := skipGetBuildkiteApi(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthGetBuildkiteApi
				}
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
	ErrInvalidLengthGetBuildkiteApi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGetBuildkiteApi   = fmt.Errorf("proto: integer overflow")
)
