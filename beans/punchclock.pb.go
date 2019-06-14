// Code generated by protoc-gen-go.
// source: punchclock.proto
// DO NOT EDIT!

/*
Package beans is a generated protocol buffer package.

It is generated from these files:
	punchclock.proto
	redis.proto
	sendMail.proto

It has these top-level messages:
	Punchclock
	TimeStruct
	Redis
	SendMail
*/
package beans

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Punchclock struct {
	Begin *TimeStruct `protobuf:"bytes,7,opt,name=Begin"    json:"begin"`
	End   *TimeStruct `protobuf:"bytes,8,opt,name=End"      json:"end"`
}

func (m *Punchclock) Reset()                    { *m = Punchclock{} }
func (m *Punchclock) String() string            { return proto.CompactTextString(m) }
func (*Punchclock) ProtoMessage()               {}
func (*Punchclock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Punchclock) GetBegin() *TimeStruct {
	if m != nil {
		return m.Begin
	}
	return nil
}

func (m *Punchclock) GetEnd() *TimeStruct {
	if m != nil {
		return m.End
	}
	return nil
}

type TimeStruct struct {
	Year   string `protobuf:"bytes,1,opt,name=Year"     json:"year"`
	Month  string `protobuf:"bytes,2,opt,name=Month"    json:"month"`
	Day    string `protobuf:"bytes,3,opt,name=Day"      json:"day"`
	Hour   string `protobuf:"bytes,4,opt,name=Hour"     json:"hour"`
	Minute string `protobuf:"bytes,5,opt,name=Minute"   json:"minute"`
	Second string `protobuf:"bytes,6,opt,name=Second"   json:"second"`
}

func (m *TimeStruct) Reset()                    { *m = TimeStruct{} }
func (m *TimeStruct) String() string            { return proto.CompactTextString(m) }
func (*TimeStruct) ProtoMessage()               {}
func (*TimeStruct) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TimeStruct) GetYear() string {
	if m != nil {
		return m.Year
	}
	return ""
}

func (m *TimeStruct) GetMonth() string {
	if m != nil {
		return m.Month
	}
	return ""
}

func (m *TimeStruct) GetDay() string {
	if m != nil {
		return m.Day
	}
	return ""
}

func (m *TimeStruct) GetHour() string {
	if m != nil {
		return m.Hour
	}
	return ""
}

func (m *TimeStruct) GetMinute() string {
	if m != nil {
		return m.Minute
	}
	return ""
}

func (m *TimeStruct) GetSecond() string {
	if m != nil {
		return m.Second
	}
	return ""
}

func init() {
	proto.RegisterType((*Punchclock)(nil), "beans.Punchclock")
	proto.RegisterType((*TimeStruct)(nil), "beans.TimeStruct")
}

func init() { proto.RegisterFile("punchclock.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xcd, 0x4b,
	0xce, 0x48, 0xce, 0xc9, 0x4f, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x4a,
	0x4d, 0xcc, 0x2b, 0x56, 0x8a, 0xe2, 0xe2, 0x0a, 0x80, 0x4b, 0x09, 0xa9, 0x73, 0xb1, 0x3a, 0xa5,
	0xa6, 0x67, 0xe6, 0x49, 0xb0, 0x2b, 0x30, 0x6a, 0x70, 0x1b, 0x09, 0xea, 0x81, 0x15, 0xe9, 0x85,
	0x64, 0xe6, 0xa6, 0x06, 0x97, 0x14, 0x95, 0x26, 0x97, 0x04, 0x41, 0xe4, 0x85, 0x94, 0xb9, 0x98,
	0x5d, 0xf3, 0x52, 0x24, 0x38, 0x70, 0x29, 0x03, 0xc9, 0x2a, 0xf5, 0x30, 0x72, 0x71, 0x21, 0xc4,
	0x84, 0x84, 0xb8, 0x58, 0x22, 0x53, 0x13, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0,
	0x6c, 0x21, 0x11, 0x2e, 0x56, 0xdf, 0xfc, 0xbc, 0x92, 0x0c, 0x09, 0x26, 0xb0, 0x20, 0x84, 0x23,
	0x24, 0xc0, 0xc5, 0xec, 0x92, 0x58, 0x29, 0xc1, 0x0c, 0x16, 0x03, 0x31, 0x41, 0x7a, 0x3d, 0xf2,
	0x4b, 0x8b, 0x24, 0x58, 0x20, 0x7a, 0x41, 0x6c, 0x21, 0x31, 0x2e, 0x36, 0xdf, 0xcc, 0xbc, 0xd2,
	0x92, 0x54, 0x09, 0x56, 0xb0, 0x28, 0x94, 0x07, 0x12, 0x0f, 0x4e, 0x4d, 0xce, 0xcf, 0x4b, 0x91,
	0x60, 0x83, 0x88, 0x43, 0x78, 0x49, 0x6c, 0x60, 0x8f, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x05, 0x73, 0xab, 0x0e, 0x0c, 0x01, 0x00, 0x00,
}
