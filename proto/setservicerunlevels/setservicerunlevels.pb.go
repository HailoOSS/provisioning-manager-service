// Code generated by protoc-gen-go.
// source: github.com/HailoOSS/provisioning-manager-service/proto/setservicerunlevels/setservicerunlevels.proto
// DO NOT EDIT!

/*
Package com_HailoOSS_kernel_provisioningmanager_setservicerunlevels is a generated protocol buffer package.

It is generated from these files:
	github.com/HailoOSS/provisioning-manager-service/proto/setservicerunlevels/setservicerunlevels.proto

It has these top-level messages:
	Request
	Response
*/
package com_HailoOSS_kernel_provisioningmanager_setservicerunlevels

import proto "github.com/HailoOSS/protobuf/proto"
import json "encoding/json"
import math "math"
import com_HailoOSS_kernel_provisioningmanager "github.com/HailoOSS/provisioning-manager-service/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	ServiceName      *string                                         `protobuf:"bytes,1,req,name=serviceName" json:"serviceName,omitempty"`
	Levels           []com_HailoOSS_kernel_provisioningmanager.Level `protobuf:"varint,2,rep,name=levels,enum=com.HailoOSS.kernel.provisioningmanager.Level" json:"levels,omitempty"`
	XXX_unrecognized []byte                                          `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetServiceName() string {
	if m != nil && m.ServiceName != nil {
		return *m.ServiceName
	}
	return ""
}

func (m *Request) GetLevels() []com_HailoOSS_kernel_provisioningmanager.Level {
	if m != nil {
		return m.Levels
	}
	return nil
}

type Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func init() {
}
