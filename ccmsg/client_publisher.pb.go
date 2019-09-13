// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client_publisher.proto

package ccmsg

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type ContentRequest_ClientCacheStatus_Status int32

const (
	// No special information from the client about the cache
	ContentRequest_ClientCacheStatus_DEFAULT ContentRequest_ClientCacheStatus_Status = 0
	// The client cannot use the cache, at this time we're not
	// discriminating why but possible reasons include cannot establish
	// TLS connections, routing failures, PMTUd black holes, or
	// potentially some form of fraud detection (but thats hard!). The
	// publisher cannot trust this enum to make value judgements about
	// the cache, but must honour it as far as delivering service to the
	// client.
	ContentRequest_ClientCacheStatus_UNUSABLE ContentRequest_ClientCacheStatus_Status = 1
)

var ContentRequest_ClientCacheStatus_Status_name = map[int32]string{
	0: "DEFAULT",
	1: "UNUSABLE",
}

var ContentRequest_ClientCacheStatus_Status_value = map[string]int32{
	"DEFAULT":  0,
	"UNUSABLE": 1,
}

func (x ContentRequest_ClientCacheStatus_Status) String() string {
	return proto.EnumName(ContentRequest_ClientCacheStatus_Status_name, int32(x))
}

func (ContentRequest_ClientCacheStatus_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_28618e6c1963c4b4, []int{0, 0, 0}
}

// Client-to-publisher
type ContentRequest struct {
	// TODO: This does not actually need to be repeated with each request, if we are using a connection-oriented
	// transport.
	ClientPublicKey *PublicKey `protobuf:"bytes,1,opt,name=client_public_key,json=clientPublicKey,proto3" json:"client_public_key,omitempty"`
	Path            string     `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// In bytes; the publisher can map that to chunks however it pleases.  range_begin is inclusive and range_end is
	// exclusive.  A value of zero for range_end means "the end of the object".
	RangeBegin uint64 `protobuf:"varint,5,opt,name=range_begin,json=rangeBegin,proto3" json:"range_begin,omitempty"`
	RangeEnd   uint64 `protobuf:"varint,6,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
	SequenceNo uint64 `protobuf:"varint,4,opt,name=sequence_no,json=sequenceNo,proto3" json:"sequence_no,omitempty"`
	// keyed by cache public key
	CacheStatus          map[string]*ContentRequest_ClientCacheStatus `protobuf:"bytes,7,rep,name=cache_status,json=cacheStatus,proto3" json:"cache_status,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *ContentRequest) Reset()         { *m = ContentRequest{} }
func (m *ContentRequest) String() string { return proto.CompactTextString(m) }
func (*ContentRequest) ProtoMessage()    {}
func (*ContentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_28618e6c1963c4b4, []int{0}
}
func (m *ContentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentRequest.Merge(m, src)
}
func (m *ContentRequest) XXX_Size() int {
	return m.Size()
}
func (m *ContentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContentRequest proto.InternalMessageInfo

func (m *ContentRequest) GetClientPublicKey() *PublicKey {
	if m != nil {
		return m.ClientPublicKey
	}
	return nil
}

func (m *ContentRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ContentRequest) GetRangeBegin() uint64 {
	if m != nil {
		return m.RangeBegin
	}
	return 0
}

func (m *ContentRequest) GetRangeEnd() uint64 {
	if m != nil {
		return m.RangeEnd
	}
	return 0
}

func (m *ContentRequest) GetSequenceNo() uint64 {
	if m != nil {
		return m.SequenceNo
	}
	return 0
}

func (m *ContentRequest) GetCacheStatus() map[string]*ContentRequest_ClientCacheStatus {
	if m != nil {
		return m.CacheStatus
	}
	return nil
}

type ContentRequest_ClientCacheStatus struct {
	BacklogDepth         uint64                                  `protobuf:"varint,1,opt,name=backlog_depth,json=backlogDepth,proto3" json:"backlog_depth,omitempty"`
	Status               ContentRequest_ClientCacheStatus_Status `protobuf:"varint,2,opt,name=status,proto3,enum=ccmsg.ContentRequest_ClientCacheStatus_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *ContentRequest_ClientCacheStatus) Reset()         { *m = ContentRequest_ClientCacheStatus{} }
func (m *ContentRequest_ClientCacheStatus) String() string { return proto.CompactTextString(m) }
func (*ContentRequest_ClientCacheStatus) ProtoMessage()    {}
func (*ContentRequest_ClientCacheStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_28618e6c1963c4b4, []int{0, 0}
}
func (m *ContentRequest_ClientCacheStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContentRequest_ClientCacheStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContentRequest_ClientCacheStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContentRequest_ClientCacheStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentRequest_ClientCacheStatus.Merge(m, src)
}
func (m *ContentRequest_ClientCacheStatus) XXX_Size() int {
	return m.Size()
}
func (m *ContentRequest_ClientCacheStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentRequest_ClientCacheStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ContentRequest_ClientCacheStatus proto.InternalMessageInfo

func (m *ContentRequest_ClientCacheStatus) GetBacklogDepth() uint64 {
	if m != nil {
		return m.BacklogDepth
	}
	return 0
}

func (m *ContentRequest_ClientCacheStatus) GetStatus() ContentRequest_ClientCacheStatus_Status {
	if m != nil {
		return m.Status
	}
	return ContentRequest_ClientCacheStatus_DEFAULT
}

// Response to a ContentRequest.
type ContentResponse struct {
	// Identifies the request that this message is a response to.
	RequestSequenceNo uint64 `protobuf:"varint,1,opt,name=request_sequence_no,json=requestSequenceNo,proto3" json:"request_sequence_no,omitempty"`
	// Exactly one of these fields may be present.
	Error                *Error          `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Bundles              []*TicketBundle `protobuf:"bytes,3,rep,name=bundles,proto3" json:"bundles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ContentResponse) Reset()         { *m = ContentResponse{} }
func (m *ContentResponse) String() string { return proto.CompactTextString(m) }
func (*ContentResponse) ProtoMessage()    {}
func (*ContentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_28618e6c1963c4b4, []int{1}
}
func (m *ContentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentResponse.Merge(m, src)
}
func (m *ContentResponse) XXX_Size() int {
	return m.Size()
}
func (m *ContentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContentResponse proto.InternalMessageInfo

func (m *ContentResponse) GetRequestSequenceNo() uint64 {
	if m != nil {
		return m.RequestSequenceNo
	}
	return 0
}

func (m *ContentResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ContentResponse) GetBundles() []*TicketBundle {
	if m != nil {
		return m.Bundles
	}
	return nil
}

func init() {
	proto.RegisterEnum("ccmsg.ContentRequest_ClientCacheStatus_Status", ContentRequest_ClientCacheStatus_Status_name, ContentRequest_ClientCacheStatus_Status_value)
	proto.RegisterType((*ContentRequest)(nil), "ccmsg.ContentRequest")
	proto.RegisterMapType((map[string]*ContentRequest_ClientCacheStatus)(nil), "ccmsg.ContentRequest.CacheStatusEntry")
	proto.RegisterType((*ContentRequest_ClientCacheStatus)(nil), "ccmsg.ContentRequest.ClientCacheStatus")
	proto.RegisterType((*ContentResponse)(nil), "ccmsg.ContentResponse")
}

func init() { proto.RegisterFile("client_publisher.proto", fileDescriptor_28618e6c1963c4b4) }

var fileDescriptor_28618e6c1963c4b4 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x6f, 0xda, 0x30,
	0x14, 0xc5, 0xe5, 0xab, 0xdc, 0xb0, 0x12, 0x5c, 0xad, 0x8a, 0x98, 0xc4, 0x10, 0x95, 0x36, 0x5e,
	0x9a, 0x4a, 0xec, 0x65, 0x9a, 0xd6, 0x87, 0x42, 0xe9, 0x34, 0xad, 0xaa, 0x2a, 0x53, 0x9e, 0xa3,
	0xc4, 0x78, 0x01, 0x11, 0xec, 0x2c, 0x71, 0x26, 0xf1, 0x3f, 0xf6, 0xb0, 0x97, 0x3d, 0xee, 0xbf,
	0xec, 0x71, 0x3f, 0x61, 0x62, 0x7f, 0x64, 0x8a, 0x6d, 0x0a, 0x43, 0x9b, 0xd4, 0x27, 0x7c, 0xcf,
	0x3d, 0x87, 0x7b, 0x8e, 0x6f, 0x0c, 0x27, 0x34, 0x9a, 0x33, 0x2e, 0xbd, 0x38, 0x0b, 0xa2, 0x79,
	0x3a, 0x63, 0x89, 0x1b, 0x27, 0x42, 0x0a, 0x5c, 0xa6, 0x74, 0x99, 0x86, 0xad, 0xb3, 0x70, 0x2e,
	0x67, 0x59, 0xe0, 0x52, 0xb1, 0x3c, 0x0f, 0x45, 0x28, 0xce, 0x55, 0x37, 0xc8, 0x3e, 0xaa, 0x4a,
	0x15, 0xea, 0xa4, 0x55, 0xad, 0x3a, 0x15, 0xcb, 0xa5, 0xe0, 0xba, 0xea, 0x7e, 0x2b, 0xc1, 0xd1,
	0x50, 0x70, 0xc9, 0xb8, 0x24, 0xec, 0x53, 0xc6, 0x52, 0x89, 0xdf, 0x42, 0x73, 0x77, 0x20, 0xf5,
	0x16, 0x6c, 0xe5, 0xa0, 0x0e, 0xea, 0x59, 0x7d, 0xdb, 0x55, 0x23, 0xdd, 0x3b, 0xd5, 0xf8, 0xc0,
	0x56, 0xa4, 0xa1, 0xa9, 0x0f, 0x00, 0xc6, 0x50, 0x8a, 0x7d, 0x39, 0x73, 0x0e, 0x3a, 0xa8, 0x57,
	0x23, 0xea, 0x8c, 0x9f, 0x83, 0x95, 0xf8, 0x3c, 0x64, 0x5e, 0xc0, 0xc2, 0x39, 0x77, 0xca, 0x1d,
	0xd4, 0x2b, 0x11, 0x50, 0xd0, 0x20, 0x47, 0xf0, 0x33, 0xa8, 0x69, 0x02, 0xe3, 0x53, 0xa7, 0xa2,
	0xda, 0x87, 0x0a, 0x18, 0xf1, 0x69, 0xae, 0x4e, 0x73, 0x6b, 0x9c, 0x32, 0x8f, 0x0b, 0xa7, 0xa4,
	0xd5, 0x1b, 0xe8, 0x56, 0xe0, 0xf7, 0x50, 0xa7, 0x3e, 0x9d, 0x31, 0x2f, 0x95, 0xbe, 0xcc, 0x52,
	0xa7, 0xda, 0x29, 0xf6, 0xac, 0xfe, 0x0b, 0xe3, 0xf5, 0xef, 0x74, 0xee, 0x30, 0x67, 0x8e, 0x15,
	0x71, 0xc4, 0x65, 0xb2, 0x22, 0x16, 0xdd, 0x22, 0xad, 0xef, 0x08, 0x9a, 0x43, 0x95, 0x68, 0x87,
	0x87, 0x4f, 0xe1, 0x49, 0xe0, 0xd3, 0x45, 0x24, 0x42, 0x6f, 0xca, 0x62, 0x39, 0x53, 0xb7, 0x51,
	0x22, 0x75, 0x03, 0x5e, 0xe5, 0x18, 0xbe, 0x86, 0x8a, 0x99, 0x9f, 0x47, 0x3f, 0xea, 0xbb, 0xff,
	0x99, 0xbf, 0xff, 0xef, 0xae, 0xfe, 0x21, 0x46, 0xdd, 0x3d, 0x85, 0x8a, 0x19, 0x6b, 0x41, 0xf5,
	0x6a, 0x74, 0x7d, 0x39, 0xb9, 0xb9, 0xb7, 0x0b, 0xb8, 0x0e, 0x87, 0x93, 0xdb, 0xc9, 0xf8, 0x72,
	0x70, 0x33, 0xb2, 0x51, 0x2b, 0x04, 0x7b, 0x3f, 0x08, 0xb6, 0xa1, 0xb8, 0xd9, 0x54, 0x8d, 0xe4,
	0x47, 0x7c, 0x01, 0xe5, 0xcf, 0x7e, 0x94, 0x31, 0xe5, 0xc8, 0xea, 0xbf, 0x7c, 0xa4, 0x23, 0xa2,
	0x55, 0x6f, 0x0e, 0x5e, 0xa3, 0xee, 0x17, 0x04, 0x8d, 0x07, 0x7e, 0x1a, 0x0b, 0x9e, 0x32, 0xec,
	0xc2, 0x71, 0xa2, 0xb5, 0xde, 0xee, 0x62, 0xf4, 0xa5, 0x34, 0x4d, 0x6b, 0xbc, 0xdd, 0x4f, 0x17,
	0xca, 0x2c, 0x49, 0x44, 0x62, 0x6c, 0xd4, 0x8d, 0x8d, 0x51, 0x8e, 0x11, 0xdd, 0xc2, 0x67, 0x50,
	0x0d, 0x32, 0x3e, 0x8d, 0x58, 0xea, 0x14, 0xd5, 0xfa, 0x8e, 0x0d, 0xeb, 0x7e, 0x4e, 0x17, 0x4c,
	0x0e, 0x54, 0x8f, 0x6c, 0x38, 0xfd, 0x3b, 0x68, 0x0c, 0xb7, 0x1f, 0x5e, 0xfe, 0x26, 0xf0, 0x05,
	0xc0, 0x3b, 0x26, 0x8d, 0x57, 0xfc, 0xf4, 0x9f, 0x59, 0x5b, 0x27, 0xfb, 0xb0, 0x8e, 0xd4, 0x2d,
	0x0c, 0xec, 0x1f, 0xeb, 0x36, 0xfa, 0xb9, 0x6e, 0xa3, 0x5f, 0xeb, 0x36, 0xfa, 0xfa, 0xbb, 0x5d,
	0x08, 0x2a, 0xea, 0x85, 0xbc, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0x34, 0x2f, 0x34, 0x52, 0x7f,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientPublisherClient is the client API for ClientPublisher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientPublisherClient interface {
	GetContent(ctx context.Context, in *ContentRequest, opts ...grpc.CallOption) (*ContentResponse, error)
}

type clientPublisherClient struct {
	cc *grpc.ClientConn
}

func NewClientPublisherClient(cc *grpc.ClientConn) ClientPublisherClient {
	return &clientPublisherClient{cc}
}

func (c *clientPublisherClient) GetContent(ctx context.Context, in *ContentRequest, opts ...grpc.CallOption) (*ContentResponse, error) {
	out := new(ContentResponse)
	err := c.cc.Invoke(ctx, "/ccmsg.ClientPublisher/GetContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientPublisherServer is the server API for ClientPublisher service.
type ClientPublisherServer interface {
	GetContent(context.Context, *ContentRequest) (*ContentResponse, error)
}

func RegisterClientPublisherServer(s *grpc.Server, srv ClientPublisherServer) {
	s.RegisterService(&_ClientPublisher_serviceDesc, srv)
}

func _ClientPublisher_GetContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientPublisherServer).GetContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ccmsg.ClientPublisher/GetContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientPublisherServer).GetContent(ctx, req.(*ContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientPublisher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ccmsg.ClientPublisher",
	HandlerType: (*ClientPublisherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContent",
			Handler:    _ClientPublisher_GetContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client_publisher.proto",
}

func (m *ContentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContentRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ClientPublicKey != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintClientPublisher(dAtA, i, uint64(m.ClientPublicKey.Size()))
		n1, err1 := m.ClientPublicKey.MarshalTo(dAtA[i:])
		if err1 != nil {
			return 0, err1
		}
		i += n1
	}
	if len(m.Path) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintClientPublisher(dAtA, i, uint64(len(m.Path)))
		i += copy(dAtA[i:], m.Path)
	}
	if m.SequenceNo != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintClientPublisher(dAtA, i, uint64(m.SequenceNo))
	}
	if m.RangeBegin != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintClientPublisher(dAtA, i, uint64(m.RangeBegin))
	}
	if m.RangeEnd != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintClientPublisher(dAtA, i, uint64(m.RangeEnd))
	}
	if len(m.CacheStatus) > 0 {
		for k, _ := range m.CacheStatus {
			dAtA[i] = 0x3a
			i++
			v := m.CacheStatus[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovClientPublisher(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovClientPublisher(uint64(len(k))) + msgSize
			i = encodeVarintClientPublisher(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintClientPublisher(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintClientPublisher(dAtA, i, uint64(v.Size()))
				n2, err2 := v.MarshalTo(dAtA[i:])
				if err2 != nil {
					return 0, err2
				}
				i += n2
			}
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ContentRequest_ClientCacheStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContentRequest_ClientCacheStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BacklogDepth != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintClientPublisher(dAtA, i, uint64(m.BacklogDepth))
	}
	if m.Status != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintClientPublisher(dAtA, i, uint64(m.Status))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ContentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContentResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RequestSequenceNo != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintClientPublisher(dAtA, i, uint64(m.RequestSequenceNo))
	}
	if m.Error != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintClientPublisher(dAtA, i, uint64(m.Error.Size()))
		n3, err3 := m.Error.MarshalTo(dAtA[i:])
		if err3 != nil {
			return 0, err3
		}
		i += n3
	}
	if len(m.Bundles) > 0 {
		for _, msg := range m.Bundles {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintClientPublisher(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintClientPublisher(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ContentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClientPublicKey != nil {
		l = m.ClientPublicKey.Size()
		n += 1 + l + sovClientPublisher(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovClientPublisher(uint64(l))
	}
	if m.SequenceNo != 0 {
		n += 1 + sovClientPublisher(uint64(m.SequenceNo))
	}
	if m.RangeBegin != 0 {
		n += 1 + sovClientPublisher(uint64(m.RangeBegin))
	}
	if m.RangeEnd != 0 {
		n += 1 + sovClientPublisher(uint64(m.RangeEnd))
	}
	if len(m.CacheStatus) > 0 {
		for k, v := range m.CacheStatus {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovClientPublisher(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovClientPublisher(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovClientPublisher(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ContentRequest_ClientCacheStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BacklogDepth != 0 {
		n += 1 + sovClientPublisher(uint64(m.BacklogDepth))
	}
	if m.Status != 0 {
		n += 1 + sovClientPublisher(uint64(m.Status))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ContentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestSequenceNo != 0 {
		n += 1 + sovClientPublisher(uint64(m.RequestSequenceNo))
	}
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovClientPublisher(uint64(l))
	}
	if len(m.Bundles) > 0 {
		for _, e := range m.Bundles {
			l = e.Size()
			n += 1 + l + sovClientPublisher(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovClientPublisher(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozClientPublisher(x uint64) (n int) {
	return sovClientPublisher(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClientPublisher
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
			return fmt.Errorf("proto: ContentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientPublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientPublisher
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
				return ErrInvalidLengthClientPublisher
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClientPublisher
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClientPublicKey == nil {
				m.ClientPublicKey = &PublicKey{}
			}
			if err := m.ClientPublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientPublisher
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
				return ErrInvalidLengthClientPublisher
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClientPublisher
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceNo", wireType)
			}
			m.SequenceNo = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientPublisher
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SequenceNo |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeBegin", wireType)
			}
			m.RangeBegin = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientPublisher
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RangeBegin |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeEnd", wireType)
			}
			m.RangeEnd = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientPublisher
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RangeEnd |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientPublisher
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
				return ErrInvalidLengthClientPublisher
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClientPublisher
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CacheStatus == nil {
				m.CacheStatus = make(map[string]*ContentRequest_ClientCacheStatus)
			}
			var mapkey string
			var mapvalue *ContentRequest_ClientCacheStatus
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClientPublisher
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClientPublisher
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthClientPublisher
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthClientPublisher
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClientPublisher
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthClientPublisher
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthClientPublisher
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &ContentRequest_ClientCacheStatus{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipClientPublisher(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthClientPublisher
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.CacheStatus[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClientPublisher(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClientPublisher
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClientPublisher
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
func (m *ContentRequest_ClientCacheStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClientPublisher
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
			return fmt.Errorf("proto: ClientCacheStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientCacheStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BacklogDepth", wireType)
			}
			m.BacklogDepth = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientPublisher
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BacklogDepth |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientPublisher
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ContentRequest_ClientCacheStatus_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClientPublisher(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClientPublisher
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClientPublisher
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
func (m *ContentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClientPublisher
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
			return fmt.Errorf("proto: ContentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestSequenceNo", wireType)
			}
			m.RequestSequenceNo = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientPublisher
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestSequenceNo |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientPublisher
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
				return ErrInvalidLengthClientPublisher
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClientPublisher
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &Error{}
			}
			if err := m.Error.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bundles", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientPublisher
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
				return ErrInvalidLengthClientPublisher
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClientPublisher
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bundles = append(m.Bundles, &TicketBundle{})
			if err := m.Bundles[len(m.Bundles)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClientPublisher(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClientPublisher
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthClientPublisher
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
func skipClientPublisher(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClientPublisher
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
					return 0, ErrIntOverflowClientPublisher
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
					return 0, ErrIntOverflowClientPublisher
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
				return 0, ErrInvalidLengthClientPublisher
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthClientPublisher
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowClientPublisher
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
				next, err := skipClientPublisher(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthClientPublisher
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
	ErrInvalidLengthClientPublisher = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClientPublisher   = fmt.Errorf("proto: integer overflow")
)