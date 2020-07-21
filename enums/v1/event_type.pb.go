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
// source: temporal/api/enums/v1/event_type.proto

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

// Whenever this list of events is changed do change the function shouldBufferEvent in mutableStateBuilder.go to make sure to do the correct event ordering.
type EventType int32

const (
	EVENT_TYPE_UNSPECIFIED                                          EventType = 0
	EVENT_TYPE_WORKFLOW_EXECUTION_STARTED                           EventType = 1
	EVENT_TYPE_WORKFLOW_EXECUTION_COMPLETED                         EventType = 2
	EVENT_TYPE_WORKFLOW_EXECUTION_FAILED                            EventType = 3
	EVENT_TYPE_WORKFLOW_EXECUTION_TIMED_OUT                         EventType = 4
	EVENT_TYPE_WORKFLOW_TASK_SCHEDULED                              EventType = 5
	EVENT_TYPE_WORKFLOW_TASK_STARTED                                EventType = 6
	EVENT_TYPE_WORKFLOW_TASK_COMPLETED                              EventType = 7
	EVENT_TYPE_WORKFLOW_TASK_TIMED_OUT                              EventType = 8
	EVENT_TYPE_WORKFLOW_TASK_FAILED                                 EventType = 9
	EVENT_TYPE_ACTIVITY_TASK_SCHEDULED                              EventType = 10
	EVENT_TYPE_ACTIVITY_TASK_STARTED                                EventType = 11
	EVENT_TYPE_ACTIVITY_TASK_COMPLETED                              EventType = 12
	EVENT_TYPE_ACTIVITY_TASK_FAILED                                 EventType = 13
	EVENT_TYPE_ACTIVITY_TASK_TIMED_OUT                              EventType = 14
	EVENT_TYPE_ACTIVITY_TASK_CANCEL_REQUESTED                       EventType = 15
	EVENT_TYPE_ACTIVITY_TASK_CANCELED                               EventType = 16
	EVENT_TYPE_TIMER_STARTED                                        EventType = 17
	EVENT_TYPE_TIMER_FIRED                                          EventType = 18
	EVENT_TYPE_TIMER_CANCELED                                       EventType = 19
	EVENT_TYPE_WORKFLOW_EXECUTION_CANCEL_REQUESTED                  EventType = 20
	EVENT_TYPE_WORKFLOW_EXECUTION_CANCELED                          EventType = 21
	EVENT_TYPE_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED EventType = 22
	EVENT_TYPE_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED    EventType = 23
	EVENT_TYPE_EXTERNAL_WORKFLOW_EXECUTION_CANCEL_REQUESTED         EventType = 24
	EVENT_TYPE_MARKER_RECORDED                                      EventType = 25
	EVENT_TYPE_WORKFLOW_EXECUTION_SIGNALED                          EventType = 26
	EVENT_TYPE_WORKFLOW_EXECUTION_TERMINATED                        EventType = 27
	EVENT_TYPE_WORKFLOW_EXECUTION_CONTINUED_AS_NEW                  EventType = 28
	EVENT_TYPE_START_CHILD_WORKFLOW_EXECUTION_INITIATED             EventType = 29
	EVENT_TYPE_START_CHILD_WORKFLOW_EXECUTION_FAILED                EventType = 30
	EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_STARTED                     EventType = 31
	EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_COMPLETED                   EventType = 32
	EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_FAILED                      EventType = 33
	EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_CANCELED                    EventType = 34
	EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_TIMED_OUT                   EventType = 35
	EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_TERMINATED                  EventType = 36
	EVENT_TYPE_SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED         EventType = 37
	EVENT_TYPE_SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED            EventType = 38
	EVENT_TYPE_EXTERNAL_WORKFLOW_EXECUTION_SIGNALED                 EventType = 39
	EVENT_TYPE_UPSERT_WORKFLOW_SEARCH_ATTRIBUTES                    EventType = 40
)

var EventType_name = map[int32]string{
	0:  "Unspecified",
	1:  "WorkflowExecutionStarted",
	2:  "WorkflowExecutionCompleted",
	3:  "WorkflowExecutionFailed",
	4:  "WorkflowExecutionTimedOut",
	5:  "WorkflowTaskScheduled",
	6:  "WorkflowTaskStarted",
	7:  "WorkflowTaskCompleted",
	8:  "WorkflowTaskTimedOut",
	9:  "WorkflowTaskFailed",
	10: "ActivityTaskScheduled",
	11: "ActivityTaskStarted",
	12: "ActivityTaskCompleted",
	13: "ActivityTaskFailed",
	14: "ActivityTaskTimedOut",
	15: "ActivityTaskCancelRequested",
	16: "ActivityTaskCanceled",
	17: "TimerStarted",
	18: "TimerFired",
	19: "TimerCanceled",
	20: "WorkflowExecutionCancelRequested",
	21: "WorkflowExecutionCanceled",
	22: "RequestCancelExternalWorkflowExecutionInitiated",
	23: "RequestCancelExternalWorkflowExecutionFailed",
	24: "ExternalWorkflowExecutionCancelRequested",
	25: "MarkerRecorded",
	26: "WorkflowExecutionSignaled",
	27: "WorkflowExecutionTerminated",
	28: "WorkflowExecutionContinuedAsNew",
	29: "StartChildWorkflowExecutionInitiated",
	30: "StartChildWorkflowExecutionFailed",
	31: "ChildWorkflowExecutionStarted",
	32: "ChildWorkflowExecutionCompleted",
	33: "ChildWorkflowExecutionFailed",
	34: "ChildWorkflowExecutionCanceled",
	35: "ChildWorkflowExecutionTimedOut",
	36: "ChildWorkflowExecutionTerminated",
	37: "SignalExternalWorkflowExecutionInitiated",
	38: "SignalExternalWorkflowExecutionFailed",
	39: "ExternalWorkflowExecutionSignaled",
	40: "UpsertWorkflowSearchAttributes",
}

var EventType_value = map[string]int32{
	"Unspecified":                                     0,
	"WorkflowExecutionStarted":                        1,
	"WorkflowExecutionCompleted":                      2,
	"WorkflowExecutionFailed":                         3,
	"WorkflowExecutionTimedOut":                       4,
	"WorkflowTaskScheduled":                           5,
	"WorkflowTaskStarted":                             6,
	"WorkflowTaskCompleted":                           7,
	"WorkflowTaskTimedOut":                            8,
	"WorkflowTaskFailed":                              9,
	"ActivityTaskScheduled":                           10,
	"ActivityTaskStarted":                             11,
	"ActivityTaskCompleted":                           12,
	"ActivityTaskFailed":                              13,
	"ActivityTaskTimedOut":                            14,
	"ActivityTaskCancelRequested":                     15,
	"ActivityTaskCanceled":                            16,
	"TimerStarted":                                    17,
	"TimerFired":                                      18,
	"TimerCanceled":                                   19,
	"WorkflowExecutionCancelRequested":                20,
	"WorkflowExecutionCanceled":                       21,
	"RequestCancelExternalWorkflowExecutionInitiated": 22,
	"RequestCancelExternalWorkflowExecutionFailed":    23,
	"ExternalWorkflowExecutionCancelRequested":        24,
	"MarkerRecorded":                                  25,
	"WorkflowExecutionSignaled":                       26,
	"WorkflowExecutionTerminated":                     27,
	"WorkflowExecutionContinuedAsNew":                 28,
	"StartChildWorkflowExecutionInitiated":            29,
	"StartChildWorkflowExecutionFailed":               30,
	"ChildWorkflowExecutionStarted":                   31,
	"ChildWorkflowExecutionCompleted":                 32,
	"ChildWorkflowExecutionFailed":                    33,
	"ChildWorkflowExecutionCanceled":                  34,
	"ChildWorkflowExecutionTimedOut":                  35,
	"ChildWorkflowExecutionTerminated":                36,
	"SignalExternalWorkflowExecutionInitiated":        37,
	"SignalExternalWorkflowExecutionFailed":           38,
	"ExternalWorkflowExecutionSignaled":               39,
	"UpsertWorkflowSearchAttributes":                  40,
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b482d2737d9259e4, []int{0}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.EventType", EventType_name, EventType_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/event_type.proto", fileDescriptor_b482d2737d9259e4)
}

var fileDescriptor_b482d2737d9259e4 = []byte{
	// 706 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0x4d, 0x4f, 0x13, 0x5f,
	0x14, 0xc6, 0x3b, 0xfc, 0xff, 0xa2, 0x1c, 0x15, 0xc7, 0xa3, 0x60, 0x41, 0xb8, 0xbc, 0x97, 0xf7,
	0x96, 0x0a, 0x09, 0xa6, 0x98, 0xe8, 0x30, 0x73, 0x2a, 0x37, 0xb4, 0xd3, 0x3a, 0x73, 0x87, 0x17,
	0x37, 0x13, 0x4c, 0x1a, 0xd3, 0x44, 0x68, 0x83, 0x95, 0x84, 0x9d, 0x1f, 0xc1, 0x4f, 0xe0, 0xda,
	0x8f, 0xe2, 0x92, 0x9d, 0x2c, 0xa5, 0x6c, 0x8c, 0x2b, 0x3e, 0x82, 0x99, 0x32, 0xed, 0x0c, 0x03,
	0x9d, 0x8e, 0xee, 0x9a, 0xde, 0xdf, 0xf3, 0xdc, 0xf3, 0x9c, 0x7b, 0x6e, 0xee, 0x40, 0xa2, 0x56,
	0xda, 0xaf, 0x56, 0x0e, 0xf7, 0x3e, 0xa4, 0xf6, 0xaa, 0xe5, 0x54, 0xe9, 0xe0, 0xd3, 0xfe, 0xc7,
	0xd4, 0x51, 0x3a, 0x55, 0x3a, 0x2a, 0x1d, 0xd4, 0xec, 0xda, 0x71, 0xb5, 0x94, 0xac, 0x1e, 0x56,
	0x6a, 0x15, 0xec, 0x6b, 0x72, 0xc9, 0xbd, 0x6a, 0x39, 0xd9, 0xe0, 0x92, 0x47, 0xe9, 0xb9, 0x1f,
	0xbd, 0xd0, 0x43, 0x0e, 0x2b, 0x8e, 0xab, 0x25, 0x1c, 0x84, 0x7e, 0xda, 0x22, 0x5d, 0xd8, 0x62,
	0xb7, 0x48, 0xb6, 0xa5, 0x9b, 0x45, 0x52, 0x79, 0x96, 0x93, 0x26, 0xc7, 0x70, 0x16, 0xa6, 0x7c,
	0x6b, 0xdb, 0x05, 0x63, 0x33, 0x9b, 0x2b, 0x6c, 0xdb, 0xb4, 0x43, 0xaa, 0x25, 0x78, 0x41, 0xb7,
	0x4d, 0xa1, 0x18, 0x82, 0x34, 0x59, 0xc2, 0x79, 0x98, 0x0e, 0x47, 0xd5, 0x42, 0xbe, 0x98, 0x23,
	0x07, 0xee, 0xc2, 0x19, 0x98, 0x0c, 0x87, 0xb3, 0x0a, 0xcf, 0x91, 0x26, 0xff, 0xd7, 0xd9, 0x56,
	0xf0, 0x3c, 0x69, 0x76, 0xc1, 0x12, 0xf2, 0xff, 0x98, 0x80, 0xf1, 0x9b, 0x60, 0xa1, 0x98, 0x9b,
	0xb6, 0xa9, 0x6e, 0x90, 0x66, 0x39, 0xa6, 0xb7, 0x70, 0x12, 0x46, 0xdb, 0x73, 0x6e, 0xa2, 0xee,
	0x50, 0x37, 0x2f, 0xcc, 0xed, 0x50, 0xce, 0xab, 0xee, 0x0e, 0x4e, 0xc0, 0x48, 0x5b, 0xce, 0xcd,
	0xdb, 0x13, 0x30, 0x53, 0x54, 0xc1, 0xb7, 0xb8, 0xd8, 0x0d, 0x46, 0x80, 0x40, 0x84, 0x00, 0xe7,
	0x46, 0xb8, 0x1b, 0xea, 0xe6, 0x45, 0xb8, 0x17, 0x28, 0xed, 0x2a, 0xe7, 0x96, 0x76, 0x3f, 0xd4,
	0xcc, 0xcb, 0xd9, 0x8b, 0x8b, 0x30, 0xdb, 0x7e, 0x53, 0x45, 0x57, 0x29, 0x67, 0x1b, 0xf4, 0xc6,
	0x22, 0xd3, 0xd9, 0xfb, 0x01, 0x4e, 0xc1, 0x58, 0x07, 0x9c, 0x34, 0x59, 0xc6, 0x21, 0x88, 0xfb,
	0x30, 0x67, 0x3f, 0xa3, 0x15, 0xf4, 0x61, 0x60, 0x88, 0x2f, 0x57, 0xb3, 0xdc, 0x20, 0x4d, 0x46,
	0x1c, 0x86, 0x81, 0x6b, 0x6b, 0x2d, 0xe3, 0x47, 0xf8, 0x0c, 0x92, 0x1d, 0x06, 0x37, 0x58, 0xf3,
	0x63, 0x9c, 0x83, 0x44, 0x14, 0x0d, 0x69, 0x72, 0x1f, 0xaa, 0xf0, 0xd2, 0xc7, 0xba, 0x2e, 0x4d,
	0x53, 0xda, 0x11, 0x64, 0xe8, 0x4a, 0xee, 0x26, 0x0f, 0xae, 0x73, 0xc1, 0x15, 0x67, 0xc3, 0x7e,
	0x7c, 0x05, 0x2f, 0xfe, 0xcd, 0xc4, 0x3d, 0xbd, 0x27, 0xb8, 0x06, 0xab, 0x3e, 0x87, 0x30, 0xc9,
	0xb5, 0xbc, 0x71, 0x64, 0x30, 0xe8, 0x13, 0xe7, 0x15, 0x63, 0x93, 0x0c, 0xdb, 0x20, 0xb5, 0x60,
	0x68, 0xa4, 0xc9, 0x03, 0x9d, 0xfb, 0x61, 0xf2, 0xd7, 0xba, 0xe2, 0x14, 0x32, 0x88, 0x0b, 0x30,
	0xd3, 0xe1, 0x46, 0x93, 0x91, 0xe7, 0x7a, 0x23, 0xf8, 0xd3, 0x08, 0xa7, 0x53, 0xd0, 0x05, 0xd7,
	0x2d, 0xd2, 0x6c, 0xc5, 0xb4, 0x75, 0xda, 0x96, 0x87, 0x70, 0x15, 0x96, 0x7d, 0x9a, 0xc6, 0x90,
	0xd8, 0xea, 0x06, 0xcf, 0x69, 0xe1, 0x5d, 0x1e, 0xc6, 0x15, 0x58, 0x8a, 0x2e, 0x74, 0x3b, 0xcb,
	0x30, 0x05, 0xf3, 0x3e, 0x55, 0x5b, 0xbe, 0x39, 0xac, 0x23, 0x98, 0x86, 0xc5, 0x28, 0x02, 0xef,
	0x82, 0x8e, 0x62, 0x12, 0xe6, 0xa2, 0x48, 0xdc, 0x9a, 0xc6, 0x70, 0x09, 0x16, 0x22, 0x6d, 0xd1,
	0x1c, 0xd3, 0xf1, 0xa8, 0x45, 0x79, 0x17, 0x7d, 0x22, 0x70, 0x36, 0xed, 0x25, 0xde, 0x79, 0x4e,
	0x06, 0xc6, 0xf0, 0x72, 0x2c, 0x22, 0xde, 0x82, 0x29, 0x7c, 0x0e, 0x2b, 0x7f, 0x27, 0x76, 0xfb,
	0x91, 0xc0, 0x65, 0x48, 0x45, 0x9c, 0xfe, 0xd6, 0xa4, 0x4e, 0x07, 0x9a, 0x68, 0x15, 0x4d, 0x32,
	0x84, 0x27, 0x31, 0x49, 0x31, 0xd4, 0x0d, 0x5b, 0x11, 0xc2, 0xe0, 0xeb, 0x96, 0x20, 0x53, 0x9e,
	0x59, 0xff, 0x2a, 0x9d, 0x9c, 0xb1, 0xd8, 0xe9, 0x19, 0x8b, 0x5d, 0x9c, 0x31, 0xe9, 0x73, 0x9d,
	0x49, 0xdf, 0xea, 0x4c, 0xfa, 0x5e, 0x67, 0xd2, 0x49, 0x9d, 0x49, 0x3f, 0xeb, 0x4c, 0xfa, 0x55,
	0x67, 0xb1, 0x8b, 0x3a, 0x93, 0xbe, 0x9c, 0xb3, 0xd8, 0xc9, 0x39, 0x8b, 0x9d, 0x9e, 0xb3, 0x18,
	0xc4, 0xcb, 0x95, 0xe4, 0x8d, 0x4f, 0xf5, 0x7a, 0x6f, 0xeb, 0x9d, 0x2e, 0x3a, 0x2f, 0x7a, 0x51,
	0x7a, 0x3b, 0xf6, 0xde, 0xc7, 0x96, 0x2b, 0x57, 0xbe, 0x00, 0xd6, 0x1a, 0x3f, 0x7e, 0x77, 0xc5,
	0x85, 0x0b, 0x64, 0x32, 0x4a, 0xb5, 0x9c, 0xc9, 0x90, 0xf3, 0x77, 0x26, 0xb3, 0x95, 0x7e, 0xd7,
	0xdd, 0xf8, 0x30, 0x58, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xfc, 0xa8, 0xea, 0x42, 0x08,
	0x00, 0x00,
}

func (x EventType) String() string {
	s, ok := EventType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
