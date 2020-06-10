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
// source: tasklist/enum.proto

package tasklist

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
	return fileDescriptor_59d009b09d589e82, []int{0}
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
	return fileDescriptor_59d009b09d589e82, []int{1}
}

func init() {
	proto.RegisterEnum("tasklist.TaskListKind", TaskListKind_name, TaskListKind_value)
	proto.RegisterEnum("tasklist.TaskListType", TaskListType_name, TaskListType_value)
}

func init() { proto.RegisterFile("tasklist/enum.proto", fileDescriptor_59d009b09d589e82) }

var fileDescriptor_59d009b09d589e82 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x49, 0x2c, 0xce,
	0xce, 0xc9, 0x2c, 0x2e, 0xd1, 0x4f, 0xcd, 0x2b, 0xcd, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x80, 0x09, 0x6a, 0xa5, 0x70, 0xf1, 0x84, 0x24, 0x16, 0x67, 0xfb, 0x64, 0x16, 0x97, 0x78,
	0x67, 0xe6, 0xa5, 0x08, 0xc9, 0x71, 0x49, 0x85, 0x38, 0x06, 0x7b, 0xc7, 0xfb, 0x78, 0x06, 0x87,
	0xc4, 0x7b, 0x7b, 0xfa, 0xb9, 0xc4, 0x87, 0xfa, 0x05, 0x07, 0xb8, 0x3a, 0x7b, 0xba, 0x79, 0xba,
	0xba, 0x08, 0x30, 0x08, 0x49, 0x72, 0x89, 0xa2, 0xc9, 0xfb, 0xf9, 0x07, 0xf9, 0x3a, 0xfa, 0x08,
	0x30, 0x62, 0x91, 0x0a, 0x0e, 0xf1, 0x74, 0xf6, 0x8e, 0x14, 0x60, 0xd2, 0xca, 0x40, 0xd8, 0x12,
	0x52, 0x59, 0x90, 0x8a, 0x6a, 0x4b, 0x48, 0x64, 0x80, 0x2b, 0x9a, 0x2d, 0xd2, 0x5c, 0xe2, 0x68,
	0xf2, 0x2e, 0xae, 0xce, 0x9e, 0xc1, 0x9e, 0xfe, 0x7e, 0x02, 0x8c, 0x58, 0x24, 0x1d, 0x9d, 0x43,
	0x3c, 0xc3, 0x3c, 0x43, 0x22, 0x05, 0x98, 0x9c, 0xaa, 0x2e, 0x3c, 0x94, 0x63, 0xb8, 0xf1, 0x50,
	0x8e, 0xe1, 0xc3, 0x43, 0x39, 0xc6, 0x86, 0x47, 0x72, 0x8c, 0x2b, 0x1e, 0xc9, 0x31, 0x9e, 0x78,
	0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x2f, 0x1e, 0xc9, 0x31, 0x7c,
	0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31,
	0x70, 0x49, 0x65, 0xe6, 0xeb, 0x95, 0xa4, 0xe6, 0x16, 0xe4, 0x17, 0x25, 0xe6, 0x40, 0x82, 0x48,
	0x0f, 0x16, 0x42, 0x01, 0x8c, 0x51, 0x6a, 0xe9, 0x48, 0xb2, 0x99, 0xf9, 0xfa, 0x30, 0xb6, 0x2e,
	0x58, 0xa5, 0x3e, 0x4c, 0x65, 0x12, 0x1b, 0x98, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc6,
	0x81, 0xf8, 0x36, 0x73, 0x01, 0x00, 0x00,
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
