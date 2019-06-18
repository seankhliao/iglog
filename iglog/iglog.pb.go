// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iglog.proto

package iglog

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type EventType int32

const (
	EventType_Unknown         EventType = 0
	EventType_FollowerGained  EventType = 1
	EventType_FollowerLost    EventType = 2
	EventType_FollowingGained EventType = 3
	EventType_FollowingLost   EventType = 4
)

var EventType_name = map[int32]string{
	0: "Unknown",
	1: "FollowerGained",
	2: "FollowerLost",
	3: "FollowingGained",
	4: "FollowingLost",
}

var EventType_value = map[string]int32{
	"Unknown":         0,
	"FollowerGained":  1,
	"FollowerLost":    2,
	"FollowingGained": 3,
	"FollowingLost":   4,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e72d68d161447c8, []int{0}
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e72d68d161447c8, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Events struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Events) Reset()         { *m = Events{} }
func (m *Events) String() string { return proto.CompactTextString(m) }
func (*Events) ProtoMessage()    {}
func (*Events) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e72d68d161447c8, []int{1}
}

func (m *Events) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Events.Unmarshal(m, b)
}
func (m *Events) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Events.Marshal(b, m, deterministic)
}
func (m *Events) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Events.Merge(m, src)
}
func (m *Events) XXX_Size() int {
	return xxx_messageInfo_Events.Size(m)
}
func (m *Events) XXX_DiscardUnknown() {
	xxx_messageInfo_Events.DiscardUnknown(m)
}

var xxx_messageInfo_Events proto.InternalMessageInfo

func (m *Events) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type Event struct {
	Time                 string    `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Type                 EventType `protobuf:"varint,2,opt,name=type,proto3,enum=iglog.EventType" json:"type,omitempty"`
	User                 *User     `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e72d68d161447c8, []int{2}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *Event) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_Unknown
}

func (m *Event) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Users struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e72d68d161447c8, []int{3}
}

func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (m *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(m, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Displayname          string   `protobuf:"bytes,3,opt,name=displayname,proto3" json:"displayname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e72d68d161447c8, []int{4}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetDisplayname() string {
	if m != nil {
		return m.Displayname
	}
	return ""
}

func init() {
	proto.RegisterEnum("iglog.EventType", EventType_name, EventType_value)
	proto.RegisterType((*Request)(nil), "iglog.Request")
	proto.RegisterType((*Events)(nil), "iglog.Events")
	proto.RegisterType((*Event)(nil), "iglog.Event")
	proto.RegisterType((*Users)(nil), "iglog.Users")
	proto.RegisterType((*User)(nil), "iglog.User")
}

func init() { proto.RegisterFile("iglog.proto", fileDescriptor_0e72d68d161447c8) }

var fileDescriptor_0e72d68d161447c8 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x72, 0xb2, 0x30,
	0x14, 0xc5, 0xbf, 0xf0, 0x47, 0xe5, 0xa2, 0x7c, 0xf4, 0x76, 0xc3, 0xb8, 0x29, 0x65, 0x5c, 0xa0,
	0x0b, 0x17, 0xf4, 0x19, 0xda, 0x6e, 0x5c, 0x65, 0xf4, 0x01, 0xb0, 0xa4, 0x34, 0x53, 0x4c, 0x28,
	0x89, 0x75, 0x7c, 0x88, 0xbe, 0x73, 0x87, 0x80, 0x0e, 0x9d, 0xce, 0x74, 0x97, 0x73, 0xce, 0x0f,
	0xee, 0xe1, 0x12, 0xf0, 0x79, 0x59, 0xc9, 0x72, 0x5d, 0x37, 0x52, 0x4b, 0x74, 0x8d, 0x48, 0x3c,
	0x18, 0x53, 0xf6, 0x71, 0x64, 0x4a, 0x27, 0x6b, 0x18, 0x3d, 0x7e, 0x32, 0xa1, 0x15, 0x2e, 0x60,
	0xc4, 0xcc, 0x29, 0x22, 0xb1, 0x9d, 0xfa, 0xd9, 0x74, 0xdd, 0x3d, 0x69, 0x62, 0xda, 0x67, 0xc9,
	0x1e, 0x5c, 0x63, 0x20, 0x82, 0xa3, 0xf9, 0x81, 0x45, 0x24, 0x26, 0xa9, 0x47, 0xcd, 0x19, 0x17,
	0xe0, 0xe8, 0x73, 0xcd, 0x22, 0x2b, 0x26, 0x69, 0x90, 0x85, 0xc3, 0x17, 0x6c, 0xcf, 0x35, 0xa3,
	0x26, 0xc5, 0x3b, 0x70, 0x8e, 0x8a, 0x35, 0x91, 0x1d, 0x93, 0xd4, 0xcf, 0xfc, 0x9e, 0xda, 0x29,
	0xd6, 0x50, 0x13, 0x24, 0x2b, 0x70, 0x5b, 0xa5, 0xf0, 0x1e, 0xdc, 0xd6, 0xb8, 0x34, 0xfa, 0x81,
	0x76, 0x49, 0xb2, 0x05, 0xa7, 0x95, 0x18, 0x80, 0xc5, 0x0b, 0x53, 0xc6, 0xa6, 0x16, 0x2f, 0x70,
	0x0e, 0x93, 0x16, 0x10, 0xf9, 0xa1, 0xab, 0xe3, 0xd1, 0xab, 0xc6, 0x18, 0xfc, 0x82, 0xab, 0xba,
	0xca, 0xcf, 0x26, 0xb6, 0x4d, 0x3c, 0xb4, 0x56, 0xaf, 0xe0, 0x5d, 0x5b, 0xa3, 0x0f, 0xe3, 0x9d,
	0x78, 0x17, 0xf2, 0x24, 0xc2, 0x7f, 0x88, 0x10, 0x3c, 0xc9, 0xaa, 0x92, 0x27, 0xd6, 0x3c, 0xe7,
	0x5c, 0xb0, 0x22, 0x24, 0x18, 0xc2, 0xf4, 0xe2, 0x6d, 0xa4, 0xd2, 0xa1, 0x85, 0xb7, 0xf0, 0xbf,
	0x73, 0xb8, 0x28, 0x7b, 0xcc, 0xc6, 0x1b, 0x98, 0x5d, 0x4d, 0xc3, 0x39, 0xd9, 0x17, 0x01, 0xe8,
	0xbc, 0x5c, 0xbf, 0xbc, 0xe1, 0x12, 0x26, 0x66, 0xec, 0x46, 0x96, 0x18, 0xf4, 0x1f, 0xdb, 0xff,
	0xa8, 0xf9, 0x6c, 0xb8, 0x4d, 0x85, 0x4b, 0xf0, 0x2e, 0x33, 0xd5, 0x2f, 0x76, 0x3a, 0x58, 0xd4,
	0x00, 0xe5, 0xa2, 0xfc, 0x1b, 0xdd, 0x8f, 0xcc, 0x35, 0x79, 0xf8, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0xde, 0x2a, 0xe8, 0x79, 0x35, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FollowatchClient is the client API for Followatch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FollowatchClient interface {
	EventLog(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Events, error)
	Followers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Users, error)
	Following(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Users, error)
}

type followatchClient struct {
	cc *grpc.ClientConn
}

func NewFollowatchClient(cc *grpc.ClientConn) FollowatchClient {
	return &followatchClient{cc}
}

func (c *followatchClient) EventLog(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Events, error) {
	out := new(Events)
	err := c.cc.Invoke(ctx, "/iglog.Followatch/EventLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followatchClient) Followers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/iglog.Followatch/Followers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followatchClient) Following(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/iglog.Followatch/Following", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowatchServer is the server API for Followatch service.
type FollowatchServer interface {
	EventLog(context.Context, *Request) (*Events, error)
	Followers(context.Context, *Request) (*Users, error)
	Following(context.Context, *Request) (*Users, error)
}

// UnimplementedFollowatchServer can be embedded to have forward compatible implementations.
type UnimplementedFollowatchServer struct {
}

func (*UnimplementedFollowatchServer) EventLog(ctx context.Context, req *Request) (*Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventLog not implemented")
}
func (*UnimplementedFollowatchServer) Followers(ctx context.Context, req *Request) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Followers not implemented")
}
func (*UnimplementedFollowatchServer) Following(ctx context.Context, req *Request) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Following not implemented")
}

func RegisterFollowatchServer(s *grpc.Server, srv FollowatchServer) {
	s.RegisterService(&_Followatch_serviceDesc, srv)
}

func _Followatch_EventLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowatchServer).EventLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iglog.Followatch/EventLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowatchServer).EventLog(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followatch_Followers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowatchServer).Followers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iglog.Followatch/Followers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowatchServer).Followers(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Followatch_Following_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowatchServer).Following(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iglog.Followatch/Following",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowatchServer).Following(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Followatch_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iglog.Followatch",
	HandlerType: (*FollowatchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EventLog",
			Handler:    _Followatch_EventLog_Handler,
		},
		{
			MethodName: "Followers",
			Handler:    _Followatch_Followers_Handler,
		},
		{
			MethodName: "Following",
			Handler:    _Followatch_Following_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iglog.proto",
}