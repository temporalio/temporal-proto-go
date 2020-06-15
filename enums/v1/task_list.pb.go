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
// source: temporal/enums/v1/task_list.proto

package enums

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
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

type TaskListKind int32

const (
	TASK_LIST_KIND_UNSPECIFIED TaskListKind = 0
	TASK_LIST_KIND_NORMAL      TaskListKind = 1
	TASK_LIST_KIND_STICKY      TaskListKind = 2
)

var TaskListKind_name = map[int32]string{
	0: "Unspecified",
	1: "Normal",
	2: "Sticky",
}

var TaskListKind_value = map[string]int32{
	"Unspecified": 0,
	"Normal":      1,
	"Sticky":      2,
}

func (TaskListKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_69f5a5ec4dadfb09, []int{0}
}

type TaskListType int32

const (
	TASK_LIST_TYPE_UNSPECIFIED TaskListType = 0
	// Decision type of task list.
	TASK_LIST_TYPE_DECISION TaskListType = 1
	// Activity type of task list.
	TASK_LIST_TYPE_ACTIVITY TaskListType = 2
)

var TaskListType_name = map[int32]string{
	0: "Unspecified",
	1: "Decision",
	2: "Activity",
}

var TaskListType_value = map[string]int32{
	"Unspecified": 0,
	"Decision":    1,
	"Activity":    2,
}

func (TaskListType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_69f5a5ec4dadfb09, []int{1}
}

func init() {
	proto.RegisterEnum("temporal.enums.v1.TaskListKind", TaskListKind_name, TaskListKind_value)
	proto.RegisterEnum("temporal.enums.v1.TaskListType", TaskListType_name, TaskListType_value)
}

func init() { proto.RegisterFile("temporal/enums/v1/task_list.proto", fileDescriptor_69f5a5ec4dadfb09) }

var fileDescriptor_69f5a5ec4dadfb09 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x49, 0xcd, 0x2d,
	0xc8, 0x2f, 0x4a, 0xcc, 0xd1, 0x4f, 0xcd, 0x2b, 0xcd, 0x2d, 0xd6, 0x2f, 0x33, 0xd4, 0x2f, 0x49,
	0x2c, 0xce, 0x8e, 0xcf, 0xc9, 0x2c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84,
	0x29, 0xd1, 0x03, 0x2b, 0xd1, 0x2b, 0x33, 0xd4, 0x4a, 0xe1, 0xe2, 0x09, 0x49, 0x2c, 0xce, 0xf6,
	0xc9, 0x2c, 0x2e, 0xf1, 0xce, 0xcc, 0x4b, 0x11, 0x92, 0xe3, 0x92, 0x0a, 0x71, 0x0c, 0xf6, 0x8e,
	0xf7, 0xf1, 0x0c, 0x0e, 0x89, 0xf7, 0xf6, 0xf4, 0x73, 0x89, 0x0f, 0xf5, 0x0b, 0x0e, 0x70, 0x75,
	0xf6, 0x74, 0xf3, 0x74, 0x75, 0x11, 0x60, 0x10, 0x92, 0xe4, 0x12, 0x45, 0x93, 0xf7, 0xf3, 0x0f,
	0xf2, 0x75, 0xf4, 0x11, 0x60, 0xc4, 0x22, 0x15, 0x1c, 0xe2, 0xe9, 0xec, 0x1d, 0x29, 0xc0, 0xa4,
	0x95, 0x81, 0xb0, 0x25, 0xa4, 0xb2, 0x20, 0x15, 0xd5, 0x96, 0x90, 0xc8, 0x00, 0x57, 0x34, 0x5b,
	0xa4, 0xb9, 0xc4, 0xd1, 0xe4, 0x5d, 0x5c, 0x9d, 0x3d, 0x83, 0x3d, 0xfd, 0xfd, 0x04, 0x18, 0xb1,
	0x48, 0x3a, 0x3a, 0x87, 0x78, 0x86, 0x79, 0x86, 0x44, 0x0a, 0x30, 0x39, 0x75, 0x32, 0x5e, 0x78,
	0x28, 0xc7, 0x70, 0xe3, 0xa1, 0x1c, 0xc3, 0x87, 0x87, 0x72, 0x8c, 0x0d, 0x8f, 0xe4, 0x18, 0x57,
	0x3c, 0x92, 0x63, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18,
	0x5f, 0x3c, 0x92, 0x63, 0xf8, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5,
	0x18, 0x6e, 0x3c, 0x96, 0x63, 0xe0, 0x12, 0xc9, 0xcc, 0xd7, 0xc3, 0x08, 0x1c, 0x27, 0x5e, 0x98,
	0xa3, 0x03, 0x40, 0xc1, 0x17, 0xc0, 0x18, 0xa5, 0x93, 0x8e, 0xa4, 0x2c, 0x33, 0x5f, 0x1f, 0xc6,
	0xd6, 0x05, 0x87, 0x2f, 0x3c, 0xe0, 0xad, 0xc1, 0x8c, 0x24, 0x36, 0xb0, 0xa8, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x0b, 0x5f, 0x65, 0x21, 0x9a, 0x01, 0x00, 0x00,
}

func (x TaskListKind) String() string {
	s, ok := TaskListKind_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x TaskListType) String() string {
	s, ok := TaskListType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
