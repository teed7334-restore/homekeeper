// Code generated by protoc-gen-go.
// source: sendMail.proto
// DO NOT EDIT!

package beans

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SendMail struct {
	To      string ` protobuf:"bytes,1,opt,name=To"            json:"To,omitempty"            form:"to"`
	Cc      string ` protobuf:"bytes,2,opt,name=Cc"            json:"Cc,omitempty"            form:"cc"`
	Subject string ` protobuf:"bytes,3,opt,name=Subject"       json:"Subject,omitempty"       form:"subject"`
	Content string ` protobuf:"bytes,4,opt,name=Content"       json:"Content,omitempty"       form:"content"`
}

func (m *SendMail) Reset()                    { *m = SendMail{} }
func (m *SendMail) String() string            { return proto.CompactTextString(m) }
func (*SendMail) ProtoMessage()               {}
func (*SendMail) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *SendMail) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *SendMail) GetCc() string {
	if m != nil {
		return m.Cc
	}
	return ""
}

func (m *SendMail) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *SendMail) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*SendMail)(nil), "beans.SendMail")
}

func init() { proto.RegisterFile("sendMail.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4e, 0xcd, 0x4b,
	0xf1, 0x4d, 0xcc, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x4a, 0x4d, 0xcc,
	0x2b, 0x56, 0x8a, 0xe3, 0xe2, 0x08, 0x86, 0x4a, 0x08, 0xf1, 0x71, 0x31, 0x85, 0xe4, 0x4b, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x85, 0xe4, 0x83, 0xf8, 0xce, 0xc9, 0x12, 0x4c, 0x10, 0xbe,
	0x73, 0xb2, 0x90, 0x04, 0x17, 0x7b, 0x70, 0x69, 0x52, 0x56, 0x6a, 0x72, 0x89, 0x04, 0x33, 0x58,
	0x10, 0xc6, 0x05, 0xc9, 0x38, 0xe7, 0xe7, 0x95, 0xa4, 0xe6, 0x95, 0x48, 0xb0, 0x40, 0x64, 0xa0,
	0xdc, 0x24, 0x36, 0xb0, 0x6d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x28, 0xd8, 0x50,
	0x7f, 0x00, 0x00, 0x00,
}
