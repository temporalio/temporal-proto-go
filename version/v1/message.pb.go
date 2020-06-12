// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/version/v1/message.proto

package version

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SupportedSDKVersions contains the support versions for SDK.
type SupportedSDKVersions struct {
	GoSdk   string `protobuf:"bytes,1,opt,name=go_sdk,json=goSdk,proto3" json:"go_sdk,omitempty"`
	JavaSdk string `protobuf:"bytes,2,opt,name=java_sdk,json=javaSdk,proto3" json:"java_sdk,omitempty"`
}

func (m *SupportedSDKVersions) Reset()      { *m = SupportedSDKVersions{} }
func (*SupportedSDKVersions) ProtoMessage() {}
func (*SupportedSDKVersions) Descriptor() ([]byte, []int) {
	return fileDescriptor_642dc4375f676645, []int{0}
}
func (m *SupportedSDKVersions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SupportedSDKVersions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SupportedSDKVersions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SupportedSDKVersions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupportedSDKVersions.Merge(m, src)
}
func (m *SupportedSDKVersions) XXX_Size() int {
	return m.Size()
}
func (m *SupportedSDKVersions) XXX_DiscardUnknown() {
	xxx_messageInfo_SupportedSDKVersions.DiscardUnknown(m)
}

var xxx_messageInfo_SupportedSDKVersions proto.InternalMessageInfo

func (m *SupportedSDKVersions) GetGoSdk() string {
	if m != nil {
		return m.GoSdk
	}
	return ""
}

func (m *SupportedSDKVersions) GetJavaSdk() string {
	if m != nil {
		return m.JavaSdk
	}
	return ""
}

type WorkerVersionInfo struct {
	Implementation string `protobuf:"bytes,1,opt,name=implementation,proto3" json:"implementation,omitempty"`
	FeatureVersion string `protobuf:"bytes,2,opt,name=feature_version,json=featureVersion,proto3" json:"feature_version,omitempty"`
}

func (m *WorkerVersionInfo) Reset()      { *m = WorkerVersionInfo{} }
func (*WorkerVersionInfo) ProtoMessage() {}
func (*WorkerVersionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_642dc4375f676645, []int{1}
}
func (m *WorkerVersionInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkerVersionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkerVersionInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkerVersionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkerVersionInfo.Merge(m, src)
}
func (m *WorkerVersionInfo) XXX_Size() int {
	return m.Size()
}
func (m *WorkerVersionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkerVersionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_WorkerVersionInfo proto.InternalMessageInfo

func (m *WorkerVersionInfo) GetImplementation() string {
	if m != nil {
		return m.Implementation
	}
	return ""
}

func (m *WorkerVersionInfo) GetFeatureVersion() string {
	if m != nil {
		return m.FeatureVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*SupportedSDKVersions)(nil), "temporal.version.v1.SupportedSDKVersions")
	proto.RegisterType((*WorkerVersionInfo)(nil), "temporal.version.v1.WorkerVersionInfo")
}

func init() { proto.RegisterFile("temporal/version/v1/message.proto", fileDescriptor_642dc4375f676645) }

var fileDescriptor_642dc4375f676645 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x49, 0xcd, 0x2d,
	0xc8, 0x2f, 0x4a, 0xcc, 0xd1, 0x2f, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0xd3, 0x2f, 0x33, 0xd4,
	0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x86,
	0x29, 0xd1, 0x83, 0x2a, 0xd1, 0x2b, 0x33, 0x54, 0xf2, 0xe0, 0x12, 0x09, 0x2e, 0x2d, 0x28, 0xc8,
	0x2f, 0x2a, 0x49, 0x4d, 0x09, 0x76, 0xf1, 0x0e, 0x83, 0xc8, 0x14, 0x0b, 0x89, 0x72, 0xb1, 0xa5,
	0xe7, 0xc7, 0x17, 0xa7, 0x64, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xb1, 0xa6, 0xe7, 0x07,
	0xa7, 0x64, 0x0b, 0x49, 0x72, 0x71, 0x64, 0x25, 0x96, 0x25, 0x82, 0x25, 0x98, 0xc0, 0x12, 0xec,
	0x20, 0x7e, 0x70, 0x4a, 0xb6, 0x52, 0x0a, 0x97, 0x60, 0x78, 0x7e, 0x51, 0x76, 0x6a, 0x11, 0xd4,
	0x0c, 0xcf, 0xbc, 0xb4, 0x7c, 0x21, 0x35, 0x2e, 0xbe, 0xcc, 0xdc, 0x82, 0x9c, 0xd4, 0xdc, 0xd4,
	0xbc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0x3c, 0xa8, 0x71, 0x68, 0xa2, 0x42, 0xea, 0x5c, 0xfc, 0x69,
	0xa9, 0x89, 0x25, 0xa5, 0x45, 0xa9, 0xf1, 0x50, 0xc7, 0x41, 0x8d, 0xe7, 0x83, 0x0a, 0x43, 0x0d,
	0x75, 0x9a, 0xc2, 0x78, 0xe1, 0xa1, 0x1c, 0xc3, 0x8d, 0x87, 0x72, 0x0c, 0x1f, 0x1e, 0xca, 0x31,
	0x36, 0x3c, 0x92, 0x63, 0x5c, 0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4,
	0x18, 0x1f, 0x3c, 0x92, 0x63, 0x7c, 0xf1, 0x48, 0x8e, 0xe1, 0xc3, 0x23, 0x39, 0xc6, 0x09, 0x8f,
	0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x81, 0x4b, 0x26, 0x33, 0x5f, 0x0f,
	0xee, 0x7d, 0x70, 0x70, 0x20, 0x05, 0x82, 0x13, 0x8f, 0x2f, 0x24, 0xa0, 0x02, 0x40, 0x12, 0x01,
	0x8c, 0x51, 0x06, 0xe9, 0x48, 0xaa, 0x33, 0xf3, 0xf5, 0x61, 0x6c, 0x5d, 0xb0, 0x4e, 0xa4, 0x10,
	0xb6, 0x86, 0x32, 0x93, 0xd8, 0xc0, 0x32, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x81,
	0x82, 0xf6, 0x87, 0x01, 0x00, 0x00,
}

func (this *SupportedSDKVersions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SupportedSDKVersions)
	if !ok {
		that2, ok := that.(SupportedSDKVersions)
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
	if this.GoSdk != that1.GoSdk {
		return false
	}
	if this.JavaSdk != that1.JavaSdk {
		return false
	}
	return true
}
func (this *WorkerVersionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorkerVersionInfo)
	if !ok {
		that2, ok := that.(WorkerVersionInfo)
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
	if this.Implementation != that1.Implementation {
		return false
	}
	if this.FeatureVersion != that1.FeatureVersion {
		return false
	}
	return true
}
func (this *SupportedSDKVersions) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&version.SupportedSDKVersions{")
	s = append(s, "GoSdk: "+fmt.Sprintf("%#v", this.GoSdk)+",\n")
	s = append(s, "JavaSdk: "+fmt.Sprintf("%#v", this.JavaSdk)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *WorkerVersionInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&version.WorkerVersionInfo{")
	s = append(s, "Implementation: "+fmt.Sprintf("%#v", this.Implementation)+",\n")
	s = append(s, "FeatureVersion: "+fmt.Sprintf("%#v", this.FeatureVersion)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMessage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *SupportedSDKVersions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SupportedSDKVersions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SupportedSDKVersions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.JavaSdk) > 0 {
		i -= len(m.JavaSdk)
		copy(dAtA[i:], m.JavaSdk)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.JavaSdk)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.GoSdk) > 0 {
		i -= len(m.GoSdk)
		copy(dAtA[i:], m.GoSdk)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.GoSdk)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WorkerVersionInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkerVersionInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkerVersionInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeatureVersion) > 0 {
		i -= len(m.FeatureVersion)
		copy(dAtA[i:], m.FeatureVersion)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.FeatureVersion)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Implementation) > 0 {
		i -= len(m.Implementation)
		copy(dAtA[i:], m.Implementation)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Implementation)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SupportedSDKVersions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GoSdk)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.JavaSdk)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *WorkerVersionInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Implementation)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.FeatureVersion)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SupportedSDKVersions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SupportedSDKVersions{`,
		`GoSdk:` + fmt.Sprintf("%v", this.GoSdk) + `,`,
		`JavaSdk:` + fmt.Sprintf("%v", this.JavaSdk) + `,`,
		`}`,
	}, "")
	return s
}
func (this *WorkerVersionInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WorkerVersionInfo{`,
		`Implementation:` + fmt.Sprintf("%v", this.Implementation) + `,`,
		`FeatureVersion:` + fmt.Sprintf("%v", this.FeatureVersion) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SupportedSDKVersions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: SupportedSDKVersions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SupportedSDKVersions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoSdk", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoSdk = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JavaSdk", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JavaSdk = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
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
func (m *WorkerVersionInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: WorkerVersionInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkerVersionInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Implementation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Implementation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeatureVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeatureVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)
