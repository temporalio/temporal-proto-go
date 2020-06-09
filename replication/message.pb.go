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
// source: replication/message.proto

package replication

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

type ClusterReplicationConfiguration struct {
	ClusterName string `protobuf:"bytes,1,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
}

func (m *ClusterReplicationConfiguration) Reset()      { *m = ClusterReplicationConfiguration{} }
func (*ClusterReplicationConfiguration) ProtoMessage() {}
func (*ClusterReplicationConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_80dec7ec196c3419, []int{0}
}
func (m *ClusterReplicationConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterReplicationConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterReplicationConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterReplicationConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterReplicationConfiguration.Merge(m, src)
}
func (m *ClusterReplicationConfiguration) XXX_Size() int {
	return m.Size()
}
func (m *ClusterReplicationConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterReplicationConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterReplicationConfiguration proto.InternalMessageInfo

func (m *ClusterReplicationConfiguration) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

type NamespaceReplicationConfiguration struct {
	ActiveClusterName string                             `protobuf:"bytes,1,opt,name=activeClusterName,proto3" json:"activeClusterName,omitempty"`
	Clusters          []*ClusterReplicationConfiguration `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (m *NamespaceReplicationConfiguration) Reset()      { *m = NamespaceReplicationConfiguration{} }
func (*NamespaceReplicationConfiguration) ProtoMessage() {}
func (*NamespaceReplicationConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_80dec7ec196c3419, []int{1}
}
func (m *NamespaceReplicationConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NamespaceReplicationConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NamespaceReplicationConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NamespaceReplicationConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespaceReplicationConfiguration.Merge(m, src)
}
func (m *NamespaceReplicationConfiguration) XXX_Size() int {
	return m.Size()
}
func (m *NamespaceReplicationConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespaceReplicationConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_NamespaceReplicationConfiguration proto.InternalMessageInfo

func (m *NamespaceReplicationConfiguration) GetActiveClusterName() string {
	if m != nil {
		return m.ActiveClusterName
	}
	return ""
}

func (m *NamespaceReplicationConfiguration) GetClusters() []*ClusterReplicationConfiguration {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterReplicationConfiguration)(nil), "replication.ClusterReplicationConfiguration")
	proto.RegisterType((*NamespaceReplicationConfiguration)(nil), "replication.NamespaceReplicationConfiguration")
}

func init() { proto.RegisterFile("replication/message.proto", fileDescriptor_80dec7ec196c3419) }

var fileDescriptor_80dec7ec196c3419 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x4a, 0x2d, 0xc8,
	0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46, 0x92, 0x52, 0x72, 0xe6, 0x92, 0x77, 0xce, 0x29,
	0x2d, 0x2e, 0x49, 0x2d, 0x0a, 0x42, 0x88, 0x3a, 0xe7, 0xe7, 0xa5, 0x65, 0xa6, 0x97, 0x16, 0x81,
	0x39, 0x42, 0x0a, 0x5c, 0xdc, 0xc9, 0x10, 0x25, 0x7e, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0xc8, 0x42, 0x4a, 0xb3, 0x19, 0xb9, 0x14, 0x41, 0x8c, 0xe2, 0x82, 0xc4, 0xe4,
	0x54, 0x9c, 0xe6, 0xe8, 0x70, 0x09, 0x26, 0x26, 0x97, 0x64, 0x96, 0xa5, 0x3a, 0x63, 0x98, 0x86,
	0x29, 0x21, 0xe4, 0xc1, 0xc5, 0x01, 0xb5, 0xa2, 0x58, 0x82, 0x49, 0x81, 0x59, 0x83, 0xdb, 0x48,
	0x47, 0x0f, 0xc9, 0xe1, 0x7a, 0x04, 0x5c, 0x1d, 0x04, 0xd7, 0xed, 0xd4, 0xc0, 0x78, 0xe1, 0xa1,
	0x1c, 0xc3, 0x8d, 0x87, 0x72, 0x0c, 0x1f, 0x1e, 0xca, 0x31, 0x36, 0x3c, 0x92, 0x63, 0x5c, 0xf1,
	0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x7c,
	0xf1, 0x48, 0x8e, 0xe1, 0xc3, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63,
	0xb8, 0xf1, 0x58, 0x8e, 0x81, 0x4b, 0x36, 0x33, 0x5f, 0xaf, 0x24, 0x35, 0xb7, 0x20, 0xbf, 0x28,
	0x31, 0x07, 0x12, 0x72, 0xc8, 0xf6, 0x07, 0x30, 0x46, 0x69, 0xa6, 0x23, 0x29, 0xc8, 0xcc, 0xd7,
	0x87, 0xb1, 0x75, 0xc1, 0x8a, 0xf5, 0x91, 0x14, 0x27, 0xb1, 0x81, 0x85, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x72, 0x62, 0x8c, 0xcd, 0x96, 0x01, 0x00, 0x00,
}

func (this *ClusterReplicationConfiguration) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClusterReplicationConfiguration)
	if !ok {
		that2, ok := that.(ClusterReplicationConfiguration)
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
	if this.ClusterName != that1.ClusterName {
		return false
	}
	return true
}
func (this *NamespaceReplicationConfiguration) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NamespaceReplicationConfiguration)
	if !ok {
		that2, ok := that.(NamespaceReplicationConfiguration)
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
	if this.ActiveClusterName != that1.ActiveClusterName {
		return false
	}
	if len(this.Clusters) != len(that1.Clusters) {
		return false
	}
	for i := range this.Clusters {
		if !this.Clusters[i].Equal(that1.Clusters[i]) {
			return false
		}
	}
	return true
}
func (this *ClusterReplicationConfiguration) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&replication.ClusterReplicationConfiguration{")
	s = append(s, "ClusterName: "+fmt.Sprintf("%#v", this.ClusterName)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *NamespaceReplicationConfiguration) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&replication.NamespaceReplicationConfiguration{")
	s = append(s, "ActiveClusterName: "+fmt.Sprintf("%#v", this.ActiveClusterName)+",\n")
	if this.Clusters != nil {
		s = append(s, "Clusters: "+fmt.Sprintf("%#v", this.Clusters)+",\n")
	}
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
func (m *ClusterReplicationConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterReplicationConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterReplicationConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClusterName) > 0 {
		i -= len(m.ClusterName)
		copy(dAtA[i:], m.ClusterName)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.ClusterName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NamespaceReplicationConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceReplicationConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NamespaceReplicationConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Clusters) > 0 {
		for iNdEx := len(m.Clusters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Clusters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ActiveClusterName) > 0 {
		i -= len(m.ActiveClusterName)
		copy(dAtA[i:], m.ActiveClusterName)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.ActiveClusterName)))
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
func (m *ClusterReplicationConfiguration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClusterName)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *NamespaceReplicationConfiguration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ActiveClusterName)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if len(m.Clusters) > 0 {
		for _, e := range m.Clusters {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ClusterReplicationConfiguration) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClusterReplicationConfiguration{`,
		`ClusterName:` + fmt.Sprintf("%v", this.ClusterName) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NamespaceReplicationConfiguration) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForClusters := "[]*ClusterReplicationConfiguration{"
	for _, f := range this.Clusters {
		repeatedStringForClusters += strings.Replace(f.String(), "ClusterReplicationConfiguration", "ClusterReplicationConfiguration", 1) + ","
	}
	repeatedStringForClusters += "}"
	s := strings.Join([]string{`&NamespaceReplicationConfiguration{`,
		`ActiveClusterName:` + fmt.Sprintf("%v", this.ActiveClusterName) + `,`,
		`Clusters:` + repeatedStringForClusters + `,`,
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
func (m *ClusterReplicationConfiguration) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ClusterReplicationConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterReplicationConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterName", wireType)
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
			m.ClusterName = string(dAtA[iNdEx:postIndex])
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
func (m *NamespaceReplicationConfiguration) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: NamespaceReplicationConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceReplicationConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveClusterName", wireType)
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
			m.ActiveClusterName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clusters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Clusters = append(m.Clusters, &ClusterReplicationConfiguration{})
			if err := m.Clusters[len(m.Clusters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
