// Code generated by protoc-gen-go.
// source: github.com/HailoOSS/provisioning-manager-service/proto/listprovisioners/listprovisioners.proto
// DO NOT EDIT!

/*
Package com_HailoOSS_kernel_provisioningmanager_provisioners is a generated protocol buffer package.

It is generated from these files:
	github.com/HailoOSS/provisioning-manager-service/proto/listprovisioners/listprovisioners.proto

It has these top-level messages:
	Request
	Response
*/
package com_HailoOSS_kernel_provisioningmanager_provisioners

import proto "github.com/HailoOSS/protobuf/proto"
import json "encoding/json"
import math "math"
import com_HailoOSS_kernel_provisioningmanager "github.com/HailoOSS/provisioning-manager-service/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	MachineClass     *string `protobuf:"bytes,1,opt,name=machineClass" json:"machineClass,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetMachineClass() string {
	if m != nil && m.MachineClass != nil {
		return *m.MachineClass
	}
	return ""
}

type Response struct {
	Provisioners     []*com_HailoOSS_kernel_provisioningmanager.Provisioner `protobuf:"bytes,1,rep,name=provisioners" json:"provisioners,omitempty"`
	XXX_unrecognized []byte                                                 `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetProvisioners() []*com_HailoOSS_kernel_provisioningmanager.Provisioner {
	if m != nil {
		return m.Provisioners
	}
	return nil
}

func init() {
}